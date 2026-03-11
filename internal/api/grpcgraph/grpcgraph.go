package grpcgraph

import (
	"context"
	"errors"
	"fmt"
	"sync"

	echoer "github.com/alexfalkowski/go-monolith/api/echoer/v1"
	greeter "github.com/alexfalkowski/go-monolith/api/greeter/v1"
	"github.com/alexfalkowski/go-service/v2/transport/grpc/meta"
)

// Register the local gRPC services and their allowed direct calls.
//
// The current graph is:
//
//	greeter.v1.Service -> echoer.v1.Service
func Register(graph *Graph) error {
	if err := graph.Register(ServiceName(greeter.Service_ServiceDesc.ServiceName)); err != nil {
		return err
	}

	if err := graph.Register(ServiceName(echoer.Service_ServiceDesc.ServiceName)); err != nil {
		return err
	}

	if err := graph.AllowCall(
		ServiceName(greeter.Service_ServiceDesc.ServiceName),
		ServiceName(echoer.Service_ServiceDesc.ServiceName),
	); err != nil {
		return err
	}

	return nil
}

// ServiceName identifies a gRPC service by descriptor name.
type ServiceName string

const callerMetadataKey = "x-grpc-caller-service"

// New creates a new in-memory service graph.
func New() *Graph {
	return &Graph{
		services: make(map[ServiceName]struct{}),
		outgoing: make(map[ServiceName]map[ServiceName]struct{}),
		incoming: make(map[ServiceName]map[ServiceName]struct{}),
	}
}

// Graph is an in-memory registry of allowed gRPC service-to-service calls.
type Graph struct {
	services map[ServiceName]struct{}
	outgoing map[ServiceName]map[ServiceName]struct{}
	incoming map[ServiceName]map[ServiceName]struct{}
	mu       sync.RWMutex
}

// Register adds a service to the graph.
func (g *Graph) Register(service ServiceName) error {
	if service == "" {
		return errors.New("grpcgraph: service name is required")
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	g.services[service] = struct{}{}

	if _, ok := g.outgoing[service]; !ok {
		g.outgoing[service] = make(map[ServiceName]struct{})
	}

	if _, ok := g.incoming[service]; !ok {
		g.incoming[service] = make(map[ServiceName]struct{})
	}

	return nil
}

// AllowCall registers a direct allowed call edge from caller to callee.
func (g *Graph) AllowCall(caller, callee ServiceName) error {
	if caller == "" {
		return errors.New("grpcgraph: caller service name is required")
	}

	if callee == "" {
		return errors.New("grpcgraph: callee service name is required")
	}

	if caller == callee {
		return fmt.Errorf("grpcgraph: service %q cannot call itself", caller)
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	if _, ok := g.services[caller]; !ok {
		return fmt.Errorf("grpcgraph: caller service %q is not registered", caller)
	}

	if _, ok := g.services[callee]; !ok {
		return fmt.Errorf("grpcgraph: callee service %q is not registered", callee)
	}

	g.outgoing[caller][callee] = struct{}{}
	g.incoming[callee][caller] = struct{}{}

	return nil
}

// CanCall reports whether caller is allowed to call callee directly.
func (g *Graph) CanCall(caller, callee ServiceName) bool {
	g.mu.RLock()
	defer g.mu.RUnlock()

	callees, ok := g.outgoing[caller]
	if !ok {
		return false
	}

	_, ok = callees[callee]
	return ok
}

// CanBeCalledBy reports whether target may be called by caller directly.
func (g *Graph) CanBeCalledBy(target, caller ServiceName) bool {
	g.mu.RLock()
	defer g.mu.RUnlock()

	callers, ok := g.incoming[target]
	if !ok {
		return false
	}

	_, ok = callers[caller]
	return ok
}

// WithCaller adds the caller service name to outgoing gRPC metadata.
func WithCaller(ctx context.Context, caller ServiceName) context.Context {
	if caller == "" {
		return ctx
	}

	md := meta.ExtractOutgoing(ctx)
	md.Append(callerMetadataKey, string(caller))

	return meta.NewOutgoingContext(ctx, md)
}

// Caller extracts the caller service name from incoming metadata.
func Caller(ctx context.Context) (ServiceName, bool) {
	md := meta.ExtractIncoming(ctx)

	values := md.Get(callerMetadataKey)
	if len(values) == 0 {
		return "", false
	}

	caller := ServiceName(values[0])
	if caller == "" {
		return "", false
	}

	return caller, true
}
