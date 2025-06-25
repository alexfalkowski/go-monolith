# frozen_string_literal: true

When('I send the name {string} to greeter') do |name|
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Example-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }

  @response = Example::V1.greeter.greet(name, opts)
end

Then('I should receive {string} from greeter') do |name|
  expect(@response.code).to eq(200)
  expect(JSON.parse(@response.body)).to eq('message' => name)
end
