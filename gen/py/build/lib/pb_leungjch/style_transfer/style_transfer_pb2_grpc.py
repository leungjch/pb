# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import style_transfer_pb2 as style__transfer__pb2


class StyleTransferServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.StyleTransfer = channel.unary_unary(
                '/StyleTransferService/StyleTransfer',
                request_serializer=style__transfer__pb2.StyleTransferRequest.SerializeToString,
                response_deserializer=style__transfer__pb2.StyleTransferResponse.FromString,
                )
        self.StyleTransferURL = channel.unary_unary(
                '/StyleTransferService/StyleTransferURL',
                request_serializer=style__transfer__pb2.StyleTransferURLRequest.SerializeToString,
                response_deserializer=style__transfer__pb2.StyleTransferResponse.FromString,
                )


class StyleTransferServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def StyleTransfer(self, request, context):
        """Performs arbitrary style transfer given a style image, content, and style strength (alpha)
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def StyleTransferURL(self, request, context):
        """Performs arbitrary style transfer given image URLs
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_StyleTransferServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'StyleTransfer': grpc.unary_unary_rpc_method_handler(
                    servicer.StyleTransfer,
                    request_deserializer=style__transfer__pb2.StyleTransferRequest.FromString,
                    response_serializer=style__transfer__pb2.StyleTransferResponse.SerializeToString,
            ),
            'StyleTransferURL': grpc.unary_unary_rpc_method_handler(
                    servicer.StyleTransferURL,
                    request_deserializer=style__transfer__pb2.StyleTransferURLRequest.FromString,
                    response_serializer=style__transfer__pb2.StyleTransferResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'StyleTransferService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class StyleTransferService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def StyleTransfer(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/StyleTransferService/StyleTransfer',
            style__transfer__pb2.StyleTransferRequest.SerializeToString,
            style__transfer__pb2.StyleTransferResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def StyleTransferURL(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/StyleTransferService/StyleTransferURL',
            style__transfer__pb2.StyleTransferURLRequest.SerializeToString,
            style__transfer__pb2.StyleTransferResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
