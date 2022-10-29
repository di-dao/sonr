// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bucket/bucket.proto

// Package Motor is used for defining a Motor node and its properties.

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// BucketVisibility is the visibility of the bucket to authorized users of an application
type BucketVisibility int32

const (
	// Bucket does not have visibility set.
	BucketVisibility_PUBLIC BucketVisibility = 0
	// Bucket is visible to anyone.
	BucketVisibility_USER BucketVisibility = 1
	// Bucket is visible to anyone who has access token.
	BucketVisibility_APPLICATION BucketVisibility = 2
)

var BucketVisibility_name = map[int32]string{
	0: "PUBLIC",
	1: "USER",
	2: "APPLICATION",
}

var BucketVisibility_value = map[string]int32{
	"PUBLIC":      0,
	"USER":        1,
	"APPLICATION": 2,
}

func (x BucketVisibility) String() string {
	return proto.EnumName(BucketVisibility_name, int32(x))
}

func (BucketVisibility) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_748678784c4bd53f, []int{0}
}

//
//Wraps items within a bucket. Items will be one of the following
//DID -> reference to another bucket (WhereIs)
//CID -> reference to content (map[string]interface{})
type BucketContent struct {
	// Raw content serialized to bytes
	Item []byte `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	// Content id a CID, DID, or unspecified
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *BucketContent) Reset()         { *m = BucketContent{} }
func (m *BucketContent) String() string { return proto.CompactTextString(m) }
func (*BucketContent) ProtoMessage()    {}
func (*BucketContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_748678784c4bd53f, []int{0}
}
func (m *BucketContent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BucketContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BucketContent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BucketContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketContent.Merge(m, src)
}
func (m *BucketContent) XXX_Size() int {
	return m.Size()
}
func (m *BucketContent) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketContent.DiscardUnknown(m)
}

var xxx_messageInfo_BucketContent proto.InternalMessageInfo

func (m *BucketContent) GetItem() []byte {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *BucketContent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type BucketConfig struct {
	// DID of the created bucket.
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Creator of the new bucket
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	// Name of the new bucket.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Visibility of the new bucket.
	Visibility BucketVisibility `protobuf:"varint,4,opt,name=visibility,proto3,enum=sonrio.sonr.bucket.BucketVisibility" json:"visibility,omitempty"`
	// IsActive flag of the new bucket.
	IsActive bool `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// Timestamp of the new bucket.
	Size_ int64 `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
}

func (m *BucketConfig) Reset()         { *m = BucketConfig{} }
func (m *BucketConfig) String() string { return proto.CompactTextString(m) }
func (*BucketConfig) ProtoMessage()    {}
func (*BucketConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_748678784c4bd53f, []int{1}
}
func (m *BucketConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BucketConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BucketConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BucketConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketConfig.Merge(m, src)
}
func (m *BucketConfig) XXX_Size() int {
	return m.Size()
}
func (m *BucketConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BucketConfig proto.InternalMessageInfo

func (m *BucketConfig) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *BucketConfig) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *BucketConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BucketConfig) GetVisibility() BucketVisibility {
	if m != nil {
		return m.Visibility
	}
	return BucketVisibility_PUBLIC
}

func (m *BucketConfig) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *BucketConfig) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func init() {
	proto.RegisterEnum("sonrio.sonr.bucket.BucketVisibility", BucketVisibility_name, BucketVisibility_value)
	proto.RegisterType((*BucketContent)(nil), "sonrio.sonr.bucket.BucketContent")
	proto.RegisterType((*BucketConfig)(nil), "sonrio.sonr.bucket.BucketConfig")
}

func init() { proto.RegisterFile("bucket/bucket.proto", fileDescriptor_748678784c4bd53f) }

var fileDescriptor_748678784c4bd53f = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x41, 0x4b, 0xeb, 0x40,
	0x10, 0xc7, 0xb3, 0x69, 0xe9, 0x6b, 0xe6, 0xf5, 0xf5, 0x85, 0xed, 0x25, 0x20, 0x84, 0x50, 0x44,
	0x82, 0x60, 0x02, 0xf6, 0xe4, 0xb1, 0xad, 0x1e, 0x0a, 0x45, 0xc3, 0x6a, 0x15, 0xbc, 0x48, 0x93,
	0xae, 0x75, 0xd0, 0x64, 0x4b, 0xb2, 0x29, 0xd6, 0x4f, 0xe1, 0xb7, 0xd2, 0x63, 0x8f, 0x1e, 0xa5,
	0xfd, 0x22, 0x92, 0x4d, 0x5b, 0x44, 0x4f, 0xff, 0xff, 0x30, 0xff, 0xdf, 0xcc, 0xc0, 0x40, 0x2b,
	0xcc, 0xa3, 0x47, 0x2e, 0xfd, 0x52, 0xbc, 0x59, 0x2a, 0xa4, 0xa0, 0x34, 0x13, 0x49, 0x8a, 0xc2,
	0x2b, 0xc4, 0x2b, 0x3b, 0xed, 0x0e, 0xfc, 0xeb, 0x29, 0xd7, 0x17, 0x89, 0xe4, 0x89, 0xa4, 0x14,
	0xaa, 0x28, 0x79, 0x6c, 0x11, 0x87, 0xb8, 0x0d, 0xa6, 0x3c, 0x6d, 0x82, 0x8e, 0x13, 0x4b, 0x77,
	0x88, 0x6b, 0x30, 0x1d, 0x27, 0xed, 0x37, 0x02, 0x8d, 0x1d, 0x75, 0x8f, 0xd3, 0x02, 0xca, 0x73,
	0x9c, 0x28, 0xc8, 0x60, 0xca, 0x53, 0x0b, 0xfe, 0x44, 0x29, 0x1f, 0x4b, 0x91, 0x6e, 0xc8, 0x6d,
	0x59, 0xa4, 0x93, 0x71, 0xcc, 0xad, 0x4a, 0x99, 0x2e, 0x3c, 0x3d, 0x05, 0x98, 0x63, 0x86, 0x21,
	0x3e, 0xa1, 0x5c, 0x58, 0x55, 0x87, 0xb8, 0xcd, 0xe3, 0x7d, 0xef, 0xf7, 0xc1, 0x5e, 0xb9, 0xf7,
	0x7a, 0x97, 0x65, 0xdf, 0x38, 0xba, 0x07, 0x06, 0x66, 0x77, 0xe3, 0x48, 0xe2, 0x9c, 0x5b, 0x35,
	0x87, 0xb8, 0x75, 0x56, 0xc7, 0xac, 0xab, 0xea, 0x62, 0x6d, 0x86, 0x2f, 0xdc, 0x32, 0x1c, 0xe2,
	0x56, 0x98, 0xf2, 0x87, 0x27, 0x60, 0xfe, 0x1c, 0x48, 0x01, 0x6a, 0xc1, 0xa8, 0x37, 0x1c, 0xf4,
	0x4d, 0x8d, 0xd6, 0xa1, 0x3a, 0xba, 0x3c, 0x63, 0x26, 0xa1, 0xff, 0xe1, 0x6f, 0x37, 0x08, 0x86,
	0x83, 0x7e, 0xf7, 0x6a, 0x70, 0x71, 0x6e, 0xea, 0xbd, 0x9b, 0xf7, 0x95, 0x4d, 0x96, 0x2b, 0x9b,
	0x7c, 0xae, 0x6c, 0xf2, 0xba, 0xb6, 0xb5, 0xe5, 0xda, 0xd6, 0x3e, 0xd6, 0xb6, 0x06, 0xad, 0xed,
	0xc9, 0x72, 0x31, 0xe3, 0x99, 0x17, 0x0b, 0x29, 0xd2, 0x80, 0xdc, 0x1e, 0x4c, 0x51, 0x3e, 0xe4,
	0xa1, 0x17, 0x89, 0xd8, 0x2f, 0xda, 0x47, 0x28, 0x94, 0xfa, 0xcf, 0x9b, 0x37, 0xf9, 0x2a, 0x1f,
	0xd6, 0xd4, 0xb7, 0x3a, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x4e, 0xff, 0x67, 0xc4, 0x01,
	0x00, 0x00,
}

func (m *BucketContent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BucketContent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BucketContent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintBucket(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Item) > 0 {
		i -= len(m.Item)
		copy(dAtA[i:], m.Item)
		i = encodeVarintBucket(dAtA, i, uint64(len(m.Item)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BucketConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BucketConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BucketConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Size_ != 0 {
		i = encodeVarintBucket(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x48
	}
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.Visibility != 0 {
		i = encodeVarintBucket(dAtA, i, uint64(m.Visibility))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintBucket(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintBucket(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Uuid) > 0 {
		i -= len(m.Uuid)
		copy(dAtA[i:], m.Uuid)
		i = encodeVarintBucket(dAtA, i, uint64(len(m.Uuid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBucket(dAtA []byte, offset int, v uint64) int {
	offset -= sovBucket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BucketContent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Item)
	if l > 0 {
		n += 1 + l + sovBucket(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovBucket(uint64(l))
	}
	return n
}

func (m *BucketConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovBucket(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovBucket(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovBucket(uint64(l))
	}
	if m.Visibility != 0 {
		n += 1 + sovBucket(uint64(m.Visibility))
	}
	if m.IsActive {
		n += 2
	}
	if m.Size_ != 0 {
		n += 1 + sovBucket(uint64(m.Size_))
	}
	return n
}

func sovBucket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBucket(x uint64) (n int) {
	return sovBucket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BucketContent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBucket
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
			return fmt.Errorf("proto: BucketContent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BucketContent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Item", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBucket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBucket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBucket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Item = append(m.Item[:0], dAtA[iNdEx:postIndex]...)
			if m.Item == nil {
				m.Item = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBucket
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
				return ErrInvalidLengthBucket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBucket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBucket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBucket
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
func (m *BucketConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBucket
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
			return fmt.Errorf("proto: BucketConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BucketConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBucket
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
				return ErrInvalidLengthBucket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBucket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBucket
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
				return ErrInvalidLengthBucket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBucket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBucket
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
				return ErrInvalidLengthBucket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBucket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Visibility", wireType)
			}
			m.Visibility = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBucket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Visibility |= BucketVisibility(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBucket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsActive = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBucket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBucket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBucket
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
func skipBucket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBucket
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
					return 0, ErrIntOverflowBucket
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
					return 0, ErrIntOverflowBucket
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
				return 0, ErrInvalidLengthBucket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBucket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBucket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBucket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBucket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBucket = fmt.Errorf("proto: unexpected end of group")
)
