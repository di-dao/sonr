// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/common/event.proto

package common

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

// Traffic Direction of Session
type CompleteEvent_Direction int32

const (
	CompleteEvent_DEFAULT  CompleteEvent_Direction = 0 // From Memory/Viewing
	CompleteEvent_INCOMING CompleteEvent_Direction = 1 // Incoming Transfer
	CompleteEvent_OUTGOING CompleteEvent_Direction = 2 // Outgoing Transfer
)

// Enum value maps for CompleteEvent_Direction.
var (
	CompleteEvent_Direction_name = map[int32]string{
		0: "DEFAULT",
		1: "INCOMING",
		2: "OUTGOING",
	}
	CompleteEvent_Direction_value = map[string]int32{
		"DEFAULT":  0,
		"INCOMING": 1,
		"OUTGOING": 2,
	}
)

func (x CompleteEvent_Direction) Enum() *CompleteEvent_Direction {
	p := new(CompleteEvent_Direction)
	*p = x
	return p
}

func (x CompleteEvent_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompleteEvent_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_common_event_proto_enumTypes[0].Descriptor()
}

func (CompleteEvent_Direction) Type() protoreflect.EnumType {
	return &file_proto_common_event_proto_enumTypes[0]
}

func (x CompleteEvent_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompleteEvent_Direction.Descriptor instead.
func (CompleteEvent_Direction) EnumDescriptor() ([]byte, []int) {
	return file_proto_common_event_proto_rawDescGZIP(), []int{5, 0}
}

// DecisionEvent is emitted when a decision is made by Peer.
type DecisionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Decision bool  `protobuf:"varint,1,opt,name=decision,proto3" json:"decision,omitempty"` // true = accept, false = reject
	From     *Peer `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`          // Peer that made decision
	Received int64 `protobuf:"varint,3,opt,name=received,proto3" json:"received,omitempty"` // Timestamp
}

func (x *DecisionEvent) Reset() {
	*x = DecisionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecisionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecisionEvent) ProtoMessage() {}

func (x *DecisionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecisionEvent.ProtoReflect.Descriptor instead.
func (*DecisionEvent) Descriptor() ([]byte, []int) {
	return file_proto_common_event_proto_rawDescGZIP(), []int{0}
}

func (x *DecisionEvent) GetDecision() bool {
	if x != nil {
		return x.Decision
	}
	return false
}

func (x *DecisionEvent) GetFrom() *Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *DecisionEvent) GetReceived() int64 {
	if x != nil {
		return x.Received
	}
	return 0
}

// Message Sent when peer messages Lobby
type RefreshEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Olc      string  `protobuf:"bytes,1,opt,name=olc,proto3" json:"olc,omitempty"`            // OLC Code of Topic
	Peers    []*Peer `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`        // User Information
	Received int64   `protobuf:"varint,3,opt,name=received,proto3" json:"received,omitempty"` // Invite received Timestamp
}

func (x *RefreshEvent) Reset() {
	*x = RefreshEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshEvent) ProtoMessage() {}

func (x *RefreshEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshEvent.ProtoReflect.Descriptor instead.
func (*RefreshEvent) Descriptor() ([]byte, []int) {
	return file_proto_common_event_proto_rawDescGZIP(), []int{1}
}

func (x *RefreshEvent) GetOlc() string {
	if x != nil {
		return x.Olc
	}
	return ""
}

func (x *RefreshEvent) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *RefreshEvent) GetReceived() int64 {
	if x != nil {
		return x.Received
	}
	return 0
}

// InviteEvent notifies Peer that an Invite has been received
type InviteEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Received int64    `protobuf:"varint,1,opt,name=received,proto3" json:"received,omitempty"` // Invite received Timestamp
	From     *Peer    `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`          // Peer that sent the Invite
	Payload  *Payload `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`    // Attached Data
}

func (x *InviteEvent) Reset() {
	*x = InviteEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteEvent) ProtoMessage() {}

func (x *InviteEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteEvent.ProtoReflect.Descriptor instead.
func (*InviteEvent) Descriptor() ([]byte, []int) {
	return file_proto_common_event_proto_rawDescGZIP(), []int{2}
}

func (x *InviteEvent) GetReceived() int64 {
	if x != nil {
		return x.Received
	}
	return 0
}

func (x *InviteEvent) GetFrom() *Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *InviteEvent) GetPayload() *Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

// Received Mail Event
type MailEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`             // ID is the Message ID
	Buffer   []byte    `protobuf:"bytes,2,opt,name=buffer,proto3" json:"buffer,omitempty"`     // Buffer is the message data
	From     *Profile  `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`         // Users Peer Data
	To       *Profile  `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`             // Receivers Peer Data
	Metadata *Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"` // Metadata
}

func (x *MailEvent) Reset() {
	*x = MailEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailEvent) ProtoMessage() {}

func (x *MailEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailEvent.ProtoReflect.Descriptor instead.
func (*MailEvent) Descriptor() ([]byte, []int) {
	return file_proto_common_event_proto_rawDescGZIP(), []int{3}
}

func (x *MailEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MailEvent) GetBuffer() []byte {
	if x != nil {
		return x.Buffer
	}
	return nil
}

func (x *MailEvent) GetFrom() *Profile {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *MailEvent) GetTo() *Profile {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *MailEvent) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Transfer Progress Event
type ProgressEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current  int32   `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`    // Current Transfer Item
	Total    int32   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`        // Total Transfer Progress
	Progress float64 `protobuf:"fixed64,3,opt,name=progress,proto3" json:"progress,omitempty"` // Current Transfer Progress
	Received int64   `protobuf:"varint,4,opt,name=received,proto3" json:"received,omitempty"`  // Timestamp
}

func (x *ProgressEvent) Reset() {
	*x = ProgressEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressEvent) ProtoMessage() {}

func (x *ProgressEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressEvent.ProtoReflect.Descriptor instead.
func (*ProgressEvent) Descriptor() ([]byte, []int) {
	return file_proto_common_event_proto_rawDescGZIP(), []int{4}
}

func (x *ProgressEvent) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *ProgressEvent) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ProgressEvent) GetProgress() float64 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *ProgressEvent) GetReceived() int64 {
	if x != nil {
		return x.Received
	}
	return 0
}

// Message Sent after Completed Transfer
type CompleteEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Direction  CompleteEvent_Direction `protobuf:"varint,1,opt,name=direction,proto3,enum=sonr.core.CompleteEvent_Direction" json:"direction,omitempty"` // Transfer Direction
	Payload    *Payload                `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`                                             // Transfer Data
	From       *Peer                   `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`                                                   // Peer that sent the Complete Event
	To         *Peer                   `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`                                                       // Peer that received the Complete Event
	CreatedAt  int64                   `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`                                        // Transfer Created Timestamp
	ReceivedAt int64                   `protobuf:"varint,6,opt,name=receivedAt,proto3" json:"receivedAt,omitempty"`                                      // Transfer Received Timestamp
}

func (x *CompleteEvent) Reset() {
	*x = CompleteEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteEvent) ProtoMessage() {}

func (x *CompleteEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteEvent.ProtoReflect.Descriptor instead.
func (*CompleteEvent) Descriptor() ([]byte, []int) {
	return file_proto_common_event_proto_rawDescGZIP(), []int{5}
}

func (x *CompleteEvent) GetDirection() CompleteEvent_Direction {
	if x != nil {
		return x.Direction
	}
	return CompleteEvent_DEFAULT
}

func (x *CompleteEvent) GetPayload() *Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *CompleteEvent) GetFrom() *Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *CompleteEvent) GetTo() *Peer {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *CompleteEvent) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *CompleteEvent) GetReceivedAt() int64 {
	if x != nil {
		return x.ReceivedAt
	}
	return 0
}

var File_proto_common_event_proto protoreflect.FileDescriptor

var file_proto_common_event_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x6f, 0x6e, 0x72,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0d, 0x44, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x22, 0x63, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x6c, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6f, 0x6c, 0x63, 0x12, 0x25, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x22, 0x7c, 0x0a, 0x0b, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2c, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x6f,
	0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xb0, 0x01, 0x0a, 0x09, 0x4d, 0x61, 0x69,
	0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x22, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x2f, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x77, 0x0a, 0x0d, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x22, 0xb9, 0x02, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1f, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x22, 0x34, 0x0a, 0x09, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55,
	0x4c, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x55, 0x54, 0x47, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_common_event_proto_rawDescOnce sync.Once
	file_proto_common_event_proto_rawDescData = file_proto_common_event_proto_rawDesc
)

func file_proto_common_event_proto_rawDescGZIP() []byte {
	file_proto_common_event_proto_rawDescOnce.Do(func() {
		file_proto_common_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_common_event_proto_rawDescData)
	})
	return file_proto_common_event_proto_rawDescData
}

var file_proto_common_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_common_event_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_common_event_proto_goTypes = []interface{}{
	(CompleteEvent_Direction)(0), // 0: sonr.core.CompleteEvent.Direction
	(*DecisionEvent)(nil),        // 1: sonr.core.DecisionEvent
	(*RefreshEvent)(nil),         // 2: sonr.core.RefreshEvent
	(*InviteEvent)(nil),          // 3: sonr.core.InviteEvent
	(*MailEvent)(nil),            // 4: sonr.core.MailEvent
	(*ProgressEvent)(nil),        // 5: sonr.core.ProgressEvent
	(*CompleteEvent)(nil),        // 6: sonr.core.CompleteEvent
	(*Peer)(nil),                 // 7: sonr.core.Peer
	(*Payload)(nil),              // 8: sonr.core.Payload
	(*Profile)(nil),              // 9: sonr.core.Profile
	(*Metadata)(nil),             // 10: sonr.core.Metadata
}
var file_proto_common_event_proto_depIdxs = []int32{
	7,  // 0: sonr.core.DecisionEvent.from:type_name -> sonr.core.Peer
	7,  // 1: sonr.core.RefreshEvent.peers:type_name -> sonr.core.Peer
	7,  // 2: sonr.core.InviteEvent.from:type_name -> sonr.core.Peer
	8,  // 3: sonr.core.InviteEvent.payload:type_name -> sonr.core.Payload
	9,  // 4: sonr.core.MailEvent.from:type_name -> sonr.core.Profile
	9,  // 5: sonr.core.MailEvent.to:type_name -> sonr.core.Profile
	10, // 6: sonr.core.MailEvent.metadata:type_name -> sonr.core.Metadata
	0,  // 7: sonr.core.CompleteEvent.direction:type_name -> sonr.core.CompleteEvent.Direction
	8,  // 8: sonr.core.CompleteEvent.payload:type_name -> sonr.core.Payload
	7,  // 9: sonr.core.CompleteEvent.from:type_name -> sonr.core.Peer
	7,  // 10: sonr.core.CompleteEvent.to:type_name -> sonr.core.Peer
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_proto_common_event_proto_init() }
func file_proto_common_event_proto_init() {
	if File_proto_common_event_proto != nil {
		return
	}
	file_proto_common_core_proto_init()
	file_proto_common_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_common_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecisionEvent); i {
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
		file_proto_common_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshEvent); i {
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
		file_proto_common_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteEvent); i {
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
		file_proto_common_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailEvent); i {
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
		file_proto_common_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressEvent); i {
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
		file_proto_common_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteEvent); i {
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
			RawDescriptor: file_proto_common_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_common_event_proto_goTypes,
		DependencyIndexes: file_proto_common_event_proto_depIdxs,
		EnumInfos:         file_proto_common_event_proto_enumTypes,
		MessageInfos:      file_proto_common_event_proto_msgTypes,
	}.Build()
	File_proto_common_event_proto = out.File
	file_proto_common_event_proto_rawDesc = nil
	file_proto_common_event_proto_goTypes = nil
	file_proto_common_event_proto_depIdxs = nil
}
