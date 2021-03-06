// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: style_transfer.proto

package style_transfer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StyleTransferServiceClient is the client API for StyleTransferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StyleTransferServiceClient interface {
	// Performs arbitrary style transfer given a style image, content, and style strength (alpha)
	StyleTransfer(ctx context.Context, in *StyleTransferRequest, opts ...grpc.CallOption) (*StyleTransferResponse, error)
	// Performs arbitrary style transfer given image URLs
	StyleTransferURL(ctx context.Context, in *StyleTransferURLRequest, opts ...grpc.CallOption) (*StyleTransferResponse, error)
}

type styleTransferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStyleTransferServiceClient(cc grpc.ClientConnInterface) StyleTransferServiceClient {
	return &styleTransferServiceClient{cc}
}

func (c *styleTransferServiceClient) StyleTransfer(ctx context.Context, in *StyleTransferRequest, opts ...grpc.CallOption) (*StyleTransferResponse, error) {
	out := new(StyleTransferResponse)
	err := c.cc.Invoke(ctx, "/StyleTransferService/StyleTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *styleTransferServiceClient) StyleTransferURL(ctx context.Context, in *StyleTransferURLRequest, opts ...grpc.CallOption) (*StyleTransferResponse, error) {
	out := new(StyleTransferResponse)
	err := c.cc.Invoke(ctx, "/StyleTransferService/StyleTransferURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StyleTransferServiceServer is the server API for StyleTransferService service.
// All implementations must embed UnimplementedStyleTransferServiceServer
// for forward compatibility
type StyleTransferServiceServer interface {
	// Performs arbitrary style transfer given a style image, content, and style strength (alpha)
	StyleTransfer(context.Context, *StyleTransferRequest) (*StyleTransferResponse, error)
	// Performs arbitrary style transfer given image URLs
	StyleTransferURL(context.Context, *StyleTransferURLRequest) (*StyleTransferResponse, error)
	mustEmbedUnimplementedStyleTransferServiceServer()
}

// UnimplementedStyleTransferServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStyleTransferServiceServer struct {
}

func (UnimplementedStyleTransferServiceServer) StyleTransfer(context.Context, *StyleTransferRequest) (*StyleTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StyleTransfer not implemented")
}
func (UnimplementedStyleTransferServiceServer) StyleTransferURL(context.Context, *StyleTransferURLRequest) (*StyleTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StyleTransferURL not implemented")
}
func (UnimplementedStyleTransferServiceServer) mustEmbedUnimplementedStyleTransferServiceServer() {}

// UnsafeStyleTransferServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StyleTransferServiceServer will
// result in compilation errors.
type UnsafeStyleTransferServiceServer interface {
	mustEmbedUnimplementedStyleTransferServiceServer()
}

func RegisterStyleTransferServiceServer(s grpc.ServiceRegistrar, srv StyleTransferServiceServer) {
	s.RegisterService(&StyleTransferService_ServiceDesc, srv)
}

func _StyleTransferService_StyleTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StyleTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StyleTransferServiceServer).StyleTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StyleTransferService/StyleTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StyleTransferServiceServer).StyleTransfer(ctx, req.(*StyleTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StyleTransferService_StyleTransferURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StyleTransferURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StyleTransferServiceServer).StyleTransferURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StyleTransferService/StyleTransferURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StyleTransferServiceServer).StyleTransferURL(ctx, req.(*StyleTransferURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StyleTransferService_ServiceDesc is the grpc.ServiceDesc for StyleTransferService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StyleTransferService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StyleTransferService",
	HandlerType: (*StyleTransferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StyleTransfer",
			Handler:    _StyleTransferService_StyleTransfer_Handler,
		},
		{
			MethodName: "StyleTransferURL",
			Handler:    _StyleTransferService_StyleTransferURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "style_transfer.proto",
}
