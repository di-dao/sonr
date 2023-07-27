// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/identity/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Params defines the parameters for the module.
type Params struct {
	AccountDidMethodName    string   `protobuf:"bytes,1,opt,name=account_did_method_name,json=accountDidMethodName,proto3" json:"account_did_method_name,omitempty"`
	AccountDidMethodContext string   `protobuf:"bytes,2,opt,name=account_did_method_context,json=accountDidMethodContext,proto3" json:"account_did_method_context,omitempty"`
	AcccountDiscoveryReward int64    `protobuf:"varint,3,opt,name=acccount_discovery_reward,json=acccountDiscoveryReward,proto3" json:"acccount_discovery_reward,omitempty"`
	DidBaseContext          string   `protobuf:"bytes,4,opt,name=did_base_context,json=didBaseContext,proto3" json:"did_base_context,omitempty"`
	MaximumIdentityAliases  int32    `protobuf:"varint,5,opt,name=maximum_identity_aliases,json=maximumIdentityAliases,proto3" json:"maximum_identity_aliases,omitempty"`
	SupportedDidMethods     []string `protobuf:"bytes,6,rep,name=supported_did_methods,json=supportedDidMethods,proto3" json:"supported_did_methods,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b3a47d8d2dc969e, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAccountDidMethodName() string {
	if m != nil {
		return m.AccountDidMethodName
	}
	return ""
}

func (m *Params) GetAccountDidMethodContext() string {
	if m != nil {
		return m.AccountDidMethodContext
	}
	return ""
}

func (m *Params) GetAcccountDiscoveryReward() int64 {
	if m != nil {
		return m.AcccountDiscoveryReward
	}
	return 0
}

func (m *Params) GetDidBaseContext() string {
	if m != nil {
		return m.DidBaseContext
	}
	return ""
}

func (m *Params) GetMaximumIdentityAliases() int32 {
	if m != nil {
		return m.MaximumIdentityAliases
	}
	return 0
}

func (m *Params) GetSupportedDidMethods() []string {
	if m != nil {
		return m.SupportedDidMethods
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "core.identity.Params")
}

func init() { proto.RegisterFile("core/identity/params.proto", fileDescriptor_2b3a47d8d2dc969e) }

var fileDescriptor_2b3a47d8d2dc969e = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd1, 0xbf, 0x4b, 0xc3, 0x40,
	0x14, 0xc0, 0xf1, 0x5c, 0x7f, 0x41, 0x0f, 0x14, 0x89, 0xd5, 0xc6, 0x0e, 0x31, 0xb8, 0x98, 0xa9,
	0x01, 0x45, 0x90, 0x3a, 0xb5, 0xba, 0x38, 0x28, 0x92, 0xd1, 0x25, 0x5c, 0xef, 0x1e, 0xed, 0x81,
	0x97, 0x8b, 0x77, 0x17, 0x6d, 0xff, 0x0b, 0x47, 0x47, 0xff, 0x1c, 0xc7, 0x4e, 0xe2, 0x28, 0xed,
	0x3f, 0x22, 0xbd, 0x26, 0x51, 0xc4, 0x2d, 0xf0, 0xc9, 0x97, 0xc7, 0xbb, 0x87, 0x7b, 0x54, 0x2a,
	0x88, 0x38, 0x83, 0xd4, 0x70, 0x33, 0x8f, 0x32, 0xa2, 0x88, 0xd0, 0xfd, 0x4c, 0x49, 0x23, 0xdd,
	0xad, 0xb5, 0xf5, 0x4b, 0xeb, 0x75, 0x26, 0x72, 0x22, 0xad, 0x44, 0xeb, 0xaf, 0xcd, 0x4f, 0x47,
	0x1f, 0x35, 0xdc, 0xba, 0xb3, 0x95, 0x7b, 0x86, 0xbb, 0x84, 0x52, 0x99, 0xa7, 0x26, 0x61, 0x9c,
	0x25, 0x02, 0xcc, 0x54, 0xb2, 0x24, 0x25, 0x02, 0x3c, 0x14, 0xa0, 0xb0, 0x1d, 0x77, 0x0a, 0xbe,
	0xe2, 0xec, 0xc6, 0xe2, 0x2d, 0x11, 0xe0, 0x5e, 0xe0, 0xde, 0x3f, 0x19, 0x95, 0xa9, 0x81, 0x99,
	0xf1, 0x6a, 0xb6, 0xec, 0xfe, 0x2d, 0x2f, 0x37, 0xec, 0x0e, 0xf0, 0x01, 0xa1, 0x55, 0xad, 0xa9,
	0x7c, 0x02, 0x35, 0x4f, 0x14, 0x3c, 0x13, 0xc5, 0xbc, 0x7a, 0x80, 0xc2, 0xba, 0x6d, 0x8b, 0xb8,
	0xf0, 0xd8, 0xb2, 0x1b, 0xe2, 0x9d, 0xf5, 0xc0, 0x31, 0xd1, 0x50, 0x8d, 0x6b, 0xd8, 0x71, 0xdb,
	0x8c, 0xb3, 0x11, 0xd1, 0x50, 0x4e, 0x39, 0xc7, 0x9e, 0x20, 0x33, 0x2e, 0x72, 0x91, 0x94, 0xcf,
	0x91, 0x90, 0x07, 0x4e, 0x34, 0x68, 0xaf, 0x19, 0xa0, 0xb0, 0x19, 0xef, 0x17, 0x7e, 0x5d, 0xf0,
	0x70, 0xa3, 0xee, 0x09, 0xde, 0xd3, 0x79, 0x96, 0x49, 0x65, 0x80, 0xfd, 0x5a, 0x4f, 0x7b, 0xad,
	0xa0, 0x1e, 0xb6, 0xe3, 0xdd, 0x0a, 0xab, 0xcd, 0xf4, 0xa0, 0xf1, 0xfa, 0x76, 0xe8, 0x8c, 0x86,
	0xef, 0x4b, 0x1f, 0x2d, 0x96, 0x3e, 0xfa, 0x5a, 0xfa, 0xe8, 0x65, 0xe5, 0x3b, 0x8b, 0x95, 0xef,
	0x7c, 0xae, 0x7c, 0xe7, 0xfe, 0x78, 0xc2, 0xcd, 0x34, 0x1f, 0xf7, 0xa9, 0x14, 0x91, 0x96, 0xa9,
	0x9a, 0x3e, 0x46, 0xf6, 0x8a, 0xb3, 0x9f, 0x3b, 0x9a, 0x79, 0x06, 0x7a, 0xdc, 0xb2, 0x27, 0x3a,
	0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x3a, 0xed, 0x24, 0xe5, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SupportedDidMethods) > 0 {
		for iNdEx := len(m.SupportedDidMethods) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SupportedDidMethods[iNdEx])
			copy(dAtA[i:], m.SupportedDidMethods[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.SupportedDidMethods[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.MaximumIdentityAliases != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaximumIdentityAliases))
		i--
		dAtA[i] = 0x28
	}
	if len(m.DidBaseContext) > 0 {
		i -= len(m.DidBaseContext)
		copy(dAtA[i:], m.DidBaseContext)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DidBaseContext)))
		i--
		dAtA[i] = 0x22
	}
	if m.AcccountDiscoveryReward != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AcccountDiscoveryReward))
		i--
		dAtA[i] = 0x18
	}
	if len(m.AccountDidMethodContext) > 0 {
		i -= len(m.AccountDidMethodContext)
		copy(dAtA[i:], m.AccountDidMethodContext)
		i = encodeVarintParams(dAtA, i, uint64(len(m.AccountDidMethodContext)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AccountDidMethodName) > 0 {
		i -= len(m.AccountDidMethodName)
		copy(dAtA[i:], m.AccountDidMethodName)
		i = encodeVarintParams(dAtA, i, uint64(len(m.AccountDidMethodName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AccountDidMethodName)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.AccountDidMethodContext)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.AcccountDiscoveryReward != 0 {
		n += 1 + sovParams(uint64(m.AcccountDiscoveryReward))
	}
	l = len(m.DidBaseContext)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MaximumIdentityAliases != 0 {
		n += 1 + sovParams(uint64(m.MaximumIdentityAliases))
	}
	if len(m.SupportedDidMethods) > 0 {
		for _, s := range m.SupportedDidMethods {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountDidMethodName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountDidMethodName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountDidMethodContext", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountDidMethodContext = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AcccountDiscoveryReward", wireType)
			}
			m.AcccountDiscoveryReward = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AcccountDiscoveryReward |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidBaseContext", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DidBaseContext = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaximumIdentityAliases", wireType)
			}
			m.MaximumIdentityAliases = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaximumIdentityAliases |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupportedDidMethods", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SupportedDidMethods = append(m.SupportedDidMethods, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)