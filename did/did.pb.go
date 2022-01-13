// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: did/v1/did.proto

package did

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

// Did represents a string that has been parsed and validated as a DID. The parts are stored
// in the individual fields.
type Did struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Method is the method used to create the DID. For the Sonr network it is "sonr".
	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// Network is the network the DID is on. For testnet it is "testnet". i.e "did:sonr:testnet:".
	Network string `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	// id is the trailing identifier after the network. i.e. "did:sonr:testnet:abc123"
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Paths is a list of paths that the DID is valid for. This is used to identify the Service.
	Paths []string `protobuf:"bytes,4,rep,name=paths,proto3" json:"paths,omitempty"`
	// Query is the query string that was used to create the DID. This is followed by a '?'.
	Query string `protobuf:"bytes,5,opt,name=query,proto3" json:"query,omitempty"`
	// Fragment is the fragment string that was used to create the DID. This is followed by a '#'.
	Fragment string `protobuf:"bytes,6,opt,name=fragment,proto3" json:"fragment,omitempty"`
}

func (x *Did) Reset() {
	*x = Did{}
	if protoimpl.UnsafeEnabled {
		mi := &file_did_v1_did_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Did) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Did) ProtoMessage() {}

func (x *Did) ProtoReflect() protoreflect.Message {
	mi := &file_did_v1_did_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Did.ProtoReflect.Descriptor instead.
func (*Did) Descriptor() ([]byte, []int) {
	return file_did_v1_did_proto_rawDescGZIP(), []int{0}
}

func (x *Did) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Did) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Did) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Did) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *Did) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Did) GetFragment() string {
	if x != nil {
		return x.Fragment
	}
	return ""
}

// DidDocument is the document that describes a DID. This document is stored on the blockchain.
type DidDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Context is the context of the DID document. This is used to identify the Service.
	Context []string `protobuf:"bytes,1,rep,name=context,proto3" json:"context,omitempty"` // optional
	// Id is the DID of the document.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Controller is the DID of the controller of the document. This will be the individual user devices and mailboxes.
	Controller []string `protobuf:"bytes,3,rep,name=controller,proto3" json:"controller,omitempty"` //optional
	// VerificationMethod is the list of verification methods for the user.
	VerificationMethod []*VerificationMethod `protobuf:"bytes,4,rep,name=verification_method,json=verificationMethod,proto3" json:"verification_method,omitempty"` // optional
	// Authentication is the list of authentication methods for the user.
	Authentication []string `protobuf:"bytes,5,rep,name=authentication,proto3" json:"authentication,omitempty"` // optional
	// AssertionMethod is the list of assertion methods for the user.
	AssertionMethod []string `protobuf:"bytes,6,rep,name=assertion_method,json=assertionMethod,proto3" json:"assertion_method,omitempty"` // optional
	// CapabilityInvocation is the list of capability invocation methods for the user.
	CapabilityInvocation []string `protobuf:"bytes,7,rep,name=capability_invocation,json=capabilityInvocation,proto3" json:"capability_invocation,omitempty"` // optional
	// CapabilityDelegation is the list of capability delegation methods for the user.
	CapabilityDelegation []string `protobuf:"bytes,8,rep,name=capability_delegation,json=capabilityDelegation,proto3" json:"capability_delegation,omitempty"` // optional
	// KeyAgreement is the list of key agreement methods for the user.
	KeyAgreement []string `protobuf:"bytes,9,rep,name=key_agreement,json=keyAgreement,proto3" json:"key_agreement,omitempty"` // optional
	// Service is the list of services or DApps that the user has access to.
	Service []*Service `protobuf:"bytes,10,rep,name=service,proto3" json:"service,omitempty"` // optional
	// AlsoKnownAs is the list of ".snr" aliases for the user.
	AlsoKnownAs []string `protobuf:"bytes,11,rep,name=also_known_as,json=alsoKnownAs,proto3" json:"also_known_as,omitempty"` // optional
}

func (x *DidDocument) Reset() {
	*x = DidDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_did_v1_did_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DidDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DidDocument) ProtoMessage() {}

func (x *DidDocument) ProtoReflect() protoreflect.Message {
	mi := &file_did_v1_did_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DidDocument.ProtoReflect.Descriptor instead.
func (*DidDocument) Descriptor() ([]byte, []int) {
	return file_did_v1_did_proto_rawDescGZIP(), []int{1}
}

func (x *DidDocument) GetContext() []string {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *DidDocument) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DidDocument) GetController() []string {
	if x != nil {
		return x.Controller
	}
	return nil
}

func (x *DidDocument) GetVerificationMethod() []*VerificationMethod {
	if x != nil {
		return x.VerificationMethod
	}
	return nil
}

func (x *DidDocument) GetAuthentication() []string {
	if x != nil {
		return x.Authentication
	}
	return nil
}

func (x *DidDocument) GetAssertionMethod() []string {
	if x != nil {
		return x.AssertionMethod
	}
	return nil
}

func (x *DidDocument) GetCapabilityInvocation() []string {
	if x != nil {
		return x.CapabilityInvocation
	}
	return nil
}

func (x *DidDocument) GetCapabilityDelegation() []string {
	if x != nil {
		return x.CapabilityDelegation
	}
	return nil
}

func (x *DidDocument) GetKeyAgreement() []string {
	if x != nil {
		return x.KeyAgreement
	}
	return nil
}

func (x *DidDocument) GetService() []*Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *DidDocument) GetAlsoKnownAs() []string {
	if x != nil {
		return x.AlsoKnownAs
	}
	return nil
}

// Service is a Application that runs on the Sonr network.
type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the DID of the service.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Type is the type of the service.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// ServiceEndpoint is the endpoint of the service.
	ServiceEndpoint string `protobuf:"bytes,3,opt,name=service_endpoint,json=serviceEndpoint,proto3" json:"service_endpoint,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_did_v1_did_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_did_v1_did_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_did_v1_did_proto_rawDescGZIP(), []int{2}
}

func (x *Service) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Service) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Service) GetServiceEndpoint() string {
	if x != nil {
		return x.ServiceEndpoint
	}
	return ""
}

// VerificationMethod is a method that can be used to verify the DID.
type VerificationMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the DID of the verification method.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Type is the type of the verification method.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Controller is the DID of the controller of the verification method.
	Controller string `protobuf:"bytes,3,opt,name=controller,proto3" json:"controller,omitempty"`
	// PublicKeyHex is the public key of the verification method in hexidecimal.
	PublicKeyHex string `protobuf:"bytes,4,opt,name=public_key_hex,json=publicKeyHex,proto3" json:"public_key_hex,omitempty"` // optional
	// PublicKeyBase58 is the public key of the verification method in base58.
	PublicKeyBase58 string `protobuf:"bytes,5,opt,name=public_key_base58,json=publicKeyBase58,proto3" json:"public_key_base58,omitempty"` // optional
	// BlockchainAccountId is the blockchain account id of the verification method.
	BlockchainAccountId string `protobuf:"bytes,6,opt,name=blockchain_account_id,json=blockchainAccountId,proto3" json:"blockchain_account_id,omitempty"` // optional
}

func (x *VerificationMethod) Reset() {
	*x = VerificationMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_did_v1_did_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerificationMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerificationMethod) ProtoMessage() {}

func (x *VerificationMethod) ProtoReflect() protoreflect.Message {
	mi := &file_did_v1_did_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerificationMethod.ProtoReflect.Descriptor instead.
func (*VerificationMethod) Descriptor() ([]byte, []int) {
	return file_did_v1_did_proto_rawDescGZIP(), []int{3}
}

func (x *VerificationMethod) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VerificationMethod) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *VerificationMethod) GetController() string {
	if x != nil {
		return x.Controller
	}
	return ""
}

func (x *VerificationMethod) GetPublicKeyHex() string {
	if x != nil {
		return x.PublicKeyHex
	}
	return ""
}

func (x *VerificationMethod) GetPublicKeyBase58() string {
	if x != nil {
		return x.PublicKeyBase58
	}
	return ""
}

func (x *VerificationMethod) GetBlockchainAccountId() string {
	if x != nil {
		return x.BlockchainAccountId
	}
	return ""
}

var File_did_v1_did_proto protoreflect.FileDescriptor

var file_did_v1_did_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x64, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x8f, 0x01, 0x0a, 0x03, 0x44,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xd5, 0x03, 0x0a,
	0x0b, 0x44, 0x69, 0x64, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x13, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52,
	0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x61,
	0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x33, 0x0a, 0x15, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x15, 0x63,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x63, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6b, 0x65, 0x79, 0x41, 0x67, 0x72, 0x65,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x22, 0x0a, 0x0d, 0x61, 0x6c, 0x73, 0x6f, 0x5f, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x61,
	0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6c, 0x73, 0x6f, 0x4b, 0x6e, 0x6f,
	0x77, 0x6e, 0x41, 0x73, 0x22, 0x58, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0xde,
	0x01, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x68, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x48, 0x65, 0x78, 0x12,
	0x2a, 0x0a, 0x11, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x62, 0x61,
	0x73, 0x65, 0x35, 0x38, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x61, 0x73, 0x65, 0x35, 0x38, 0x12, 0x32, 0x0a, 0x15, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x42,
	0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x64, 0x69, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_did_v1_did_proto_rawDescOnce sync.Once
	file_did_v1_did_proto_rawDescData = file_did_v1_did_proto_rawDesc
)

func file_did_v1_did_proto_rawDescGZIP() []byte {
	file_did_v1_did_proto_rawDescOnce.Do(func() {
		file_did_v1_did_proto_rawDescData = protoimpl.X.CompressGZIP(file_did_v1_did_proto_rawDescData)
	})
	return file_did_v1_did_proto_rawDescData
}

var file_did_v1_did_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_did_v1_did_proto_goTypes = []interface{}{
	(*Did)(nil),                // 0: did.v1.Did
	(*DidDocument)(nil),        // 1: did.v1.DidDocument
	(*Service)(nil),            // 2: did.v1.Service
	(*VerificationMethod)(nil), // 3: did.v1.VerificationMethod
}
var file_did_v1_did_proto_depIdxs = []int32{
	3, // 0: did.v1.DidDocument.verification_method:type_name -> did.v1.VerificationMethod
	2, // 1: did.v1.DidDocument.service:type_name -> did.v1.Service
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_did_v1_did_proto_init() }
func file_did_v1_did_proto_init() {
	if File_did_v1_did_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_did_v1_did_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Did); i {
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
		file_did_v1_did_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DidDocument); i {
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
		file_did_v1_did_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_did_v1_did_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerificationMethod); i {
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
			RawDescriptor: file_did_v1_did_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_did_v1_did_proto_goTypes,
		DependencyIndexes: file_did_v1_did_proto_depIdxs,
		MessageInfos:      file_did_v1_did_proto_msgTypes,
	}.Build()
	File_did_v1_did_proto = out.File
	file_did_v1_did_proto_rawDesc = nil
	file_did_v1_did_proto_goTypes = nil
	file_did_v1_did_proto_depIdxs = nil
}
