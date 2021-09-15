// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: authzed/api/v0/namespace_service.proto

package v0

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ReadConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace  string  `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	AtRevision *Zookie `protobuf:"bytes,2,opt,name=at_revision,json=atRevision,proto3" json:"at_revision,omitempty"`
}

func (x *ReadConfigRequest) Reset() {
	*x = ReadConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v0_namespace_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadConfigRequest) ProtoMessage() {}

func (x *ReadConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v0_namespace_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadConfigRequest.ProtoReflect.Descriptor instead.
func (*ReadConfigRequest) Descriptor() ([]byte, []int) {
	return file_authzed_api_v0_namespace_service_proto_rawDescGZIP(), []int{0}
}

func (x *ReadConfigRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ReadConfigRequest) GetAtRevision() *Zookie {
	if x != nil {
		return x.AtRevision
	}
	return nil
}

type ReadConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string               `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Config    *NamespaceDefinition `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	Revision  *Zookie              `protobuf:"bytes,4,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *ReadConfigResponse) Reset() {
	*x = ReadConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v0_namespace_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadConfigResponse) ProtoMessage() {}

func (x *ReadConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v0_namespace_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadConfigResponse.ProtoReflect.Descriptor instead.
func (*ReadConfigResponse) Descriptor() ([]byte, []int) {
	return file_authzed_api_v0_namespace_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReadConfigResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ReadConfigResponse) GetConfig() *NamespaceDefinition {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ReadConfigResponse) GetRevision() *Zookie {
	if x != nil {
		return x.Revision
	}
	return nil
}

type WriteConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configs []*NamespaceDefinition `protobuf:"bytes,2,rep,name=configs,proto3" json:"configs,omitempty"`
}

func (x *WriteConfigRequest) Reset() {
	*x = WriteConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v0_namespace_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteConfigRequest) ProtoMessage() {}

func (x *WriteConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v0_namespace_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteConfigRequest.ProtoReflect.Descriptor instead.
func (*WriteConfigRequest) Descriptor() ([]byte, []int) {
	return file_authzed_api_v0_namespace_service_proto_rawDescGZIP(), []int{2}
}

func (x *WriteConfigRequest) GetConfigs() []*NamespaceDefinition {
	if x != nil {
		return x.Configs
	}
	return nil
}

type WriteConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Revision *Zookie `protobuf:"bytes,1,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *WriteConfigResponse) Reset() {
	*x = WriteConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v0_namespace_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteConfigResponse) ProtoMessage() {}

func (x *WriteConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v0_namespace_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteConfigResponse.ProtoReflect.Descriptor instead.
func (*WriteConfigResponse) Descriptor() ([]byte, []int) {
	return file_authzed_api_v0_namespace_service_proto_rawDescGZIP(), []int{3}
}

func (x *WriteConfigResponse) GetRevision() *Zookie {
	if x != nil {
		return x.Revision
	}
	return nil
}

var File_authzed_api_v0_namespace_service_proto protoreflect.FileDescriptor

var file_authzed_api_v0_namespace_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30,
	0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x30, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a,
	0x11, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x66, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x48, 0xfa, 0x42, 0x45, 0x72, 0x43, 0x28, 0x80, 0x01, 0x32,
	0x3e, 0x5e, 0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f,
	0x5d, 0x7b, 0x32, 0x2c, 0x36, 0x32, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2f,
	0x29, 0x3f, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d,
	0x7b, 0x32, 0x2c, 0x36, 0x32, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x61, 0x74,
	0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30,
	0x2e, 0x5a, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0xa3, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x5a, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52,
	0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x64, 0x0a, 0x12, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x4e, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x30, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x92, 0x01, 0x09, 0x08, 0x01, 0x22,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22,
	0x49, 0x0a, 0x13, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x5a, 0x6f, 0x6f, 0x6b, 0x69, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0xc3, 0x01, 0x0a, 0x10, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x55, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x30, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0xa4, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x42, 0x15, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x65, 0x64, 0x2f, 0x73, 0x70, 0x69, 0x63, 0x65, 0x64, 0x62, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0xa2, 0x02, 0x03, 0x41, 0x41, 0x56,
	0xaa, 0x02, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56,
	0x30, 0xca, 0x02, 0x10, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x5c, 0x5c, 0x41, 0x70, 0x69,
	0x5c, 0x5c, 0x56, 0x30, 0xea, 0x02, 0x10, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x3a, 0x3a,
	0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authzed_api_v0_namespace_service_proto_rawDescOnce sync.Once
	file_authzed_api_v0_namespace_service_proto_rawDescData = file_authzed_api_v0_namespace_service_proto_rawDesc
)

func file_authzed_api_v0_namespace_service_proto_rawDescGZIP() []byte {
	file_authzed_api_v0_namespace_service_proto_rawDescOnce.Do(func() {
		file_authzed_api_v0_namespace_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_authzed_api_v0_namespace_service_proto_rawDescData)
	})
	return file_authzed_api_v0_namespace_service_proto_rawDescData
}

var file_authzed_api_v0_namespace_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_authzed_api_v0_namespace_service_proto_goTypes = []interface{}{
	(*ReadConfigRequest)(nil),   // 0: authzed.api.v0.ReadConfigRequest
	(*ReadConfigResponse)(nil),  // 1: authzed.api.v0.ReadConfigResponse
	(*WriteConfigRequest)(nil),  // 2: authzed.api.v0.WriteConfigRequest
	(*WriteConfigResponse)(nil), // 3: authzed.api.v0.WriteConfigResponse
	(*Zookie)(nil),              // 4: authzed.api.v0.Zookie
	(*NamespaceDefinition)(nil), // 5: authzed.api.v0.NamespaceDefinition
}
var file_authzed_api_v0_namespace_service_proto_depIdxs = []int32{
	4, // 0: authzed.api.v0.ReadConfigRequest.at_revision:type_name -> authzed.api.v0.Zookie
	5, // 1: authzed.api.v0.ReadConfigResponse.config:type_name -> authzed.api.v0.NamespaceDefinition
	4, // 2: authzed.api.v0.ReadConfigResponse.revision:type_name -> authzed.api.v0.Zookie
	5, // 3: authzed.api.v0.WriteConfigRequest.configs:type_name -> authzed.api.v0.NamespaceDefinition
	4, // 4: authzed.api.v0.WriteConfigResponse.revision:type_name -> authzed.api.v0.Zookie
	0, // 5: authzed.api.v0.NamespaceService.ReadConfig:input_type -> authzed.api.v0.ReadConfigRequest
	2, // 6: authzed.api.v0.NamespaceService.WriteConfig:input_type -> authzed.api.v0.WriteConfigRequest
	1, // 7: authzed.api.v0.NamespaceService.ReadConfig:output_type -> authzed.api.v0.ReadConfigResponse
	3, // 8: authzed.api.v0.NamespaceService.WriteConfig:output_type -> authzed.api.v0.WriteConfigResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_authzed_api_v0_namespace_service_proto_init() }
func file_authzed_api_v0_namespace_service_proto_init() {
	if File_authzed_api_v0_namespace_service_proto != nil {
		return
	}
	file_authzed_api_v0_core_proto_init()
	file_authzed_api_v0_namespace_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_authzed_api_v0_namespace_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadConfigRequest); i {
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
		file_authzed_api_v0_namespace_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadConfigResponse); i {
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
		file_authzed_api_v0_namespace_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteConfigRequest); i {
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
		file_authzed_api_v0_namespace_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteConfigResponse); i {
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
			RawDescriptor: file_authzed_api_v0_namespace_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authzed_api_v0_namespace_service_proto_goTypes,
		DependencyIndexes: file_authzed_api_v0_namespace_service_proto_depIdxs,
		MessageInfos:      file_authzed_api_v0_namespace_service_proto_msgTypes,
	}.Build()
	File_authzed_api_v0_namespace_service_proto = out.File
	file_authzed_api_v0_namespace_service_proto_rawDesc = nil
	file_authzed_api_v0_namespace_service_proto_goTypes = nil
	file_authzed_api_v0_namespace_service_proto_depIdxs = nil
}
