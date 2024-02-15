// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: sonr/service/module/v1/state.proto

package modulev1

import (
	_ "cosmossdk.io/api/cosmos/orm/v1"
	_ "cosmossdk.io/api/cosmos/orm/v1alpha1"
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

// ServicePermissions is the specified permissions the application has with the
// user's account.
type ServicePermissions int32

const (
	// SERVICE_PERMISSIONS_BASE - off chain visibility
	ServicePermissions_SERVICE_PERMISSIONS_BASE ServicePermissions = 0
	// SERVICE_PERMISSIONS_READ - read access
	ServicePermissions_SERVICE_PERMISSIONS_READ ServicePermissions = 1
	// SERVICE_PERMISSIONS_WRITE - write access
	ServicePermissions_SERVICE_PERMISSIONS_WRITE ServicePermissions = 2
	// SERVICE_PERMISSIONS_OWN - ownership
	ServicePermissions_SERVICE_PERMISSIONS_OWN ServicePermissions = 3
)

// Enum value maps for ServicePermissions.
var (
	ServicePermissions_name = map[int32]string{
		0: "SERVICE_PERMISSIONS_BASE",
		1: "SERVICE_PERMISSIONS_READ",
		2: "SERVICE_PERMISSIONS_WRITE",
		3: "SERVICE_PERMISSIONS_OWN",
	}
	ServicePermissions_value = map[string]int32{
		"SERVICE_PERMISSIONS_BASE":  0,
		"SERVICE_PERMISSIONS_READ":  1,
		"SERVICE_PERMISSIONS_WRITE": 2,
		"SERVICE_PERMISSIONS_OWN":   3,
	}
)

func (x ServicePermissions) Enum() *ServicePermissions {
	p := new(ServicePermissions)
	*p = x
	return p
}

func (x ServicePermissions) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServicePermissions) Descriptor() protoreflect.EnumDescriptor {
	return file_sonr_service_module_v1_state_proto_enumTypes[0].Descriptor()
}

func (ServicePermissions) Type() protoreflect.EnumType {
	return &file_sonr_service_module_v1_state_proto_enumTypes[0]
}

func (x ServicePermissions) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServicePermissions.Descriptor instead.
func (ServicePermissions) EnumDescriptor() ([]byte, []int) {
	return file_sonr_service_module_v1_state_proto_rawDescGZIP(), []int{0}
}

// Module is the app config object of the module.
// Learn more: https://docs.cosmos.network/main/building-modules/depinject
type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_service_module_v1_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_service_module_v1_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_sonr_service_module_v1_state_proto_rawDescGZIP(), []int{0}
}

// ServiceRecord is the balance of an account.
type ServiceRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Origin      string             `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Name        string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string             `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Authority   string             `protobuf:"bytes,5,opt,name=authority,proto3" json:"authority,omitempty"`
	Permissions ServicePermissions `protobuf:"varint,6,opt,name=permissions,proto3,enum=sonr.service.module.v1.ServicePermissions" json:"permissions,omitempty"`
}

func (x *ServiceRecord) Reset() {
	*x = ServiceRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_service_module_v1_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceRecord) ProtoMessage() {}

func (x *ServiceRecord) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_service_module_v1_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceRecord.ProtoReflect.Descriptor instead.
func (*ServiceRecord) Descriptor() ([]byte, []int) {
	return file_sonr_service_module_v1_state_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceRecord) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ServiceRecord) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *ServiceRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceRecord) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ServiceRecord) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

func (x *ServiceRecord) GetPermissions() ServicePermissions {
	if x != nil {
		return x.Permissions
	}
	return ServicePermissions_SERVICE_PERMISSIONS_BASE
}

// Credential is the total supply of the module.
type Credential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sequence        uint64   `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Handle          string   `protobuf:"bytes,2,opt,name=handle,proto3" json:"handle,omitempty"`
	Transport       []string `protobuf:"bytes,3,rep,name=transport,proto3" json:"transport,omitempty"`
	Authority       string   `protobuf:"bytes,4,opt,name=authority,proto3" json:"authority,omitempty"`
	AttestationType string   `protobuf:"bytes,5,opt,name=attestation_type,json=attestationType,proto3" json:"attestation_type,omitempty"`
	Id              []byte   `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	Origin          string   `protobuf:"bytes,7,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (x *Credential) Reset() {
	*x = Credential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_service_module_v1_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credential) ProtoMessage() {}

func (x *Credential) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_service_module_v1_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credential.ProtoReflect.Descriptor instead.
func (*Credential) Descriptor() ([]byte, []int) {
	return file_sonr_service_module_v1_state_proto_rawDescGZIP(), []int{2}
}

func (x *Credential) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *Credential) GetHandle() string {
	if x != nil {
		return x.Handle
	}
	return ""
}

func (x *Credential) GetTransport() []string {
	if x != nil {
		return x.Transport
	}
	return nil
}

func (x *Credential) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

func (x *Credential) GetAttestationType() string {
	if x != nil {
		return x.AttestationType
	}
	return ""
}

func (x *Credential) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Credential) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

// Identifier is a psuedo-anonomyous representation of a unique id on the Sonr blockchain. Used as
// authorizer to the underlying wallet interface.
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index     uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Origin    string `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Key       string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value     string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Authority string `protobuf:"bytes,5,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_service_module_v1_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_service_module_v1_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_sonr_service_module_v1_state_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *User) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *User) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *User) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *User) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

// BaseParams is the total supply of the module.
type BaseParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions              ServicePermissions `protobuf:"varint,1,opt,name=permissions,proto3,enum=sonr.service.module.v1.ServicePermissions" json:"permissions,omitempty"`
	ResidentKey              string             `protobuf:"bytes,2,opt,name=resident_key,json=residentKey,proto3" json:"resident_key,omitempty"`
	Algorithm                int32              `protobuf:"varint,3,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	AuthenticationAttachment string             `protobuf:"bytes,4,opt,name=authentication_attachment,json=authenticationAttachment,proto3" json:"authentication_attachment,omitempty"`
}

func (x *BaseParams) Reset() {
	*x = BaseParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_service_module_v1_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseParams) ProtoMessage() {}

func (x *BaseParams) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_service_module_v1_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseParams.ProtoReflect.Descriptor instead.
func (*BaseParams) Descriptor() ([]byte, []int) {
	return file_sonr_service_module_v1_state_proto_rawDescGZIP(), []int{4}
}

func (x *BaseParams) GetPermissions() ServicePermissions {
	if x != nil {
		return x.Permissions
	}
	return ServicePermissions_SERVICE_PERMISSIONS_BASE
}

func (x *BaseParams) GetResidentKey() string {
	if x != nil {
		return x.ResidentKey
	}
	return ""
}

func (x *BaseParams) GetAlgorithm() int32 {
	if x != nil {
		return x.Algorithm
	}
	return 0
}

func (x *BaseParams) GetAuthenticationAttachment() string {
	if x != nil {
		return x.AuthenticationAttachment
	}
	return ""
}

// ReadParams is the total supply of the module.
type ReadParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions              ServicePermissions `protobuf:"varint,1,opt,name=permissions,proto3,enum=sonr.service.module.v1.ServicePermissions" json:"permissions,omitempty"`
	ResidentKey              string             `protobuf:"bytes,2,opt,name=resident_key,json=residentKey,proto3" json:"resident_key,omitempty"`
	Algorithm                int32              `protobuf:"varint,3,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	AuthenticationAttachment string             `protobuf:"bytes,4,opt,name=authentication_attachment,json=authenticationAttachment,proto3" json:"authentication_attachment,omitempty"`
}

func (x *ReadParams) Reset() {
	*x = ReadParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_service_module_v1_state_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadParams) ProtoMessage() {}

func (x *ReadParams) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_service_module_v1_state_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadParams.ProtoReflect.Descriptor instead.
func (*ReadParams) Descriptor() ([]byte, []int) {
	return file_sonr_service_module_v1_state_proto_rawDescGZIP(), []int{5}
}

func (x *ReadParams) GetPermissions() ServicePermissions {
	if x != nil {
		return x.Permissions
	}
	return ServicePermissions_SERVICE_PERMISSIONS_BASE
}

func (x *ReadParams) GetResidentKey() string {
	if x != nil {
		return x.ResidentKey
	}
	return ""
}

func (x *ReadParams) GetAlgorithm() int32 {
	if x != nil {
		return x.Algorithm
	}
	return 0
}

func (x *ReadParams) GetAuthenticationAttachment() string {
	if x != nil {
		return x.AuthenticationAttachment
	}
	return ""
}

// WriteParams is the total supply of the module.
type WriteParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions              ServicePermissions `protobuf:"varint,1,opt,name=permissions,proto3,enum=sonr.service.module.v1.ServicePermissions" json:"permissions,omitempty"`
	ResidentKey              string             `protobuf:"bytes,2,opt,name=resident_key,json=residentKey,proto3" json:"resident_key,omitempty"`
	Algorithm                int32              `protobuf:"varint,3,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	AuthenticationAttachment string             `protobuf:"bytes,4,opt,name=authentication_attachment,json=authenticationAttachment,proto3" json:"authentication_attachment,omitempty"`
}

func (x *WriteParams) Reset() {
	*x = WriteParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_service_module_v1_state_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteParams) ProtoMessage() {}

func (x *WriteParams) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_service_module_v1_state_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteParams.ProtoReflect.Descriptor instead.
func (*WriteParams) Descriptor() ([]byte, []int) {
	return file_sonr_service_module_v1_state_proto_rawDescGZIP(), []int{6}
}

func (x *WriteParams) GetPermissions() ServicePermissions {
	if x != nil {
		return x.Permissions
	}
	return ServicePermissions_SERVICE_PERMISSIONS_BASE
}

func (x *WriteParams) GetResidentKey() string {
	if x != nil {
		return x.ResidentKey
	}
	return ""
}

func (x *WriteParams) GetAlgorithm() int32 {
	if x != nil {
		return x.Algorithm
	}
	return 0
}

func (x *WriteParams) GetAuthenticationAttachment() string {
	if x != nil {
		return x.AuthenticationAttachment
	}
	return ""
}

// OwnParams is the total supply of the module.
type OwnParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions              ServicePermissions `protobuf:"varint,1,opt,name=permissions,proto3,enum=sonr.service.module.v1.ServicePermissions" json:"permissions,omitempty"`
	ResidentKey              string             `protobuf:"bytes,2,opt,name=resident_key,json=residentKey,proto3" json:"resident_key,omitempty"`
	Algorithm                int32              `protobuf:"varint,3,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	AuthenticationAttachment string             `protobuf:"bytes,4,opt,name=authentication_attachment,json=authenticationAttachment,proto3" json:"authentication_attachment,omitempty"`
}

func (x *OwnParams) Reset() {
	*x = OwnParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_service_module_v1_state_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OwnParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnParams) ProtoMessage() {}

func (x *OwnParams) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_service_module_v1_state_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OwnParams.ProtoReflect.Descriptor instead.
func (*OwnParams) Descriptor() ([]byte, []int) {
	return file_sonr_service_module_v1_state_proto_rawDescGZIP(), []int{7}
}

func (x *OwnParams) GetPermissions() ServicePermissions {
	if x != nil {
		return x.Permissions
	}
	return ServicePermissions_SERVICE_PERMISSIONS_BASE
}

func (x *OwnParams) GetResidentKey() string {
	if x != nil {
		return x.ResidentKey
	}
	return ""
}

func (x *OwnParams) GetAlgorithm() int32 {
	if x != nil {
		return x.Algorithm
	}
	return 0
}

func (x *OwnParams) GetAuthenticationAttachment() string {
	if x != nil {
		return x.AuthenticationAttachment
	}
	return ""
}

var File_sonr_service_module_v1_state_proto protoreflect.FileDescriptor

var file_sonr_service_module_v1_state_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72,
	0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x3a, 0x2a, 0x82, 0x9f, 0xd3, 0x8e, 0x03, 0x24, 0x0a, 0x22, 0x08, 0x01, 0x12, 0x1e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02, 0x0a,
	0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x4c, 0x0a, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2a, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x2d, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x27,
	0x0a, 0x06, 0x0a, 0x02, 0x69, 0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x10, 0x01, 0x18, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x10, 0x02, 0x18, 0x01, 0x22, 0xaa, 0x02, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x3a, 0x59, 0xf2, 0x9e, 0xd3, 0x8e, 0x03,
	0x53, 0x0a, 0x0c, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0d, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x2c, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x10, 0x02, 0x18, 0x01,
	0x12, 0x08, 0x0a, 0x02, 0x69, 0x64, 0x10, 0x03, 0x18, 0x01, 0x12, 0x16, 0x0a, 0x10, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x2c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x10, 0x04,
	0x18, 0x01, 0x18, 0x02, 0x22, 0xa7, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x3a, 0x2b, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x25, 0x0a, 0x09, 0x0a, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x2c, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x10, 0x01, 0x18, 0x01, 0x18, 0x03, 0x22, 0xe2,
	0x01, 0x0a, 0x0a, 0x42, 0x61, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x4c, 0x0a,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x3b, 0x0a, 0x19,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x18, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x08, 0xfa, 0x9e, 0xd3, 0x8e, 0x03,
	0x02, 0x08, 0x04, 0x22, 0xe2, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x4c, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x12, 0x3b, 0x0a, 0x19, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x08,
	0xfa, 0x9e, 0xd3, 0x8e, 0x03, 0x02, 0x08, 0x05, 0x22, 0xe3, 0x01, 0x0a, 0x0b, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x4c, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65,
	0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x3b, 0x0a, 0x19, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x08, 0xfa, 0x9e, 0xd3, 0x8e, 0x03, 0x02, 0x08, 0x06, 0x22, 0xe1,
	0x01, 0x0a, 0x09, 0x4f, 0x77, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x4c, 0x0a, 0x0b,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2a, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x3b, 0x0a, 0x19, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x08, 0xfa, 0x9e, 0xd3, 0x8e, 0x03, 0x02,
	0x08, 0x07, 0x2a, 0x8c, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53,
	0x5f, 0x42, 0x41, 0x53, 0x45, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x52,
	0x45, 0x41, 0x44, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45,
	0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x57, 0x52, 0x49,
	0x54, 0x45, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x4f, 0x57, 0x4e, 0x10,
	0x03, 0x42, 0xea, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x68,
	0x71, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x78, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x4d, 0xaa, 0x02, 0x16, 0x53, 0x6f,
	0x6e, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16, 0x53, 0x6f, 0x6e, 0x72, 0x5c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x22,
	0x53, 0x6f, 0x6e, 0x72, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x19, 0x53, 0x6f, 0x6e, 0x72, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sonr_service_module_v1_state_proto_rawDescOnce sync.Once
	file_sonr_service_module_v1_state_proto_rawDescData = file_sonr_service_module_v1_state_proto_rawDesc
)

func file_sonr_service_module_v1_state_proto_rawDescGZIP() []byte {
	file_sonr_service_module_v1_state_proto_rawDescOnce.Do(func() {
		file_sonr_service_module_v1_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_sonr_service_module_v1_state_proto_rawDescData)
	})
	return file_sonr_service_module_v1_state_proto_rawDescData
}

var file_sonr_service_module_v1_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sonr_service_module_v1_state_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_sonr_service_module_v1_state_proto_goTypes = []interface{}{
	(ServicePermissions)(0), // 0: sonr.service.module.v1.ServicePermissions
	(*State)(nil),           // 1: sonr.service.module.v1.State
	(*ServiceRecord)(nil),   // 2: sonr.service.module.v1.ServiceRecord
	(*Credential)(nil),      // 3: sonr.service.module.v1.Credential
	(*User)(nil),            // 4: sonr.service.module.v1.User
	(*BaseParams)(nil),      // 5: sonr.service.module.v1.BaseParams
	(*ReadParams)(nil),      // 6: sonr.service.module.v1.ReadParams
	(*WriteParams)(nil),     // 7: sonr.service.module.v1.WriteParams
	(*OwnParams)(nil),       // 8: sonr.service.module.v1.OwnParams
}
var file_sonr_service_module_v1_state_proto_depIdxs = []int32{
	0, // 0: sonr.service.module.v1.ServiceRecord.permissions:type_name -> sonr.service.module.v1.ServicePermissions
	0, // 1: sonr.service.module.v1.BaseParams.permissions:type_name -> sonr.service.module.v1.ServicePermissions
	0, // 2: sonr.service.module.v1.ReadParams.permissions:type_name -> sonr.service.module.v1.ServicePermissions
	0, // 3: sonr.service.module.v1.WriteParams.permissions:type_name -> sonr.service.module.v1.ServicePermissions
	0, // 4: sonr.service.module.v1.OwnParams.permissions:type_name -> sonr.service.module.v1.ServicePermissions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_sonr_service_module_v1_state_proto_init() }
func file_sonr_service_module_v1_state_proto_init() {
	if File_sonr_service_module_v1_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sonr_service_module_v1_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
		file_sonr_service_module_v1_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceRecord); i {
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
		file_sonr_service_module_v1_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credential); i {
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
		file_sonr_service_module_v1_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_sonr_service_module_v1_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseParams); i {
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
		file_sonr_service_module_v1_state_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadParams); i {
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
		file_sonr_service_module_v1_state_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteParams); i {
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
		file_sonr_service_module_v1_state_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OwnParams); i {
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
			RawDescriptor: file_sonr_service_module_v1_state_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sonr_service_module_v1_state_proto_goTypes,
		DependencyIndexes: file_sonr_service_module_v1_state_proto_depIdxs,
		EnumInfos:         file_sonr_service_module_v1_state_proto_enumTypes,
		MessageInfos:      file_sonr_service_module_v1_state_proto_msgTypes,
	}.Build()
	File_sonr_service_module_v1_state_proto = out.File
	file_sonr_service_module_v1_state_proto_rawDesc = nil
	file_sonr_service_module_v1_state_proto_goTypes = nil
	file_sonr_service_module_v1_state_proto_depIdxs = nil
}
