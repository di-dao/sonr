// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: proto/node/event.proto

package api

import (
	common "github.com/sonr-io/core/common"
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

// DecisionEvent is emitted when a decision is made by Peer.
type DecisionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Decision bool         `protobuf:"varint,1,opt,name=decision,proto3" json:"decision,omitempty"` // true = accept, false = reject
	From     *common.Peer `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`          // Peer that made decision
	Received int64        `protobuf:"varint,3,opt,name=received,proto3" json:"received,omitempty"` // Timestamp
}

func (x *DecisionEvent) Reset() {
	*x = DecisionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecisionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecisionEvent) ProtoMessage() {}

func (x *DecisionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_event_proto_msgTypes[0]
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
	return file_proto_node_event_proto_rawDescGZIP(), []int{0}
}

func (x *DecisionEvent) GetDecision() bool {
	if x != nil {
		return x.Decision
	}
	return false
}

func (x *DecisionEvent) GetFrom() *common.Peer {
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

	Olc      string         `protobuf:"bytes,1,opt,name=olc,proto3" json:"olc,omitempty"`            // OLC Code of Topic
	Peers    []*common.Peer `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`        // User Information
	Received int64          `protobuf:"varint,3,opt,name=received,proto3" json:"received,omitempty"` // Invite received Timestamp
}

func (x *RefreshEvent) Reset() {
	*x = RefreshEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshEvent) ProtoMessage() {}

func (x *RefreshEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_event_proto_msgTypes[1]
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
	return file_proto_node_event_proto_rawDescGZIP(), []int{1}
}

func (x *RefreshEvent) GetOlc() string {
	if x != nil {
		return x.Olc
	}
	return ""
}

func (x *RefreshEvent) GetPeers() []*common.Peer {
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

	Received int64           `protobuf:"varint,1,opt,name=received,proto3" json:"received,omitempty"` // Invite received Timestamp
	From     *common.Peer    `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`          // Peer that sent the Invite
	Payload  *common.Payload `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`    // Attached Data
}

func (x *InviteEvent) Reset() {
	*x = InviteEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteEvent) ProtoMessage() {}

func (x *InviteEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_event_proto_msgTypes[2]
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
	return file_proto_node_event_proto_rawDescGZIP(), []int{2}
}

func (x *InviteEvent) GetReceived() int64 {
	if x != nil {
		return x.Received
	}
	return 0
}

func (x *InviteEvent) GetFrom() *common.Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *InviteEvent) GetPayload() *common.Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

// Received Mail Event
type MailboxEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`             // ID is the Message ID
	Buffer   []byte           `protobuf:"bytes,2,opt,name=buffer,proto3" json:"buffer,omitempty"`     // Buffer is the message data
	From     *common.Profile  `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`         // Users Peer Data
	To       *common.Profile  `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`             // Receivers Peer Data
	Metadata *common.Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"` // Metadata
}

func (x *MailboxEvent) Reset() {
	*x = MailboxEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailboxEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxEvent) ProtoMessage() {}

func (x *MailboxEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxEvent.ProtoReflect.Descriptor instead.
func (*MailboxEvent) Descriptor() ([]byte, []int) {
	return file_proto_node_event_proto_rawDescGZIP(), []int{3}
}

func (x *MailboxEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MailboxEvent) GetBuffer() []byte {
	if x != nil {
		return x.Buffer
	}
	return nil
}

func (x *MailboxEvent) GetFrom() *common.Profile {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *MailboxEvent) GetTo() *common.Profile {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *MailboxEvent) GetMetadata() *common.Metadata {
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

	Progress  float64          `protobuf:"fixed64,1,opt,name=progress,proto3" json:"progress,omitempty"`                           // Current Transfer Progress
	Received  int64            `protobuf:"varint,2,opt,name=received,proto3" json:"received,omitempty"`                            // Timestamp
	Current   int32            `protobuf:"varint,3,opt,name=current,proto3" json:"current,omitempty"`                              // Current position of item in list
	Total     int32            `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`                                  // Total number of items in list
	Direction common.Direction `protobuf:"varint,5,opt,name=direction,proto3,enum=sonr.core.Direction" json:"direction,omitempty"` // Direction of Transfer
}

func (x *ProgressEvent) Reset() {
	*x = ProgressEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressEvent) ProtoMessage() {}

func (x *ProgressEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_event_proto_msgTypes[4]
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
	return file_proto_node_event_proto_rawDescGZIP(), []int{4}
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

func (x *ProgressEvent) GetDirection() common.Direction {
	if x != nil {
		return x.Direction
	}
	return common.Direction(0)
}

// Message Sent after Completed Transfer
type CompleteEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Direction  common.Direction `protobuf:"varint,1,opt,name=direction,proto3,enum=sonr.core.Direction" json:"direction,omitempty"`                                                             // Direction of Transfer
	Payload    *common.Payload  `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`                                                                                           // Transfer Data
	From       *common.Peer     `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`                                                                                                 // Peer that sent the Complete Event
	To         *common.Peer     `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`                                                                                                     // Peer that received the Complete Event
	CreatedAt  int64            `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`                                                                                      // Transfer Created Timestamp
	ReceivedAt int64            `protobuf:"varint,6,opt,name=receivedAt,proto3" json:"receivedAt,omitempty"`                                                                                    // Transfer Received Timestamp
	Results    map[int32]bool   `protobuf:"bytes,7,rep,name=results,proto3" json:"results,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // Transfer Success
}

func (x *CompleteEvent) Reset() {
	*x = CompleteEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteEvent) ProtoMessage() {}

func (x *CompleteEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_event_proto_msgTypes[5]
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
	return file_proto_node_event_proto_rawDescGZIP(), []int{5}
}

func (x *CompleteEvent) GetDirection() common.Direction {
	if x != nil {
		return x.Direction
	}
	return common.Direction(0)
}

func (x *CompleteEvent) GetPayload() *common.Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *CompleteEvent) GetFrom() *common.Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *CompleteEvent) GetTo() *common.Peer {
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

func (x *CompleteEvent) GetResults() map[int32]bool {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_proto_node_event_proto protoreflect.FileDescriptor

var file_proto_node_event_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0d, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x65, 0x65, 0x72,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x22, 0x63, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x6c, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6f, 0x6c, 0x63, 0x12, 0x25, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x22, 0x7c, 0x0a, 0x0b, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2c, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xb3, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f,
	0x78, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x22, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x2f, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xab, 0x01, 0x0a, 0x0d,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf1, 0x02, 0x0a, 0x0d, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x6f,
	0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x1f, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x3e, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x22, 0x5a,
	0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72,
	0x2d, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_node_event_proto_rawDescOnce sync.Once
	file_proto_node_event_proto_rawDescData = file_proto_node_event_proto_rawDesc
)

func file_proto_node_event_proto_rawDescGZIP() []byte {
	file_proto_node_event_proto_rawDescOnce.Do(func() {
		file_proto_node_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_node_event_proto_rawDescData)
	})
	return file_proto_node_event_proto_rawDescData
}

var file_proto_node_event_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_node_event_proto_goTypes = []interface{}{
	(*DecisionEvent)(nil),   // 0: sonr.api.DecisionEvent
	(*RefreshEvent)(nil),    // 1: sonr.api.RefreshEvent
	(*InviteEvent)(nil),     // 2: sonr.api.InviteEvent
	(*MailboxEvent)(nil),    // 3: sonr.api.MailboxEvent
	(*ProgressEvent)(nil),   // 4: sonr.api.ProgressEvent
	(*CompleteEvent)(nil),   // 5: sonr.api.CompleteEvent
	nil,                     // 6: sonr.api.CompleteEvent.ResultsEntry
	(*common.Peer)(nil),     // 7: sonr.core.Peer
	(*common.Payload)(nil),  // 8: sonr.core.Payload
	(*common.Profile)(nil),  // 9: sonr.core.Profile
	(*common.Metadata)(nil), // 10: sonr.core.Metadata
	(common.Direction)(0),   // 11: sonr.core.Direction
}
var file_proto_node_event_proto_depIdxs = []int32{
	7,  // 0: sonr.api.DecisionEvent.from:type_name -> sonr.core.Peer
	7,  // 1: sonr.api.RefreshEvent.peers:type_name -> sonr.core.Peer
	7,  // 2: sonr.api.InviteEvent.from:type_name -> sonr.core.Peer
	8,  // 3: sonr.api.InviteEvent.payload:type_name -> sonr.core.Payload
	9,  // 4: sonr.api.MailboxEvent.from:type_name -> sonr.core.Profile
	9,  // 5: sonr.api.MailboxEvent.to:type_name -> sonr.core.Profile
	10, // 6: sonr.api.MailboxEvent.metadata:type_name -> sonr.core.Metadata
	11, // 7: sonr.api.ProgressEvent.direction:type_name -> sonr.core.Direction
	11, // 8: sonr.api.CompleteEvent.direction:type_name -> sonr.core.Direction
	8,  // 9: sonr.api.CompleteEvent.payload:type_name -> sonr.core.Payload
	7,  // 10: sonr.api.CompleteEvent.from:type_name -> sonr.core.Peer
	7,  // 11: sonr.api.CompleteEvent.to:type_name -> sonr.core.Peer
	6,  // 12: sonr.api.CompleteEvent.results:type_name -> sonr.api.CompleteEvent.ResultsEntry
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_proto_node_event_proto_init() }
func file_proto_node_event_proto_init() {
	if File_proto_node_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_node_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_node_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_node_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_node_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailboxEvent); i {
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
		file_proto_node_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_node_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_proto_node_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_node_event_proto_goTypes,
		DependencyIndexes: file_proto_node_event_proto_depIdxs,
		MessageInfos:      file_proto_node_event_proto_msgTypes,
	}.Build()
	File_proto_node_event_proto = out.File
	file_proto_node_event_proto_rawDesc = nil
	file_proto_node_event_proto_goTypes = nil
	file_proto_node_event_proto_depIdxs = nil
}
