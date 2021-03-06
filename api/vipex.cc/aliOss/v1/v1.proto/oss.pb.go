// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/vipex.cc/aliOss/v1/v1.proto/oss.proto

package v1_proto // import "codeup.aliyun.com/vipex/go-grpc/api/vipex.cc/aliOss/v1/v1.proto"

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

type OssResult struct {
	Result               *v1.Result   `protobuf:"bytes,1,opt,name=result,proto3" json:"result"`
	Data                 bool         `protobuf:"varint,2,opt,name=data,proto3" json:"data"`
	ErrorDes             *v1.ErrorDes `protobuf:"bytes,3,opt,name=errorDes,proto3" json:"errorDes"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *OssResult) Reset()         { *m = OssResult{} }
func (m *OssResult) String() string { return proto.CompactTextString(m) }
func (*OssResult) ProtoMessage()    {}
func (*OssResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_oss_5e7409456591c558, []int{0}
}
func (m *OssResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OssResult.Unmarshal(m, b)
}
func (m *OssResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OssResult.Marshal(b, m, deterministic)
}
func (dst *OssResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OssResult.Merge(dst, src)
}
func (m *OssResult) XXX_Size() int {
	return xxx_messageInfo_OssResult.Size(m)
}
func (m *OssResult) XXX_DiscardUnknown() {
	xxx_messageInfo_OssResult.DiscardUnknown(m)
}

var xxx_messageInfo_OssResult proto.InternalMessageInfo

func (m *OssResult) GetResult() *v1.Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *OssResult) GetData() bool {
	if m != nil {
		return m.Data
	}
	return false
}

func (m *OssResult) GetErrorDes() *v1.ErrorDes {
	if m != nil {
		return m.ErrorDes
	}
	return nil
}

type Oss struct {
	Test1                string   `protobuf:"bytes,1,opt,name=test1,proto3" json:"test1"`
	Test                 string   `protobuf:"bytes,2,opt,name=test,proto3" json:"test"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Oss) Reset()         { *m = Oss{} }
func (m *Oss) String() string { return proto.CompactTextString(m) }
func (*Oss) ProtoMessage()    {}
func (*Oss) Descriptor() ([]byte, []int) {
	return fileDescriptor_oss_5e7409456591c558, []int{1}
}
func (m *Oss) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Oss.Unmarshal(m, b)
}
func (m *Oss) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Oss.Marshal(b, m, deterministic)
}
func (dst *Oss) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Oss.Merge(dst, src)
}
func (m *Oss) XXX_Size() int {
	return xxx_messageInfo_Oss.Size(m)
}
func (m *Oss) XXX_DiscardUnknown() {
	xxx_messageInfo_Oss.DiscardUnknown(m)
}

var xxx_messageInfo_Oss proto.InternalMessageInfo

func (m *Oss) GetTest1() string {
	if m != nil {
		return m.Test1
	}
	return ""
}

func (m *Oss) GetTest() string {
	if m != nil {
		return m.Test
	}
	return ""
}

func init() {
	proto.RegisterType((*OssResult)(nil), "v1.proto.OssResult")
	proto.RegisterType((*Oss)(nil), "v1.proto.Oss")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OssGrpcClient is the client API for OssGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OssGrpcClient interface {
	Test(ctx context.Context, in *Oss, opts ...grpc.CallOption) (*OssResult, error)
}

type ossGrpcClient struct {
	cc *grpc.ClientConn
}

func NewOssGrpcClient(cc *grpc.ClientConn) OssGrpcClient {
	return &ossGrpcClient{cc}
}

func (c *ossGrpcClient) Test(ctx context.Context, in *Oss, opts ...grpc.CallOption) (*OssResult, error) {
	out := new(OssResult)
	err := c.cc.Invoke(ctx, "/v1.proto.OssGrpc/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OssGrpcServer is the server API for OssGrpc service.
type OssGrpcServer interface {
	Test(context.Context, *Oss) (*OssResult, error)
}

func RegisterOssGrpcServer(s *grpc.Server, srv OssGrpcServer) {
	s.RegisterService(&_OssGrpc_serviceDesc, srv)
}

func _OssGrpc_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Oss)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OssGrpcServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.proto.OssGrpc/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OssGrpcServer).Test(ctx, req.(*Oss))
	}
	return interceptor(ctx, in, info, handler)
}

var _OssGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.proto.OssGrpc",
	HandlerType: (*OssGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _OssGrpc_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/vipex.cc/aliOss/v1/v1.proto/oss.proto",
}

func init() {
	proto.RegisterFile("api/vipex.cc/aliOss/v1/v1.proto/oss.proto", fileDescriptor_oss_5e7409456591c558)
}

var fileDescriptor_oss_5e7409456591c558 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xad, 0x9b, 0xb3, 0x8d, 0x08, 0xf2, 0xf4, 0x50, 0x76, 0x1a, 0x3d, 0x55, 0xd0, 0x3c,
	0x3a, 0x0f, 0x1e, 0xbc, 0xc8, 0x50, 0x3c, 0x16, 0x82, 0x27, 0x2f, 0x92, 0xc5, 0x50, 0x02, 0x95,
	0x84, 0xbc, 0xac, 0xe8, 0x7f, 0x2f, 0x4b, 0x36, 0x47, 0x4f, 0x3b, 0xe5, 0xf3, 0x7e, 0x7c, 0xf8,
	0x86, 0xc7, 0x6e, 0xa5, 0x33, 0x38, 0x18, 0xa7, 0x7f, 0xb8, 0x52, 0x28, 0x7b, 0xd3, 0x12, 0xe1,
	0xd0, 0xe0, 0xd0, 0x70, 0xe7, 0x6d, 0xb0, 0x68, 0x89, 0x12, 0x41, 0xbe, 0xef, 0xcd, 0x71, 0x24,
	0xb9, 0xcd, 0xba, 0x37, 0x6a, 0x24, 0x49, 0x67, 0x84, 0xa6, 0x4d, 0x1f, 0x52, 0x5d, 0xfd, 0xb2,
	0xa2, 0x25, 0x4a, 0x2d, 0xa8, 0xd9, 0xcc, 0x47, 0x2a, 0xb3, 0x45, 0x56, 0x5f, 0x2c, 0xaf, 0xf8,
	0xde, 0xe3, 0x69, 0x43, 0xec, 0xe6, 0x00, 0x6c, 0xfa, 0x25, 0x83, 0x2c, 0x4f, 0x17, 0x59, 0x9d,
	0x8b, 0xc8, 0xc0, 0x59, 0xae, 0xbd, 0xb7, 0xfe, 0x45, 0x53, 0x39, 0x89, 0x3e, 0x1c, 0xfc, 0xd7,
	0xdd, 0x44, 0xfc, 0xef, 0x54, 0xc8, 0x26, 0x2d, 0x11, 0xdc, 0xb0, 0xb3, 0xa0, 0x29, 0x34, 0x31,
	0xb3, 0x10, 0xa9, 0xd8, 0x06, 0x6c, 0x21, 0x06, 0x14, 0x22, 0xf2, 0xf2, 0x91, 0x9d, 0xb7, 0x44,
	0x6f, 0xde, 0x29, 0xb8, 0x63, 0xd3, 0x77, 0x4d, 0x01, 0x2e, 0x0f, 0x09, 0x2d, 0xd1, 0xfc, 0x7a,
	0x54, 0xa6, 0x3f, 0x57, 0x27, 0xab, 0xd5, 0xc7, 0x73, 0x67, 0x82, 0xd6, 0x5c, 0xd9, 0xef, 0x74,
	0x1d, 0xec, 0xec, 0x7d, 0xe7, 0x9d, 0xc2, 0x23, 0x47, 0x7e, 0x1a, 0x9a, 0xcf, 0x08, 0xeb, 0x59,
	0x7c, 0x1e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x6c, 0xfc, 0x6a, 0x97, 0x01, 0x00, 0x00,
}
