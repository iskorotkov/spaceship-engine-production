// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: api/report-printer.proto

package api

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

type PrintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filepath    string    `protobuf:"bytes,1,opt,name=filepath,proto3" json:"filepath,omitempty"`
	TotalOrders int32     `protobuf:"varint,2,opt,name=totalOrders,proto3" json:"totalOrders,omitempty"`
	TopClients  []*Client `protobuf:"bytes,3,rep,name=topClients,proto3" json:"topClients,omitempty"`
	TopEngines  []*Engine `protobuf:"bytes,4,rep,name=topEngines,proto3" json:"topEngines,omitempty"`
}

func (x *PrintRequest) Reset() {
	*x = PrintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_report_printer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintRequest) ProtoMessage() {}

func (x *PrintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_report_printer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintRequest.ProtoReflect.Descriptor instead.
func (*PrintRequest) Descriptor() ([]byte, []int) {
	return file_api_report_printer_proto_rawDescGZIP(), []int{0}
}

func (x *PrintRequest) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *PrintRequest) GetTotalOrders() int32 {
	if x != nil {
		return x.TotalOrders
	}
	return 0
}

func (x *PrintRequest) GetTopClients() []*Client {
	if x != nil {
		return x.TopClients
	}
	return nil
}

func (x *PrintRequest) GetTopEngines() []*Engine {
	if x != nil {
		return x.TopEngines
	}
	return nil
}

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Orders int32  `protobuf:"varint,2,opt,name=orders,proto3" json:"orders,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_report_printer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_api_report_printer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_api_report_printer_proto_rawDescGZIP(), []int{1}
}

func (x *Client) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Client) GetOrders() int32 {
	if x != nil {
		return x.Orders
	}
	return 0
}

type Engine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Power  float64 `protobuf:"fixed64,2,opt,name=power,proto3" json:"power,omitempty"`
	Orders int32   `protobuf:"varint,3,opt,name=orders,proto3" json:"orders,omitempty"`
}

func (x *Engine) Reset() {
	*x = Engine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_report_printer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Engine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Engine) ProtoMessage() {}

func (x *Engine) ProtoReflect() protoreflect.Message {
	mi := &file_api_report_printer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Engine.ProtoReflect.Descriptor instead.
func (*Engine) Descriptor() ([]byte, []int) {
	return file_api_report_printer_proto_rawDescGZIP(), []int{2}
}

func (x *Engine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Engine) GetPower() float64 {
	if x != nil {
		return x.Power
	}
	return 0
}

func (x *Engine) GetOrders() int32 {
	if x != nil {
		return x.Orders
	}
	return 0
}

type PrintResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PrintResponse) Reset() {
	*x = PrintResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_report_printer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrintResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintResponse) ProtoMessage() {}

func (x *PrintResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_report_printer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintResponse.ProtoReflect.Descriptor instead.
func (*PrintResponse) Descriptor() ([]byte, []int) {
	return file_api_report_printer_proto_rawDescGZIP(), []int{3}
}

var File_api_report_printer_proto protoreflect.FileDescriptor

var file_api_report_printer_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x0c, 0x50,
	0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x0a, 0x74, 0x6f, 0x70,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x74, 0x6f, 0x70, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x27, 0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52,
	0x0a, 0x74, 0x6f, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x06, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x22, 0x4a, 0x0a, 0x06, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x70, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x0f, 0x0a,
	0x0d, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x39,
	0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x28, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x73, 0x6b, 0x6f, 0x72, 0x6f, 0x74, 0x6b,
	0x6f, 0x76, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x68, 0x69, 0x70, 0x2d, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_report_printer_proto_rawDescOnce sync.Once
	file_api_report_printer_proto_rawDescData = file_api_report_printer_proto_rawDesc
)

func file_api_report_printer_proto_rawDescGZIP() []byte {
	file_api_report_printer_proto_rawDescOnce.Do(func() {
		file_api_report_printer_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_report_printer_proto_rawDescData)
	})
	return file_api_report_printer_proto_rawDescData
}

var file_api_report_printer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_report_printer_proto_goTypes = []interface{}{
	(*PrintRequest)(nil),  // 0: PrintRequest
	(*Client)(nil),        // 1: Client
	(*Engine)(nil),        // 2: Engine
	(*PrintResponse)(nil), // 3: PrintResponse
}
var file_api_report_printer_proto_depIdxs = []int32{
	1, // 0: PrintRequest.topClients:type_name -> Client
	2, // 1: PrintRequest.topEngines:type_name -> Engine
	0, // 2: ReportPrinter.Print:input_type -> PrintRequest
	3, // 3: ReportPrinter.Print:output_type -> PrintResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_report_printer_proto_init() }
func file_api_report_printer_proto_init() {
	if File_api_report_printer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_report_printer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrintRequest); i {
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
		file_api_report_printer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_api_report_printer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Engine); i {
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
		file_api_report_printer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrintResponse); i {
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
			RawDescriptor: file_api_report_printer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_report_printer_proto_goTypes,
		DependencyIndexes: file_api_report_printer_proto_depIdxs,
		MessageInfos:      file_api_report_printer_proto_msgTypes,
	}.Build()
	File_api_report_printer_proto = out.File
	file_api_report_printer_proto_rawDesc = nil
	file_api_report_printer_proto_goTypes = nil
	file_api_report_printer_proto_depIdxs = nil
}
