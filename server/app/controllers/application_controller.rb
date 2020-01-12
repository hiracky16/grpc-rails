require 'echoService_services_pb.rb'
require 'echoService_pb.rb'

class ApplicationController < ActionController::API
  def echo
    echo_stub = Echo::Stub.new('localhost:2525', :this_channel_is_insecure)
    res = echo_stub.get_echo(Echo::Service::GetEchoMessage.new(target_echo: 'test'))
    render json: { message: res.input }
  end
end
