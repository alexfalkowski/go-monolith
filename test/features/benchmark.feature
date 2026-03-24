@benchmark
Feature: Benchmark API
  Make sure these endpoints perform at their best.

  Scenario: Echo in good time frame and memory.
    When I send the message to echoer which performs in 15 ms
    And the process 'server' should consume less than '70mb' of memory

  Scenario: Greet in good time frame and memory.
    When I send the message to greeter which performs in 15 ms
    And the process 'server' should consume less than '70mb' of memory
