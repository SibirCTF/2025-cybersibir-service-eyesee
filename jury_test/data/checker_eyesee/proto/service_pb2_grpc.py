# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

import service_pb2 as service__pb2

GRPC_GENERATED_VERSION = '1.71.0'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in service_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class EyeSeeServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ServiceHealthCheck = channel.unary_unary(
                '/service.proto.EyeSeeService/ServiceHealthCheck',
                request_serializer=service__pb2.ServiceHealthCheckRequest.SerializeToString,
                response_deserializer=service__pb2.ServiceHealthCheckResponse.FromString,
                _registered_method=True)
        self.AddPrescription = channel.unary_unary(
                '/service.proto.EyeSeeService/AddPrescription',
                request_serializer=service__pb2.AddPrescriptionRequest.SerializeToString,
                response_deserializer=service__pb2.AddPrescriptionResponse.FromString,
                _registered_method=True)
        self.GetPrescriptionIDs = channel.unary_unary(
                '/service.proto.EyeSeeService/GetPrescriptionIDs',
                request_serializer=service__pb2.GetPrescriptionIDsRequest.SerializeToString,
                response_deserializer=service__pb2.GetPrescriptionIDsResponse.FromString,
                _registered_method=True)
        self.CheckPrescription = channel.unary_unary(
                '/service.proto.EyeSeeService/CheckPrescription',
                request_serializer=service__pb2.CheckPrescriptionRequest.SerializeToString,
                response_deserializer=service__pb2.CheckPrescriptionResponse.FromString,
                _registered_method=True)


class EyeSeeServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ServiceHealthCheck(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddPrescription(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPrescriptionIDs(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CheckPrescription(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_EyeSeeServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ServiceHealthCheck': grpc.unary_unary_rpc_method_handler(
                    servicer.ServiceHealthCheck,
                    request_deserializer=service__pb2.ServiceHealthCheckRequest.FromString,
                    response_serializer=service__pb2.ServiceHealthCheckResponse.SerializeToString,
            ),
            'AddPrescription': grpc.unary_unary_rpc_method_handler(
                    servicer.AddPrescription,
                    request_deserializer=service__pb2.AddPrescriptionRequest.FromString,
                    response_serializer=service__pb2.AddPrescriptionResponse.SerializeToString,
            ),
            'GetPrescriptionIDs': grpc.unary_unary_rpc_method_handler(
                    servicer.GetPrescriptionIDs,
                    request_deserializer=service__pb2.GetPrescriptionIDsRequest.FromString,
                    response_serializer=service__pb2.GetPrescriptionIDsResponse.SerializeToString,
            ),
            'CheckPrescription': grpc.unary_unary_rpc_method_handler(
                    servicer.CheckPrescription,
                    request_deserializer=service__pb2.CheckPrescriptionRequest.FromString,
                    response_serializer=service__pb2.CheckPrescriptionResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'service.proto.EyeSeeService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('service.proto.EyeSeeService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class EyeSeeService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ServiceHealthCheck(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/service.proto.EyeSeeService/ServiceHealthCheck',
            service__pb2.ServiceHealthCheckRequest.SerializeToString,
            service__pb2.ServiceHealthCheckResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def AddPrescription(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/service.proto.EyeSeeService/AddPrescription',
            service__pb2.AddPrescriptionRequest.SerializeToString,
            service__pb2.AddPrescriptionResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetPrescriptionIDs(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/service.proto.EyeSeeService/GetPrescriptionIDs',
            service__pb2.GetPrescriptionIDsRequest.SerializeToString,
            service__pb2.GetPrescriptionIDsResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def CheckPrescription(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/service.proto.EyeSeeService/CheckPrescription',
            service__pb2.CheckPrescriptionRequest.SerializeToString,
            service__pb2.CheckPrescriptionResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
