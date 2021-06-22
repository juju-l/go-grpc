// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/vipex.cc/oauth2/v1/v1.proto/user.proto

package v1_proto // import "codeup.aliyun.com/vipex/go-grpc/api/vipex.cc/oauth2/v1/v1.proto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v1 "codeup.aliyun.com/vipex/go-grpc/api/vipex.cc/public/v1/v1.proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserLoginReq struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user"`
	Pswd                 string   `protobuf:"bytes,2,opt,name=pswd,proto3" json:"pswd"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginReq) Reset()         { *m = UserLoginReq{} }
func (m *UserLoginReq) String() string { return proto.CompactTextString(m) }
func (*UserLoginReq) ProtoMessage()    {}
func (*UserLoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_9d015886710acee9, []int{0}
}
func (m *UserLoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginReq.Unmarshal(m, b)
}
func (m *UserLoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginReq.Marshal(b, m, deterministic)
}
func (dst *UserLoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginReq.Merge(dst, src)
}
func (m *UserLoginReq) XXX_Size() int {
	return xxx_messageInfo_UserLoginReq.Size(m)
}
func (m *UserLoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginReq proto.InternalMessageInfo

func (m *UserLoginReq) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserLoginReq) GetPswd() string {
	if m != nil {
		return m.Pswd
	}
	return ""
}

type UserResult struct {
	Result               *v1.Result   `protobuf:"bytes,1,opt,name=result,proto3" json:"result"`
	Data                 *User        `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
	ErrorDes             *v1.ErrorDes `protobuf:"bytes,3,opt,name=errorDes,proto3" json:"errorDes"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UserResult) Reset()         { *m = UserResult{} }
func (m *UserResult) String() string { return proto.CompactTextString(m) }
func (*UserResult) ProtoMessage()    {}
func (*UserResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_9d015886710acee9, []int{1}
}
func (m *UserResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResult.Unmarshal(m, b)
}
func (m *UserResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResult.Marshal(b, m, deterministic)
}
func (dst *UserResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResult.Merge(dst, src)
}
func (m *UserResult) XXX_Size() int {
	return xxx_messageInfo_UserResult.Size(m)
}
func (m *UserResult) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResult.DiscardUnknown(m)
}

var xxx_messageInfo_UserResult proto.InternalMessageInfo

func (m *UserResult) GetResult() *v1.Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *UserResult) GetData() *User {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UserResult) GetErrorDes() *v1.ErrorDes {
	if m != nil {
		return m.ErrorDes
	}
	return nil
}

type User struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_9d015886710acee9, []int{2}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func init() {
	proto.RegisterType((*UserLoginReq)(nil), "v1.proto.UserLoginReq")
	proto.RegisterType((*UserResult)(nil), "v1.proto.UserResult")
	proto.RegisterType((*User)(nil), "v1.proto.User")
}

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
	Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserResult, error)
}

type userGrpcClient struct {
	cc *grpc.ClientConn
}

func NewUserGrpcClient(cc *grpc.ClientConn) UserGrpcClient {
	return &userGrpcClient{cc}
}

func (c *userGrpcClient) Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserResult, error) {
	out := new(UserResult)
	err := c.cc.Invoke(ctx, "/v1.proto.UserGrpc/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserGrpcServer is the server API for UserGrpc service.
type UserGrpcServer interface {
	Login(context.Context, *UserLoginReq) (*UserResult, error)
}

func RegisterUserGrpcServer(s *grpc.Server, srv UserGrpcServer) {
	s.RegisterService(&_UserGrpc_serviceDesc, srv)
}

func _UserGrpc_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.proto.UserGrpc/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServer).Login(ctx, req.(*UserLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.proto.UserGrpc",
	HandlerType: (*UserGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserGrpc_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/vipex.cc/oauth2/v1/v1.proto/user.proto",
}

func init() {
	proto.RegisterFile("api/vipex.cc/oauth2/v1/v1.proto/user.proto", fileDescriptor_user_9d015886710acee9)
}

var fileDescriptor_user_9d015886710acee9 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xb6, 0x5a, 0x47, 0x7d, 0x13, 0x91, 0x20, 0x52, 0x7a, 0x92, 0x9e, 0x86, 0x60, 0x42, 0x2b,
	0xe8, 0xc1, 0x8b, 0x4c, 0xc5, 0x8b, 0xa7, 0xc2, 0x2e, 0x5e, 0xa4, 0xcb, 0x42, 0x0d, 0x4c, 0x13,
	0x93, 0xb4, 0xfa, 0x37, 0xf8, 0x57, 0x4b, 0x5e, 0x5a, 0x47, 0x77, 0xd9, 0xa9, 0x5f, 0xdf, 0xf7,
	0xe3, 0x7d, 0xaf, 0x85, 0xcb, 0x5a, 0x4b, 0xd6, 0x49, 0x2d, 0x7e, 0x28, 0xe7, 0x4c, 0xd5, 0xad,
	0x7b, 0x2f, 0x59, 0x57, 0xb0, 0xae, 0xa0, 0xda, 0x28, 0xa7, 0x58, 0x6b, 0x85, 0x09, 0x90, 0x24,
	0xc3, 0x30, 0x63, 0x23, 0x97, 0x6e, 0x97, 0x6b, 0xc9, 0x47, 0xae, 0x5a, 0xcb, 0x4a, 0xd8, 0x76,
	0xed, 0xc2, 0x7b, 0x7e, 0x03, 0xc7, 0x0b, 0x2b, 0xcc, 0x8b, 0x6a, 0xe4, 0x67, 0x25, 0xbe, 0x08,
	0x81, 0xd8, 0x07, 0xa7, 0xd1, 0x45, 0x34, 0x3b, 0xaa, 0x10, 0xfb, 0x99, 0xb6, 0xdf, 0xab, 0x74,
	0x3f, 0xcc, 0x3c, 0xce, 0x7f, 0x23, 0x00, 0x6f, 0x0c, 0x61, 0x64, 0x06, 0x13, 0x83, 0x08, 0x8d,
	0xd3, 0xf2, 0x94, 0x0e, 0x1b, 0x69, 0x50, 0x54, 0x3d, 0x4f, 0x72, 0x88, 0x57, 0xb5, 0xab, 0x31,
	0x6c, 0x5a, 0x9e, 0x6c, 0x74, 0x98, 0x86, 0x1c, 0xa1, 0x90, 0x08, 0x63, 0x94, 0x79, 0x14, 0x36,
	0x3d, 0x40, 0x1d, 0xd9, 0xe8, 0x9e, 0x7a, 0xa6, 0xfa, 0xd7, 0xe4, 0x19, 0xc4, 0x8b, 0xbe, 0xe8,
	0x76, 0xf9, 0xf2, 0x01, 0x12, 0xcf, 0x3d, 0x1b, 0xcd, 0xc9, 0x2d, 0x1c, 0xe2, 0xa1, 0xe4, 0x7c,
	0xbc, 0x76, 0xb8, 0x3e, 0x3b, 0xdb, 0xaa, 0x83, 0x95, 0xf3, 0xbd, 0xf9, 0xfc, 0xf5, 0xbe, 0x91,
	0x4e, 0x08, 0xca, 0xd5, 0x47, 0xf8, 0xbc, 0xac, 0x51, 0x57, 0x8d, 0xd1, 0x9c, 0xed, 0xf8, 0x4d,
	0x77, 0x5d, 0xf1, 0x86, 0x60, 0x39, 0xc1, 0xc7, 0xf5, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31,
	0x95, 0xc2, 0xe9, 0xd9, 0x01, 0x00, 0x00,
}
