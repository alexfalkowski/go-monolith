# frozen_string_literal: true

When('I send the message to echoer which performs in {int} ms') do |time|
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Example-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }

  expect { Example::V1.echoer.echo('test', opts) }.to perform_under(time).ms
end

When('I send the message to greeter which performs in {int} ms') do |time|
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Example-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }

  expect { Example::V1.greeter.greet('test', opts) }.to perform_under(time).ms
end
