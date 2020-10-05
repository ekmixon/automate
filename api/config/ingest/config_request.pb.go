// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: config/ingest/config_request.proto

package ingest

import (
	shared "github.com/chef/automate/api/config/shared"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type ConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V1 *ConfigRequest_V1 `protobuf:"bytes,3,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
}

func (x *ConfigRequest) Reset() {
	*x = ConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_ingest_config_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest) ProtoMessage() {}

func (x *ConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_ingest_config_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest.ProtoReflect.Descriptor instead.
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_ingest_config_request_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if x != nil {
		return x.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sys *ConfigRequest_V1_System  `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
	Svc *ConfigRequest_V1_Service `protobuf:"bytes,2,opt,name=svc,proto3" json:"svc,omitempty" toml:"svc,omitempty" mapstructure:"svc,omitempty"`
}

func (x *ConfigRequest_V1) Reset() {
	*x = ConfigRequest_V1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_ingest_config_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1) ProtoMessage() {}

func (x *ConfigRequest_V1) ProtoReflect() protoreflect.Message {
	mi := &file_config_ingest_config_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return file_config_ingest_config_request_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if x != nil {
		return x.Sys
	}
	return nil
}

func (x *ConfigRequest_V1) GetSvc() *ConfigRequest_V1_Service {
	if x != nil {
		return x.Svc
	}
	return nil
}

type ConfigRequest_V1_System struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mlsa    *shared.Mlsa                     `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Service *ConfigRequest_V1_System_Service `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Tls     *shared.TLSCredentials           `protobuf:"bytes,3,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Log     *ConfigRequest_V1_System_Log     `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
}

func (x *ConfigRequest_V1_System) Reset() {
	*x = ConfigRequest_V1_System{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_ingest_config_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System) ProtoMessage() {}

func (x *ConfigRequest_V1_System) ProtoReflect() protoreflect.Message {
	mi := &file_config_ingest_config_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return file_config_ingest_config_request_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if x != nil {
		return x.Mlsa
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if x != nil {
		return x.Tls
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetLog() *ConfigRequest_V1_System_Log {
	if x != nil {
		return x.Log
	}
	return nil
}

type ConfigRequest_V1_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigRequest_V1_Service) Reset() {
	*x = ConfigRequest_V1_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_ingest_config_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_Service) ProtoMessage() {}

func (x *ConfigRequest_V1_Service) ProtoReflect() protoreflect.Message {
	mi := &file_config_ingest_config_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_Service.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_Service) Descriptor() ([]byte, []int) {
	return file_config_ingest_config_request_proto_rawDescGZIP(), []int{0, 0, 1}
}

type ConfigRequest_V1_System_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Do not use.
	Host *wrappers.StringValue `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"` // The listen host is no longer setable(localhost only)
	Port *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	// NOTE: purge actions are no longer configurable via config
	// and are now managed with gRPC endpoints. They are not
	// reserved so we can migrate the initial values from config to
	// to the purge cereal workflows.
	// Setting these values is prevented in the Validate() callback.
	//
	// Deprecated: Do not use.
	PurgeConvergeHistoryAfterDays *wrappers.Int32Value `protobuf:"bytes,3,opt,name=purge_converge_history_after_days,json=purgeConvergeHistoryAfterDays,proto3" json:"purge_converge_history_after_days,omitempty" toml:"purge_converge_history_after_days,omitempty" mapstructure:"purge_converge_history_after_days,omitempty"`
	// Deprecated: Do not use.
	PurgeActionsAfterDays                 *wrappers.Int32Value `protobuf:"bytes,4,opt,name=purge_actions_after_days,json=purgeActionsAfterDays,proto3" json:"purge_actions_after_days,omitempty" toml:"purge_actions_after_days,omitempty" mapstructure:"purge_actions_after_days,omitempty"`
	MaxNumberOfBundledRunMsgs             *wrappers.Int32Value `protobuf:"bytes,6,opt,name=max_number_of_bundled_run_msgs,json=maxNumberOfBundledRunMsgs,proto3" json:"max_number_of_bundled_run_msgs,omitempty" toml:"max_number_of_bundled_run_msgs,omitempty" mapstructure:"max_number_of_bundled_run_msgs,omitempty"`
	MaxNumberOfBundledActionMsgs          *wrappers.Int32Value `protobuf:"bytes,7,opt,name=max_number_of_bundled_action_msgs,json=maxNumberOfBundledActionMsgs,proto3" json:"max_number_of_bundled_action_msgs,omitempty" toml:"max_number_of_bundled_action_msgs,omitempty" mapstructure:"max_number_of_bundled_action_msgs,omitempty"`
	NumberOfRunMsgsTransformers           *wrappers.Int32Value `protobuf:"bytes,8,opt,name=number_of_run_msgs_transformers,json=numberOfRunMsgsTransformers,proto3" json:"number_of_run_msgs_transformers,omitempty" toml:"number_of_run_msgs_transformers,omitempty" mapstructure:"number_of_run_msgs_transformers,omitempty"`
	NumberOfRunMsgPublishers              *wrappers.Int32Value `protobuf:"bytes,9,opt,name=number_of_run_msg_publishers,json=numberOfRunMsgPublishers,proto3" json:"number_of_run_msg_publishers,omitempty" toml:"number_of_run_msg_publishers,omitempty" mapstructure:"number_of_run_msg_publishers,omitempty"`
	MessageBufferSize                     *wrappers.Int32Value `protobuf:"bytes,10,opt,name=message_buffer_size,json=messageBufferSize,proto3" json:"message_buffer_size,omitempty" toml:"message_buffer_size,omitempty" mapstructure:"message_buffer_size,omitempty"`
	NodesMissingRunningDefault            *wrappers.BoolValue  `protobuf:"bytes,11,opt,name=nodes_missing_running_default,json=nodesMissingRunningDefault,proto3" json:"nodes_missing_running_default,omitempty" toml:"nodes_missing_running_default,omitempty" mapstructure:"nodes_missing_running_default,omitempty"`
	MissingNodesForDeletionRunningDefault *wrappers.BoolValue  `protobuf:"bytes,12,opt,name=missing_nodes_for_deletion_running_default,json=missingNodesForDeletionRunningDefault,proto3" json:"missing_nodes_for_deletion_running_default,omitempty" toml:"missing_nodes_for_deletion_running_default,omitempty" mapstructure:"missing_nodes_for_deletion_running_default,omitempty"`
	NumberOfNodemanagerPublishers         *wrappers.Int32Value `protobuf:"bytes,13,opt,name=number_of_nodemanager_publishers,json=numberOfNodemanagerPublishers,proto3" json:"number_of_nodemanager_publishers,omitempty" toml:"number_of_nodemanager_publishers,omitempty" mapstructure:"number_of_nodemanager_publishers,omitempty"`
}

func (x *ConfigRequest_V1_System_Service) Reset() {
	*x = ConfigRequest_V1_System_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_ingest_config_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System_Service) ProtoMessage() {}

func (x *ConfigRequest_V1_System_Service) ProtoReflect() protoreflect.Message {
	mi := &file_config_ingest_config_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System_Service.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return file_config_ingest_config_request_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

// Deprecated: Do not use.
func (x *ConfigRequest_V1_System_Service) GetHost() *wrappers.StringValue {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *ConfigRequest_V1_System_Service) GetPort() *wrappers.Int32Value {
	if x != nil {
		return x.Port
	}
	return nil
}

// Deprecated: Do not use.
func (x *ConfigRequest_V1_System_Service) GetPurgeConvergeHistoryAfterDays() *wrappers.Int32Value {
	if x != nil {
		return x.PurgeConvergeHistoryAfterDays
	}
	return nil
}

// Deprecated: Do not use.
func (x *ConfigRequest_V1_System_Service) GetPurgeActionsAfterDays() *wrappers.Int32Value {
	if x != nil {
		return x.PurgeActionsAfterDays
	}
	return nil
}

func (x *ConfigRequest_V1_System_Service) GetMaxNumberOfBundledRunMsgs() *wrappers.Int32Value {
	if x != nil {
		return x.MaxNumberOfBundledRunMsgs
	}
	return nil
}

func (x *ConfigRequest_V1_System_Service) GetMaxNumberOfBundledActionMsgs() *wrappers.Int32Value {
	if x != nil {
		return x.MaxNumberOfBundledActionMsgs
	}
	return nil
}

func (x *ConfigRequest_V1_System_Service) GetNumberOfRunMsgsTransformers() *wrappers.Int32Value {
	if x != nil {
		return x.NumberOfRunMsgsTransformers
	}
	return nil
}

func (x *ConfigRequest_V1_System_Service) GetNumberOfRunMsgPublishers() *wrappers.Int32Value {
	if x != nil {
		return x.NumberOfRunMsgPublishers
	}
	return nil
}

func (x *ConfigRequest_V1_System_Service) GetMessageBufferSize() *wrappers.Int32Value {
	if x != nil {
		return x.MessageBufferSize
	}
	return nil
}

func (x *ConfigRequest_V1_System_Service) GetNodesMissingRunningDefault() *wrappers.BoolValue {
	if x != nil {
		return x.NodesMissingRunningDefault
	}
	return nil
}

func (x *ConfigRequest_V1_System_Service) GetMissingNodesForDeletionRunningDefault() *wrappers.BoolValue {
	if x != nil {
		return x.MissingNodesForDeletionRunningDefault
	}
	return nil
}

func (x *ConfigRequest_V1_System_Service) GetNumberOfNodemanagerPublishers() *wrappers.Int32Value {
	if x != nil {
		return x.NumberOfNodemanagerPublishers
	}
	return nil
}

type ConfigRequest_V1_System_Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Format *wrappers.StringValue `protobuf:"bytes,1,opt,name=format,proto3" json:"format,omitempty" toml:"format,omitempty" mapstructure:"format,omitempty"`
	Level  *wrappers.StringValue `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
}

func (x *ConfigRequest_V1_System_Log) Reset() {
	*x = ConfigRequest_V1_System_Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_ingest_config_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System_Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System_Log) ProtoMessage() {}

func (x *ConfigRequest_V1_System_Log) ProtoReflect() protoreflect.Message {
	mi := &file_config_ingest_config_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System_Log.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System_Log) Descriptor() ([]byte, []int) {
	return file_config_ingest_config_request_proto_rawDescGZIP(), []int{0, 0, 0, 1}
}

func (x *ConfigRequest_V1_System_Log) GetFormat() *wrappers.StringValue {
	if x != nil {
		return x.Format
	}
	return nil
}

func (x *ConfigRequest_V1_System_Log) GetLevel() *wrappers.StringValue {
	if x != nil {
		return x.Level
	}
	return nil
}

var File_config_ingest_config_request_proto protoreflect.FileDescriptor

var file_config_ingest_config_request_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x74, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x61, 0x32, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x32, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x0e, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x02, 0x76, 0x31, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x67,
	0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x56, 0x31, 0x52, 0x02, 0x76, 0x31, 0x1a, 0x99, 0x0d, 0x0a, 0x02, 0x56, 0x31, 0x12,
	0x46, 0x0a, 0x03, 0x73, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x52, 0x03, 0x73, 0x79, 0x73, 0x12, 0x47, 0x0a, 0x03, 0x73, 0x76, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x56, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x03, 0x73, 0x76, 0x63,
	0x1a, 0xf6, 0x0b, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x34, 0x0a, 0x04, 0x6d,
	0x6c, 0x73, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x6c, 0x73, 0x61, 0x52, 0x04, 0x6d, 0x6c, 0x73,
	0x61, 0x12, 0x56, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56,
	0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x03, 0x74, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x12, 0x4a, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03,
	0x6c, 0x6f, 0x67, 0x1a, 0xe2, 0x08, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x34, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x16, 0xc2, 0xf3, 0x18, 0x12, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x10,
	0x8a, 0x4f, 0x1a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x69,
	0x0a, 0x21, 0x70, 0x75, 0x72, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x67, 0x65,
	0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x64,
	0x61, 0x79, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x18, 0x01, 0x52, 0x1d, 0x70, 0x75, 0x72, 0x67,
	0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x67, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x41, 0x66, 0x74, 0x65, 0x72, 0x44, 0x61, 0x79, 0x73, 0x12, 0x58, 0x0a, 0x18, 0x70, 0x75, 0x72,
	0x67, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x18, 0x01, 0x52, 0x15, 0x70, 0x75,
	0x72, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x66, 0x74, 0x65, 0x72, 0x44,
	0x61, 0x79, 0x73, 0x12, 0x5e, 0x0a, 0x1e, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x5f, 0x72, 0x75, 0x6e,
	0x5f, 0x6d, 0x73, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x19, 0x6d, 0x61, 0x78, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x4f, 0x66, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x52, 0x75, 0x6e, 0x4d,
	0x73, 0x67, 0x73, 0x12, 0x64, 0x0a, 0x21, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x1c, 0x6d, 0x61, 0x78,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x67, 0x73, 0x12, 0x61, 0x0a, 0x1f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x1b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x52, 0x75, 0x6e, 0x4d, 0x73, 0x67, 0x73,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x5b, 0x0a, 0x1c,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x6d, 0x73,
	0x67, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x18, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x52, 0x75, 0x6e, 0x4d, 0x73, 0x67, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x73, 0x12, 0x4b, 0x0a, 0x13, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x5d, 0x0a, 0x1d, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x5f,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x1a, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x75, 0x0a, 0x2a, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x25, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x64, 0x0a, 0x20,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x73,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x1d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x4e, 0x6f, 0x64,
	0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x72, 0x73, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x1a, 0x6f, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12,
	0x34, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x1a, 0x09, 0x0a, 0x07, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x3a, 0x14, 0xc2, 0xf3, 0x18, 0x10, 0x0a, 0x0e, 0x69, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x03,
	0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_ingest_config_request_proto_rawDescOnce sync.Once
	file_config_ingest_config_request_proto_rawDescData = file_config_ingest_config_request_proto_rawDesc
)

func file_config_ingest_config_request_proto_rawDescGZIP() []byte {
	file_config_ingest_config_request_proto_rawDescOnce.Do(func() {
		file_config_ingest_config_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_ingest_config_request_proto_rawDescData)
	})
	return file_config_ingest_config_request_proto_rawDescData
}

var file_config_ingest_config_request_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_config_ingest_config_request_proto_goTypes = []interface{}{
	(*ConfigRequest)(nil),                   // 0: chef.automate.domain.ingest.ConfigRequest
	(*ConfigRequest_V1)(nil),                // 1: chef.automate.domain.ingest.ConfigRequest.V1
	(*ConfigRequest_V1_System)(nil),         // 2: chef.automate.domain.ingest.ConfigRequest.V1.System
	(*ConfigRequest_V1_Service)(nil),        // 3: chef.automate.domain.ingest.ConfigRequest.V1.Service
	(*ConfigRequest_V1_System_Service)(nil), // 4: chef.automate.domain.ingest.ConfigRequest.V1.System.Service
	(*ConfigRequest_V1_System_Log)(nil),     // 5: chef.automate.domain.ingest.ConfigRequest.V1.System.Log
	(*shared.Mlsa)(nil),                     // 6: chef.automate.infra.config.Mlsa
	(*shared.TLSCredentials)(nil),           // 7: chef.automate.infra.config.TLSCredentials
	(*wrappers.StringValue)(nil),            // 8: google.protobuf.StringValue
	(*wrappers.Int32Value)(nil),             // 9: google.protobuf.Int32Value
	(*wrappers.BoolValue)(nil),              // 10: google.protobuf.BoolValue
}
var file_config_ingest_config_request_proto_depIdxs = []int32{
	1,  // 0: chef.automate.domain.ingest.ConfigRequest.v1:type_name -> chef.automate.domain.ingest.ConfigRequest.V1
	2,  // 1: chef.automate.domain.ingest.ConfigRequest.V1.sys:type_name -> chef.automate.domain.ingest.ConfigRequest.V1.System
	3,  // 2: chef.automate.domain.ingest.ConfigRequest.V1.svc:type_name -> chef.automate.domain.ingest.ConfigRequest.V1.Service
	6,  // 3: chef.automate.domain.ingest.ConfigRequest.V1.System.mlsa:type_name -> chef.automate.infra.config.Mlsa
	4,  // 4: chef.automate.domain.ingest.ConfigRequest.V1.System.service:type_name -> chef.automate.domain.ingest.ConfigRequest.V1.System.Service
	7,  // 5: chef.automate.domain.ingest.ConfigRequest.V1.System.tls:type_name -> chef.automate.infra.config.TLSCredentials
	5,  // 6: chef.automate.domain.ingest.ConfigRequest.V1.System.log:type_name -> chef.automate.domain.ingest.ConfigRequest.V1.System.Log
	8,  // 7: chef.automate.domain.ingest.ConfigRequest.V1.System.Service.host:type_name -> google.protobuf.StringValue
	9,  // 8: chef.automate.domain.ingest.ConfigRequest.V1.System.Service.port:type_name -> google.protobuf.Int32Value
	9,  // 9: chef.automate.domain.ingest.ConfigRequest.V1.System.Service.purge_converge_history_after_days:type_name -> google.protobuf.Int32Value
	9,  // 10: chef.automate.domain.ingest.ConfigRequest.V1.System.Service.purge_actions_after_days:type_name -> google.protobuf.Int32Value
	9,  // 11: chef.automate.domain.ingest.ConfigRequest.V1.System.Service.max_number_of_bundled_run_msgs:type_name -> google.protobuf.Int32Value
	9,  // 12: chef.automate.domain.ingest.ConfigRequest.V1.System.Service.max_number_of_bundled_action_msgs:type_name -> google.protobuf.Int32Value
	9,  // 13: chef.automate.domain.ingest.ConfigRequest.V1.System.Service.number_of_run_msgs_transformers:type_name -> google.protobuf.Int32Value
	9,  // 14: chef.automate.domain.ingest.ConfigRequest.V1.System.Service.number_of_run_msg_publishers:type_name -> google.protobuf.Int32Value
	9,  // 15: chef.automate.domain.ingest.ConfigRequest.V1.System.Service.message_buffer_size:type_name -> google.protobuf.Int32Value
	10, // 16: chef.automate.domain.ingest.ConfigRequest.V1.System.Service.nodes_missing_running_default:type_name -> google.protobuf.BoolValue
	10, // 17: chef.automate.domain.ingest.ConfigRequest.V1.System.Service.missing_nodes_for_deletion_running_default:type_name -> google.protobuf.BoolValue
	9,  // 18: chef.automate.domain.ingest.ConfigRequest.V1.System.Service.number_of_nodemanager_publishers:type_name -> google.protobuf.Int32Value
	8,  // 19: chef.automate.domain.ingest.ConfigRequest.V1.System.Log.format:type_name -> google.protobuf.StringValue
	8,  // 20: chef.automate.domain.ingest.ConfigRequest.V1.System.Log.level:type_name -> google.protobuf.StringValue
	21, // [21:21] is the sub-list for method output_type
	21, // [21:21] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_config_ingest_config_request_proto_init() }
func file_config_ingest_config_request_proto_init() {
	if File_config_ingest_config_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_ingest_config_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest); i {
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
		file_config_ingest_config_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1); i {
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
		file_config_ingest_config_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System); i {
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
		file_config_ingest_config_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_Service); i {
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
		file_config_ingest_config_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System_Service); i {
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
		file_config_ingest_config_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System_Log); i {
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
			RawDescriptor: file_config_ingest_config_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_ingest_config_request_proto_goTypes,
		DependencyIndexes: file_config_ingest_config_request_proto_depIdxs,
		MessageInfos:      file_config_ingest_config_request_proto_msgTypes,
	}.Build()
	File_config_ingest_config_request_proto = out.File
	file_config_ingest_config_request_proto_rawDesc = nil
	file_config_ingest_config_request_proto_goTypes = nil
	file_config_ingest_config_request_proto_depIdxs = nil
}
