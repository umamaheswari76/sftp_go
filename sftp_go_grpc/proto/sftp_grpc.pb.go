// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/sftp.proto

package sftp_go

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

// SftpServiceClient is the client API for SftpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SftpServiceClient interface {
	List(ctx context.Context, in *RemotePathRequest, opts ...grpc.CallOption) (*Response, error)
}

type sftpServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSftpServiceClient(cc grpc.ClientConnInterface) SftpServiceClient {
	return &sftpServiceClient{cc}
}

func (c *sftpServiceClient) List(ctx context.Context, in *RemotePathRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/sftp.SftpService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SftpServiceServer is the server API for SftpService service.
// All implementations must embed UnimplementedSftpServiceServer
// for forward compatibility
type SftpServiceServer interface {
	List(context.Context, *RemotePathRequest) (*Response, error)
	mustEmbedUnimplementedSftpServiceServer()
}

// UnimplementedSftpServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSftpServiceServer struct {
}

func (UnimplementedSftpServiceServer) List(context.Context, *RemotePathRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSftpServiceServer) mustEmbedUnimplementedSftpServiceServer() {}

// UnsafeSftpServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SftpServiceServer will
// result in compilation errors.
type UnsafeSftpServiceServer interface {
	mustEmbedUnimplementedSftpServiceServer()
}

func RegisterSftpServiceServer(s grpc.ServiceRegistrar, srv SftpServiceServer) {
	s.RegisterService(&SftpService_ServiceDesc, srv)
}

func _SftpService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemotePathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SftpServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sftp.SftpService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SftpServiceServer).List(ctx, req.(*RemotePathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SftpService_ServiceDesc is the grpc.ServiceDesc for SftpService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SftpService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sftp.SftpService",
	HandlerType: (*SftpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _SftpService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/sftp.proto",
}
