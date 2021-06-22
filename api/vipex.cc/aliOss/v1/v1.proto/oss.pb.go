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

type OssPutReq struct {
	ReqFile              []byte   `protobuf:"bytes,1,opt,name=reqFile,proto3" json:"reqFile"`
	ObjectKey            string   `protobuf:"bytes,2,opt,name=objectKey,proto3" json:"objectKey"`
	BucketName           string   `protobuf:"bytes,3,opt,name=bucketName,proto3" json:"bucketName"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OssPutReq) Reset()         { *m = OssPutReq{} }
func (m *OssPutReq) String() string { return proto.CompactTextString(m) }
func (*OssPutReq) ProtoMessage()    {}
func (*OssPutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_oss_bdb8401631e4c41f, []int{0}
}
func (m *OssPutReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OssPutReq.Unmarshal(m, b)
}
func (m *OssPutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OssPutReq.Marshal(b, m, deterministic)
}
func (dst *OssPutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OssPutReq.Merge(dst, src)
}
func (m *OssPutReq) XXX_Size() int {
	return xxx_messageInfo_OssPutReq.Size(m)
}
func (m *OssPutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_OssPutReq.DiscardUnknown(m)
}

var xxx_messageInfo_OssPutReq proto.InternalMessageInfo

func (m *OssPutReq) GetReqFile() []byte {
	if m != nil {
		return m.ReqFile
	}
	return nil
}

func (m *OssPutReq) GetObjectKey() string {
	if m != nil {
		return m.ObjectKey
	}
	return ""
}

func (m *OssPutReq) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

type OssGetReq struct {
	BucketName           string   `protobuf:"bytes,1,opt,name=bucketName,proto3" json:"bucketName"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OssGetReq) Reset()         { *m = OssGetReq{} }
func (m *OssGetReq) String() string { return proto.CompactTextString(m) }
func (*OssGetReq) ProtoMessage()    {}
func (*OssGetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_oss_bdb8401631e4c41f, []int{1}
}
func (m *OssGetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OssGetReq.Unmarshal(m, b)
}
func (m *OssGetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OssGetReq.Marshal(b, m, deterministic)
}
func (dst *OssGetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OssGetReq.Merge(dst, src)
}
func (m *OssGetReq) XXX_Size() int {
	return xxx_messageInfo_OssGetReq.Size(m)
}
func (m *OssGetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_OssGetReq.DiscardUnknown(m)
}

var xxx_messageInfo_OssGetReq proto.InternalMessageInfo

func (m *OssGetReq) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

type OssPutRst struct {
	Result               *v1.Result   `protobuf:"bytes,1,opt,name=result,proto3" json:"result"`
	Data                 bool         `protobuf:"varint,2,opt,name=data,proto3" json:"data"`
	ErrorDes             *v1.ErrorDes `protobuf:"bytes,3,opt,name=errorDes,proto3" json:"errorDes"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *OssPutRst) Reset()         { *m = OssPutRst{} }
func (m *OssPutRst) String() string { return proto.CompactTextString(m) }
func (*OssPutRst) ProtoMessage()    {}
func (*OssPutRst) Descriptor() ([]byte, []int) {
	return fileDescriptor_oss_bdb8401631e4c41f, []int{2}
}
func (m *OssPutRst) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OssPutRst.Unmarshal(m, b)
}
func (m *OssPutRst) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OssPutRst.Marshal(b, m, deterministic)
}
func (dst *OssPutRst) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OssPutRst.Merge(dst, src)
}
func (m *OssPutRst) XXX_Size() int {
	return xxx_messageInfo_OssPutRst.Size(m)
}
func (m *OssPutRst) XXX_DiscardUnknown() {
	xxx_messageInfo_OssPutRst.DiscardUnknown(m)
}

var xxx_messageInfo_OssPutRst proto.InternalMessageInfo

func (m *OssPutRst) GetResult() *v1.Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *OssPutRst) GetData() bool {
	if m != nil {
		return m.Data
	}
	return false
}

func (m *OssPutRst) GetErrorDes() *v1.ErrorDes {
	if m != nil {
		return m.ErrorDes
	}
	return nil
}

type OssResult struct {
	Result               *v1.Result   `protobuf:"bytes,1,opt,name=result,proto3" json:"result"`
	Data                 *Oss         `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
	ErrorDes             *v1.ErrorDes `protobuf:"bytes,3,opt,name=errorDes,proto3" json:"errorDes"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *OssResult) Reset()         { *m = OssResult{} }
func (m *OssResult) String() string { return proto.CompactTextString(m) }
func (*OssResult) ProtoMessage()    {}
func (*OssResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_oss_bdb8401631e4c41f, []int{3}
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

func (m *OssResult) GetData() *Oss {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *OssResult) GetErrorDes() *v1.ErrorDes {
	if m != nil {
		return m.ErrorDes
	}
	return nil
}

type Oss struct {
	Url                  []string `protobuf:"bytes,1,rep,name=url,proto3" json:"url"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Oss) Reset()         { *m = Oss{} }
func (m *Oss) String() string { return proto.CompactTextString(m) }
func (*Oss) ProtoMessage()    {}
func (*Oss) Descriptor() ([]byte, []int) {
	return fileDescriptor_oss_bdb8401631e4c41f, []int{4}
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

func (m *Oss) GetUrl() []string {
	if m != nil {
		return m.Url
	}
	return nil
}

func init() {
	proto.RegisterType((*OssPutReq)(nil), "v1.proto.OssPutReq")
	proto.RegisterType((*OssGetReq)(nil), "v1.proto.OssGetReq")
	proto.RegisterType((*OssPutRst)(nil), "v1.proto.OssPutRst")
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
	Put(ctx context.Context, in *OssPutReq, opts ...grpc.CallOption) (*OssPutRst, error)
	Get(ctx context.Context, in *OssGetReq, opts ...grpc.CallOption) (*OssResult, error)
}

type ossGrpcClient struct {
	cc *grpc.ClientConn
}

func NewOssGrpcClient(cc *grpc.ClientConn) OssGrpcClient {
	return &ossGrpcClient{cc}
}

func (c *ossGrpcClient) Put(ctx context.Context, in *OssPutReq, opts ...grpc.CallOption) (*OssPutRst, error) {
	out := new(OssPutRst)
	err := c.cc.Invoke(ctx, "/v1.proto.OssGrpc/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ossGrpcClient) Get(ctx context.Context, in *OssGetReq, opts ...grpc.CallOption) (*OssResult, error) {
	out := new(OssResult)
	err := c.cc.Invoke(ctx, "/v1.proto.OssGrpc/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OssGrpcServer is the server API for OssGrpc service.
type OssGrpcServer interface {
	Put(context.Context, *OssPutReq) (*OssPutRst, error)
	Get(context.Context, *OssGetReq) (*OssResult, error)
}

func RegisterOssGrpcServer(s *grpc.Server, srv OssGrpcServer) {
	s.RegisterService(&_OssGrpc_serviceDesc, srv)
}

func _OssGrpc_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OssPutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OssGrpcServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.proto.OssGrpc/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OssGrpcServer).Put(ctx, req.(*OssPutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OssGrpc_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OssGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OssGrpcServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.proto.OssGrpc/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OssGrpcServer).Get(ctx, req.(*OssGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _OssGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.proto.OssGrpc",
	HandlerType: (*OssGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _OssGrpc_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _OssGrpc_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/vipex.cc/aliOss/v1/v1.proto/oss.proto",
}

func init() {
	proto.RegisterFile("api/vipex.cc/aliOss/v1/v1.proto/oss.proto", fileDescriptor_oss_bdb8401631e4c41f)
}

var fileDescriptor_oss_bdb8401631e4c41f = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcd, 0x4e, 0xf3, 0x30,
	0x10, 0x6c, 0xbe, 0x7c, 0x6a, 0x1b, 0x17, 0xa4, 0xca, 0x1c, 0x88, 0x2a, 0x84, 0x4a, 0x4e, 0x41,
	0x88, 0x58, 0x09, 0x47, 0x2e, 0xa8, 0x02, 0x7a, 0x40, 0xa2, 0x95, 0x8f, 0x5c, 0x50, 0x62, 0x56,
	0x55, 0x20, 0x95, 0x5d, 0xdb, 0x89, 0xe8, 0x1b, 0xf0, 0xd8, 0x28, 0x76, 0xff, 0xd2, 0x1e, 0x10,
	0x9c, 0xb2, 0x9e, 0x9d, 0xf1, 0xcc, 0xc6, 0x8b, 0x2e, 0x53, 0x91, 0x93, 0x2a, 0x17, 0xf0, 0x19,
	0x31, 0x46, 0xd2, 0x22, 0x9f, 0x28, 0x45, 0xaa, 0x98, 0x54, 0x71, 0x24, 0x24, 0xd7, 0x9c, 0x70,
	0xa5, 0x6c, 0x85, 0xbb, 0x6b, 0x6c, 0x40, 0x1a, 0x22, 0x51, 0x66, 0x45, 0xce, 0x1a, 0xa2, 0x54,
	0xe4, 0x14, 0x54, 0x59, 0x68, 0x7b, 0x0e, 0x18, 0xf2, 0x26, 0x4a, 0x4d, 0x4b, 0x4d, 0x61, 0x81,
	0x7d, 0xd4, 0x91, 0xb0, 0x78, 0xcc, 0x0b, 0xf0, 0x9d, 0xa1, 0x13, 0x1e, 0xd1, 0xf5, 0x11, 0x9f,
	0x21, 0x8f, 0x67, 0xef, 0xc0, 0xf4, 0x13, 0x2c, 0xfd, 0x7f, 0x43, 0x27, 0xf4, 0xe8, 0x16, 0xc0,
	0xe7, 0x08, 0x65, 0x25, 0xfb, 0x00, 0xfd, 0x9c, 0xce, 0xc1, 0x77, 0x4d, 0x7b, 0x07, 0x09, 0xae,
	0x8c, 0xc9, 0x18, 0x8c, 0x49, 0x93, 0xec, 0x1c, 0x90, 0x97, 0x9b, 0x44, 0x4a, 0xe3, 0x10, 0xb5,
	0xa5, 0x89, 0x6b, 0x88, 0xbd, 0xa4, 0x1f, 0xad, 0x27, 0x89, 0xec, 0x18, 0x74, 0xd5, 0xc7, 0x18,
	0xfd, 0x7f, 0x4b, 0x75, 0x6a, 0xc2, 0x75, 0xa9, 0xa9, 0x71, 0x84, 0xba, 0x20, 0x25, 0x97, 0xf7,
	0xa0, 0x4c, 0xaa, 0x5e, 0x82, 0xb7, 0xfa, 0x87, 0x55, 0x87, 0x6e, 0x38, 0xc1, 0x97, 0x63, 0xbc,
	0xed, 0xcd, 0xbf, 0xf0, 0xbe, 0xd8, 0xf1, 0xee, 0x25, 0xc7, 0x5b, 0x5e, 0x7d, 0xd9, 0xdf, 0xa2,
	0x9c, 0x22, 0x77, 0xa2, 0x14, 0xee, 0x23, 0xb7, 0x94, 0x85, 0xef, 0x0c, 0xdd, 0xd0, 0xa3, 0x75,
	0x99, 0x70, 0xd4, 0xa9, 0xff, 0xa5, 0x14, 0x0c, 0xc7, 0xc8, 0x9d, 0x96, 0x1a, 0x9f, 0x34, 0xfc,
	0xec, 0x53, 0x0e, 0x0e, 0x41, 0xa5, 0x83, 0x56, 0x2d, 0x19, 0xc3, 0xbe, 0xc4, 0x3e, 0xcc, 0x9e,
	0xc4, 0x8e, 0x18, 0xb4, 0x46, 0xa3, 0x97, 0xbb, 0x59, 0xae, 0x01, 0x22, 0xc6, 0xe7, 0x76, 0xb5,
	0xc8, 0x8c, 0x5f, 0xcf, 0xa4, 0x60, 0xe4, 0x87, 0x0d, 0xbd, 0xad, 0xe2, 0x57, 0x53, 0x64, 0x6d,
	0xf3, 0xb9, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x85, 0x11, 0x3a, 0x1d, 0xd4, 0x02, 0x00, 0x00,
}
