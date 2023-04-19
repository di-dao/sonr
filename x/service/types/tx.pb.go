// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/service/tx.proto

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

type MsgCreateServiceRecord struct {
	Controller string         `protobuf:"bytes,1,opt,name=controller,proto3" json:"controller,omitempty"`
	Record     *ServiceRecord `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
}

func (m *MsgCreateServiceRecord) Reset()         { *m = MsgCreateServiceRecord{} }
func (m *MsgCreateServiceRecord) String() string { return proto.CompactTextString(m) }
func (*MsgCreateServiceRecord) ProtoMessage()    {}
func (*MsgCreateServiceRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_376dd44ebb86aa2f, []int{0}
}
func (m *MsgCreateServiceRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateServiceRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateServiceRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateServiceRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateServiceRecord.Merge(m, src)
}
func (m *MsgCreateServiceRecord) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateServiceRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateServiceRecord.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateServiceRecord proto.InternalMessageInfo

func (m *MsgCreateServiceRecord) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func (m *MsgCreateServiceRecord) GetRecord() *ServiceRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

type MsgCreateServiceRecordResponse struct {
}

func (m *MsgCreateServiceRecordResponse) Reset()         { *m = MsgCreateServiceRecordResponse{} }
func (m *MsgCreateServiceRecordResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateServiceRecordResponse) ProtoMessage()    {}
func (*MsgCreateServiceRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_376dd44ebb86aa2f, []int{1}
}
func (m *MsgCreateServiceRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateServiceRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateServiceRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateServiceRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateServiceRecordResponse.Merge(m, src)
}
func (m *MsgCreateServiceRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateServiceRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateServiceRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateServiceRecordResponse proto.InternalMessageInfo

type MsgUpdateServiceRecord struct {
	Controller string         `protobuf:"bytes,1,opt,name=controller,proto3" json:"controller,omitempty"`
	Record     *ServiceRecord `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
}

func (m *MsgUpdateServiceRecord) Reset()         { *m = MsgUpdateServiceRecord{} }
func (m *MsgUpdateServiceRecord) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateServiceRecord) ProtoMessage()    {}
func (*MsgUpdateServiceRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_376dd44ebb86aa2f, []int{2}
}
func (m *MsgUpdateServiceRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateServiceRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateServiceRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateServiceRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateServiceRecord.Merge(m, src)
}
func (m *MsgUpdateServiceRecord) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateServiceRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateServiceRecord.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateServiceRecord proto.InternalMessageInfo

func (m *MsgUpdateServiceRecord) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func (m *MsgUpdateServiceRecord) GetRecord() *ServiceRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

type MsgUpdateServiceRecordResponse struct {
}

func (m *MsgUpdateServiceRecordResponse) Reset()         { *m = MsgUpdateServiceRecordResponse{} }
func (m *MsgUpdateServiceRecordResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateServiceRecordResponse) ProtoMessage()    {}
func (*MsgUpdateServiceRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_376dd44ebb86aa2f, []int{3}
}
func (m *MsgUpdateServiceRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateServiceRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateServiceRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateServiceRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateServiceRecordResponse.Merge(m, src)
}
func (m *MsgUpdateServiceRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateServiceRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateServiceRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateServiceRecordResponse proto.InternalMessageInfo

type MsgDeleteServiceRecord struct {
	Controller string `protobuf:"bytes,1,opt,name=controller,proto3" json:"controller,omitempty"`
	Id         string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgDeleteServiceRecord) Reset()         { *m = MsgDeleteServiceRecord{} }
func (m *MsgDeleteServiceRecord) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteServiceRecord) ProtoMessage()    {}
func (*MsgDeleteServiceRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_376dd44ebb86aa2f, []int{4}
}
func (m *MsgDeleteServiceRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteServiceRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteServiceRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteServiceRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteServiceRecord.Merge(m, src)
}
func (m *MsgDeleteServiceRecord) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteServiceRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteServiceRecord.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteServiceRecord proto.InternalMessageInfo

func (m *MsgDeleteServiceRecord) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func (m *MsgDeleteServiceRecord) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MsgDeleteServiceRecordResponse struct {
}

func (m *MsgDeleteServiceRecordResponse) Reset()         { *m = MsgDeleteServiceRecordResponse{} }
func (m *MsgDeleteServiceRecordResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteServiceRecordResponse) ProtoMessage()    {}
func (*MsgDeleteServiceRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_376dd44ebb86aa2f, []int{5}
}
func (m *MsgDeleteServiceRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteServiceRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteServiceRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteServiceRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteServiceRecordResponse.Merge(m, src)
}
func (m *MsgDeleteServiceRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteServiceRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteServiceRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteServiceRecordResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateServiceRecord)(nil), "sonrhq.core.service.MsgCreateServiceRecord")
	proto.RegisterType((*MsgCreateServiceRecordResponse)(nil), "sonrhq.core.service.MsgCreateServiceRecordResponse")
	proto.RegisterType((*MsgUpdateServiceRecord)(nil), "sonrhq.core.service.MsgUpdateServiceRecord")
	proto.RegisterType((*MsgUpdateServiceRecordResponse)(nil), "sonrhq.core.service.MsgUpdateServiceRecordResponse")
	proto.RegisterType((*MsgDeleteServiceRecord)(nil), "sonrhq.core.service.MsgDeleteServiceRecord")
	proto.RegisterType((*MsgDeleteServiceRecordResponse)(nil), "sonrhq.core.service.MsgDeleteServiceRecordResponse")
}

func init() { proto.RegisterFile("core/service/tx.proto", fileDescriptor_376dd44ebb86aa2f) }

var fileDescriptor_376dd44ebb86aa2f = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0x2f, 0x4a,
	0xd5, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x2e, 0xce, 0xcf, 0x2b, 0xca, 0x28, 0xd4, 0x03, 0xc9, 0xea, 0x41, 0x65, 0xa5,
	0x24, 0x51, 0xd4, 0x16, 0xa5, 0x26, 0xe7, 0x17, 0xa5, 0x40, 0xd4, 0x2b, 0x95, 0x70, 0x89, 0xf9,
	0x16, 0xa7, 0x3b, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0x06, 0x43, 0x14, 0x04, 0x81, 0xe5, 0x85, 0xe4,
	0xb8, 0xb8, 0x92, 0xf3, 0xf3, 0x4a, 0x8a, 0xf2, 0x73, 0x72, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x90, 0x44, 0x84, 0xac, 0xb8, 0xd8, 0x20, 0x26, 0x49, 0x30, 0x29, 0x30, 0x6a,
	0x70, 0x1b, 0x29, 0xe9, 0x61, 0xb1, 0x5a, 0x0f, 0xc5, 0xcc, 0x20, 0xa8, 0x0e, 0x25, 0x05, 0x2e,
	0x39, 0xec, 0xb6, 0x06, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x42, 0xdd, 0x15, 0x5a, 0x90,
	0x32, 0x00, 0xee, 0xc2, 0x62, 0x2b, 0xdc, 0x5d, 0x1e, 0x60, 0x77, 0xb9, 0xa4, 0xe6, 0xa4, 0x92,
	0xea, 0x2e, 0x3e, 0x2e, 0xa6, 0x4c, 0x88, 0x9b, 0x38, 0x83, 0x98, 0x32, 0x61, 0x76, 0x61, 0x31,
	0x09, 0x66, 0x97, 0xd1, 0x07, 0x26, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0xa1, 0x72, 0x2e, 0x61, 0x6c,
	0x11, 0xa4, 0x8d, 0xd5, 0x63, 0xd8, 0xc3, 0x55, 0xca, 0x98, 0x04, 0xc5, 0x30, 0x07, 0x80, 0x2c,
	0xc6, 0x16, 0x03, 0x38, 0x2d, 0xc6, 0xa2, 0x18, 0xb7, 0xc5, 0x78, 0x42, 0x19, 0x64, 0x31, 0xb6,
	0x20, 0xc6, 0x69, 0x31, 0x16, 0xc5, 0xb8, 0x2d, 0xc6, 0x13, 0xe4, 0x4e, 0x0e, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e,
	0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x96, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97,
	0x9c, 0x9f, 0xab, 0x0f, 0x31, 0x58, 0x1f, 0x9c, 0xab, 0x2a, 0x10, 0x79, 0xb0, 0xb2, 0x20, 0xb5,
	0x38, 0x89, 0x0d, 0x9c, 0xaf, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xa6, 0x89, 0x3b,
	0xa0, 0x03, 0x00, 0x00,
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
	CreateServiceRecord(ctx context.Context, in *MsgCreateServiceRecord, opts ...grpc.CallOption) (*MsgCreateServiceRecordResponse, error)
	UpdateServiceRecord(ctx context.Context, in *MsgUpdateServiceRecord, opts ...grpc.CallOption) (*MsgUpdateServiceRecordResponse, error)
	DeleteServiceRecord(ctx context.Context, in *MsgDeleteServiceRecord, opts ...grpc.CallOption) (*MsgDeleteServiceRecordResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateServiceRecord(ctx context.Context, in *MsgCreateServiceRecord, opts ...grpc.CallOption) (*MsgCreateServiceRecordResponse, error) {
	out := new(MsgCreateServiceRecordResponse)
	err := c.cc.Invoke(ctx, "/sonrhq.core.service.Msg/CreateServiceRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateServiceRecord(ctx context.Context, in *MsgUpdateServiceRecord, opts ...grpc.CallOption) (*MsgUpdateServiceRecordResponse, error) {
	out := new(MsgUpdateServiceRecordResponse)
	err := c.cc.Invoke(ctx, "/sonrhq.core.service.Msg/UpdateServiceRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteServiceRecord(ctx context.Context, in *MsgDeleteServiceRecord, opts ...grpc.CallOption) (*MsgDeleteServiceRecordResponse, error) {
	out := new(MsgDeleteServiceRecordResponse)
	err := c.cc.Invoke(ctx, "/sonrhq.core.service.Msg/DeleteServiceRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateServiceRecord(context.Context, *MsgCreateServiceRecord) (*MsgCreateServiceRecordResponse, error)
	UpdateServiceRecord(context.Context, *MsgUpdateServiceRecord) (*MsgUpdateServiceRecordResponse, error)
	DeleteServiceRecord(context.Context, *MsgDeleteServiceRecord) (*MsgDeleteServiceRecordResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateServiceRecord(ctx context.Context, req *MsgCreateServiceRecord) (*MsgCreateServiceRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServiceRecord not implemented")
}
func (*UnimplementedMsgServer) UpdateServiceRecord(ctx context.Context, req *MsgUpdateServiceRecord) (*MsgUpdateServiceRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServiceRecord not implemented")
}
func (*UnimplementedMsgServer) DeleteServiceRecord(ctx context.Context, req *MsgDeleteServiceRecord) (*MsgDeleteServiceRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServiceRecord not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateServiceRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateServiceRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateServiceRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrhq.core.service.Msg/CreateServiceRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateServiceRecord(ctx, req.(*MsgCreateServiceRecord))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateServiceRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateServiceRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateServiceRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrhq.core.service.Msg/UpdateServiceRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateServiceRecord(ctx, req.(*MsgUpdateServiceRecord))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteServiceRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteServiceRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteServiceRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrhq.core.service.Msg/DeleteServiceRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteServiceRecord(ctx, req.(*MsgDeleteServiceRecord))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonrhq.core.service.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateServiceRecord",
			Handler:    _Msg_CreateServiceRecord_Handler,
		},
		{
			MethodName: "UpdateServiceRecord",
			Handler:    _Msg_UpdateServiceRecord_Handler,
		},
		{
			MethodName: "DeleteServiceRecord",
			Handler:    _Msg_DeleteServiceRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/service/tx.proto",
}

func (m *MsgCreateServiceRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateServiceRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateServiceRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Record != nil {
		{
			size, err := m.Record.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateServiceRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateServiceRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateServiceRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateServiceRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateServiceRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateServiceRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Record != nil {
		{
			size, err := m.Record.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateServiceRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateServiceRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateServiceRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteServiceRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteServiceRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteServiceRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteServiceRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteServiceRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteServiceRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgCreateServiceRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Record != nil {
		l = m.Record.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateServiceRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateServiceRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Record != nil {
		l = m.Record.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateServiceRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteServiceRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeleteServiceRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateServiceRecord) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateServiceRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateServiceRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
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
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
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
			if m.Record == nil {
				m.Record = &ServiceRecord{}
			}
			if err := m.Record.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgCreateServiceRecordResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateServiceRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateServiceRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *MsgUpdateServiceRecord) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateServiceRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateServiceRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
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
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
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
			if m.Record == nil {
				m.Record = &ServiceRecord{}
			}
			if err := m.Record.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgUpdateServiceRecordResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateServiceRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateServiceRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *MsgDeleteServiceRecord) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeleteServiceRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteServiceRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
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
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
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
func (m *MsgDeleteServiceRecordResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeleteServiceRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteServiceRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
