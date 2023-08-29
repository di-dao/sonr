// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/domain/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgCreateUsernameRecords struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index   string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Method  string `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
}

func (m *MsgCreateUsernameRecords) Reset()         { *m = MsgCreateUsernameRecords{} }
func (m *MsgCreateUsernameRecords) String() string { return proto.CompactTextString(m) }
func (*MsgCreateUsernameRecords) ProtoMessage()    {}
func (*MsgCreateUsernameRecords) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e2442c81dc82e6, []int{0}
}
func (m *MsgCreateUsernameRecords) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateUsernameRecords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateUsernameRecords.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateUsernameRecords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateUsernameRecords.Merge(m, src)
}
func (m *MsgCreateUsernameRecords) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateUsernameRecords) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateUsernameRecords.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateUsernameRecords proto.InternalMessageInfo

func (m *MsgCreateUsernameRecords) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateUsernameRecords) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *MsgCreateUsernameRecords) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

type MsgCreateUsernameRecordsResponse struct {
}

func (m *MsgCreateUsernameRecordsResponse) Reset()         { *m = MsgCreateUsernameRecordsResponse{} }
func (m *MsgCreateUsernameRecordsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateUsernameRecordsResponse) ProtoMessage()    {}
func (*MsgCreateUsernameRecordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e2442c81dc82e6, []int{1}
}
func (m *MsgCreateUsernameRecordsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateUsernameRecordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateUsernameRecordsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateUsernameRecordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateUsernameRecordsResponse.Merge(m, src)
}
func (m *MsgCreateUsernameRecordsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateUsernameRecordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateUsernameRecordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateUsernameRecordsResponse proto.InternalMessageInfo

type MsgUpdateUsernameRecords struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index   string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *MsgUpdateUsernameRecords) Reset()         { *m = MsgUpdateUsernameRecords{} }
func (m *MsgUpdateUsernameRecords) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateUsernameRecords) ProtoMessage()    {}
func (*MsgUpdateUsernameRecords) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e2442c81dc82e6, []int{2}
}
func (m *MsgUpdateUsernameRecords) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateUsernameRecords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateUsernameRecords.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateUsernameRecords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateUsernameRecords.Merge(m, src)
}
func (m *MsgUpdateUsernameRecords) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateUsernameRecords) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateUsernameRecords.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateUsernameRecords proto.InternalMessageInfo

func (m *MsgUpdateUsernameRecords) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateUsernameRecords) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type MsgUpdateUsernameRecordsResponse struct {
}

func (m *MsgUpdateUsernameRecordsResponse) Reset()         { *m = MsgUpdateUsernameRecordsResponse{} }
func (m *MsgUpdateUsernameRecordsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateUsernameRecordsResponse) ProtoMessage()    {}
func (*MsgUpdateUsernameRecordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e2442c81dc82e6, []int{3}
}
func (m *MsgUpdateUsernameRecordsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateUsernameRecordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateUsernameRecordsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateUsernameRecordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateUsernameRecordsResponse.Merge(m, src)
}
func (m *MsgUpdateUsernameRecordsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateUsernameRecordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateUsernameRecordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateUsernameRecordsResponse proto.InternalMessageInfo

type MsgDeleteUsernameRecords struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index   string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *MsgDeleteUsernameRecords) Reset()         { *m = MsgDeleteUsernameRecords{} }
func (m *MsgDeleteUsernameRecords) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteUsernameRecords) ProtoMessage()    {}
func (*MsgDeleteUsernameRecords) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e2442c81dc82e6, []int{4}
}
func (m *MsgDeleteUsernameRecords) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteUsernameRecords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteUsernameRecords.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteUsernameRecords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteUsernameRecords.Merge(m, src)
}
func (m *MsgDeleteUsernameRecords) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteUsernameRecords) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteUsernameRecords.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteUsernameRecords proto.InternalMessageInfo

func (m *MsgDeleteUsernameRecords) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeleteUsernameRecords) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type MsgDeleteUsernameRecordsResponse struct {
}

func (m *MsgDeleteUsernameRecordsResponse) Reset()         { *m = MsgDeleteUsernameRecordsResponse{} }
func (m *MsgDeleteUsernameRecordsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteUsernameRecordsResponse) ProtoMessage()    {}
func (*MsgDeleteUsernameRecordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e2442c81dc82e6, []int{5}
}
func (m *MsgDeleteUsernameRecordsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteUsernameRecordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteUsernameRecordsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteUsernameRecordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteUsernameRecordsResponse.Merge(m, src)
}
func (m *MsgDeleteUsernameRecordsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteUsernameRecordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteUsernameRecordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteUsernameRecordsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateUsernameRecords)(nil), "core.domain.MsgCreateUsernameRecords")
	proto.RegisterType((*MsgCreateUsernameRecordsResponse)(nil), "core.domain.MsgCreateUsernameRecordsResponse")
	proto.RegisterType((*MsgUpdateUsernameRecords)(nil), "core.domain.MsgUpdateUsernameRecords")
	proto.RegisterType((*MsgUpdateUsernameRecordsResponse)(nil), "core.domain.MsgUpdateUsernameRecordsResponse")
	proto.RegisterType((*MsgDeleteUsernameRecords)(nil), "core.domain.MsgDeleteUsernameRecords")
	proto.RegisterType((*MsgDeleteUsernameRecordsResponse)(nil), "core.domain.MsgDeleteUsernameRecordsResponse")
}

func init() { proto.RegisterFile("core/domain/tx.proto", fileDescriptor_d1e2442c81dc82e6) }

var fileDescriptor_d1e2442c81dc82e6 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xce, 0x2f, 0x4a,
	0xd5, 0x4f, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0xd3, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x06, 0x89, 0xea, 0x41, 0x44, 0xa5, 0x24, 0x90, 0x95, 0x14, 0xa5, 0x26, 0xe7, 0x17,
	0xa5, 0x40, 0x94, 0x29, 0x25, 0x71, 0x49, 0xf8, 0x16, 0xa7, 0x3b, 0x17, 0xa5, 0x26, 0x96, 0xa4,
	0x86, 0x16, 0xa7, 0x16, 0xe5, 0x25, 0xe6, 0xa6, 0x06, 0x81, 0x15, 0x14, 0x0b, 0x49, 0x70, 0xb1,
	0x27, 0x83, 0x24, 0xf2, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x11,
	0x2e, 0xd6, 0xcc, 0xbc, 0x94, 0xd4, 0x0a, 0x09, 0x26, 0xb0, 0x38, 0x84, 0x23, 0x24, 0xc6, 0xc5,
	0x96, 0x9b, 0x5a, 0x92, 0x91, 0x9f, 0x22, 0xc1, 0x0c, 0x16, 0x86, 0xf2, 0x94, 0x94, 0xb8, 0x14,
	0x70, 0xd9, 0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0xaa, 0xe4, 0x05, 0x76, 0x47, 0x68,
	0x41, 0x0a, 0xe5, 0xee, 0x80, 0xda, 0x87, 0xd5, 0x2c, 0x34, 0xfb, 0x5c, 0x52, 0x73, 0x52, 0xa9,
	0x65, 0x1f, 0x56, 0xb3, 0x60, 0xf6, 0x19, 0x9d, 0x67, 0xe2, 0x62, 0xf6, 0x2d, 0x4e, 0x17, 0xca,
	0xe1, 0x12, 0xc1, 0x16, 0x10, 0x42, 0xaa, 0x7a, 0x48, 0xf1, 0xa5, 0x87, 0x2b, 0xb8, 0xa4, 0x74,
	0x89, 0x52, 0x06, 0xb3, 0x15, 0x64, 0x1b, 0xb6, 0x60, 0xc0, 0xb4, 0x0d, 0x6b, 0x60, 0x61, 0xda,
	0x86, 0x37, 0x4c, 0x41, 0xb6, 0x61, 0x0b, 0x04, 0x4c, 0xdb, 0xb0, 0x06, 0x15, 0xa6, 0x6d, 0x78,
	0x43, 0xd4, 0xc9, 0xfe, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63,
	0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x54, 0xd3,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x8b, 0xf3, 0xf3, 0x8a, 0x32, 0x0a,
	0xf5, 0xc1, 0xe9, 0xbf, 0x02, 0x9e, 0x49, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x39, 0xc0,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x18, 0x05, 0x6b, 0x40, 0x03, 0x00, 0x00,
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
	CreateUsernameRecord(ctx context.Context, in *MsgCreateUsernameRecords, opts ...grpc.CallOption) (*MsgCreateUsernameRecordsResponse, error)
	UpdateUsernameRecord(ctx context.Context, in *MsgUpdateUsernameRecords, opts ...grpc.CallOption) (*MsgUpdateUsernameRecordsResponse, error)
	DeleteUsernameRecord(ctx context.Context, in *MsgDeleteUsernameRecords, opts ...grpc.CallOption) (*MsgDeleteUsernameRecordsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateUsernameRecord(ctx context.Context, in *MsgCreateUsernameRecords, opts ...grpc.CallOption) (*MsgCreateUsernameRecordsResponse, error) {
	out := new(MsgCreateUsernameRecordsResponse)
	err := c.cc.Invoke(ctx, "/core.domain.Msg/CreateUsernameRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateUsernameRecord(ctx context.Context, in *MsgUpdateUsernameRecords, opts ...grpc.CallOption) (*MsgUpdateUsernameRecordsResponse, error) {
	out := new(MsgUpdateUsernameRecordsResponse)
	err := c.cc.Invoke(ctx, "/core.domain.Msg/UpdateUsernameRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteUsernameRecord(ctx context.Context, in *MsgDeleteUsernameRecords, opts ...grpc.CallOption) (*MsgDeleteUsernameRecordsResponse, error) {
	out := new(MsgDeleteUsernameRecordsResponse)
	err := c.cc.Invoke(ctx, "/core.domain.Msg/DeleteUsernameRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateUsernameRecord(context.Context, *MsgCreateUsernameRecords) (*MsgCreateUsernameRecordsResponse, error)
	UpdateUsernameRecord(context.Context, *MsgUpdateUsernameRecords) (*MsgUpdateUsernameRecordsResponse, error)
	DeleteUsernameRecord(context.Context, *MsgDeleteUsernameRecords) (*MsgDeleteUsernameRecordsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateUsernameRecord(ctx context.Context, req *MsgCreateUsernameRecords) (*MsgCreateUsernameRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUsernameRecord not implemented")
}
func (*UnimplementedMsgServer) UpdateUsernameRecord(ctx context.Context, req *MsgUpdateUsernameRecords) (*MsgUpdateUsernameRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUsernameRecord not implemented")
}
func (*UnimplementedMsgServer) DeleteUsernameRecord(ctx context.Context, req *MsgDeleteUsernameRecords) (*MsgDeleteUsernameRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUsernameRecord not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateUsernameRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateUsernameRecords)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateUsernameRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.domain.Msg/CreateUsernameRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateUsernameRecord(ctx, req.(*MsgCreateUsernameRecords))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateUsernameRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateUsernameRecords)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateUsernameRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.domain.Msg/UpdateUsernameRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateUsernameRecord(ctx, req.(*MsgUpdateUsernameRecords))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteUsernameRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteUsernameRecords)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteUsernameRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.domain.Msg/DeleteUsernameRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteUsernameRecord(ctx, req.(*MsgDeleteUsernameRecords))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "core.domain.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUsernameRecord",
			Handler:    _Msg_CreateUsernameRecord_Handler,
		},
		{
			MethodName: "UpdateUsernameRecord",
			Handler:    _Msg_UpdateUsernameRecord_Handler,
		},
		{
			MethodName: "DeleteUsernameRecord",
			Handler:    _Msg_DeleteUsernameRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/domain/tx.proto",
}

func (m *MsgCreateUsernameRecords) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateUsernameRecords) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateUsernameRecords) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Method) > 0 {
		i -= len(m.Method)
		copy(dAtA[i:], m.Method)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Method)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Index)))
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

func (m *MsgCreateUsernameRecordsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateUsernameRecordsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateUsernameRecordsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateUsernameRecords) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateUsernameRecords) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateUsernameRecords) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Index)))
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

func (m *MsgUpdateUsernameRecordsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateUsernameRecordsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateUsernameRecordsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteUsernameRecords) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteUsernameRecords) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteUsernameRecords) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Index)))
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

func (m *MsgDeleteUsernameRecordsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteUsernameRecordsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteUsernameRecordsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCreateUsernameRecords) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateUsernameRecordsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateUsernameRecords) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateUsernameRecordsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteUsernameRecords) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeleteUsernameRecordsResponse) Size() (n int) {
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
func (m *MsgCreateUsernameRecords) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateUsernameRecords: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateUsernameRecords: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
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
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
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
			m.Method = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCreateUsernameRecordsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateUsernameRecordsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateUsernameRecordsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgUpdateUsernameRecords) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateUsernameRecords: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateUsernameRecords: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
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
			m.Index = string(dAtA[iNdEx:postIndex])
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
func (m *MsgUpdateUsernameRecordsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateUsernameRecordsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateUsernameRecordsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgDeleteUsernameRecords) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeleteUsernameRecords: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteUsernameRecords: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
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
			m.Index = string(dAtA[iNdEx:postIndex])
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
func (m *MsgDeleteUsernameRecordsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeleteUsernameRecordsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteUsernameRecordsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
