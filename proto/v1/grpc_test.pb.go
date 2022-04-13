// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1--rc1
// source: grpc_test.proto

package grpc_test

import (
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

type TestVersion int32

const (
	// 未知版本
	TestVersion_UnKNOW TestVersion = 0
	// 版本1
	TestVersion_VERSION_1 TestVersion = 1
)

// Enum value maps for TestVersion.
var (
	TestVersion_name = map[int32]string{
		0: "UnKNOW",
		1: "VERSION_1",
	}
	TestVersion_value = map[string]int32{
		"UnKNOW":    0,
		"VERSION_1": 1,
	}
)

func (x TestVersion) Enum() *TestVersion {
	p := new(TestVersion)
	*p = x
	return p
}

func (x TestVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_test_proto_enumTypes[0].Descriptor()
}

func (TestVersion) Type() protoreflect.EnumType {
	return &file_grpc_test_proto_enumTypes[0]
}

func (x TestVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestVersion.Descriptor instead.
func (TestVersion) EnumDescriptor() ([]byte, []int) {
	return file_grpc_test_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 姓名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 年龄
	Age int32 `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	// 其他参数
	Others []byte `protobuf:"bytes,100,opt,name=others,proto3" json:"others,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_grpc_test_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Request) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Request) GetOthers() []byte {
	if x != nil {
		return x.Others
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前版本
	Version TestVersion `protobuf:"varint,1,opt,name=version,proto3,enum=grpc_test.v1.TestVersion" json:"version,omitempty"`
	// 状态码
	Code int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	// 数据集
	Data map[string]string `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_grpc_test_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetVersion() TestVersion {
	if x != nil {
		return x.Version
	}
	return TestVersion_UnKNOW
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_grpc_test_proto protoreflect.FileDescriptor

var file_grpc_test_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x22,
	0x47, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x22, 0xc2, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x34,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x28, 0x0a,
	0x0b, 0x54, 0x65, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06,
	0x55, 0x6e, 0x4b, 0x4e, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x56, 0x45, 0x52, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x31, 0x10, 0x01, 0x32, 0x41, 0x0a, 0x08, 0x47, 0x72, 0x70, 0x63, 0x54,
	0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f,
	0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_grpc_test_proto_rawDescOnce sync.Once
	file_grpc_test_proto_rawDescData = file_grpc_test_proto_rawDesc
)

func file_grpc_test_proto_rawDescGZIP() []byte {
	file_grpc_test_proto_rawDescOnce.Do(func() {
		file_grpc_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_test_proto_rawDescData)
	})
	return file_grpc_test_proto_rawDescData
}

var file_grpc_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpc_test_proto_goTypes = []interface{}{
	(TestVersion)(0), // 0: grpc_test.v1.TestVersion
	(*Request)(nil),  // 1: grpc_test.v1.Request
	(*Response)(nil), // 2: grpc_test.v1.Response
	nil,              // 3: grpc_test.v1.Response.DataEntry
}
var file_grpc_test_proto_depIdxs = []int32{
	0, // 0: grpc_test.v1.Response.version:type_name -> grpc_test.v1.TestVersion
	3, // 1: grpc_test.v1.Response.data:type_name -> grpc_test.v1.Response.DataEntry
	1, // 2: grpc_test.v1.GrpcTest.Test:input_type -> grpc_test.v1.Request
	2, // 3: grpc_test.v1.GrpcTest.Test:output_type -> grpc_test.v1.Response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_test_proto_init() }
func file_grpc_test_proto_init() {
	if File_grpc_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_grpc_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_grpc_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_test_proto_goTypes,
		DependencyIndexes: file_grpc_test_proto_depIdxs,
		EnumInfos:         file_grpc_test_proto_enumTypes,
		MessageInfos:      file_grpc_test_proto_msgTypes,
	}.Build()
	File_grpc_test_proto = out.File
	file_grpc_test_proto_rawDesc = nil
	file_grpc_test_proto_goTypes = nil
	file_grpc_test_proto_depIdxs = nil
}