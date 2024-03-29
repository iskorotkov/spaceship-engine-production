// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: api/data-loader/data-loader.proto

package data_loader

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

type ClientReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ClientReq) Reset() {
	*x = ClientReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_loader_data_loader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientReq) ProtoMessage() {}

func (x *ClientReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_loader_data_loader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientReq.ProtoReflect.Descriptor instead.
func (*ClientReq) Descriptor() ([]byte, []int) {
	return file_api_data_loader_data_loader_proto_rawDescGZIP(), []int{0}
}

func (x *ClientReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ClientReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ClientResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientResp) Reset() {
	*x = ClientResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_loader_data_loader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientResp) ProtoMessage() {}

func (x *ClientResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_loader_data_loader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientResp.ProtoReflect.Descriptor instead.
func (*ClientResp) Descriptor() ([]byte, []int) {
	return file_api_data_loader_data_loader_proto_rawDescGZIP(), []int{1}
}

type ComponentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ComponentReq) Reset() {
	*x = ComponentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_loader_data_loader_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentReq) ProtoMessage() {}

func (x *ComponentReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_loader_data_loader_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentReq.ProtoReflect.Descriptor instead.
func (*ComponentReq) Descriptor() ([]byte, []int) {
	return file_api_data_loader_data_loader_proto_rawDescGZIP(), []int{2}
}

func (x *ComponentReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ComponentReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ComponentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ComponentResp) Reset() {
	*x = ComponentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_loader_data_loader_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentResp) ProtoMessage() {}

func (x *ComponentResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_loader_data_loader_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentResp.ProtoReflect.Descriptor instead.
func (*ComponentResp) Descriptor() ([]byte, []int) {
	return file_api_data_loader_data_loader_proto_rawDescGZIP(), []int{3}
}

type EngineReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Power float64 `protobuf:"fixed64,3,opt,name=power,proto3" json:"power,omitempty"`
}

func (x *EngineReq) Reset() {
	*x = EngineReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_loader_data_loader_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EngineReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EngineReq) ProtoMessage() {}

func (x *EngineReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_loader_data_loader_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EngineReq.ProtoReflect.Descriptor instead.
func (*EngineReq) Descriptor() ([]byte, []int) {
	return file_api_data_loader_data_loader_proto_rawDescGZIP(), []int{4}
}

func (x *EngineReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EngineReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EngineReq) GetPower() float64 {
	if x != nil {
		return x.Power
	}
	return 0
}

type EngineResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EngineResp) Reset() {
	*x = EngineResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_loader_data_loader_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EngineResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EngineResp) ProtoMessage() {}

func (x *EngineResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_loader_data_loader_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EngineResp.ProtoReflect.Descriptor instead.
func (*EngineResp) Descriptor() ([]byte, []int) {
	return file_api_data_loader_data_loader_proto_rawDescGZIP(), []int{5}
}

type OrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount      int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	CreatedAt   int64  `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	CompletedAt int64  `protobuf:"varint,4,opt,name=completedAt,proto3" json:"completedAt,omitempty"`
	ClientId    uint32 `protobuf:"varint,5,opt,name=clientId,proto3" json:"clientId,omitempty"`
	EngineId    uint32 `protobuf:"varint,6,opt,name=engineId,proto3" json:"engineId,omitempty"`
}

func (x *OrderReq) Reset() {
	*x = OrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_loader_data_loader_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderReq) ProtoMessage() {}

func (x *OrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_loader_data_loader_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderReq.ProtoReflect.Descriptor instead.
func (*OrderReq) Descriptor() ([]byte, []int) {
	return file_api_data_loader_data_loader_proto_rawDescGZIP(), []int{6}
}

func (x *OrderReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderReq) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *OrderReq) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *OrderReq) GetCompletedAt() int64 {
	if x != nil {
		return x.CompletedAt
	}
	return 0
}

func (x *OrderReq) GetClientId() uint32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *OrderReq) GetEngineId() uint32 {
	if x != nil {
		return x.EngineId
	}
	return 0
}

type OrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderResp) Reset() {
	*x = OrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_loader_data_loader_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResp) ProtoMessage() {}

func (x *OrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_loader_data_loader_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResp.ProtoReflect.Descriptor instead.
func (*OrderResp) Descriptor() ([]byte, []int) {
	return file_api_data_loader_data_loader_proto_rawDescGZIP(), []int{7}
}

type ProviderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ProviderReq) Reset() {
	*x = ProviderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_loader_data_loader_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderReq) ProtoMessage() {}

func (x *ProviderReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_loader_data_loader_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderReq.ProtoReflect.Descriptor instead.
func (*ProviderReq) Descriptor() ([]byte, []int) {
	return file_api_data_loader_data_loader_proto_rawDescGZIP(), []int{8}
}

func (x *ProviderReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProviderReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ProviderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProviderResp) Reset() {
	*x = ProviderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_loader_data_loader_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderResp) ProtoMessage() {}

func (x *ProviderResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_loader_data_loader_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderResp.ProtoReflect.Descriptor instead.
func (*ProviderResp) Descriptor() ([]byte, []int) {
	return file_api_data_loader_data_loader_proto_rawDescGZIP(), []int{9}
}

type RequirementReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EngineId    uint32 `protobuf:"varint,2,opt,name=engineId,proto3" json:"engineId,omitempty"`
	ComponentId uint32 `protobuf:"varint,3,opt,name=componentId,proto3" json:"componentId,omitempty"`
}

func (x *RequirementReq) Reset() {
	*x = RequirementReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_loader_data_loader_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequirementReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequirementReq) ProtoMessage() {}

func (x *RequirementReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_loader_data_loader_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequirementReq.ProtoReflect.Descriptor instead.
func (*RequirementReq) Descriptor() ([]byte, []int) {
	return file_api_data_loader_data_loader_proto_rawDescGZIP(), []int{10}
}

func (x *RequirementReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RequirementReq) GetEngineId() uint32 {
	if x != nil {
		return x.EngineId
	}
	return 0
}

func (x *RequirementReq) GetComponentId() uint32 {
	if x != nil {
		return x.ComponentId
	}
	return 0
}

type RequirementResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequirementResp) Reset() {
	*x = RequirementResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_loader_data_loader_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequirementResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequirementResp) ProtoMessage() {}

func (x *RequirementResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_loader_data_loader_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequirementResp.ProtoReflect.Descriptor instead.
func (*RequirementResp) Descriptor() ([]byte, []int) {
	return file_api_data_loader_data_loader_proto_rawDescGZIP(), []int{11}
}

type SupplyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ComponentId uint32 `protobuf:"varint,2,opt,name=componentId,proto3" json:"componentId,omitempty"`
	ProviderId  uint32 `protobuf:"varint,3,opt,name=providerId,proto3" json:"providerId,omitempty"`
}

func (x *SupplyReq) Reset() {
	*x = SupplyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_loader_data_loader_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplyReq) ProtoMessage() {}

func (x *SupplyReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_loader_data_loader_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplyReq.ProtoReflect.Descriptor instead.
func (*SupplyReq) Descriptor() ([]byte, []int) {
	return file_api_data_loader_data_loader_proto_rawDescGZIP(), []int{12}
}

func (x *SupplyReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SupplyReq) GetComponentId() uint32 {
	if x != nil {
		return x.ComponentId
	}
	return 0
}

func (x *SupplyReq) GetProviderId() uint32 {
	if x != nil {
		return x.ProviderId
	}
	return 0
}

type SupplyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SupplyResp) Reset() {
	*x = SupplyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_loader_data_loader_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplyResp) ProtoMessage() {}

func (x *SupplyResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_loader_data_loader_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplyResp.ProtoReflect.Descriptor instead.
func (*SupplyResp) Descriptor() ([]byte, []int) {
	return file_api_data_loader_data_loader_proto_rawDescGZIP(), []int{13}
}

var File_api_data_loader_data_loader_proto protoreflect.FileDescriptor

var file_api_data_loader_data_loader_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x6c, 0x6f, 0x61, 0x64, 0x65,
	0x72, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x32, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x45, 0x0a, 0x09, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x77, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x22, 0x0c,
	0x0a, 0x0a, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0xaa, 0x01, 0x0a,
	0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x22, 0x0b, 0x0a, 0x09, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x31, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x5e, 0x0a, 0x0e, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x5d, 0x0a, 0x09,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x0c, 0x0a, 0x0a, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x32, 0xc6, 0x02, 0x0a, 0x0a, 0x44, 0x61,
	0x74, 0x61, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0a, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x30, 0x0a, 0x0d, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x0e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x12, 0x0a, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x09,
	0x53, 0x61, 0x76, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x12, 0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x0d, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x36, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x0a, 0x53, 0x61, 0x76,
	0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x0a, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79,
	0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x73, 0x6b, 0x6f, 0x72, 0x6f, 0x74, 0x6b, 0x6f, 0x76, 0x2f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x68, 0x69, 0x70, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2d, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2d, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_data_loader_data_loader_proto_rawDescOnce sync.Once
	file_api_data_loader_data_loader_proto_rawDescData = file_api_data_loader_data_loader_proto_rawDesc
)

func file_api_data_loader_data_loader_proto_rawDescGZIP() []byte {
	file_api_data_loader_data_loader_proto_rawDescOnce.Do(func() {
		file_api_data_loader_data_loader_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_data_loader_data_loader_proto_rawDescData)
	})
	return file_api_data_loader_data_loader_proto_rawDescData
}

var file_api_data_loader_data_loader_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_data_loader_data_loader_proto_goTypes = []interface{}{
	(*ClientReq)(nil),       // 0: ClientReq
	(*ClientResp)(nil),      // 1: ClientResp
	(*ComponentReq)(nil),    // 2: ComponentReq
	(*ComponentResp)(nil),   // 3: ComponentResp
	(*EngineReq)(nil),       // 4: EngineReq
	(*EngineResp)(nil),      // 5: EngineResp
	(*OrderReq)(nil),        // 6: OrderReq
	(*OrderResp)(nil),       // 7: OrderResp
	(*ProviderReq)(nil),     // 8: ProviderReq
	(*ProviderResp)(nil),    // 9: ProviderResp
	(*RequirementReq)(nil),  // 10: RequirementReq
	(*RequirementResp)(nil), // 11: RequirementResp
	(*SupplyReq)(nil),       // 12: SupplyReq
	(*SupplyResp)(nil),      // 13: SupplyResp
}
var file_api_data_loader_data_loader_proto_depIdxs = []int32{
	0,  // 0: DataLoader.SaveClient:input_type -> ClientReq
	2,  // 1: DataLoader.SaveComponent:input_type -> ComponentReq
	4,  // 2: DataLoader.SaveEngine:input_type -> EngineReq
	6,  // 3: DataLoader.SaveOrder:input_type -> OrderReq
	8,  // 4: DataLoader.SaveProvider:input_type -> ProviderReq
	10, // 5: DataLoader.SaveRequirement:input_type -> RequirementReq
	12, // 6: DataLoader.SaveSupply:input_type -> SupplyReq
	1,  // 7: DataLoader.SaveClient:output_type -> ClientResp
	3,  // 8: DataLoader.SaveComponent:output_type -> ComponentResp
	5,  // 9: DataLoader.SaveEngine:output_type -> EngineResp
	7,  // 10: DataLoader.SaveOrder:output_type -> OrderResp
	9,  // 11: DataLoader.SaveProvider:output_type -> ProviderResp
	11, // 12: DataLoader.SaveRequirement:output_type -> RequirementResp
	13, // 13: DataLoader.SaveSupply:output_type -> SupplyResp
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_data_loader_data_loader_proto_init() }
func file_api_data_loader_data_loader_proto_init() {
	if File_api_data_loader_data_loader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_data_loader_data_loader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientReq); i {
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
		file_api_data_loader_data_loader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientResp); i {
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
		file_api_data_loader_data_loader_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentReq); i {
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
		file_api_data_loader_data_loader_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentResp); i {
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
		file_api_data_loader_data_loader_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EngineReq); i {
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
		file_api_data_loader_data_loader_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EngineResp); i {
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
		file_api_data_loader_data_loader_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderReq); i {
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
		file_api_data_loader_data_loader_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResp); i {
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
		file_api_data_loader_data_loader_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderReq); i {
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
		file_api_data_loader_data_loader_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderResp); i {
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
		file_api_data_loader_data_loader_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequirementReq); i {
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
		file_api_data_loader_data_loader_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequirementResp); i {
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
		file_api_data_loader_data_loader_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplyReq); i {
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
		file_api_data_loader_data_loader_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplyResp); i {
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
			RawDescriptor: file_api_data_loader_data_loader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_data_loader_data_loader_proto_goTypes,
		DependencyIndexes: file_api_data_loader_data_loader_proto_depIdxs,
		MessageInfos:      file_api_data_loader_data_loader_proto_msgTypes,
	}.Build()
	File_api_data_loader_data_loader_proto = out.File
	file_api_data_loader_data_loader_proto_rawDesc = nil
	file_api_data_loader_data_loader_proto_goTypes = nil
	file_api_data_loader_data_loader_proto_depIdxs = nil
}
