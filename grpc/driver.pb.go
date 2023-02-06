// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: driver.proto

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

type FindDriverReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string   `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Type   TaxiType `protobuf:"varint,2,opt,name=Type,proto3,enum=grpc.TaxiType" json:"Type,omitempty"`
}

func (x *FindDriverReq) Reset() {
	*x = FindDriverReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindDriverReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindDriverReq) ProtoMessage() {}

func (x *FindDriverReq) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindDriverReq.ProtoReflect.Descriptor instead.
func (*FindDriverReq) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{0}
}

func (x *FindDriverReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FindDriverReq) GetType() TaxiType {
	if x != nil {
		return x.Type
	}
	return TaxiType_UndefinedTaxi
}

type FindDriverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId string `protobuf:"bytes,1,opt,name=DriverId,proto3" json:"DriverId,omitempty"`
}

func (x *FindDriverResponse) Reset() {
	*x = FindDriverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindDriverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindDriverResponse) ProtoMessage() {}

func (x *FindDriverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindDriverResponse.ProtoReflect.Descriptor instead.
func (*FindDriverResponse) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{1}
}

func (x *FindDriverResponse) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

var File_driver_proto protoreflect.FileDescriptor

var file_driver_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x67, 0x72, 0x70, 0x63, 0x1a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4b, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x54, 0x61, 0x78, 0x69, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x30,
	0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x32, 0x47, 0x0a, 0x06, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0a, 0x46, 0x69,
	0x6e, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x2d, 0x6b, 0x6f,
	0x72, 0x6f, 0x73, 0x74, 0x65, 0x6c, 0x65, 0x76, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_driver_proto_rawDescOnce sync.Once
	file_driver_proto_rawDescData = file_driver_proto_rawDesc
)

func file_driver_proto_rawDescGZIP() []byte {
	file_driver_proto_rawDescOnce.Do(func() {
		file_driver_proto_rawDescData = protoimpl.X.CompressGZIP(file_driver_proto_rawDescData)
	})
	return file_driver_proto_rawDescData
}

var file_driver_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_driver_proto_goTypes = []interface{}{
	(*FindDriverReq)(nil),      // 0: grpc.FindDriverReq
	(*FindDriverResponse)(nil), // 1: grpc.FindDriverResponse
	(TaxiType)(0),              // 2: grpc.TaxiType
}
var file_driver_proto_depIdxs = []int32{
	2, // 0: grpc.FindDriverReq.Type:type_name -> grpc.TaxiType
	0, // 1: grpc.Driver.FindDriver:input_type -> grpc.FindDriverReq
	1, // 2: grpc.Driver.FindDriver:output_type -> grpc.FindDriverResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_driver_proto_init() }
func file_driver_proto_init() {
	if File_driver_proto != nil {
		return
	}
	file_order_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_driver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindDriverReq); i {
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
		file_driver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindDriverResponse); i {
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
			RawDescriptor: file_driver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_driver_proto_goTypes,
		DependencyIndexes: file_driver_proto_depIdxs,
		MessageInfos:      file_driver_proto_msgTypes,
	}.Build()
	File_driver_proto = out.File
	file_driver_proto_rawDesc = nil
	file_driver_proto_goTypes = nil
	file_driver_proto_depIdxs = nil
}
