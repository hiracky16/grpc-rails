# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: echoService.proto for package ''

require 'grpc'
require 'echoService_pb'

module Echo
  class Service

    include GRPC::GenericService

    self.marshal_class_method = :encode
    self.unmarshal_class_method = :decode
    self.service_name = 'Echo'

    rpc :GetEcho, GetEchoMessage, EchoResponse
  end

  Stub = Service.rpc_stub_class
end
