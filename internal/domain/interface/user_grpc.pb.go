// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package interfacepri

import pb "gitee.com/vipex/go-grpc/api/vipex.cc/oauth2/proto"

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

// UserGrpcClient is the client API for UserGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserGrpcClient interface {
	Login(ctx context.Context, in *pb.User, opts ...grpc.CallOption) (*pb.BaseResult, error)
}

type userGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewUserGrpcClient(cc grpc.ClientConnInterface) UserGrpcClient {
	return &userGrpcClient{cc}
}

func (c *userGrpcClient) Login(ctx context.Context, in *pb.User, opts ...grpc.CallOption) (*pb.BaseResult, error) {
	out := new(pb.BaseResult)
	err := c.cc.Invoke(ctx, "/pb.UserGrpc/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserGrpcServer is the server API for UserGrpc service.
// All implementations must embed UnimplementedUserGrpcServer
// for forward compatibility
type UserGrpcServer interface {
	Login(context.Context, *pb.User) (*pb.BaseResult, error)
	mustEmbedUnimplementedUserGrpcServer()
}

// UnimplementedUserGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedUserGrpcServer struct {
}

func (UnimplementedUserGrpcServer) Login(context.Context, *pb.User) (*pb.BaseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserGrpcServer) mustEmbedUnimplementedUserGrpcServer() {}

// UnsafeUserGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserGrpcServer will
// result in compilation errors.
type UnsafeUserGrpcServer interface {
	mustEmbedUnimplementedUserGrpcServer()
}

func RegisterUserGrpcServer(s grpc.ServiceRegistrar, srv UserGrpcServer) {
	s.RegisterService(&UserGrpc_ServiceDesc, srv)
}

func _UserGrpc_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserGrpc/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServer).Login(ctx, req.(*pb.User))
	}
	return interceptor(ctx, in, info, handler)
}

// UserGrpc_ServiceDesc is the grpc.ServiceDesc for UserGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserGrpc",
	HandlerType: (*pb.UserGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserGrpc_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
