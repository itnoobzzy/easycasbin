// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: role.proto

package proto

import (
	context "context"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddDomainRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainName string `protobuf:"bytes,1,opt,name=domainName,proto3" json:"domainName,omitempty"`
	RoleName   string `protobuf:"bytes,2,opt,name=roleName,proto3" json:"roleName,omitempty"`
}

func (x *AddDomainRoleReq) Reset() {
	*x = AddDomainRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDomainRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDomainRoleReq) ProtoMessage() {}

func (x *AddDomainRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDomainRoleReq.ProtoReflect.Descriptor instead.
func (*AddDomainRoleReq) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{0}
}

func (x *AddDomainRoleReq) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

func (x *AddDomainRoleReq) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type AddDomainRoleRpl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32                 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *AddDomainRoleRpl_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AddDomainRoleRpl) Reset() {
	*x = AddDomainRoleRpl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDomainRoleRpl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDomainRoleRpl) ProtoMessage() {}

func (x *AddDomainRoleRpl) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDomainRoleRpl.ProtoReflect.Descriptor instead.
func (*AddDomainRoleRpl) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{1}
}

func (x *AddDomainRoleRpl) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AddDomainRoleRpl) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *AddDomainRoleRpl) GetData() *AddDomainRoleRpl_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateDomainRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainName  string `protobuf:"bytes,1,opt,name=domainName,proto3" json:"domainName,omitempty"`
	RoleName    string `protobuf:"bytes,2,opt,name=roleName,proto3" json:"roleName,omitempty"`
	NewRoleName string `protobuf:"bytes,3,opt,name=newRoleName,proto3" json:"newRoleName,omitempty"`
}

func (x *UpdateDomainRoleReq) Reset() {
	*x = UpdateDomainRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDomainRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDomainRoleReq) ProtoMessage() {}

func (x *UpdateDomainRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDomainRoleReq.ProtoReflect.Descriptor instead.
func (*UpdateDomainRoleReq) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateDomainRoleReq) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

func (x *UpdateDomainRoleReq) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *UpdateDomainRoleReq) GetNewRoleName() string {
	if x != nil {
		return x.NewRoleName
	}
	return ""
}

type UpdateDomainRoleRpl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32                    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *UpdateDomainRoleRpl_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateDomainRoleRpl) Reset() {
	*x = UpdateDomainRoleRpl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDomainRoleRpl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDomainRoleRpl) ProtoMessage() {}

func (x *UpdateDomainRoleRpl) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDomainRoleRpl.ProtoReflect.Descriptor instead.
func (*UpdateDomainRoleRpl) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDomainRoleRpl) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UpdateDomainRoleRpl) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *UpdateDomainRoleRpl) GetData() *UpdateDomainRoleRpl_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteDomainRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainName string `protobuf:"bytes,1,opt,name=domainName,proto3" json:"domainName,omitempty"`
	RoleName   string `protobuf:"bytes,2,opt,name=roleName,proto3" json:"roleName,omitempty"`
}

func (x *DeleteDomainRoleReq) Reset() {
	*x = DeleteDomainRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDomainRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDomainRoleReq) ProtoMessage() {}

func (x *DeleteDomainRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDomainRoleReq.ProtoReflect.Descriptor instead.
func (*DeleteDomainRoleReq) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteDomainRoleReq) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

func (x *DeleteDomainRoleReq) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type DeleteDomainRoleRpl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32                    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *DeleteDomainRoleRpl_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DeleteDomainRoleRpl) Reset() {
	*x = DeleteDomainRoleRpl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDomainRoleRpl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDomainRoleRpl) ProtoMessage() {}

func (x *DeleteDomainRoleRpl) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDomainRoleRpl.ProtoReflect.Descriptor instead.
func (*DeleteDomainRoleRpl) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteDomainRoleRpl) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DeleteDomainRoleRpl) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *DeleteDomainRoleRpl) GetData() *DeleteDomainRoleRpl_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddDomainRoleRpl_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=creatTime,proto3" json:"creatTime,omitempty"`
	Name      string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Domain    string                 `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *AddDomainRoleRpl_Data) Reset() {
	*x = AddDomainRoleRpl_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDomainRoleRpl_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDomainRoleRpl_Data) ProtoMessage() {}

func (x *AddDomainRoleRpl_Data) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDomainRoleRpl_Data.ProtoReflect.Descriptor instead.
func (*AddDomainRoleRpl_Data) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AddDomainRoleRpl_Data) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddDomainRoleRpl_Data) GetCreatTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatTime
	}
	return nil
}

func (x *AddDomainRoleRpl_Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddDomainRoleRpl_Data) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type UpdateDomainRoleRpl_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	Name       string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Domain     string                 `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *UpdateDomainRoleRpl_Data) Reset() {
	*x = UpdateDomainRoleRpl_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDomainRoleRpl_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDomainRoleRpl_Data) ProtoMessage() {}

func (x *UpdateDomainRoleRpl_Data) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDomainRoleRpl_Data.ProtoReflect.Descriptor instead.
func (*UpdateDomainRoleRpl_Data) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{3, 0}
}

func (x *UpdateDomainRoleRpl_Data) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateDomainRoleRpl_Data) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *UpdateDomainRoleRpl_Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateDomainRoleRpl_Data) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type DeleteDomainRoleRpl_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=deleteTime,proto3" json:"deleteTime,omitempty"`
}

func (x *DeleteDomainRoleRpl_Data) Reset() {
	*x = DeleteDomainRoleRpl_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDomainRoleRpl_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDomainRoleRpl_Data) ProtoMessage() {}

func (x *DeleteDomainRoleRpl_Data) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDomainRoleRpl_Data.ProtoReflect.Descriptor instead.
func (*DeleteDomainRoleRpl_Data) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{5, 0}
}

func (x *DeleteDomainRoleRpl_Data) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

var File_role_proto protoreflect.FileDescriptor

var file_role_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01,
	0x0a, 0x10, 0x41, 0x64, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xfa, 0x42, 0x16, 0x72, 0x14, 0x32, 0x12, 0x5e,
	0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x7d,
	0x24, 0x52, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a,
	0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x19, 0xfa, 0x42, 0x16, 0x72, 0x14, 0x32, 0x12, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x7a,
	0x41, 0x2d, 0x5a, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x7d, 0x24, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xe2, 0x01, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x70, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x41, 0x64, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x70, 0x6c,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x7c, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xc4, 0x01, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xfa, 0x42, 0x16, 0x72, 0x14, 0x32, 0x12, 0x5e, 0x5b,
	0x30, 0x2d, 0x39, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x7d, 0x24,
	0x52, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08,
	0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19,
	0xfa, 0x42, 0x16, 0x72, 0x14, 0x32, 0x12, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x7a, 0x41,
	0x2d, 0x5a, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x7d, 0x24, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xfa, 0x42, 0x16, 0x72, 0x14, 0x32,
	0x12, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5f, 0x5d, 0x7b, 0x31,
	0x2c, 0x7d, 0x24, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xea, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x70, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2d,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x70, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x7e, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x87, 0x01,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xfa, 0x42, 0x16, 0x72, 0x14,
	0x32, 0x12, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5f, 0x5d, 0x7b,
	0x31, 0x2c, 0x7d, 0x24, 0x52, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x35, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x19, 0xfa, 0x42, 0x16, 0x72, 0x14, 0x32, 0x12, 0x5e, 0x5b, 0x30, 0x2d, 0x39,
	0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x7d, 0x24, 0x52, 0x08, 0x72,
	0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xae, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x70, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x70, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x42, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xbb, 0x01, 0x0a, 0x0a, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x41, 0x64,
	0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x70, 0x6c, 0x12, 0x3c,
	0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x14, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x70, 0x6c, 0x12, 0x38, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x14, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x70, 0x6c, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x2e, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_proto_rawDescOnce sync.Once
	file_role_proto_rawDescData = file_role_proto_rawDesc
)

func file_role_proto_rawDescGZIP() []byte {
	file_role_proto_rawDescOnce.Do(func() {
		file_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_proto_rawDescData)
	})
	return file_role_proto_rawDescData
}

var file_role_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_role_proto_goTypes = []interface{}{
	(*AddDomainRoleReq)(nil),         // 0: AddDomainRoleReq
	(*AddDomainRoleRpl)(nil),         // 1: AddDomainRoleRpl
	(*UpdateDomainRoleReq)(nil),      // 2: UpdateDomainRoleReq
	(*UpdateDomainRoleRpl)(nil),      // 3: UpdateDomainRoleRpl
	(*DeleteDomainRoleReq)(nil),      // 4: DeleteDomainRoleReq
	(*DeleteDomainRoleRpl)(nil),      // 5: DeleteDomainRoleRpl
	(*AddDomainRoleRpl_Data)(nil),    // 6: AddDomainRoleRpl.Data
	(*UpdateDomainRoleRpl_Data)(nil), // 7: UpdateDomainRoleRpl.Data
	(*DeleteDomainRoleRpl_Data)(nil), // 8: DeleteDomainRoleRpl.Data
	(*timestamppb.Timestamp)(nil),    // 9: google.protobuf.Timestamp
}
var file_role_proto_depIdxs = []int32{
	6, // 0: AddDomainRoleRpl.data:type_name -> AddDomainRoleRpl.Data
	7, // 1: UpdateDomainRoleRpl.data:type_name -> UpdateDomainRoleRpl.Data
	8, // 2: DeleteDomainRoleRpl.data:type_name -> DeleteDomainRoleRpl.Data
	9, // 3: AddDomainRoleRpl.Data.creatTime:type_name -> google.protobuf.Timestamp
	9, // 4: UpdateDomainRoleRpl.Data.updateTime:type_name -> google.protobuf.Timestamp
	9, // 5: DeleteDomainRoleRpl.Data.deleteTime:type_name -> google.protobuf.Timestamp
	0, // 6: DomainRole.AddDomainRole:input_type -> AddDomainRoleReq
	2, // 7: DomainRole.UpdateRoleInfo:input_type -> UpdateDomainRoleReq
	4, // 8: DomainRole.DeleteRole:input_type -> DeleteDomainRoleReq
	1, // 9: DomainRole.AddDomainRole:output_type -> AddDomainRoleRpl
	3, // 10: DomainRole.UpdateRoleInfo:output_type -> UpdateDomainRoleRpl
	5, // 11: DomainRole.DeleteRole:output_type -> DeleteDomainRoleRpl
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_role_proto_init() }
func file_role_proto_init() {
	if File_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDomainRoleReq); i {
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
		file_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDomainRoleRpl); i {
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
		file_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDomainRoleReq); i {
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
		file_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDomainRoleRpl); i {
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
		file_role_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDomainRoleReq); i {
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
		file_role_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDomainRoleRpl); i {
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
		file_role_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDomainRoleRpl_Data); i {
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
		file_role_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDomainRoleRpl_Data); i {
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
		file_role_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDomainRoleRpl_Data); i {
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
			RawDescriptor: file_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_proto_goTypes,
		DependencyIndexes: file_role_proto_depIdxs,
		MessageInfos:      file_role_proto_msgTypes,
	}.Build()
	File_role_proto = out.File
	file_role_proto_rawDesc = nil
	file_role_proto_goTypes = nil
	file_role_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DomainRoleClient is the client API for DomainRole service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DomainRoleClient interface {
	// 添加域角色
	AddDomainRole(ctx context.Context, in *AddDomainRoleReq, opts ...grpc.CallOption) (*AddDomainRoleRpl, error)
	// 更新域角色信息
	UpdateRoleInfo(ctx context.Context, in *UpdateDomainRoleReq, opts ...grpc.CallOption) (*UpdateDomainRoleRpl, error)
	// DeleteRole 删除对应域的角色
	DeleteRole(ctx context.Context, in *DeleteDomainRoleReq, opts ...grpc.CallOption) (*DeleteDomainRoleRpl, error)
}

type domainRoleClient struct {
	cc grpc.ClientConnInterface
}

func NewDomainRoleClient(cc grpc.ClientConnInterface) DomainRoleClient {
	return &domainRoleClient{cc}
}

func (c *domainRoleClient) AddDomainRole(ctx context.Context, in *AddDomainRoleReq, opts ...grpc.CallOption) (*AddDomainRoleRpl, error) {
	out := new(AddDomainRoleRpl)
	err := c.cc.Invoke(ctx, "/DomainRole/AddDomainRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainRoleClient) UpdateRoleInfo(ctx context.Context, in *UpdateDomainRoleReq, opts ...grpc.CallOption) (*UpdateDomainRoleRpl, error) {
	out := new(UpdateDomainRoleRpl)
	err := c.cc.Invoke(ctx, "/DomainRole/UpdateRoleInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainRoleClient) DeleteRole(ctx context.Context, in *DeleteDomainRoleReq, opts ...grpc.CallOption) (*DeleteDomainRoleRpl, error) {
	out := new(DeleteDomainRoleRpl)
	err := c.cc.Invoke(ctx, "/DomainRole/DeleteRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DomainRoleServer is the server API for DomainRole service.
type DomainRoleServer interface {
	// 添加域角色
	AddDomainRole(context.Context, *AddDomainRoleReq) (*AddDomainRoleRpl, error)
	// 更新域角色信息
	UpdateRoleInfo(context.Context, *UpdateDomainRoleReq) (*UpdateDomainRoleRpl, error)
	// DeleteRole 删除对应域的角色
	DeleteRole(context.Context, *DeleteDomainRoleReq) (*DeleteDomainRoleRpl, error)
}

// UnimplementedDomainRoleServer can be embedded to have forward compatible implementations.
type UnimplementedDomainRoleServer struct {
}

func (*UnimplementedDomainRoleServer) AddDomainRole(context.Context, *AddDomainRoleReq) (*AddDomainRoleRpl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDomainRole not implemented")
}
func (*UnimplementedDomainRoleServer) UpdateRoleInfo(context.Context, *UpdateDomainRoleReq) (*UpdateDomainRoleRpl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoleInfo not implemented")
}
func (*UnimplementedDomainRoleServer) DeleteRole(context.Context, *DeleteDomainRoleReq) (*DeleteDomainRoleRpl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}

func RegisterDomainRoleServer(s *grpc.Server, srv DomainRoleServer) {
	s.RegisterService(&_DomainRole_serviceDesc, srv)
}

func _DomainRole_AddDomainRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDomainRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainRoleServer).AddDomainRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DomainRole/AddDomainRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainRoleServer).AddDomainRole(ctx, req.(*AddDomainRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainRole_UpdateRoleInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDomainRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainRoleServer).UpdateRoleInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DomainRole/UpdateRoleInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainRoleServer).UpdateRoleInfo(ctx, req.(*UpdateDomainRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainRole_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDomainRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainRoleServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DomainRole/DeleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainRoleServer).DeleteRole(ctx, req.(*DeleteDomainRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _DomainRole_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DomainRole",
	HandlerType: (*DomainRoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDomainRole",
			Handler:    _DomainRole_AddDomainRole_Handler,
		},
		{
			MethodName: "UpdateRoleInfo",
			Handler:    _DomainRole_UpdateRoleInfo_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _DomainRole_DeleteRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role.proto",
}
