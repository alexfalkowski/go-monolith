Feature: Greeter

  Greeter wil greet your name.

  Scenario: Greet name
    When I send the name "Bob" to greeter
    Then I should receive "Hello Bob" from greeter
