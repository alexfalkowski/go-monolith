# frozen_string_literal: true

module Greeter
  module V1
    class HTTP < Nonnative::HTTPClient
      def greet(name, opts = {})
        get("/greeter/v1/hello/#{name}", opts)
      end
    end
  end
end
