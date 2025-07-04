# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'grpc/health/v1/health_services_pb'

require 'echoer/v1/http'
require 'echoer/v1/service_services_pb'
require 'greeter/v1/http'
require 'greeter/v1/service_services_pb'

module Example
  class << self
    def config
      @config ||= Nonnative.configurations('.config/server.yml')
    end

    def health_grpc
      @health_grpc ||= Grpc::Health::V1::Health::Stub.new('localhost:12000', :this_channel_is_insecure, channel_args: Example.user_agent)
    end

    def user_agent
      @user_agent ||= Nonnative::Header.grpc_user_agent('Example-ruby-client/1.0 gRPC/1.0')
    end
  end

  module V1
    class << self
      def echoer
        @echoer ||= Echoer::V1::HTTP.new('http://localhost:11000')
      end

      def greeter
        @greeter ||= Greeter::V1::HTTP.new('http://localhost:11000')
      end
    end
  end
end
