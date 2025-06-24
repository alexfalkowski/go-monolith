Feature: Echoer

  Echoer just replies with what is sent.

  Scenario: Echo text
    When I send the message "test" to echoer
    Then I should receive a message of "test" from echoer
