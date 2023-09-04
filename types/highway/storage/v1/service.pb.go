// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: highway/storage/v1/service.proto

package storagepb

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetCIDRequest struct {
	Cid string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (m *GetCIDRequest) Reset()         { *m = GetCIDRequest{} }
func (m *GetCIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetCIDRequest) ProtoMessage()    {}
func (*GetCIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_913be91a9864b42e, []int{0}
}
func (m *GetCIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetCIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetCIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetCIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCIDRequest.Merge(m, src)
}
func (m *GetCIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetCIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCIDRequest proto.InternalMessageInfo

func (m *GetCIDRequest) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

type GetCIDResponse struct {
	Cid  string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *GetCIDResponse) Reset()         { *m = GetCIDResponse{} }
func (m *GetCIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetCIDResponse) ProtoMessage()    {}
func (*GetCIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_913be91a9864b42e, []int{1}
}
func (m *GetCIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetCIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetCIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetCIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCIDResponse.Merge(m, src)
}
func (m *GetCIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetCIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCIDResponse proto.InternalMessageInfo

func (m *GetCIDResponse) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *GetCIDResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type PutDataRequest struct {
	Cid  string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PutDataRequest) Reset()         { *m = PutDataRequest{} }
func (m *PutDataRequest) String() string { return proto.CompactTextString(m) }
func (*PutDataRequest) ProtoMessage()    {}
func (*PutDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_913be91a9864b42e, []int{2}
}
func (m *PutDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PutDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PutDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PutDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutDataRequest.Merge(m, src)
}
func (m *PutDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *PutDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutDataRequest proto.InternalMessageInfo

func (m *PutDataRequest) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *PutDataRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type PutDataResponse struct {
	Cid string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (m *PutDataResponse) Reset()         { *m = PutDataResponse{} }
func (m *PutDataResponse) String() string { return proto.CompactTextString(m) }
func (*PutDataResponse) ProtoMessage()    {}
func (*PutDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_913be91a9864b42e, []int{3}
}
func (m *PutDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PutDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PutDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PutDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutDataResponse.Merge(m, src)
}
func (m *PutDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *PutDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutDataResponse proto.InternalMessageInfo

func (m *PutDataResponse) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCIDRequest)(nil), "highway.storage.v1.GetCIDRequest")
	proto.RegisterType((*GetCIDResponse)(nil), "highway.storage.v1.GetCIDResponse")
	proto.RegisterType((*PutDataRequest)(nil), "highway.storage.v1.PutDataRequest")
	proto.RegisterType((*PutDataResponse)(nil), "highway.storage.v1.PutDataResponse")
}

func init() { proto.RegisterFile("highway/storage/v1/service.proto", fileDescriptor_913be91a9864b42e) }

var fileDescriptor_913be91a9864b42e = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x97, 0x29, 0x13, 0x03, 0xdb, 0x24, 0x20, 0x8c, 0x21, 0x61, 0xeb, 0x2e, 0x9e, 0x12,
	0xa6, 0x20, 0x88, 0x37, 0x1d, 0x88, 0x37, 0x59, 0x6f, 0xde, 0xd2, 0x34, 0xb6, 0x01, 0x6d, 0xba,
	0x26, 0xad, 0x0c, 0xf1, 0xe2, 0x27, 0x10, 0xfc, 0x52, 0x1e, 0x07, 0x5e, 0x3c, 0x4a, 0xeb, 0x07,
	0x91, 0xfe, 0x51, 0x28, 0xdd, 0x76, 0x7b, 0x29, 0xbf, 0xe7, 0x7d, 0xde, 0xe7, 0x69, 0xe0, 0xc8,
	0x97, 0x9e, 0xff, 0xc4, 0x96, 0x54, 0x1b, 0x15, 0x31, 0x4f, 0xd0, 0x64, 0x4a, 0xb5, 0x88, 0x12,
	0xc9, 0x05, 0x09, 0x23, 0x65, 0x14, 0x42, 0x15, 0x41, 0x2a, 0x82, 0x24, 0xd3, 0xe1, 0x91, 0xa7,
	0x94, 0xf7, 0x20, 0x28, 0x0b, 0x25, 0x65, 0x41, 0xa0, 0x0c, 0x33, 0x52, 0x05, 0xba, 0x54, 0x58,
	0x63, 0xd8, 0xbd, 0x16, 0xe6, 0xea, 0x66, 0x36, 0x17, 0x8b, 0x58, 0x68, 0x83, 0x0e, 0xe0, 0x0e,
	0x97, 0xee, 0x00, 0x8c, 0xc0, 0xf1, 0xfe, 0x3c, 0x1f, 0xad, 0x33, 0xd8, 0xfb, 0x43, 0x74, 0xa8,
	0x02, 0x2d, 0x9a, 0x0c, 0x42, 0x70, 0xd7, 0x65, 0x86, 0x0d, 0xda, 0xc5, 0xa7, 0x62, 0xce, 0x75,
	0xb7, 0xb1, 0x99, 0x31, 0xc3, 0x36, 0xee, 0x5e, 0xab, 0x9b, 0xc0, 0xfe, 0xbf, 0x6e, 0x93, 0xe1,
	0x49, 0x06, 0x60, 0xcf, 0x2e, 0x43, 0xda, 0x65, 0x05, 0xe8, 0x1e, 0x76, 0xca, 0x3b, 0xd1, 0x98,
	0x34, 0x7b, 0x20, 0xb5, 0x98, 0x43, 0x6b, 0x1b, 0x52, 0xba, 0x5a, 0x87, 0xaf, 0x9f, 0x3f, 0xef,
	0xed, 0x3e, 0xea, 0xe6, 0x45, 0x73, 0xe9, 0xd2, 0x67, 0x2e, 0xdd, 0x17, 0x24, 0xe1, 0x5e, 0x75,
	0x1f, 0x5a, 0xbb, 0xa5, 0x1e, 0x7a, 0x38, 0xd9, 0xca, 0xd4, 0xad, 0xac, 0xba, 0xd5, 0xa5, 0xfd,
	0x91, 0x62, 0xb0, 0x4a, 0x31, 0xf8, 0x4e, 0x31, 0x78, 0xcb, 0x70, 0x6b, 0x95, 0xe1, 0xd6, 0x57,
	0x86, 0x5b, 0x77, 0xe7, 0x9e, 0x34, 0x7e, 0xec, 0x10, 0xae, 0x1e, 0xa9, 0x56, 0x41, 0xe4, 0x2f,
	0x28, 0x57, 0x91, 0xa0, 0x66, 0x19, 0x0a, 0x4d, 0x9b, 0x0f, 0xe5, 0xa2, 0x1a, 0x43, 0xc7, 0xe9,
	0x14, 0x7f, 0xfe, 0xf4, 0x37, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x0d, 0x80, 0xb4, 0x4f, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageServiceClient interface {
	GetCID(ctx context.Context, in *GetCIDRequest, opts ...grpc.CallOption) (*GetCIDResponse, error)
	PutData(ctx context.Context, in *PutDataRequest, opts ...grpc.CallOption) (*PutDataResponse, error)
}

type storageServiceClient struct {
	cc grpc1.ClientConn
}

func NewStorageServiceClient(cc grpc1.ClientConn) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) GetCID(ctx context.Context, in *GetCIDRequest, opts ...grpc.CallOption) (*GetCIDResponse, error) {
	out := new(GetCIDResponse)
	err := c.cc.Invoke(ctx, "/highway.storage.v1.StorageService/GetCID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) PutData(ctx context.Context, in *PutDataRequest, opts ...grpc.CallOption) (*PutDataResponse, error) {
	out := new(PutDataResponse)
	err := c.cc.Invoke(ctx, "/highway.storage.v1.StorageService/PutData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServiceServer is the server API for StorageService service.
type StorageServiceServer interface {
	GetCID(context.Context, *GetCIDRequest) (*GetCIDResponse, error)
	PutData(context.Context, *PutDataRequest) (*PutDataResponse, error)
}

// UnimplementedStorageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStorageServiceServer struct {
}

func (*UnimplementedStorageServiceServer) GetCID(ctx context.Context, req *GetCIDRequest) (*GetCIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCID not implemented")
}
func (*UnimplementedStorageServiceServer) PutData(ctx context.Context, req *PutDataRequest) (*PutDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutData not implemented")
}

func RegisterStorageServiceServer(s grpc1.Server, srv StorageServiceServer) {
	s.RegisterService(&_StorageService_serviceDesc, srv)
}

func _StorageService_GetCID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetCID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/highway.storage.v1.StorageService/GetCID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetCID(ctx, req.(*GetCIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_PutData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).PutData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/highway.storage.v1.StorageService/PutData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).PutData(ctx, req.(*PutDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StorageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "highway.storage.v1.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCID",
			Handler:    _StorageService_GetCID_Handler,
		},
		{
			MethodName: "PutData",
			Handler:    _StorageService_PutData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "highway/storage/v1/service.proto",
}

func (m *GetCIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetCIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetCIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintService(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetCIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetCIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetCIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintService(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintService(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PutDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PutDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PutDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintService(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintService(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PutDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PutDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PutDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintService(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetCIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *GetCIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *PutDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *PutDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetCIDRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: GetCIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *GetCIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: GetCIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *PutDataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: PutDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PutDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *PutDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: PutDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PutDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)
