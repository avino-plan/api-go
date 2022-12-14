// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: postar.proto

package postar

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

// PostarServiceClient is the client API for PostarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostarServiceClient interface {
	// SendEmail send one email.
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error)
}

type postarServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostarServiceClient(cc grpc.ClientConnInterface) PostarServiceClient {
	return &postarServiceClient{cc}
}

func (c *postarServiceClient) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error) {
	out := new(SendEmailResponse)
	err := c.cc.Invoke(ctx, "/github.com.avinoplan.api.postar.PostarService/SendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostarServiceServer is the server API for PostarService service.
// All implementations must embed UnimplementedPostarServiceServer
// for forward compatibility
type PostarServiceServer interface {
	// SendEmail send one email.
	SendEmail(context.Context, *SendEmailRequest) (*SendEmailResponse, error)
	mustEmbedUnimplementedPostarServiceServer()
}

// UnimplementedPostarServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostarServiceServer struct {
}

func (UnimplementedPostarServiceServer) SendEmail(context.Context, *SendEmailRequest) (*SendEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedPostarServiceServer) mustEmbedUnimplementedPostarServiceServer() {}

// UnsafePostarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostarServiceServer will
// result in compilation errors.
type UnsafePostarServiceServer interface {
	mustEmbedUnimplementedPostarServiceServer()
}

func RegisterPostarServiceServer(s grpc.ServiceRegistrar, srv PostarServiceServer) {
	s.RegisterService(&PostarService_ServiceDesc, srv)
}

func _PostarService_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostarServiceServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.avinoplan.api.postar.PostarService/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostarServiceServer).SendEmail(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostarService_ServiceDesc is the grpc.ServiceDesc for PostarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.avinoplan.api.postar.PostarService",
	HandlerType: (*PostarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _PostarService_SendEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "postar.proto",
}
