// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreateSchema struct {
	// Address of the creator account
	// can be an application address or user address
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// label for the schema.
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	// List of Kind Definitions for the schema
	Fields []*SchemaField `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	// Metadata is a map of key-value pairs that can be used to store additional information about the WhatIs (Schema)
	Metadata []*MetadataDefintion `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *MsgCreateSchema) Reset()         { *m = MsgCreateSchema{} }
func (m *MsgCreateSchema) String() string { return proto.CompactTextString(m) }
func (*MsgCreateSchema) ProtoMessage()    {}
func (*MsgCreateSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_710211b109a793bf, []int{0}
}
func (m *MsgCreateSchema) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateSchema.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateSchema.Merge(m, src)
}
func (m *MsgCreateSchema) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateSchema.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateSchema proto.InternalMessageInfo

func (m *MsgCreateSchema) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateSchema) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *MsgCreateSchema) GetFields() []*SchemaField {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *MsgCreateSchema) GetMetadata() []*MetadataDefintion {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type MsgCreateSchemaResponse struct {
	// Status code of the response
	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// WhatIs object created on chain
	WhatIs *WhatIs `protobuf:"bytes,3,opt,name=what_is,json=whatIs,proto3" json:"what_is,omitempty"`
}

func (m *MsgCreateSchemaResponse) Reset()         { *m = MsgCreateSchemaResponse{} }
func (m *MsgCreateSchemaResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateSchemaResponse) ProtoMessage()    {}
func (*MsgCreateSchemaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_710211b109a793bf, []int{1}
}
func (m *MsgCreateSchemaResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateSchemaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateSchemaResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateSchemaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateSchemaResponse.Merge(m, src)
}
func (m *MsgCreateSchemaResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateSchemaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateSchemaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateSchemaResponse proto.InternalMessageInfo

func (m *MsgCreateSchemaResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *MsgCreateSchemaResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *MsgCreateSchemaResponse) GetWhatIs() *WhatIs {
	if m != nil {
		return m.WhatIs
	}
	return nil
}

type MsgDeprecateSchema struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Did     string `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
}

func (m *MsgDeprecateSchema) Reset()         { *m = MsgDeprecateSchema{} }
func (m *MsgDeprecateSchema) String() string { return proto.CompactTextString(m) }
func (*MsgDeprecateSchema) ProtoMessage()    {}
func (*MsgDeprecateSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_710211b109a793bf, []int{2}
}
func (m *MsgDeprecateSchema) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeprecateSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeprecateSchema.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeprecateSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeprecateSchema.Merge(m, src)
}
func (m *MsgDeprecateSchema) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeprecateSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeprecateSchema.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeprecateSchema proto.InternalMessageInfo

func (m *MsgDeprecateSchema) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeprecateSchema) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

type MsgDeprecateSchemaResponse struct {
	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *MsgDeprecateSchemaResponse) Reset()         { *m = MsgDeprecateSchemaResponse{} }
func (m *MsgDeprecateSchemaResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeprecateSchemaResponse) ProtoMessage()    {}
func (*MsgDeprecateSchemaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_710211b109a793bf, []int{3}
}
func (m *MsgDeprecateSchemaResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeprecateSchemaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeprecateSchemaResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeprecateSchemaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeprecateSchemaResponse.Merge(m, src)
}
func (m *MsgDeprecateSchemaResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeprecateSchemaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeprecateSchemaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeprecateSchemaResponse proto.InternalMessageInfo

func (m *MsgDeprecateSchemaResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *MsgDeprecateSchemaResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgCreateSchema)(nil), "sonrio.sonr.schema.MsgCreateSchema")
	proto.RegisterType((*MsgCreateSchemaResponse)(nil), "sonrio.sonr.schema.MsgCreateSchemaResponse")
	proto.RegisterType((*MsgDeprecateSchema)(nil), "sonrio.sonr.schema.MsgDeprecateSchema")
	proto.RegisterType((*MsgDeprecateSchemaResponse)(nil), "sonrio.sonr.schema.MsgDeprecateSchemaResponse")
}

func init() { proto.RegisterFile("schema/v1/tx.proto", fileDescriptor_710211b109a793bf) }

var fileDescriptor_710211b109a793bf = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x8a, 0xdb, 0x30,
	0x10, 0xc7, 0xa3, 0x3a, 0x1f, 0xed, 0xa4, 0x90, 0x22, 0x42, 0x63, 0x7c, 0x70, 0x83, 0x4b, 0x43,
	0xa0, 0xd4, 0xa6, 0xc9, 0xa1, 0xd7, 0x7e, 0x84, 0x42, 0x0b, 0x86, 0xe2, 0x1e, 0x5a, 0x7a, 0x69,
	0x15, 0x5b, 0xb1, 0x05, 0xb1, 0x65, 0x2c, 0xb5, 0x49, 0xa1, 0x0f, 0xd1, 0x17, 0xda, 0xfb, 0x1e,
	0xc3, 0x9e, 0xf6, 0xb8, 0x24, 0x2f, 0xb2, 0x58, 0xb6, 0x13, 0xf2, 0xb1, 0x10, 0xf6, 0x34, 0x33,
	0x9a, 0xdf, 0x8c, 0xe6, 0x3f, 0x42, 0x80, 0x85, 0x1f, 0xd1, 0x98, 0x38, 0x7f, 0x5e, 0x3b, 0x72,
	0x69, 0xa7, 0x19, 0x97, 0x1c, 0x63, 0xc1, 0x93, 0x8c, 0x71, 0x3b, 0x37, 0x76, 0x91, 0x37, 0x9e,
	0xee, 0xb8, 0xc2, 0x2b, 0x58, 0xa3, 0xb7, 0x3b, 0x5f, 0x44, 0x44, 0xfe, 0x64, 0xa2, 0x48, 0x58,
	0x17, 0x08, 0x3a, 0xae, 0x08, 0x3f, 0x64, 0x94, 0x48, 0xfa, 0x55, 0x41, 0x58, 0x87, 0x96, 0x9f,
	0xc7, 0x3c, 0xd3, 0x51, 0x1f, 0x0d, 0x1f, 0x79, 0x55, 0x88, 0xbb, 0xd0, 0x98, 0x93, 0x29, 0x9d,
	0xeb, 0x0f, 0xd4, 0x79, 0x11, 0xe0, 0x37, 0xd0, 0x9c, 0x31, 0x3a, 0x0f, 0x84, 0xae, 0xf5, 0xb5,
	0x61, 0x7b, 0xf4, 0xcc, 0x3e, 0x9e, 0xcc, 0x2e, 0x7a, 0x7f, 0xcc, 0x39, 0xaf, 0xc4, 0xf1, 0x3b,
	0x78, 0x18, 0x53, 0x49, 0x02, 0x22, 0x89, 0x5e, 0x57, 0xa5, 0x2f, 0x4e, 0x95, 0xba, 0x25, 0x33,
	0xa1, 0x33, 0x96, 0x48, 0xc6, 0x13, 0x6f, 0x5b, 0x66, 0xfd, 0x83, 0xde, 0xc1, 0xf8, 0x1e, 0x15,
	0x29, 0x4f, 0x04, 0xc5, 0x18, 0xea, 0x3e, 0x0f, 0xa8, 0xd2, 0xd0, 0xf0, 0x94, 0x9f, 0x4b, 0x8b,
	0xa9, 0x10, 0x24, 0xa4, 0xa5, 0x84, 0x2a, 0xc4, 0x63, 0x68, 0x95, 0x9b, 0xd1, 0xb5, 0x3e, 0x1a,
	0xb6, 0x47, 0xc6, 0xa9, 0x51, 0xbe, 0x45, 0x44, 0x7e, 0x12, 0x5e, 0x73, 0xa1, 0xac, 0xf5, 0x16,
	0xb0, 0x2b, 0xc2, 0x09, 0x4d, 0x33, 0xea, 0x9f, 0xb3, 0xbf, 0x27, 0xa0, 0x05, 0x2c, 0x28, 0xaf,
	0xce, 0x5d, 0xeb, 0x33, 0x18, 0xc7, 0x1d, 0xee, 0x27, 0x61, 0x74, 0x85, 0x40, 0x73, 0x45, 0x88,
	0x7f, 0xc1, 0xe3, 0xbd, 0xf7, 0x7c, 0x7e, 0x72, 0xa9, 0xfb, 0x5b, 0x33, 0x5e, 0x9e, 0x01, 0x6d,
	0xe7, 0x62, 0xd0, 0x39, 0x14, 0x3d, 0xb8, 0xa3, 0xfe, 0x80, 0x33, 0xec, 0xf3, 0xb8, 0xea, 0xaa,
	0xf7, 0xdf, 0x2f, 0xd7, 0x26, 0x5a, 0xad, 0x4d, 0x74, 0xb3, 0x36, 0xd1, 0xff, 0x8d, 0x59, 0x5b,
	0x6d, 0xcc, 0xda, 0xf5, 0xc6, 0xac, 0x41, 0xb7, 0x6a, 0x22, 0xff, 0xa6, 0x54, 0x94, 0xad, 0xbe,
	0xa0, 0x1f, 0x83, 0x90, 0xc9, 0xe8, 0xf7, 0xd4, 0xf6, 0x79, 0xec, 0xe4, 0xf9, 0x57, 0x8c, 0x2b,
	0xeb, 0x2c, 0xcb, 0x3f, 0xe1, 0xa8, 0x82, 0x69, 0x53, 0xfd, 0x80, 0xf1, 0x6d, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xca, 0x41, 0x63, 0x55, 0x5c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateSchema defines a new on-chain schema.
	CreateSchema(ctx context.Context, in *MsgCreateSchema, opts ...grpc.CallOption) (*MsgCreateSchemaResponse, error)
	// DeprecateSchema deactivates a schema.
	DeprecateSchema(ctx context.Context, in *MsgDeprecateSchema, opts ...grpc.CallOption) (*MsgDeprecateSchemaResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateSchema(ctx context.Context, in *MsgCreateSchema, opts ...grpc.CallOption) (*MsgCreateSchemaResponse, error) {
	out := new(MsgCreateSchemaResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.schema.Msg/CreateSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeprecateSchema(ctx context.Context, in *MsgDeprecateSchema, opts ...grpc.CallOption) (*MsgDeprecateSchemaResponse, error) {
	out := new(MsgDeprecateSchemaResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.schema.Msg/DeprecateSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateSchema defines a new on-chain schema.
	CreateSchema(context.Context, *MsgCreateSchema) (*MsgCreateSchemaResponse, error)
	// DeprecateSchema deactivates a schema.
	DeprecateSchema(context.Context, *MsgDeprecateSchema) (*MsgDeprecateSchemaResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateSchema(ctx context.Context, req *MsgCreateSchema) (*MsgCreateSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchema not implemented")
}
func (*UnimplementedMsgServer) DeprecateSchema(ctx context.Context, req *MsgDeprecateSchema) (*MsgDeprecateSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeprecateSchema not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrio.sonr.schema.Msg/CreateSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateSchema(ctx, req.(*MsgCreateSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeprecateSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeprecateSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeprecateSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrio.sonr.schema.Msg/DeprecateSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeprecateSchema(ctx, req.(*MsgDeprecateSchema))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonrio.sonr.schema.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSchema",
			Handler:    _Msg_CreateSchema_Handler,
		},
		{
			MethodName: "DeprecateSchema",
			Handler:    _Msg_DeprecateSchema_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schema/v1/tx.proto",
}

func (m *MsgCreateSchema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateSchema) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateSchema) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		for iNdEx := len(m.Metadata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Metadata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Fields) > 0 {
		for iNdEx := len(m.Fields) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fields[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateSchemaResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateSchemaResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateSchemaResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WhatIs != nil {
		{
			size, err := m.WhatIs.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeprecateSchema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeprecateSchema) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeprecateSchema) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeprecateSchemaResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeprecateSchemaResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeprecateSchemaResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateSchema) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Fields) > 0 {
		for _, e := range m.Fields {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.Metadata) > 0 {
		for _, e := range m.Metadata {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCreateSchemaResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovTx(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.WhatIs != nil {
		l = m.WhatIs.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeprecateSchema) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeprecateSchemaResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovTx(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateSchema) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateSchema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateSchema: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fields", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fields = append(m.Fields, &SchemaField{})
			if err := m.Fields[len(m.Fields)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = append(m.Metadata, &MetadataDefintion{})
			if err := m.Metadata[len(m.Metadata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateSchemaResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateSchemaResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateSchemaResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhatIs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.WhatIs == nil {
				m.WhatIs = &WhatIs{}
			}
			if err := m.WhatIs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeprecateSchema) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeprecateSchema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeprecateSchema: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeprecateSchemaResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeprecateSchemaResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeprecateSchemaResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
