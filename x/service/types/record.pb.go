// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/service/record.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// Permissions are a bitfield of permissions that a user has for a given service
// in order to perform certain actions.
type Permissions int32

const (
	// Authenticated user with basic public info and read access to public
	// resources
	Permissions_SR_BASE Permissions = 0
	// Authenticated user with service-side access to signing privelages
	Permissions_SR_SIGN Permissions = 1
	// Authenticated user with service-side access to writing to ipfs
	Permissions_SR_WRITE Permissions = 2
	// Authenticated user with service-side access to updating user DID documents
	// and creating new ones
	Permissions_SR_CREATE Permissions = 3
	// Authenticated user with service-side access to receiving notifications
	Permissions_SR_NOTIFY Permissions = 4
)

var Permissions_name = map[int32]string{
	0: "SR_BASE",
	1: "SR_SIGN",
	2: "SR_WRITE",
	3: "SR_CREATE",
	4: "SR_NOTIFY",
}

var Permissions_value = map[string]int32{
	"SR_BASE":   0,
	"SR_SIGN":   1,
	"SR_WRITE":  2,
	"SR_CREATE": 3,
	"SR_NOTIFY": 4,
}

func (x Permissions) String() string {
	return proto.EnumName(Permissions_name, int32(x))
}

func (Permissions) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_17bfe358aaf56022, []int{0}
}

type ServiceRecord struct {
	Id          string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Controller  string      `protobuf:"bytes,2,opt,name=controller,proto3" json:"controller,omitempty"`
	Origins     []string    `protobuf:"bytes,3,rep,name=origins,proto3" json:"origins,omitempty"`
	Name        string      `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description string      `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Permissions Permissions `protobuf:"varint,6,opt,name=permissions,proto3,enum=core.service.Permissions" json:"permissions,omitempty"`
}

func (m *ServiceRecord) Reset()         { *m = ServiceRecord{} }
func (m *ServiceRecord) String() string { return proto.CompactTextString(m) }
func (*ServiceRecord) ProtoMessage()    {}
func (*ServiceRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_17bfe358aaf56022, []int{0}
}
func (m *ServiceRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ServiceRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ServiceRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceRecord.Merge(m, src)
}
func (m *ServiceRecord) XXX_Size() int {
	return m.Size()
}
func (m *ServiceRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceRecord proto.InternalMessageInfo

func (m *ServiceRecord) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceRecord) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func (m *ServiceRecord) GetOrigins() []string {
	if m != nil {
		return m.Origins
	}
	return nil
}

func (m *ServiceRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceRecord) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ServiceRecord) GetPermissions() Permissions {
	if m != nil {
		return m.Permissions
	}
	return Permissions_SR_BASE
}

// ServiceRelationship is a relationship between a service and a User entity.
// This relation is used to track the number of times a user has interacted with
// a service, utilizing a verifiable random function (VRF) to generate a random
// number that is used to determine whether or not a user has access to a
// service.
type ServiceRelationship struct {
	// Reference is the reference to the service record
	Reference string `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`
	// Did is the user's DID
	Did string `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
	// Count is the number of times the user has interacted with the service
	Count uint32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *ServiceRelationship) Reset()         { *m = ServiceRelationship{} }
func (m *ServiceRelationship) String() string { return proto.CompactTextString(m) }
func (*ServiceRelationship) ProtoMessage()    {}
func (*ServiceRelationship) Descriptor() ([]byte, []int) {
	return fileDescriptor_17bfe358aaf56022, []int{1}
}
func (m *ServiceRelationship) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceRelationship) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ServiceRelationship.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ServiceRelationship) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceRelationship.Merge(m, src)
}
func (m *ServiceRelationship) XXX_Size() int {
	return m.Size()
}
func (m *ServiceRelationship) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceRelationship.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceRelationship proto.InternalMessageInfo

func (m *ServiceRelationship) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *ServiceRelationship) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *ServiceRelationship) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterEnum("core.service.Permissions", Permissions_name, Permissions_value)
	proto.RegisterType((*ServiceRecord)(nil), "core.service.ServiceRecord")
	proto.RegisterType((*ServiceRelationship)(nil), "core.service.ServiceRelationship")
}

func init() { proto.RegisterFile("core/service/record.proto", fileDescriptor_17bfe358aaf56022) }

var fileDescriptor_17bfe358aaf56022 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x51, 0xcb, 0xae, 0xd2, 0x40,
	0x18, 0xee, 0xb4, 0x5c, 0xec, 0x5f, 0x20, 0xcd, 0xe8, 0x62, 0x48, 0x4c, 0xd3, 0xb0, 0x30, 0x8d,
	0x8b, 0x36, 0xd1, 0xa5, 0x1b, 0xc1, 0x54, 0xc3, 0x06, 0xcd, 0x94, 0xc4, 0xa8, 0x0b, 0x02, 0xd3,
	0x11, 0x26, 0x81, 0x99, 0x3a, 0x53, 0x8c, 0xbe, 0x85, 0x8f, 0xe5, 0xe2, 0x2c, 0x58, 0x9e, 0xe5,
	0x09, 0xbc, 0xc8, 0x09, 0xa5, 0x07, 0xba, 0xfb, 0x6e, 0x73, 0xf9, 0xf2, 0xc1, 0x90, 0x29, 0xcd,
	0x13, 0xc3, 0xf5, 0x6f, 0xc1, 0x78, 0xa2, 0x39, 0x53, 0x3a, 0x8f, 0x0b, 0xad, 0x4a, 0x85, 0x7b,
	0x67, 0x2b, 0xae, 0xad, 0xd1, 0x1d, 0x82, 0x7e, 0x76, 0xc1, 0xb4, 0x4a, 0xe1, 0x01, 0xd8, 0x22,
	0x27, 0x28, 0x44, 0x91, 0x4b, 0x6d, 0x91, 0xe3, 0x00, 0x80, 0x29, 0x59, 0x6a, 0xb5, 0xdd, 0x72,
	0x4d, 0xec, 0x4a, 0x6f, 0x28, 0x98, 0x40, 0x57, 0x69, 0xb1, 0x16, 0xd2, 0x10, 0x27, 0x74, 0x22,
	0x97, 0x3e, 0x51, 0x8c, 0xa1, 0x25, 0x97, 0x3b, 0x4e, 0x5a, 0xd5, 0x99, 0x0a, 0xe3, 0x10, 0xbc,
	0x9c, 0x1b, 0xa6, 0x45, 0x51, 0x0a, 0x25, 0x49, 0xbb, 0xb2, 0x9a, 0x12, 0x7e, 0x07, 0x5e, 0xc1,
	0xf5, 0x4e, 0x18, 0x23, 0x94, 0x34, 0xa4, 0x13, 0xa2, 0x68, 0xf0, 0x66, 0x18, 0x37, 0x7f, 0x1d,
	0x7f, 0xb9, 0x05, 0x68, 0x33, 0x3d, 0xfa, 0x01, 0xcf, 0xaf, 0x6d, 0xb6, 0xcb, 0xf3, 0x7d, 0x66,
	0x23, 0x0a, 0xfc, 0x12, 0x5c, 0xcd, 0x7f, 0x72, 0xcd, 0x25, 0xe3, 0x75, 0xb5, 0x9b, 0x80, 0x7d,
	0x70, 0x72, 0x91, 0xd7, 0xd5, 0xce, 0x10, 0xbf, 0x80, 0x36, 0x53, 0x7b, 0x59, 0x12, 0x27, 0x44,
	0x51, 0x9f, 0x5e, 0xc8, 0xeb, 0x0c, 0xbc, 0xc6, 0xc3, 0xd8, 0x83, 0x6e, 0x46, 0x17, 0x93, 0x71,
	0x96, 0xfa, 0x56, 0x4d, 0xb2, 0xe9, 0xa7, 0x99, 0x8f, 0x70, 0x0f, 0x9e, 0x65, 0x74, 0xf1, 0x95,
	0x4e, 0xe7, 0xa9, 0x6f, 0xe3, 0x3e, 0xb8, 0x19, 0x5d, 0x7c, 0xa0, 0xe9, 0x78, 0x9e, 0xfa, 0x4e,
	0x4d, 0x67, 0x9f, 0xe7, 0xd3, 0x8f, 0xdf, 0xfc, 0xd6, 0xe4, 0xfd, 0xff, 0x63, 0x80, 0x0e, 0xc7,
	0x00, 0x3d, 0x1c, 0x03, 0xf4, 0xef, 0x14, 0x58, 0x87, 0x53, 0x60, 0xdd, 0x9f, 0x02, 0xeb, 0xfb,
	0xab, 0xb5, 0x28, 0x37, 0xfb, 0x55, 0xcc, 0xd4, 0x2e, 0x31, 0x4a, 0xea, 0xcd, 0xaf, 0xa4, 0x5a,
	0xf5, 0xcf, 0x75, 0xd7, 0xf2, 0x6f, 0xc1, 0xcd, 0xaa, 0x53, 0xed, 0xfa, 0xf6, 0x31, 0x00, 0x00,
	0xff, 0xff, 0x0d, 0x66, 0x10, 0xf5, 0xf4, 0x01, 0x00, 0x00,
}

func (m *ServiceRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServiceRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Permissions != 0 {
		i = encodeVarintRecord(dAtA, i, uint64(m.Permissions))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Origins) > 0 {
		for iNdEx := len(m.Origins) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Origins[iNdEx])
			copy(dAtA[i:], m.Origins[iNdEx])
			i = encodeVarintRecord(dAtA, i, uint64(len(m.Origins[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ServiceRelationship) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceRelationship) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServiceRelationship) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		i = encodeVarintRecord(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Reference) > 0 {
		i -= len(m.Reference)
		copy(dAtA[i:], m.Reference)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Reference)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ServiceRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	if len(m.Origins) > 0 {
		for _, s := range m.Origins {
			l = len(s)
			n += 1 + l + sovRecord(uint64(l))
		}
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	if m.Permissions != 0 {
		n += 1 + sovRecord(uint64(m.Permissions))
	}
	return n
}

func (m *ServiceRelationship) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Reference)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovRecord(uint64(m.Count))
	}
	return n
}

func sovRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRecord(x uint64) (n int) {
	return sovRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ServiceRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecord
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
			return fmt.Errorf("proto: ServiceRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origins", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Origins = append(m.Origins, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			m.Permissions = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Permissions |= Permissions(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecord
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
func (m *ServiceRelationship) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecord
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
			return fmt.Errorf("proto: ServiceRelationship: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceRelationship: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reference", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reference = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecord
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
func skipRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRecord
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
					return 0, ErrIntOverflowRecord
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
					return 0, ErrIntOverflowRecord
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
				return 0, ErrInvalidLengthRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRecord = fmt.Errorf("proto: unexpected end of group")
)
