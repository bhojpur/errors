// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// ErrorsUIClient is the client API for ErrorsUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ErrorsUIClient interface {
	// ListEngineSpecs returns a list of Analysis Engine(s) that can be started through the UI.
	ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (ErrorsUI_ListEngineSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type errorsUIClient struct {
	cc grpc.ClientConnInterface
}

func NewErrorsUIClient(cc grpc.ClientConnInterface) ErrorsUIClient {
	return &errorsUIClient{cc}
}

func (c *errorsUIClient) ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (ErrorsUI_ListEngineSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ErrorsUI_ServiceDesc.Streams[0], "/v1.ErrorsUI/ListEngineSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &errorsUIListEngineSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ErrorsUI_ListEngineSpecsClient interface {
	Recv() (*ListEngineSpecsResponse, error)
	grpc.ClientStream
}

type errorsUIListEngineSpecsClient struct {
	grpc.ClientStream
}

func (x *errorsUIListEngineSpecsClient) Recv() (*ListEngineSpecsResponse, error) {
	m := new(ListEngineSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *errorsUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.ErrorsUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ErrorsUIServer is the server API for ErrorsUI service.
// All implementations must embed UnimplementedErrorsUIServer
// for forward compatibility
type ErrorsUIServer interface {
	// ListEngineSpecs returns a list of Analysis Engine(s) that can be started through the UI.
	ListEngineSpecs(*ListEngineSpecsRequest, ErrorsUI_ListEngineSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedErrorsUIServer()
}

// UnimplementedErrorsUIServer must be embedded to have forward compatible implementations.
type UnimplementedErrorsUIServer struct {
}

func (UnimplementedErrorsUIServer) ListEngineSpecs(*ListEngineSpecsRequest, ErrorsUI_ListEngineSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEngineSpecs not implemented")
}
func (UnimplementedErrorsUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedErrorsUIServer) mustEmbedUnimplementedErrorsUIServer() {}

// UnsafeErrorsUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ErrorsUIServer will
// result in compilation errors.
type UnsafeErrorsUIServer interface {
	mustEmbedUnimplementedErrorsUIServer()
}

func RegisterErrorsUIServer(s grpc.ServiceRegistrar, srv ErrorsUIServer) {
	s.RegisterService(&ErrorsUI_ServiceDesc, srv)
}

func _ErrorsUI_ListEngineSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEngineSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ErrorsUIServer).ListEngineSpecs(m, &errorsUIListEngineSpecsServer{stream})
}

type ErrorsUI_ListEngineSpecsServer interface {
	Send(*ListEngineSpecsResponse) error
	grpc.ServerStream
}

type errorsUIListEngineSpecsServer struct {
	grpc.ServerStream
}

func (x *errorsUIListEngineSpecsServer) Send(m *ListEngineSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ErrorsUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorsUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ErrorsUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorsUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ErrorsUI_ServiceDesc is the grpc.ServiceDesc for ErrorsUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ErrorsUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ErrorsUI",
	HandlerType: (*ErrorsUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _ErrorsUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEngineSpecs",
			Handler:       _ErrorsUI_ListEngineSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "errors-ui.proto",
}
