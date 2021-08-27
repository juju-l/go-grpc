// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/vipex.cc/aliOss/v1/v1.proto/oss.proto

package v1_proto

import (
	v1_proto "codeup.aliyun.com/vipex/go-grpc/api/vipex.cc/public/v1/v1.proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OssResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   *v1_proto.Result   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Data     bool               `protobuf:"varint,2,opt,name=data,proto3" json:"data,omitempty"`
	ErrorDes *v1_proto.ErrorDes `protobuf:"bytes,3,opt,name=errorDes,proto3" json:"errorDes,omitempty"`
}

func (x *OssResult) Reset() {
	*x = OssResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OssResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OssResult) ProtoMessage() {}

func (x *OssResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OssResult.ProtoReflect.Descriptor instead.
func (*OssResult) Descriptor() ([]byte, []int) {
	return file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_rawDescGZIP(), []int{0}
}

func (x *OssResult) GetResult() *v1_proto.Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *OssResult) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

func (x *OssResult) GetErrorDes() *v1_proto.ErrorDes {
	if x != nil {
		return x.ErrorDes
	}
	return nil
}

type Oss struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Test1 string `protobuf:"bytes,1,opt,name=test1,proto3" json:"test1,omitempty"`
	Test  string `protobuf:"bytes,2,opt,name=test,proto3" json:"test,omitempty"`
}

func (x *Oss) Reset() {
	*x = Oss{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Oss) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oss) ProtoMessage() {}

func (x *Oss) ProtoReflect() protoreflect.Message {
	mi := &file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Oss.ProtoReflect.Descriptor instead.
func (*Oss) Descriptor() ([]byte, []int) {
	return file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_rawDescGZIP(), []int{1}
}

func (x *Oss) GetTest1() string {
	if x != nil {
		return x.Test1
	}
	return ""
}

func (x *Oss) GetTest() string {
	if x != nil {
		return x.Test
	}
	return ""
}

var File_api_vipex_cc_aliOss_v1_v1_proto_oss_proto protoreflect.FileDescriptor

var file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x70, 0x65, 0x78, 0x2e, 0x63, 0x63, 0x2f, 0x61,
	0x6c, 0x69, 0x4f, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6f, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x70, 0x65, 0x78,
	0x2e, 0x63, 0x63, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x09, 0x4f, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x2e, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x73, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65,
	0x73, 0x22, 0x2f, 0x0a, 0x03, 0x4f, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74,
	0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x65, 0x73, 0x74, 0x31, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x73, 0x74, 0x32, 0x37, 0x0a, 0x07, 0x4f, 0x73, 0x73, 0x47, 0x72, 0x70, 0x63, 0x12, 0x2c, 0x0a,
	0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4f, 0x73, 0x73, 0x1a, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4f, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x4a, 0x5a, 0x48, 0x63,
	0x6f, 0x64, 0x65, 0x75, 0x70, 0x2e, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x76, 0x69, 0x70, 0x65, 0x78, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x69, 0x70, 0x65, 0x78, 0x2e, 0x63, 0x63, 0x2f, 0x61, 0x6c, 0x69, 0x4f,
	0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x76,
	0x31, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_rawDescOnce sync.Once
	file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_rawDescData = file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_rawDesc
)

func file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_rawDescGZIP() []byte {
	file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_rawDescOnce.Do(func() {
		file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_rawDescData)
	})
	return file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_rawDescData
}

var file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_goTypes = []interface{}{
	(*OssResult)(nil),         // 0: v1.proto.OssResult
	(*Oss)(nil),               // 1: v1.proto.Oss
	(*v1_proto.Result)(nil),   // 2: v1.proto.Result
	(*v1_proto.ErrorDes)(nil), // 3: v1.proto.ErrorDes
}
var file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_depIdxs = []int32{
	2, // 0: v1.proto.OssResult.result:type_name -> v1.proto.Result
	3, // 1: v1.proto.OssResult.errorDes:type_name -> v1.proto.ErrorDes
	1, // 2: v1.proto.OssGrpc.Test:input_type -> v1.proto.Oss
	0, // 3: v1.proto.OssGrpc.Test:output_type -> v1.proto.OssResult
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_init() }
func file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_init() {
	if File_api_vipex_cc_aliOss_v1_v1_proto_oss_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OssResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Oss); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_goTypes,
		DependencyIndexes: file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_depIdxs,
		MessageInfos:      file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_msgTypes,
	}.Build()
	File_api_vipex_cc_aliOss_v1_v1_proto_oss_proto = out.File
	file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_rawDesc = nil
	file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_goTypes = nil
	file_api_vipex_cc_aliOss_v1_v1_proto_oss_proto_depIdxs = nil
}
