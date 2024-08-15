// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: internal/infra/grpc/protofiles/order.proto

package pb

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

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax   float32 `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_protofiles_order_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateOrderRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateOrderRequest) GetTax() float32 {
	if x != nil {
		return x.Tax
	}
	return 0
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price      float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax        float32 `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
	FinalPrice float32 `protobuf:"fixed32,4,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_protofiles_order_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateOrderResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateOrderResponse) GetTax() float32 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *CreateOrderResponse) GetFinalPrice() float32 {
	if x != nil {
		return x.FinalPrice
	}
	return 0
}

type ListOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *ListOrderRequest) Reset() {
	*x = ListOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderRequest) ProtoMessage() {}

func (x *ListOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderRequest.ProtoReflect.Descriptor instead.
func (*ListOrderRequest) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_protofiles_order_proto_rawDescGZIP(), []int{2}
}

func (x *ListOrderRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListOrderRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ItemListOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price      float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax        float32 `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
	FinalPrice float32 `protobuf:"fixed32,4,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
}

func (x *ItemListOrderResponse) Reset() {
	*x = ItemListOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemListOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemListOrderResponse) ProtoMessage() {}

func (x *ItemListOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemListOrderResponse.ProtoReflect.Descriptor instead.
func (*ItemListOrderResponse) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_protofiles_order_proto_rawDescGZIP(), []int{3}
}

func (x *ItemListOrderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ItemListOrderResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ItemListOrderResponse) GetTax() float32 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *ItemListOrderResponse) GetFinalPrice() float32 {
	if x != nil {
		return x.FinalPrice
	}
	return 0
}

type ListOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemListOrderResponse []*ItemListOrderResponse `protobuf:"bytes,1,rep,name=itemListOrderResponse,proto3" json:"itemListOrderResponse,omitempty"`
}

func (x *ListOrderResponse) Reset() {
	*x = ListOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderResponse) ProtoMessage() {}

func (x *ListOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderResponse.ProtoReflect.Descriptor instead.
func (*ListOrderResponse) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_protofiles_order_proto_rawDescGZIP(), []int{4}
}

func (x *ListOrderResponse) GetItemListOrderResponse() []*ItemListOrderResponse {
	if x != nil {
		return x.ItemListOrderResponse
	}
	return nil
}

var File_internal_infra_grpc_protofiles_order_proto protoreflect.FileDescriptor

var file_internal_infra_grpc_protofiles_order_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x4c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x74, 0x61, 0x78, 0x22, 0x6e,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x74, 0x61, 0x78, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x42,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x22, 0x70, 0x0a, 0x15, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03,
	0x74, 0x61, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x64, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x15, 0x69, 0x74, 0x65,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x15, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x88, 0x01, 0x0a, 0x0c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_infra_grpc_protofiles_order_proto_rawDescOnce sync.Once
	file_internal_infra_grpc_protofiles_order_proto_rawDescData = file_internal_infra_grpc_protofiles_order_proto_rawDesc
)

func file_internal_infra_grpc_protofiles_order_proto_rawDescGZIP() []byte {
	file_internal_infra_grpc_protofiles_order_proto_rawDescOnce.Do(func() {
		file_internal_infra_grpc_protofiles_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_infra_grpc_protofiles_order_proto_rawDescData)
	})
	return file_internal_infra_grpc_protofiles_order_proto_rawDescData
}

var file_internal_infra_grpc_protofiles_order_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_infra_grpc_protofiles_order_proto_goTypes = []interface{}{
	(*CreateOrderRequest)(nil),    // 0: pb.CreateOrderRequest
	(*CreateOrderResponse)(nil),   // 1: pb.CreateOrderResponse
	(*ListOrderRequest)(nil),      // 2: pb.ListOrderRequest
	(*ItemListOrderResponse)(nil), // 3: pb.ItemListOrderResponse
	(*ListOrderResponse)(nil),     // 4: pb.ListOrderResponse
}
var file_internal_infra_grpc_protofiles_order_proto_depIdxs = []int32{
	3, // 0: pb.ListOrderResponse.itemListOrderResponse:type_name -> pb.ItemListOrderResponse
	0, // 1: pb.OrderService.CreateOrder:input_type -> pb.CreateOrderRequest
	2, // 2: pb.OrderService.ListOrder:input_type -> pb.ListOrderRequest
	1, // 3: pb.OrderService.CreateOrder:output_type -> pb.CreateOrderResponse
	4, // 4: pb.OrderService.ListOrder:output_type -> pb.ListOrderResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_infra_grpc_protofiles_order_proto_init() }
func file_internal_infra_grpc_protofiles_order_proto_init() {
	if File_internal_infra_grpc_protofiles_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_infra_grpc_protofiles_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_internal_infra_grpc_protofiles_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
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
		file_internal_infra_grpc_protofiles_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderRequest); i {
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
		file_internal_infra_grpc_protofiles_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemListOrderResponse); i {
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
		file_internal_infra_grpc_protofiles_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderResponse); i {
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
			RawDescriptor: file_internal_infra_grpc_protofiles_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_infra_grpc_protofiles_order_proto_goTypes,
		DependencyIndexes: file_internal_infra_grpc_protofiles_order_proto_depIdxs,
		MessageInfos:      file_internal_infra_grpc_protofiles_order_proto_msgTypes,
	}.Build()
	File_internal_infra_grpc_protofiles_order_proto = out.File
	file_internal_infra_grpc_protofiles_order_proto_rawDesc = nil
	file_internal_infra_grpc_protofiles_order_proto_goTypes = nil
	file_internal_infra_grpc_protofiles_order_proto_depIdxs = nil
}
