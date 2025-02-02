// Copyright 2024 TII (SSRC) and the Ghaf contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.1
// source: systemd.proto

package systemd

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

type UnitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UnitName      string                 `protobuf:"bytes,1,opt,name=UnitName,proto3" json:"UnitName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnitRequest) Reset() {
	*x = UnitRequest{}
	mi := &file_systemd_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitRequest) ProtoMessage() {}

func (x *UnitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_systemd_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitRequest.ProtoReflect.Descriptor instead.
func (*UnitRequest) Descriptor() ([]byte, []int) {
	return file_systemd_proto_rawDescGZIP(), []int{0}
}

func (x *UnitRequest) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

type UnitResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CmdStatus     string                 `protobuf:"bytes,1,opt,name=CmdStatus,proto3" json:"CmdStatus,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnitResponse) Reset() {
	*x = UnitResponse{}
	mi := &file_systemd_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitResponse) ProtoMessage() {}

func (x *UnitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_systemd_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitResponse.ProtoReflect.Descriptor instead.
func (*UnitResponse) Descriptor() ([]byte, []int) {
	return file_systemd_proto_rawDescGZIP(), []int{1}
}

func (x *UnitResponse) GetCmdStatus() string {
	if x != nil {
		return x.CmdStatus
	}
	return ""
}

type AppUnitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UnitName      string                 `protobuf:"bytes,1,opt,name=UnitName,proto3" json:"UnitName,omitempty"`
	Args          []string               `protobuf:"bytes,2,rep,name=Args,proto3" json:"Args,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AppUnitRequest) Reset() {
	*x = AppUnitRequest{}
	mi := &file_systemd_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppUnitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppUnitRequest) ProtoMessage() {}

func (x *AppUnitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_systemd_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppUnitRequest.ProtoReflect.Descriptor instead.
func (*AppUnitRequest) Descriptor() ([]byte, []int) {
	return file_systemd_proto_rawDescGZIP(), []int{2}
}

func (x *AppUnitRequest) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

func (x *AppUnitRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

type UnitStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	LoadState     string                 `protobuf:"bytes,3,opt,name=LoadState,proto3" json:"LoadState,omitempty"`
	ActiveState   string                 `protobuf:"bytes,4,opt,name=ActiveState,proto3" json:"ActiveState,omitempty"`
	SubState      string                 `protobuf:"bytes,5,opt,name=SubState,proto3" json:"SubState,omitempty"`
	Path          string                 `protobuf:"bytes,6,opt,name=Path,proto3" json:"Path,omitempty"`
	FreezerState  string                 `protobuf:"bytes,7,opt,name=FreezerState,proto3" json:"FreezerState,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnitStatus) Reset() {
	*x = UnitStatus{}
	mi := &file_systemd_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnitStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitStatus) ProtoMessage() {}

func (x *UnitStatus) ProtoReflect() protoreflect.Message {
	mi := &file_systemd_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitStatus.ProtoReflect.Descriptor instead.
func (*UnitStatus) Descriptor() ([]byte, []int) {
	return file_systemd_proto_rawDescGZIP(), []int{3}
}

func (x *UnitStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UnitStatus) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UnitStatus) GetLoadState() string {
	if x != nil {
		return x.LoadState
	}
	return ""
}

func (x *UnitStatus) GetActiveState() string {
	if x != nil {
		return x.ActiveState
	}
	return ""
}

func (x *UnitStatus) GetSubState() string {
	if x != nil {
		return x.SubState
	}
	return ""
}

func (x *UnitStatus) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UnitStatus) GetFreezerState() string {
	if x != nil {
		return x.FreezerState
	}
	return ""
}

type UnitStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CmdStatus     string                 `protobuf:"bytes,1,opt,name=CmdStatus,proto3" json:"CmdStatus,omitempty"`
	UnitStatus    *UnitStatus            `protobuf:"bytes,2,opt,name=UnitStatus,proto3" json:"UnitStatus,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnitStatusResponse) Reset() {
	*x = UnitStatusResponse{}
	mi := &file_systemd_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnitStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitStatusResponse) ProtoMessage() {}

func (x *UnitStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_systemd_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitStatusResponse.ProtoReflect.Descriptor instead.
func (*UnitStatusResponse) Descriptor() ([]byte, []int) {
	return file_systemd_proto_rawDescGZIP(), []int{4}
}

func (x *UnitStatusResponse) GetCmdStatus() string {
	if x != nil {
		return x.CmdStatus
	}
	return ""
}

func (x *UnitStatusResponse) GetUnitStatus() *UnitStatus {
	if x != nil {
		return x.UnitStatus
	}
	return nil
}

type UnitResourceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UnitName      string                 `protobuf:"bytes,1,opt,name=UnitName,proto3" json:"UnitName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnitResourceRequest) Reset() {
	*x = UnitResourceRequest{}
	mi := &file_systemd_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnitResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitResourceRequest) ProtoMessage() {}

func (x *UnitResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_systemd_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitResourceRequest.ProtoReflect.Descriptor instead.
func (*UnitResourceRequest) Descriptor() ([]byte, []int) {
	return file_systemd_proto_rawDescGZIP(), []int{5}
}

func (x *UnitResourceRequest) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

type UnitResourceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CpuUsage      float64                `protobuf:"fixed64,1,opt,name=CpuUsage,proto3" json:"CpuUsage,omitempty"`
	MemoryUsage   float32                `protobuf:"fixed32,2,opt,name=MemoryUsage,proto3" json:"MemoryUsage,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnitResourceResponse) Reset() {
	*x = UnitResourceResponse{}
	mi := &file_systemd_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnitResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitResourceResponse) ProtoMessage() {}

func (x *UnitResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_systemd_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitResourceResponse.ProtoReflect.Descriptor instead.
func (*UnitResourceResponse) Descriptor() ([]byte, []int) {
	return file_systemd_proto_rawDescGZIP(), []int{6}
}

func (x *UnitResourceResponse) GetCpuUsage() float64 {
	if x != nil {
		return x.CpuUsage
	}
	return 0
}

func (x *UnitResourceResponse) GetMemoryUsage() float32 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

var File_systemd_proto protoreflect.FileDescriptor

var file_systemd_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x22, 0x29, 0x0a, 0x0b, 0x55, 0x6e, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x6e, 0x69, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x6e, 0x69, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x0c, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6d, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6d, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x40, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x6e, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x6e, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x41, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x41,
	0x72, 0x67, 0x73, 0x22, 0xd6, 0x01, 0x0a, 0x0a, 0x55, 0x6e, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x61, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4c, 0x6f, 0x61,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x75, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x46, 0x72, 0x65, 0x65,
	0x7a, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x46, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x67, 0x0a, 0x12,
	0x55, 0x6e, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6d, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6d, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x33, 0x0a, 0x0a, 0x55, 0x6e, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x2e, 0x55,
	0x6e, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x55, 0x6e, 0x69, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x31, 0x0a, 0x13, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x6e, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x6e, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x54, 0x0a, 0x14, 0x55, 0x6e, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x43, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x43, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9e,
	0x04, 0x0a, 0x12, 0x55, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x69, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64,
	0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x09, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x64, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x08, 0x53, 0x74, 0x6f, 0x70, 0x55,
	0x6e, 0x69, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x2e, 0x55, 0x6e,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x64, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x39, 0x0a, 0x08, 0x4b, 0x69, 0x6c, 0x6c, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x14,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x2e, 0x55,
	0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a,
	0x0a, 0x46, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x64, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x2e, 0x55, 0x6e, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x55, 0x6e,
	0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x64, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0b, 0x4d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x64, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64,
	0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x10, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x2e, 0x41, 0x70, 0x70, 0x55, 0x6e, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64,
	0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_systemd_proto_rawDescOnce sync.Once
	file_systemd_proto_rawDescData = file_systemd_proto_rawDesc
)

func file_systemd_proto_rawDescGZIP() []byte {
	file_systemd_proto_rawDescOnce.Do(func() {
		file_systemd_proto_rawDescData = protoimpl.X.CompressGZIP(file_systemd_proto_rawDescData)
	})
	return file_systemd_proto_rawDescData
}

var file_systemd_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_systemd_proto_goTypes = []any{
	(*UnitRequest)(nil),          // 0: systemd.UnitRequest
	(*UnitResponse)(nil),         // 1: systemd.UnitResponse
	(*AppUnitRequest)(nil),       // 2: systemd.AppUnitRequest
	(*UnitStatus)(nil),           // 3: systemd.UnitStatus
	(*UnitStatusResponse)(nil),   // 4: systemd.UnitStatusResponse
	(*UnitResourceRequest)(nil),  // 5: systemd.UnitResourceRequest
	(*UnitResourceResponse)(nil), // 6: systemd.UnitResourceResponse
}
var file_systemd_proto_depIdxs = []int32{
	3, // 0: systemd.UnitStatusResponse.UnitStatus:type_name -> systemd.UnitStatus
	0, // 1: systemd.UnitControlService.GetUnitStatus:input_type -> systemd.UnitRequest
	0, // 2: systemd.UnitControlService.StartUnit:input_type -> systemd.UnitRequest
	0, // 3: systemd.UnitControlService.StopUnit:input_type -> systemd.UnitRequest
	0, // 4: systemd.UnitControlService.KillUnit:input_type -> systemd.UnitRequest
	0, // 5: systemd.UnitControlService.FreezeUnit:input_type -> systemd.UnitRequest
	0, // 6: systemd.UnitControlService.UnfreezeUnit:input_type -> systemd.UnitRequest
	5, // 7: systemd.UnitControlService.MonitorUnit:input_type -> systemd.UnitResourceRequest
	2, // 8: systemd.UnitControlService.StartApplication:input_type -> systemd.AppUnitRequest
	4, // 9: systemd.UnitControlService.GetUnitStatus:output_type -> systemd.UnitStatusResponse
	1, // 10: systemd.UnitControlService.StartUnit:output_type -> systemd.UnitResponse
	1, // 11: systemd.UnitControlService.StopUnit:output_type -> systemd.UnitResponse
	1, // 12: systemd.UnitControlService.KillUnit:output_type -> systemd.UnitResponse
	1, // 13: systemd.UnitControlService.FreezeUnit:output_type -> systemd.UnitResponse
	1, // 14: systemd.UnitControlService.UnfreezeUnit:output_type -> systemd.UnitResponse
	6, // 15: systemd.UnitControlService.MonitorUnit:output_type -> systemd.UnitResourceResponse
	1, // 16: systemd.UnitControlService.StartApplication:output_type -> systemd.UnitResponse
	9, // [9:17] is the sub-list for method output_type
	1, // [1:9] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_systemd_proto_init() }
func file_systemd_proto_init() {
	if File_systemd_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_systemd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_systemd_proto_goTypes,
		DependencyIndexes: file_systemd_proto_depIdxs,
		MessageInfos:      file_systemd_proto_msgTypes,
	}.Build()
	File_systemd_proto = out.File
	file_systemd_proto_rawDesc = nil
	file_systemd_proto_goTypes = nil
	file_systemd_proto_depIdxs = nil
}
