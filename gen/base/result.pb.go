// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: result.proto

package gen

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32  `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *Result) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

type ErrorDes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserMsg string `protobuf:"bytes,1,opt,name=userMsg,proto3" json:"userMsg"`
	CodeMsg string `protobuf:"bytes,2,opt,name=codeMsg,proto3" json:"codeMsg"`
}

func (x *ErrorDes) Reset() {
	*x = ErrorDes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDes) ProtoMessage() {}

func (x *ErrorDes) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDes.ProtoReflect.Descriptor instead.
func (*ErrorDes) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorDes) GetUserMsg() string {
	if x != nil {
		return x.UserMsg
	}
	return ""
}

func (x *ErrorDes) GetCodeMsg() string {
	if x != nil {
		return x.CodeMsg
	}
	return ""
}

type BaseResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   *Result          `protobuf:"bytes,1,opt,name=result,proto3" json:"result"`
	Data     *structpb.Struct `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
	ErrorDes *ErrorDes        `protobuf:"bytes,3,opt,name=errorDes,proto3" json:"errorDes"` // google.protobuf.Any other = 4;
}

func (x *BaseResult) Reset() {
	*x = BaseResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResult) ProtoMessage() {}

func (x *BaseResult) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResult.ProtoReflect.Descriptor instead.
func (*BaseResult) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{2}
}

func (x *BaseResult) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *BaseResult) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *BaseResult) GetErrorDes() *ErrorDes {
	if x != nil {
		return x.ErrorDes
	}
	return nil
}

var File_result_proto protoreflect.FileDescriptor

var file_result_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x22,
	0x3e, 0x0a, 0x08, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x64, 0x65, 0x4d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x64, 0x65, 0x4d, 0x73, 0x67, 0x22,
	0x8d, 0x01, 0x0a, 0x0a, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x25,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x2b, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x44, 0x65, 0x73, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x73, 0x42,
	0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x70,
	0x65, 0x78, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_result_proto_rawDescOnce sync.Once
	file_result_proto_rawDescData = file_result_proto_rawDesc
)

func file_result_proto_rawDescGZIP() []byte {
	file_result_proto_rawDescOnce.Do(func() {
		file_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_result_proto_rawDescData)
	})
	return file_result_proto_rawDescData
}

var file_result_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_result_proto_goTypes = []interface{}{
	(*Result)(nil),          // 0: proto.Result
	(*ErrorDes)(nil),        // 1: proto.ErrorDes
	(*BaseResult)(nil),      // 2: proto.BaseResult
	(*structpb.Struct)(nil), // 3: google.protobuf.Struct
}
var file_result_proto_depIdxs = []int32{
	0, // 0: proto.BaseResult.result:type_name -> proto.Result
	3, // 1: proto.BaseResult.data:type_name -> google.protobuf.Struct
	1, // 2: proto.BaseResult.errorDes:type_name -> proto.ErrorDes
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_result_proto_init() }
func file_result_proto_init() {
	if File_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDes); i {
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
		file_result_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResult); i {
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
			RawDescriptor: file_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_result_proto_goTypes,
		DependencyIndexes: file_result_proto_depIdxs,
		MessageInfos:      file_result_proto_msgTypes,
	}.Build()
	File_result_proto = out.File
	file_result_proto_rawDesc = nil
	file_result_proto_goTypes = nil
	file_result_proto_depIdxs = nil
}
