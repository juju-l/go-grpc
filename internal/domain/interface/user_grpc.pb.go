// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package interfacepri

import pb "gitee.com/vipex/go-grpc/api/vipex.cc/oauth2/proto"

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserGrpcClient is the client API for UserGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserGrpcClient interface {
	Login(ctx context.Context, in *pb.User, opts ...grpc.CallOption) (*pb.BaseResult, error)
}

type userGrpcClient struct {
	cc *grpc.ClientConn
}

func NewUserGrpcClient(cc *grpc.ClientConn) UserGrpcClient {
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
type UserGrpcServer interface {
	Login(context.Context, *pb.User) (*pb.BaseResult, error)
}

func RegisterUserGrpcServer(s *grpc.Server, srv UserGrpcServer) {
	s.RegisterService(&_UserGrpc_serviceDesc, srv)
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

var _UserGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserGrpc",
	HandlerType: (*UserGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserGrpc_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_5737d3af2bb825d1) }

var fileDescriptor_user_5737d3af2bb825d1 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x51, 0xbf, 0x4f, 0xf3, 0x30,
	0x10, 0xfd, 0xd2, 0xaf, 0x0d, 0xe1, 0x40, 0x0c, 0x1e, 0x50, 0xc4, 0x54, 0x99, 0x25, 0x0b, 0xb6,
	0x14, 0x10, 0x03, 0x48, 0x0c, 0x05, 0xc4, 0x02, 0x8b, 0x25, 0x16, 0xb6, 0xc4, 0x39, 0x85, 0x48,
	0x50, 0x5b, 0xe7, 0x04, 0xf8, 0xf3, 0xd1, 0xb9, 0x0e, 0xdd, 0xde, 0xbb, 0xf7, 0x23, 0x2f, 0x32,
	0xc0, 0x14, 0x90, 0x94, 0x27, 0x37, 0x3a, 0xb1, 0xf0, 0xad, 0xbc, 0x81, 0xdc, 0x60, 0x98, 0x3e,
	0x46, 0x51, 0xc2, 0x01, 0x12, 0xdd, 0xbb, 0x0e, 0xcb, 0x6c, 0x9d, 0x55, 0x2b, 0x33, 0x53, 0x71,
	0x0a, 0x39, 0x12, 0xbd, 0x84, 0xbe, 0x5c, 0xac, 0xb3, 0xea, 0xd0, 0x24, 0x26, 0xef, 0xa0, 0x78,
	0x24, 0x72, 0xf4, 0x80, 0x81, 0xd3, 0xdc, 0xcc, 0xa6, 0x2c, 0x9a, 0x66, 0xca, 0x8a, 0x75, 0x1d,
	0xee, 0xe3, 0x33, 0x95, 0x5b, 0x80, 0x4d, 0x13, 0x30, 0x7d, 0x5f, 0x42, 0x4e, 0x11, 0xc5, 0x82,
	0xa3, 0x1a, 0x94, 0x6f, 0xd5, 0x4e, 0x33, 0x49, 0x11, 0x02, 0x96, 0x5d, 0x33, 0x36, 0xb1, 0xa8,
	0x30, 0x11, 0x8b, 0x0a, 0x0a, 0x4c, 0x2b, 0xca, 0xff, 0x31, 0x79, 0xcc, 0xc9, 0x79, 0x99, 0xf9,
	0x53, 0xa5, 0x82, 0xe5, 0x6b, 0x40, 0xe2, 0x16, 0x1e, 0x97, 0x86, 0x46, 0xcc, 0x37, 0x1f, 0xbe,
	0xbb, 0x34, 0x31, 0xe2, 0x5a, 0x43, 0xc1, 0xfe, 0x27, 0xf2, 0x56, 0x9c, 0xc3, 0xea, 0xd9, 0xf5,
	0xc3, 0x56, 0x14, 0x5c, 0xce, 0xe7, 0xb3, 0x13, 0x46, 0xfb, 0x1f, 0x90, 0xff, 0x36, 0xd7, 0x6f,
	0x57, 0xfd, 0x30, 0x22, 0x2a, 0xeb, 0x3e, 0xf5, 0xd7, 0xe0, 0xf1, 0x47, 0xf7, 0xee, 0xa2, 0x27,
	0x6f, 0x75, 0xe3, 0x87, 0xdd, 0x45, 0x59, 0xab, 0x5d, 0x33, 0x8d, 0xef, 0xb5, 0x8e, 0xaf, 0x70,
	0xeb, 0xdb, 0x36, 0x8f, 0xe8, 0xf2, 0x37, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xc6, 0x72, 0x08, 0x9d,
	0x01, 0x00, 0x00,
}
