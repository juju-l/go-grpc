// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/vipex.cc/public/proto/apiResult.proto

package pb // import "gitee.com/vipex/go-grpc/api/vipex.cc/public/proto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Result struct {
	ErrCode              int32    `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_apiResult_c8899353613134dd, []int{0}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *Result) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type ErrorDes struct {
	UserMsg              string   `protobuf:"bytes,1,opt,name=userMsg,proto3" json:"userMsg"`
	CodeMsg              string   `protobuf:"bytes,2,opt,name=codeMsg,proto3" json:"codeMsg"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorDes) Reset()         { *m = ErrorDes{} }
func (m *ErrorDes) String() string { return proto.CompactTextString(m) }
func (*ErrorDes) ProtoMessage()    {}
func (*ErrorDes) Descriptor() ([]byte, []int) {
	return fileDescriptor_apiResult_c8899353613134dd, []int{1}
}
func (m *ErrorDes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorDes.Unmarshal(m, b)
}
func (m *ErrorDes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorDes.Marshal(b, m, deterministic)
}
func (dst *ErrorDes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorDes.Merge(dst, src)
}
func (m *ErrorDes) XXX_Size() int {
	return xxx_messageInfo_ErrorDes.Size(m)
}
func (m *ErrorDes) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorDes.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorDes proto.InternalMessageInfo

func (m *ErrorDes) GetUserMsg() string {
	if m != nil {
		return m.UserMsg
	}
	return ""
}

func (m *ErrorDes) GetCodeMsg() string {
	if m != nil {
		return m.CodeMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*Result)(nil), "pb.Result")
	proto.RegisterType((*ErrorDes)(nil), "pb.ErrorDes")
}

func init() {
	proto.RegisterFile("api/vipex.cc/public/proto/apiResult.proto", fileDescriptor_apiResult_c8899353613134dd)
}

var fileDescriptor_apiResult_c8899353613134dd = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xcf, 0x3f, 0x0b, 0xc2, 0x30,
	0x10, 0x05, 0x70, 0x52, 0xb0, 0x6a, 0xc6, 0x0e, 0xd2, 0xb1, 0x74, 0xaa, 0x83, 0xcd, 0xa0, 0x38,
	0x28, 0x38, 0xf8, 0x67, 0x74, 0xe9, 0xe8, 0xd6, 0xa4, 0x47, 0x08, 0x54, 0xee, 0xb8, 0xb4, 0xe2,
	0xc7, 0x97, 0xa6, 0x16, 0x27, 0xb7, 0xfc, 0x72, 0xef, 0xc1, 0x9d, 0x5c, 0xd7, 0xe4, 0xd4, 0xcb,
	0x11, 0xbc, 0x4b, 0x63, 0x14, 0xf5, 0xba, 0x75, 0x46, 0x11, 0x63, 0x87, 0xaa, 0x26, 0x57, 0x81,
	0xef, 0xdb, 0xae, 0x0c, 0x4e, 0x22, 0xd2, 0xf9, 0x41, 0xc6, 0xe3, 0x5f, 0x92, 0xca, 0x39, 0x30,
	0x5f, 0xb0, 0x81, 0x54, 0x64, 0xa2, 0x98, 0x55, 0x13, 0x93, 0x95, 0x8c, 0x81, 0xf9, 0xee, 0x6d,
	0x1a, 0x65, 0xa2, 0x58, 0x56, 0x5f, 0xe5, 0x27, 0xb9, 0xb8, 0x31, 0x23, 0x5f, 0xc1, 0x0f, 0xed,
	0xde, 0x43, 0x08, 0x89, 0x10, 0x9a, 0x38, 0x4c, 0x0c, 0x36, 0xf0, 0xab, 0x4f, 0x3c, 0xef, 0x1f,
	0x3b, 0xeb, 0x3a, 0x80, 0xd2, 0xe0, 0x73, 0x5c, 0x59, 0x59, 0xdc, 0x58, 0x26, 0xa3, 0xfe, 0x1e,
	0x71, 0x24, 0xad, 0xe3, 0xf0, 0xda, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xd8, 0x94, 0x28,
	0xeb, 0x00, 0x00, 0x00,
}
