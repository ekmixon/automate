// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/infra_proxy/response/databags.proto

package response

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DataBags struct {
	// List of data bags item.
	DataBags             []*DataBagListItem `protobuf:"bytes,2,rep,name=data_bags,json=dataBags,proto3" json:"data_bags,omitempty" toml:"data_bags,omitempty" mapstructure:"data_bags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32              `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DataBags) Reset()         { *m = DataBags{} }
func (m *DataBags) String() string { return proto.CompactTextString(m) }
func (*DataBags) ProtoMessage()    {}
func (*DataBags) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd99e1b7099a19b5, []int{0}
}

func (m *DataBags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBags.Unmarshal(m, b)
}
func (m *DataBags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBags.Marshal(b, m, deterministic)
}
func (m *DataBags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBags.Merge(m, src)
}
func (m *DataBags) XXX_Size() int {
	return xxx_messageInfo_DataBags.Size(m)
}
func (m *DataBags) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBags.DiscardUnknown(m)
}

var xxx_messageInfo_DataBags proto.InternalMessageInfo

func (m *DataBags) GetDataBags() []*DataBagListItem {
	if m != nil {
		return m.DataBags
	}
	return nil
}

type DataBagListItem struct {
	// Name of the data bag item.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DataBagListItem) Reset()         { *m = DataBagListItem{} }
func (m *DataBagListItem) String() string { return proto.CompactTextString(m) }
func (*DataBagListItem) ProtoMessage()    {}
func (*DataBagListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd99e1b7099a19b5, []int{1}
}

func (m *DataBagListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBagListItem.Unmarshal(m, b)
}
func (m *DataBagListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBagListItem.Marshal(b, m, deterministic)
}
func (m *DataBagListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBagListItem.Merge(m, src)
}
func (m *DataBagListItem) XXX_Size() int {
	return xxx_messageInfo_DataBagListItem.Size(m)
}
func (m *DataBagListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBagListItem.DiscardUnknown(m)
}

var xxx_messageInfo_DataBagListItem proto.InternalMessageInfo

func (m *DataBagListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DataBag struct {
	// Data bag name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Data bag item ID.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Stringified json of data bag item.
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty" toml:"data,omitempty" mapstructure:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DataBag) Reset()         { *m = DataBag{} }
func (m *DataBag) String() string { return proto.CompactTextString(m) }
func (*DataBag) ProtoMessage()    {}
func (*DataBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd99e1b7099a19b5, []int{2}
}

func (m *DataBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBag.Unmarshal(m, b)
}
func (m *DataBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBag.Marshal(b, m, deterministic)
}
func (m *DataBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBag.Merge(m, src)
}
func (m *DataBag) XXX_Size() int {
	return xxx_messageInfo_DataBag.Size(m)
}
func (m *DataBag) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBag.DiscardUnknown(m)
}

var xxx_messageInfo_DataBag proto.InternalMessageInfo

func (m *DataBag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataBag) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataBag) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type CreateDataBag struct {
	// Data bag name.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *CreateDataBag) Reset()         { *m = CreateDataBag{} }
func (m *CreateDataBag) String() string { return proto.CompactTextString(m) }
func (*CreateDataBag) ProtoMessage()    {}
func (*CreateDataBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd99e1b7099a19b5, []int{3}
}

func (m *CreateDataBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDataBag.Unmarshal(m, b)
}
func (m *CreateDataBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDataBag.Marshal(b, m, deterministic)
}
func (m *CreateDataBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDataBag.Merge(m, src)
}
func (m *CreateDataBag) XXX_Size() int {
	return xxx_messageInfo_CreateDataBag.Size(m)
}
func (m *CreateDataBag) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDataBag.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDataBag proto.InternalMessageInfo

func (m *CreateDataBag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateDataBagItem struct {
	// Data bag name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Data bag item ID.
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *CreateDataBagItem) Reset()         { *m = CreateDataBagItem{} }
func (m *CreateDataBagItem) String() string { return proto.CompactTextString(m) }
func (*CreateDataBagItem) ProtoMessage()    {}
func (*CreateDataBagItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd99e1b7099a19b5, []int{4}
}

func (m *CreateDataBagItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDataBagItem.Unmarshal(m, b)
}
func (m *CreateDataBagItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDataBagItem.Marshal(b, m, deterministic)
}
func (m *CreateDataBagItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDataBagItem.Merge(m, src)
}
func (m *CreateDataBagItem) XXX_Size() int {
	return xxx_messageInfo_CreateDataBagItem.Size(m)
}
func (m *CreateDataBagItem) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDataBagItem.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDataBagItem proto.InternalMessageInfo

func (m *CreateDataBagItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateDataBagItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateDataBagItem struct {
	// Data bag name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Data bag item ID.
	ItemId               string   `protobuf:"bytes,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty" toml:"item_id,omitempty" mapstructure:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *UpdateDataBagItem) Reset()         { *m = UpdateDataBagItem{} }
func (m *UpdateDataBagItem) String() string { return proto.CompactTextString(m) }
func (*UpdateDataBagItem) ProtoMessage()    {}
func (*UpdateDataBagItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd99e1b7099a19b5, []int{5}
}

func (m *UpdateDataBagItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDataBagItem.Unmarshal(m, b)
}
func (m *UpdateDataBagItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDataBagItem.Marshal(b, m, deterministic)
}
func (m *UpdateDataBagItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDataBagItem.Merge(m, src)
}
func (m *UpdateDataBagItem) XXX_Size() int {
	return xxx_messageInfo_UpdateDataBagItem.Size(m)
}
func (m *UpdateDataBagItem) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDataBagItem.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDataBagItem proto.InternalMessageInfo

func (m *UpdateDataBagItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateDataBagItem) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func init() {
	proto.RegisterType((*DataBags)(nil), "chef.automate.domain.infra_proxy.response.DataBags")
	proto.RegisterType((*DataBagListItem)(nil), "chef.automate.domain.infra_proxy.response.DataBagListItem")
	proto.RegisterType((*DataBag)(nil), "chef.automate.domain.infra_proxy.response.DataBag")
	proto.RegisterType((*CreateDataBag)(nil), "chef.automate.domain.infra_proxy.response.CreateDataBag")
	proto.RegisterType((*CreateDataBagItem)(nil), "chef.automate.domain.infra_proxy.response.CreateDataBagItem")
	proto.RegisterType((*UpdateDataBagItem)(nil), "chef.automate.domain.infra_proxy.response.UpdateDataBagItem")
}

func init() {
	proto.RegisterFile("api/interservice/infra_proxy/response/databags.proto", fileDescriptor_bd99e1b7099a19b5)
}

var fileDescriptor_bd99e1b7099a19b5 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4b, 0x3b, 0x31,
	0x10, 0xc5, 0xe9, 0xf6, 0x4b, 0x7f, 0xcc, 0x17, 0x95, 0xee, 0xc5, 0x3d, 0x96, 0x88, 0x50, 0x2f,
	0x09, 0xa8, 0x20, 0x78, 0x90, 0x5a, 0xbd, 0x14, 0x3c, 0x15, 0x44, 0xf0, 0x52, 0xa6, 0x9b, 0xe9,
	0x36, 0x87, 0x6c, 0x96, 0x64, 0x2a, 0xfa, 0xdf, 0x4b, 0xd6, 0xad, 0x68, 0x51, 0xd9, 0xdb, 0x24,
	0xf3, 0xde, 0x67, 0x86, 0x79, 0x70, 0x89, 0x95, 0x51, 0xa6, 0x64, 0xf2, 0x81, 0xfc, 0x8b, 0xc9,
	0x49, 0x99, 0x72, 0xed, 0x71, 0x59, 0x79, 0xf7, 0xfa, 0xa6, 0x3c, 0x85, 0xca, 0x95, 0x81, 0x94,
	0x46, 0xc6, 0x15, 0x16, 0x41, 0x56, 0xde, 0xb1, 0x4b, 0xcf, 0xf2, 0x0d, 0xad, 0x25, 0x6e, 0xd9,
	0x59, 0x64, 0x92, 0xda, 0x59, 0x34, 0xa5, 0xfc, 0xe2, 0x94, 0x3b, 0xa7, 0xc8, 0x61, 0x70, 0x8f,
	0x8c, 0x33, 0x2c, 0x42, 0xfa, 0x04, 0xc3, 0x08, 0x5a, 0x46, 0x52, 0x96, 0x8c, 0xbb, 0x93, 0xff,
	0xe7, 0xd7, 0xb2, 0x35, 0x4a, 0x36, 0x9c, 0x07, 0x13, 0x78, 0xce, 0x64, 0x17, 0x03, 0xdd, 0x80,
	0xc5, 0x29, 0x1c, 0xed, 0x35, 0xd3, 0x14, 0xfe, 0x95, 0x68, 0x29, 0xeb, 0x8c, 0x3b, 0x93, 0xe1,
	0xa2, 0xae, 0xc5, 0x2d, 0xf4, 0x1b, 0xd9, 0x4f, 0xed, 0xf4, 0x10, 0x12, 0xa3, 0xb3, 0xa4, 0xfe,
	0x49, 0x8c, 0x8e, 0x9a, 0x38, 0x21, 0xeb, 0x7e, 0x68, 0x62, 0x2d, 0x4e, 0xe0, 0xe0, 0xce, 0x13,
	0x32, 0xfd, 0x01, 0x12, 0x57, 0x30, 0xfa, 0x26, 0xfa, 0x6d, 0xa1, 0xfd, 0x89, 0x62, 0x0a, 0xa3,
	0xc7, 0x4a, 0xb7, 0x30, 0x1e, 0x43, 0xdf, 0x30, 0xd9, 0xe5, 0xa7, 0xbb, 0x17, 0x9f, 0x73, 0x3d,
	0x9b, 0x3e, 0xdf, 0x14, 0x86, 0x37, 0xdb, 0x95, 0xcc, 0x9d, 0x55, 0xf1, 0xb6, 0x6a, 0x77, 0x5b,
	0xd5, 0x2a, 0xea, 0x55, 0xaf, 0x8e, 0xf8, 0xe2, 0x3d, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xa0, 0xcd,
	0x7e, 0x1a, 0x02, 0x00, 0x00,
}
