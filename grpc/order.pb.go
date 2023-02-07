// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: order.proto

package grpc_schemas

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

type TaxiType int32

const (
	TaxiType_UndefinedTaxi TaxiType = 0
	TaxiType_Economy       TaxiType = 1
	TaxiType_Comfort       TaxiType = 2
	TaxiType_Business      TaxiType = 3
)

// Enum value maps for TaxiType.
var (
	TaxiType_name = map[int32]string{
		0: "UndefinedTaxi",
		1: "Economy",
		2: "Comfort",
		3: "Business",
	}
	TaxiType_value = map[string]int32{
		"UndefinedTaxi": 0,
		"Economy":       1,
		"Comfort":       2,
		"Business":      3,
	}
)

func (x TaxiType) Enum() *TaxiType {
	p := new(TaxiType)
	*p = x
	return p
}

func (x TaxiType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaxiType) Descriptor() protoreflect.EnumDescriptor {
	return file_order_proto_enumTypes[0].Descriptor()
}

func (TaxiType) Type() protoreflect.EnumType {
	return &file_order_proto_enumTypes[0]
}

func (x TaxiType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaxiType.Descriptor instead.
func (TaxiType) EnumDescriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

type StatusType int32

const (
	StatusType_UndefinedStatus StatusType = 0
	StatusType_inProgress      StatusType = 1
	StatusType_Finished        StatusType = 2
)

// Enum value maps for StatusType.
var (
	StatusType_name = map[int32]string{
		0: "UndefinedStatus",
		1: "inProgress",
		2: "Finished",
	}
	StatusType_value = map[string]int32{
		"UndefinedStatus": 0,
		"inProgress":      1,
		"Finished":        2,
	}
)

func (x StatusType) Enum() *StatusType {
	p := new(StatusType)
	*p = x
	return p
}

func (x StatusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusType) Descriptor() protoreflect.EnumDescriptor {
	return file_order_proto_enumTypes[1].Descriptor()
}

func (StatusType) Type() protoreflect.EnumType {
	return &file_order_proto_enumTypes[1]
}

func (x StatusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusType.Descriptor instead.
func (StatusType) EnumDescriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

type DriverRatingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mark float32 `protobuf:"fixed32,1,opt,name=mark,proto3" json:"mark,omitempty"`
	Uuid string  `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *DriverRatingReq) Reset() {
	*x = DriverRatingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverRatingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverRatingReq) ProtoMessage() {}

func (x *DriverRatingReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverRatingReq.ProtoReflect.Descriptor instead.
func (*DriverRatingReq) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *DriverRatingReq) GetMark() float32 {
	if x != nil {
		return x.Mark
	}
	return 0
}

func (x *DriverRatingReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
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
	return file_order_proto_rawDescGZIP(), []int{1}
}

type UserRatingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mark float32 `protobuf:"fixed32,1,opt,name=mark,proto3" json:"mark,omitempty"`
	Id   string  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserRatingReq) Reset() {
	*x = UserRatingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRatingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRatingReq) ProtoMessage() {}

func (x *UserRatingReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRatingReq.ProtoReflect.Descriptor instead.
func (*UserRatingReq) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *UserRatingReq) GetMark() float32 {
	if x != nil {
		return x.Mark
	}
	return 0
}

func (x *UserRatingReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type OrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	From   string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To     string   `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Type   TaxiType `protobuf:"varint,4,opt,name=type,proto3,enum=grpc.TaxiType" json:"type,omitempty"`
}

func (x *OrderReq) Reset() {
	*x = OrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderReq) ProtoMessage() {}

func (x *OrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
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
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *OrderReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderReq) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *OrderReq) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *OrderReq) GetType() TaxiType {
	if x != nil {
		return x.Type
	}
	return TaxiType_UndefinedTaxi
}

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *OrderResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type FindOrdersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId *string `protobuf:"bytes,1,opt,name=DriverId,proto3,oneof" json:"DriverId,omitempty"`
	UserId   *string `protobuf:"bytes,2,opt,name=UserId,proto3,oneof" json:"UserId,omitempty"`
	From     *string `protobuf:"bytes,3,opt,name=From,proto3,oneof" json:"From,omitempty"`
	To       *string `protobuf:"bytes,4,opt,name=To,proto3,oneof" json:"To,omitempty"`
}

func (x *FindOrdersReq) Reset() {
	*x = FindOrdersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOrdersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOrdersReq) ProtoMessage() {}

func (x *FindOrdersReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOrdersReq.ProtoReflect.Descriptor instead.
func (*FindOrdersReq) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{5}
}

func (x *FindOrdersReq) GetDriverId() string {
	if x != nil && x.DriverId != nil {
		return *x.DriverId
	}
	return ""
}

func (x *FindOrdersReq) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *FindOrdersReq) GetFrom() string {
	if x != nil && x.From != nil {
		return *x.From
	}
	return ""
}

func (x *FindOrdersReq) GetTo() string {
	if x != nil && x.To != nil {
		return *x.To
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId string     `protobuf:"bytes,1,opt,name=DriverId,proto3" json:"DriverId,omitempty"`
	UserId   string     `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	From     string     `protobuf:"bytes,3,opt,name=From,proto3" json:"From,omitempty"`
	To       string     `protobuf:"bytes,4,opt,name=To,proto3" json:"To,omitempty"`
	Taxi     TaxiType   `protobuf:"varint,5,opt,name=Taxi,proto3,enum=grpc.TaxiType" json:"Taxi,omitempty"`
	UnixTime int64      `protobuf:"varint,6,opt,name=UnixTime,proto3" json:"UnixTime,omitempty"`
	Stat     StatusType `protobuf:"varint,7,opt,name=Stat,proto3,enum=grpc.StatusType" json:"Stat,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{6}
}

func (x *Order) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Order) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Order) GetTaxi() TaxiType {
	if x != nil {
		return x.Taxi
	}
	return TaxiType_UndefinedTaxi
}

func (x *Order) GetUnixTime() int64 {
	if x != nil {
		return x.UnixTime
	}
	return 0
}

func (x *Order) GetStat() StatusType {
	if x != nil {
		return x.Stat
	}
	return StatusType_UndefinedStatus
}

type Orders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order []*Order `protobuf:"bytes,1,rep,name=order,proto3" json:"order,omitempty"`
}

func (x *Orders) Reset() {
	*x = Orders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orders) ProtoMessage() {}

func (x *Orders) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orders.ProtoReflect.Descriptor instead.
func (*Orders) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{7}
}

func (x *Orders) GetOrder() []*Order {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67,
	0x72, 0x70, 0x63, 0x22, 0x39, 0x0a, 0x0f, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x0a,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x0a, 0x0d, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x61, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x6a, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x61, 0x78,
	0x69, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa3, 0x01, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x08, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12,
	0x13, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x02, 0x54,
	0x6f, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x54, 0x6f, 0x22, 0xc5, 0x01, 0x0a,
	0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72,
	0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x54, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x22,
	0x0a, 0x04, 0x54, 0x61, 0x78, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x54, 0x61, 0x78, 0x69, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x61,
	0x78, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x55, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x53, 0x74, 0x61, 0x74, 0x22, 0x2b, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x21,
	0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2a, 0x45, 0x0a, 0x08, 0x54, 0x61, 0x78, 0x69, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a,
	0x0d, 0x55, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x54, 0x61, 0x78, 0x69, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x66, 0x6f, 0x72, 0x74, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x10, 0x03, 0x2a, 0x3f, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x6e, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x69,
	0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x10, 0x02, 0x32, 0xe1, 0x01, 0x0a, 0x0c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x52, 0x61,
	0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a,
	0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x08, 0x52, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0a, 0x46, 0x69,
	0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x00, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x6d, 0x61,
	0x6e, 0x2d, 0x6b, 0x6f, 0x72, 0x6f, 0x73, 0x74, 0x65, 0x6c, 0x65, 0x76, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_order_proto_goTypes = []interface{}{
	(TaxiType)(0),           // 0: grpc.TaxiType
	(StatusType)(0),         // 1: grpc.StatusType
	(*DriverRatingReq)(nil), // 2: grpc.DriverRatingReq
	(*Response)(nil),        // 3: grpc.Response
	(*UserRatingReq)(nil),   // 4: grpc.UserRatingReq
	(*OrderReq)(nil),        // 5: grpc.OrderReq
	(*OrderResponse)(nil),   // 6: grpc.OrderResponse
	(*FindOrdersReq)(nil),   // 7: grpc.FindOrdersReq
	(*Order)(nil),           // 8: grpc.Order
	(*Orders)(nil),          // 9: grpc.Orders
}
var file_order_proto_depIdxs = []int32{
	0, // 0: grpc.OrderReq.type:type_name -> grpc.TaxiType
	0, // 1: grpc.Order.Taxi:type_name -> grpc.TaxiType
	1, // 2: grpc.Order.Stat:type_name -> grpc.StatusType
	8, // 3: grpc.Orders.order:type_name -> grpc.Order
	2, // 4: grpc.OrderService.RateDriver:input_type -> grpc.DriverRatingReq
	4, // 5: grpc.OrderService.RateUser:input_type -> grpc.UserRatingReq
	5, // 6: grpc.OrderService.CreateOrder:input_type -> grpc.OrderReq
	7, // 7: grpc.OrderService.FindOrders:input_type -> grpc.FindOrdersReq
	3, // 8: grpc.OrderService.RateDriver:output_type -> grpc.Response
	3, // 9: grpc.OrderService.RateUser:output_type -> grpc.Response
	6, // 10: grpc.OrderService.CreateOrder:output_type -> grpc.OrderResponse
	9, // 11: grpc.OrderService.FindOrders:output_type -> grpc.Orders
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverRatingReq); i {
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
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRatingReq); i {
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
		file_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResponse); i {
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
		file_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOrdersReq); i {
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
		file_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orders); i {
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
	file_order_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		EnumInfos:         file_order_proto_enumTypes,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
