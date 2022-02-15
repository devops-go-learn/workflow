// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api/apps/approval/pb/approval.proto

package approval

import (
	request "github.com/infraboard/mcube/http/request"
	request1 "github.com/infraboard/mcube/pb/request"
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

type Provider int32

const (
	Provider_DEVCLOUD Provider = 0
	Provider_FEISHU   Provider = 1
)

// Enum value maps for Provider.
var (
	Provider_name = map[int32]string{
		0: "DEVCLOUD",
		1: "FEISHU",
	}
	Provider_value = map[string]int32{
		"DEVCLOUD": 0,
		"FEISHU":   1,
	}
)

func (x Provider) Enum() *Provider {
	p := new(Provider)
	*p = x
	return p
}

func (x Provider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Provider) Descriptor() protoreflect.EnumDescriptor {
	return file_api_apps_approval_pb_approval_proto_enumTypes[0].Descriptor()
}

func (Provider) Type() protoreflect.EnumType {
	return &file_api_apps_approval_pb_approval_proto_enumTypes[0]
}

func (x Provider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Provider.Descriptor instead.
func (Provider) EnumDescriptor() ([]byte, []int) {
	return file_api_apps_approval_pb_approval_proto_rawDescGZIP(), []int{0}
}

type Approval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Approval) Reset() {
	*x = Approval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_apps_approval_pb_approval_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Approval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Approval) ProtoMessage() {}

func (x *Approval) ProtoReflect() protoreflect.Message {
	mi := &file_api_apps_approval_pb_approval_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Approval.ProtoReflect.Descriptor instead.
func (*Approval) Descriptor() ([]byte, []int) {
	return file_api_apps_approval_pb_approval_proto_rawDescGZIP(), []int{0}
}

type ApprovalSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页时，返回总数量
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// 一页的数据
	// @gotags: json:"items"
	Items []*Approval `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ApprovalSet) Reset() {
	*x = ApprovalSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_apps_approval_pb_approval_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalSet) ProtoMessage() {}

func (x *ApprovalSet) ProtoReflect() protoreflect.Message {
	mi := &file_api_apps_approval_pb_approval_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovalSet.ProtoReflect.Descriptor instead.
func (*ApprovalSet) Descriptor() ([]byte, []int) {
	return file_api_apps_approval_pb_approval_proto_rawDescGZIP(), []int{1}
}

func (x *ApprovalSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ApprovalSet) GetItems() []*Approval {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateApprovalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 工单对接的第三方系统
	// @gotags: json:"provider" bson:"provider"
	Provider Provider `protobuf:"varint,1,opt,name=provider,proto3,enum=infraboard.workorder.approval.Provider" json:"provider,omitempty"`
	// 工单模版编号, 用于对接
	// @gotags: json:"approval_code" bson:"approval_code"
	ApprovalCode string `protobuf:"bytes,2,opt,name=approval_code,json=approvalCode,proto3" json:"approval_code,omitempty"`
	// 工单状态
	// @gotags: json:"status" bson:"status"
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateApprovalRequest) Reset() {
	*x = CreateApprovalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_apps_approval_pb_approval_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApprovalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApprovalRequest) ProtoMessage() {}

func (x *CreateApprovalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_apps_approval_pb_approval_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApprovalRequest.ProtoReflect.Descriptor instead.
func (*CreateApprovalRequest) Descriptor() ([]byte, []int) {
	return file_api_apps_approval_pb_approval_proto_rawDescGZIP(), []int{2}
}

func (x *CreateApprovalRequest) GetProvider() Provider {
	if x != nil {
		return x.Provider
	}
	return Provider_DEVCLOUD
}

func (x *CreateApprovalRequest) GetApprovalCode() string {
	if x != nil {
		return x.ApprovalCode
	}
	return ""
}

func (x *CreateApprovalRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type QueryApprovalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *QueryApprovalRequest) Reset() {
	*x = QueryApprovalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_apps_approval_pb_approval_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryApprovalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryApprovalRequest) ProtoMessage() {}

func (x *QueryApprovalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_apps_approval_pb_approval_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryApprovalRequest.ProtoReflect.Descriptor instead.
func (*QueryApprovalRequest) Descriptor() ([]byte, []int) {
	return file_api_apps_approval_pb_approval_proto_rawDescGZIP(), []int{3}
}

func (x *QueryApprovalRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

type DescribeApprovalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DescribeApprovalRequest) Reset() {
	*x = DescribeApprovalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_apps_approval_pb_approval_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeApprovalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeApprovalRequest) ProtoMessage() {}

func (x *DescribeApprovalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_apps_approval_pb_approval_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeApprovalRequest.ProtoReflect.Descriptor instead.
func (*DescribeApprovalRequest) Descriptor() ([]byte, []int) {
	return file_api_apps_approval_pb_approval_proto_rawDescGZIP(), []int{4}
}

type UpdateApprovalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 更新模式
	// @gotags: json:"update_mode"
	UpdateMode request1.UpdateMode `protobuf:"varint,1,opt,name=update_mode,json=updateMode,proto3,enum=infraboard.mcube.request.UpdateMode" json:"update_mode,omitempty"`
}

func (x *UpdateApprovalRequest) Reset() {
	*x = UpdateApprovalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_apps_approval_pb_approval_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateApprovalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateApprovalRequest) ProtoMessage() {}

func (x *UpdateApprovalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_apps_approval_pb_approval_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateApprovalRequest.ProtoReflect.Descriptor instead.
func (*UpdateApprovalRequest) Descriptor() ([]byte, []int) {
	return file_api_apps_approval_pb_approval_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateApprovalRequest) GetUpdateMode() request1.UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return request1.UpdateMode(0)
}

type DeleteApprovalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteApprovalRequest) Reset() {
	*x = DeleteApprovalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_apps_approval_pb_approval_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteApprovalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApprovalRequest) ProtoMessage() {}

func (x *DeleteApprovalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_apps_approval_pb_approval_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApprovalRequest.ProtoReflect.Descriptor instead.
func (*DeleteApprovalRequest) Descriptor() ([]byte, []int) {
	return file_api_apps_approval_pb_approval_proto_rawDescGZIP(), []int{6}
}

var File_api_apps_approval_pb_approval_proto protoreflect.FileDescriptor

var file_api_apps_approval_pb_approval_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62,
	0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62,
	0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0a, 0x0a, 0x08, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x22, 0x62, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x3d, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x99, 0x01, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4e, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x5e, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75,
	0x62, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2a, 0x24, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x56, 0x43, 0x4c, 0x4f,
	0x55, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x49, 0x53, 0x48, 0x55, 0x10, 0x01,
	0x32, 0xaf, 0x04, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x34, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c,
	0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x6c, 0x0a, 0x09, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x33, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x12, 0x6f, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x36, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x6b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x34, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x6b, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x12, 0x34, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_apps_approval_pb_approval_proto_rawDescOnce sync.Once
	file_api_apps_approval_pb_approval_proto_rawDescData = file_api_apps_approval_pb_approval_proto_rawDesc
)

func file_api_apps_approval_pb_approval_proto_rawDescGZIP() []byte {
	file_api_apps_approval_pb_approval_proto_rawDescOnce.Do(func() {
		file_api_apps_approval_pb_approval_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_apps_approval_pb_approval_proto_rawDescData)
	})
	return file_api_apps_approval_pb_approval_proto_rawDescData
}

var file_api_apps_approval_pb_approval_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_apps_approval_pb_approval_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_apps_approval_pb_approval_proto_goTypes = []interface{}{
	(Provider)(0),                   // 0: infraboard.workorder.approval.Provider
	(*Approval)(nil),                // 1: infraboard.workorder.approval.Approval
	(*ApprovalSet)(nil),             // 2: infraboard.workorder.approval.ApprovalSet
	(*CreateApprovalRequest)(nil),   // 3: infraboard.workorder.approval.CreateApprovalRequest
	(*QueryApprovalRequest)(nil),    // 4: infraboard.workorder.approval.QueryApprovalRequest
	(*DescribeApprovalRequest)(nil), // 5: infraboard.workorder.approval.DescribeApprovalRequest
	(*UpdateApprovalRequest)(nil),   // 6: infraboard.workorder.approval.UpdateApprovalRequest
	(*DeleteApprovalRequest)(nil),   // 7: infraboard.workorder.approval.DeleteApprovalRequest
	(*request.PageRequest)(nil),     // 8: infraboard.mcube.page.PageRequest
	(request1.UpdateMode)(0),        // 9: infraboard.mcube.request.UpdateMode
}
var file_api_apps_approval_pb_approval_proto_depIdxs = []int32{
	1, // 0: infraboard.workorder.approval.ApprovalSet.items:type_name -> infraboard.workorder.approval.Approval
	0, // 1: infraboard.workorder.approval.CreateApprovalRequest.provider:type_name -> infraboard.workorder.approval.Provider
	8, // 2: infraboard.workorder.approval.QueryApprovalRequest.page:type_name -> infraboard.mcube.page.PageRequest
	9, // 3: infraboard.workorder.approval.UpdateApprovalRequest.update_mode:type_name -> infraboard.mcube.request.UpdateMode
	3, // 4: infraboard.workorder.approval.Service.CreateBook:input_type -> infraboard.workorder.approval.CreateApprovalRequest
	4, // 5: infraboard.workorder.approval.Service.QueryBook:input_type -> infraboard.workorder.approval.QueryApprovalRequest
	5, // 6: infraboard.workorder.approval.Service.DescribeBook:input_type -> infraboard.workorder.approval.DescribeApprovalRequest
	6, // 7: infraboard.workorder.approval.Service.UpdateBook:input_type -> infraboard.workorder.approval.UpdateApprovalRequest
	7, // 8: infraboard.workorder.approval.Service.DeleteBook:input_type -> infraboard.workorder.approval.DeleteApprovalRequest
	1, // 9: infraboard.workorder.approval.Service.CreateBook:output_type -> infraboard.workorder.approval.Approval
	2, // 10: infraboard.workorder.approval.Service.QueryBook:output_type -> infraboard.workorder.approval.ApprovalSet
	1, // 11: infraboard.workorder.approval.Service.DescribeBook:output_type -> infraboard.workorder.approval.Approval
	1, // 12: infraboard.workorder.approval.Service.UpdateBook:output_type -> infraboard.workorder.approval.Approval
	1, // 13: infraboard.workorder.approval.Service.DeleteBook:output_type -> infraboard.workorder.approval.Approval
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_apps_approval_pb_approval_proto_init() }
func file_api_apps_approval_pb_approval_proto_init() {
	if File_api_apps_approval_pb_approval_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_apps_approval_pb_approval_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Approval); i {
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
		file_api_apps_approval_pb_approval_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalSet); i {
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
		file_api_apps_approval_pb_approval_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApprovalRequest); i {
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
		file_api_apps_approval_pb_approval_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryApprovalRequest); i {
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
		file_api_apps_approval_pb_approval_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeApprovalRequest); i {
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
		file_api_apps_approval_pb_approval_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateApprovalRequest); i {
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
		file_api_apps_approval_pb_approval_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteApprovalRequest); i {
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
			RawDescriptor: file_api_apps_approval_pb_approval_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_apps_approval_pb_approval_proto_goTypes,
		DependencyIndexes: file_api_apps_approval_pb_approval_proto_depIdxs,
		EnumInfos:         file_api_apps_approval_pb_approval_proto_enumTypes,
		MessageInfos:      file_api_apps_approval_pb_approval_proto_msgTypes,
	}.Build()
	File_api_apps_approval_pb_approval_proto = out.File
	file_api_apps_approval_pb_approval_proto_rawDesc = nil
	file_api_apps_approval_pb_approval_proto_goTypes = nil
	file_api_apps_approval_pb_approval_proto_depIdxs = nil
}
