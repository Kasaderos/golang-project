// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: api/loms/v1/service.proto

package loms

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  int64            `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Items []*OrderInfoItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *OrderCreateRequest) Reset() {
	*x = OrderCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_loms_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateRequest) ProtoMessage() {}

func (x *OrderCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_loms_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateRequest.ProtoReflect.Descriptor instead.
func (*OrderCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_loms_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *OrderCreateRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *OrderCreateRequest) GetItems() []*OrderInfoItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type OrderInfoItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku   int64  `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"` // <= uint16
}

func (x *OrderInfoItem) Reset() {
	*x = OrderInfoItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_loms_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfoItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfoItem) ProtoMessage() {}

func (x *OrderInfoItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_loms_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfoItem.ProtoReflect.Descriptor instead.
func (*OrderInfoItem) Descriptor() ([]byte, []int) {
	return file_api_loms_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *OrderInfoItem) GetSku() int64 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *OrderInfoItem) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type OrderCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *OrderCreateResponse) Reset() {
	*x = OrderCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_loms_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateResponse) ProtoMessage() {}

func (x *OrderCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_loms_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateResponse.ProtoReflect.Descriptor instead.
func (*OrderCreateResponse) Descriptor() ([]byte, []int) {
	return file_api_loms_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *OrderCreateResponse) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type GetStockInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku uint32 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *GetStockInfoRequest) Reset() {
	*x = GetStockInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_loms_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStockInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStockInfoRequest) ProtoMessage() {}

func (x *GetStockInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_loms_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStockInfoRequest.ProtoReflect.Descriptor instead.
func (*GetStockInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_loms_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetStockInfoRequest) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

type GetStockInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetStockInfoResponse) Reset() {
	*x = GetStockInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_loms_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStockInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStockInfoResponse) ProtoMessage() {}

func (x *GetStockInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_loms_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStockInfoResponse.ProtoReflect.Descriptor instead.
func (*GetStockInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_loms_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetStockInfoResponse) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CreateOrderErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateOrderErrorResponse) Reset() {
	*x = CreateOrderErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_loms_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderErrorResponse) ProtoMessage() {}

func (x *CreateOrderErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_loms_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderErrorResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderErrorResponse) Descriptor() ([]byte, []int) {
	return file_api_loms_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateOrderErrorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetStockInfoErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetStockInfoErrorResponse) Reset() {
	*x = GetStockInfoErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_loms_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStockInfoErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStockInfoErrorResponse) ProtoMessage() {}

func (x *GetStockInfoErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_loms_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStockInfoErrorResponse.ProtoReflect.Descriptor instead.
func (*GetStockInfoErrorResponse) Descriptor() ([]byte, []int) {
	return file_api_loms_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetStockInfoErrorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CancelOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *CancelOrderRequest) Reset() {
	*x = CancelOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_loms_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderRequest) ProtoMessage() {}

func (x *CancelOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_loms_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderRequest.ProtoReflect.Descriptor instead.
func (*CancelOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_loms_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *CancelOrderRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type GetOrderInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *GetOrderInfoRequest) Reset() {
	*x = GetOrderInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_loms_v1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderInfoRequest) ProtoMessage() {}

func (x *GetOrderInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_loms_v1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderInfoRequest.ProtoReflect.Descriptor instead.
func (*GetOrderInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_loms_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetOrderInfoRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type GetOrderInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string           `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	User   int64            `protobuf:"varint,2,opt,name=user,proto3" json:"user,omitempty"`
	Items  []*OrderInfoItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetOrderInfoResponse) Reset() {
	*x = GetOrderInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_loms_v1_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderInfoResponse) ProtoMessage() {}

func (x *GetOrderInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_loms_v1_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderInfoResponse.ProtoReflect.Descriptor instead.
func (*GetOrderInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_loms_v1_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetOrderInfoResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetOrderInfoResponse) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *GetOrderInfoResponse) GetItems() []*OrderInfoItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type OrderPayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *OrderPayRequest) Reset() {
	*x = OrderPayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_loms_v1_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayRequest) ProtoMessage() {}

func (x *OrderPayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_loms_v1_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayRequest.ProtoReflect.Descriptor instead.
func (*OrderPayRequest) Descriptor() ([]byte, []int) {
	return file_api_loms_v1_service_proto_rawDescGZIP(), []int{10}
}

func (x *OrderPayRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

var File_api_loms_v1_service_proto protoreflect.FileDescriptor

var file_api_loms_v1_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x38, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x67, 0x6f, 0x5f,
	0x38, 0x5f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f,
	0x6d, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x12, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x5d, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x67,
	0x6f, 0x5f, 0x38, 0x5f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x37, 0x0a, 0x0d,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x30, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x22, 0x2c, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x34, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x35, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x38, 0x0a, 0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x39,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa1, 0x01, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x5d,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x67, 0x6f, 0x5f, 0x38, 0x5f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x35, 0x0a,
	0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x32, 0xf4, 0x06, 0x0a, 0x04, 0x4c, 0x4f, 0x4d, 0x53, 0x12, 0xc4, 0x01,
	0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x4c, 0x2e,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x67, 0x6f, 0x5f, 0x38, 0x5f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4d, 0x2e, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x67, 0x6f,
	0x5f, 0x38, 0x5f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c,
	0x6f, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x4c, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a,
	0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x67, 0x6f, 0x5f, 0x38, 0x5f, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x12, 0xc5, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4d, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f,
	0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x67, 0x6f, 0x5f, 0x38, 0x5f, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x4e, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a,
	0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x67, 0x6f, 0x5f, 0x38, 0x5f, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22,
	0x0b, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x84, 0x01, 0x0a,
	0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x12, 0x49, 0x2e, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x67, 0x6f, 0x5f, 0x38,
	0x5f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x6d,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f,
	0x70, 0x61, 0x79, 0x12, 0xc5, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4d, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a,
	0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x67, 0x6f, 0x5f, 0x38, 0x5f, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x4e, 0x2e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x67, 0x6f, 0x5f, 0x38, 0x5f, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b,
	0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x4b, 0x5a, 0x49, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6b,
	0x62, 0x2e, 0x6b, 0x61, 0x6c, 0x64, 0x61, 0x72, 0x6f, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x38, 0x2d,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6c,
	0x6f, 0x6d, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x6d, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_loms_v1_service_proto_rawDescOnce sync.Once
	file_api_loms_v1_service_proto_rawDescData = file_api_loms_v1_service_proto_rawDesc
)

func file_api_loms_v1_service_proto_rawDescGZIP() []byte {
	file_api_loms_v1_service_proto_rawDescOnce.Do(func() {
		file_api_loms_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_loms_v1_service_proto_rawDescData)
	})
	return file_api_loms_v1_service_proto_rawDescData
}

var file_api_loms_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_loms_v1_service_proto_goTypes = []interface{}{
	(*OrderCreateRequest)(nil),        // 0: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.OrderCreateRequest
	(*OrderInfoItem)(nil),             // 1: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.OrderInfoItem
	(*OrderCreateResponse)(nil),       // 2: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.OrderCreateResponse
	(*GetStockInfoRequest)(nil),       // 3: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.GetStockInfoRequest
	(*GetStockInfoResponse)(nil),      // 4: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.GetStockInfoResponse
	(*CreateOrderErrorResponse)(nil),  // 5: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.CreateOrderErrorResponse
	(*GetStockInfoErrorResponse)(nil), // 6: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.GetStockInfoErrorResponse
	(*CancelOrderRequest)(nil),        // 7: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.CancelOrderRequest
	(*GetOrderInfoRequest)(nil),       // 8: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.GetOrderInfoRequest
	(*GetOrderInfoResponse)(nil),      // 9: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.GetOrderInfoResponse
	(*OrderPayRequest)(nil),           // 10: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.OrderPayRequest
	(*emptypb.Empty)(nil),             // 11: google.protobuf.Empty
}
var file_api_loms_v1_service_proto_depIdxs = []int32{
	1,  // 0: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.OrderCreateRequest.items:type_name -> gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.OrderInfoItem
	1,  // 1: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.GetOrderInfoResponse.items:type_name -> gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.OrderInfoItem
	0,  // 2: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.LOMS.OrderCreate:input_type -> gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.OrderCreateRequest
	7,  // 3: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.LOMS.CancelOrder:input_type -> gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.CancelOrderRequest
	8,  // 4: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.LOMS.GetOrderInfo:input_type -> gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.GetOrderInfoRequest
	10, // 5: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.LOMS.OrderPay:input_type -> gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.OrderPayRequest
	3,  // 6: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.LOMS.GetStockInfo:input_type -> gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.GetStockInfoRequest
	2,  // 7: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.LOMS.OrderCreate:output_type -> gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.OrderCreateResponse
	11, // 8: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.LOMS.CancelOrder:output_type -> google.protobuf.Empty
	9,  // 9: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.LOMS.GetOrderInfo:output_type -> gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.GetOrderInfoResponse
	11, // 10: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.LOMS.OrderPay:output_type -> google.protobuf.Empty
	4,  // 11: gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.LOMS.GetStockInfo:output_type -> gitlab.ozon.dev.go_8_middle_project.loms.pkg.api.loms.v1.GetStockInfoResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_loms_v1_service_proto_init() }
func file_api_loms_v1_service_proto_init() {
	if File_api_loms_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_loms_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateRequest); i {
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
		file_api_loms_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderInfoItem); i {
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
		file_api_loms_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateResponse); i {
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
		file_api_loms_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStockInfoRequest); i {
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
		file_api_loms_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStockInfoResponse); i {
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
		file_api_loms_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderErrorResponse); i {
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
		file_api_loms_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStockInfoErrorResponse); i {
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
		file_api_loms_v1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOrderRequest); i {
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
		file_api_loms_v1_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderInfoRequest); i {
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
		file_api_loms_v1_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderInfoResponse); i {
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
		file_api_loms_v1_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayRequest); i {
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
			RawDescriptor: file_api_loms_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_loms_v1_service_proto_goTypes,
		DependencyIndexes: file_api_loms_v1_service_proto_depIdxs,
		MessageInfos:      file_api_loms_v1_service_proto_msgTypes,
	}.Build()
	File_api_loms_v1_service_proto = out.File
	file_api_loms_v1_service_proto_rawDesc = nil
	file_api_loms_v1_service_proto_goTypes = nil
	file_api_loms_v1_service_proto_depIdxs = nil
}
