// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: external/cfgmgmt/response/datafeed.proto

package response

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetPaginationResponse) Reset() {
	*x = GetPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_cfgmgmt_response_datafeed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaginationResponse) ProtoMessage() {}

func (x *GetPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_response_datafeed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaginationResponse.ProtoReflect.Descriptor instead.
func (*GetPaginationResponse) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_response_datafeed_proto_rawDescGZIP(), []int{0}
}

func (x *GetPaginationResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_external_cfgmgmt_response_datafeed_proto protoreflect.FileDescriptor

var file_external_cfgmgmt_response_datafeed_proto_rawDesc = []byte{
	0x0a, 0x28, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67,
	0x6d, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x66,
	0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_cfgmgmt_response_datafeed_proto_rawDescOnce sync.Once
	file_external_cfgmgmt_response_datafeed_proto_rawDescData = file_external_cfgmgmt_response_datafeed_proto_rawDesc
)

func file_external_cfgmgmt_response_datafeed_proto_rawDescGZIP() []byte {
	file_external_cfgmgmt_response_datafeed_proto_rawDescOnce.Do(func() {
		file_external_cfgmgmt_response_datafeed_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_cfgmgmt_response_datafeed_proto_rawDescData)
	})
	return file_external_cfgmgmt_response_datafeed_proto_rawDescData
}

var file_external_cfgmgmt_response_datafeed_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_external_cfgmgmt_response_datafeed_proto_goTypes = []interface{}{
	(*GetPaginationResponse)(nil), // 0: chef.automate.api.cfgmgmt.response.GetPaginationResponse
}
var file_external_cfgmgmt_response_datafeed_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_cfgmgmt_response_datafeed_proto_init() }
func file_external_cfgmgmt_response_datafeed_proto_init() {
	if File_external_cfgmgmt_response_datafeed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_cfgmgmt_response_datafeed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaginationResponse); i {
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
			RawDescriptor: file_external_cfgmgmt_response_datafeed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_cfgmgmt_response_datafeed_proto_goTypes,
		DependencyIndexes: file_external_cfgmgmt_response_datafeed_proto_depIdxs,
		MessageInfos:      file_external_cfgmgmt_response_datafeed_proto_msgTypes,
	}.Build()
	File_external_cfgmgmt_response_datafeed_proto = out.File
	file_external_cfgmgmt_response_datafeed_proto_rawDesc = nil
	file_external_cfgmgmt_response_datafeed_proto_goTypes = nil
	file_external_cfgmgmt_response_datafeed_proto_depIdxs = nil
}
