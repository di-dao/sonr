// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// QueryServiceRecordRequest is the request type for the Query/ServiceRecord RPC method.
type QueryServiceRecordRequest struct {
	// origin defines the origin of the service.
	Origin string `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (m *QueryServiceRecordRequest) Reset()         { *m = QueryServiceRecordRequest{} }
func (m *QueryServiceRecordRequest) String() string { return proto.CompactTextString(m) }
func (*QueryServiceRecordRequest) ProtoMessage()    {}
func (*QueryServiceRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ed234bfe4ffc68, []int{0}
}
func (m *QueryServiceRecordRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryServiceRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryServiceRecordRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryServiceRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryServiceRecordRequest.Merge(m, src)
}
func (m *QueryServiceRecordRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryServiceRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryServiceRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryServiceRecordRequest proto.InternalMessageInfo

func (m *QueryServiceRecordRequest) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

// QueryServiceRecordResponse is the response type for the Query/ServiceRecord RPC method.
type QueryServiceRecordResponse struct {
	// service_record defines the service record for a given origin.
	ServiceRecord Record `protobuf:"bytes,1,opt,name=service_record,json=serviceRecord,proto3" json:"service_record"`
}

func (m *QueryServiceRecordResponse) Reset()         { *m = QueryServiceRecordResponse{} }
func (m *QueryServiceRecordResponse) String() string { return proto.CompactTextString(m) }
func (*QueryServiceRecordResponse) ProtoMessage()    {}
func (*QueryServiceRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ed234bfe4ffc68, []int{1}
}
func (m *QueryServiceRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryServiceRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryServiceRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryServiceRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryServiceRecordResponse.Merge(m, src)
}
func (m *QueryServiceRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryServiceRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryServiceRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryServiceRecordResponse proto.InternalMessageInfo

func (m *QueryServiceRecordResponse) GetServiceRecord() Record {
	if m != nil {
		return m.ServiceRecord
	}
	return Record{}
}

func init() {
	proto.RegisterType((*QueryServiceRecordRequest)(nil), "service.v1.QueryServiceRecordRequest")
	proto.RegisterType((*QueryServiceRecordResponse)(nil), "service.v1.QueryServiceRecordResponse")
}

func init() { proto.RegisterFile("service/v1/query.proto", fileDescriptor_d0ed234bfe4ffc68) }

var fileDescriptor_d0ed234bfe4ffc68 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0xcf, 0x4a, 0x33, 0x31,
	0x10, 0xdf, 0x7c, 0xf0, 0x15, 0x1a, 0xa9, 0x60, 0x90, 0x52, 0x17, 0x59, 0x65, 0x51, 0xd1, 0x82,
	0x1b, 0xda, 0xbe, 0x41, 0xf5, 0x05, 0x5c, 0x6f, 0x5e, 0x64, 0xdb, 0x86, 0x34, 0x60, 0x33, 0xdb,
	0x24, 0x2d, 0x16, 0xf1, 0xe2, 0xd1, 0x93, 0xa0, 0x0f, 0xe1, 0xd1, 0xc7, 0xe8, 0xb1, 0xe0, 0xc5,
	0x93, 0x48, 0x57, 0xf0, 0x35, 0xa4, 0xc9, 0xaa, 0x2b, 0x28, 0x5e, 0xc2, 0x64, 0x66, 0x7e, 0xff,
	0x06, 0x57, 0x35, 0x53, 0x63, 0xd1, 0x65, 0x74, 0xdc, 0xa0, 0xc3, 0x11, 0x53, 0x93, 0x28, 0x55,
	0x60, 0x80, 0xe0, 0xbc, 0x1f, 0x8d, 0x1b, 0x7e, 0xad, 0xb0, 0xc3, 0x99, 0x64, 0x5a, 0x68, 0xb7,
	0xe5, 0xaf, 0x73, 0x00, 0x7e, 0xc6, 0x68, 0x92, 0x0a, 0x9a, 0x48, 0x09, 0x26, 0x31, 0x02, 0xe4,
	0xc7, 0x74, 0x25, 0x19, 0x08, 0x09, 0xd4, 0xbe, 0x79, 0x6b, 0x95, 0x03, 0x07, 0x5b, 0xd2, 0x45,
	0xe5, 0xba, 0x61, 0x0b, 0xaf, 0x1d, 0x2d, 0xb4, 0x8f, 0x9d, 0x4e, 0xcc, 0xba, 0xa0, 0x7a, 0x31,
	0x1b, 0x8e, 0x98, 0x36, 0xa4, 0x8a, 0x4b, 0xa0, 0x04, 0x17, 0xb2, 0x86, 0x36, 0xd1, 0x6e, 0x39,
	0xce, 0x7f, 0x61, 0x07, 0xfb, 0x3f, 0x81, 0x74, 0x0a, 0x52, 0x33, 0x72, 0x88, 0x97, 0x73, 0xd7,
	0xa7, 0xca, 0x4e, 0x2c, 0x7a, 0xa9, 0x49, 0xa2, 0xaf, 0x60, 0x91, 0xc3, 0xb4, 0xcb, 0xd3, 0xe7,
	0x0d, 0xef, 0xfe, 0xed, 0xa1, 0x8e, 0xe2, 0x8a, 0x2e, 0xb2, 0x35, 0xef, 0x10, 0xfe, 0x6f, 0x45,
	0xc8, 0x35, 0xc2, 0x95, 0x6f, 0x4a, 0x64, 0xbb, 0xc8, 0xf4, 0xab, 0x7d, 0x7f, 0xe7, 0xaf, 0x35,
	0x67, 0x38, 0xac, 0x5f, 0x3d, 0xbe, 0xde, 0xfe, 0xdb, 0x22, 0x21, 0xd5, 0x20, 0x55, 0x7f, 0x48,
	0x0b, 0x47, 0x77, 0x09, 0xe8, 0x85, 0x4b, 0x7e, 0xd9, 0x3e, 0x98, 0xce, 0x03, 0x34, 0x9b, 0x07,
	0xe8, 0x65, 0x1e, 0xa0, 0x9b, 0x2c, 0xf0, 0x66, 0x59, 0xe0, 0x3d, 0x65, 0x81, 0x77, 0xb2, 0xc7,
	0x85, 0xe9, 0x8f, 0x3a, 0x51, 0x17, 0x06, 0xb4, 0x27, 0x7a, 0x09, 0xec, 0x83, 0xe2, 0x96, 0x91,
	0x9e, 0x7f, 0x32, 0x9a, 0x49, 0xca, 0x74, 0xa7, 0x64, 0x6f, 0xdf, 0x7a, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0x29, 0xa9, 0xf8, 0xa7, 0x02, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// ServiceRecord returns the service record for a given origin.
	ServiceRecord(ctx context.Context, in *QueryServiceRecordRequest, opts ...grpc.CallOption) (*QueryServiceRecordResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ServiceRecord(ctx context.Context, in *QueryServiceRecordRequest, opts ...grpc.CallOption) (*QueryServiceRecordResponse, error) {
	out := new(QueryServiceRecordResponse)
	err := c.cc.Invoke(ctx, "/service.v1.Query/ServiceRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// ServiceRecord returns the service record for a given origin.
	ServiceRecord(context.Context, *QueryServiceRecordRequest) (*QueryServiceRecordResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ServiceRecord(ctx context.Context, req *QueryServiceRecordRequest) (*QueryServiceRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceRecord not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ServiceRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryServiceRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ServiceRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.v1.Query/ServiceRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ServiceRecord(ctx, req.(*QueryServiceRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServiceRecord",
			Handler:    _Query_ServiceRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/v1/query.proto",
}

func (m *QueryServiceRecordRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryServiceRecordRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryServiceRecordRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Origin) > 0 {
		i -= len(m.Origin)
		copy(dAtA[i:], m.Origin)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Origin)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryServiceRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryServiceRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryServiceRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ServiceRecord.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryServiceRecordRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Origin)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryServiceRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ServiceRecord.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryServiceRecordRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryServiceRecordRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryServiceRecordRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Origin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryServiceRecordResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryServiceRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryServiceRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceRecord", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ServiceRecord.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
