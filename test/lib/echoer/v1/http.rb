# frozen_string_literal: true

module Echoer
  module V1
    class HTTP < Nonnative::HTTPClient
      def echo(msg, opts = {})
        get("/echoer/v1/echo/#{msg}", opts)
      end
    end
  end
end
