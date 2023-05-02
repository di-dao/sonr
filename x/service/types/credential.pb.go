// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/service/credential.proto

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

// WebauthnCredential contains all needed information about a WebAuthn credential for storage
type WebauthnCredential struct {
	// A probabilistically-unique byte sequence identifying a public key credential source and its authentication assertions.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The public key portion of a Relying Party-specific credential key pair, generated by an authenticator and returned to
	// a Relying Party at registration time (see also public key credential). The private key portion of the credential key
	// pair is known as the credential private key. Note that in the case of self attestation, the credential key pair is also
	// used as the attestation key pair, see self attestation for details.
	PublicKey []byte `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// The attestation format used (if any) by the authenticator when creating the credential.
	AttestationType string `protobuf:"bytes,3,opt,name=attestation_type,json=attestationType,proto3" json:"attestation_type,omitempty"`
	// The transports used by the authenticator when creating the credential.
	Transport []string `protobuf:"bytes,4,rep,name=transport,proto3" json:"transport,omitempty"`
	// The Authenticator information for a given certificate
	Authenticator *WebauthnAuthenticator `protobuf:"bytes,5,opt,name=authenticator,proto3" json:"authenticator,omitempty"`
	// The DID Controller of the credential
	Controller string `protobuf:"bytes,6,opt,name=controller,proto3" json:"controller,omitempty"`
}

func (m *WebauthnCredential) Reset()         { *m = WebauthnCredential{} }
func (m *WebauthnCredential) String() string { return proto.CompactTextString(m) }
func (*WebauthnCredential) ProtoMessage()    {}
func (*WebauthnCredential) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e8ca7c12592694f, []int{0}
}
func (m *WebauthnCredential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WebauthnCredential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WebauthnCredential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WebauthnCredential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebauthnCredential.Merge(m, src)
}
func (m *WebauthnCredential) XXX_Size() int {
	return m.Size()
}
func (m *WebauthnCredential) XXX_DiscardUnknown() {
	xxx_messageInfo_WebauthnCredential.DiscardUnknown(m)
}

var xxx_messageInfo_WebauthnCredential proto.InternalMessageInfo

func (m *WebauthnCredential) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *WebauthnCredential) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *WebauthnCredential) GetAttestationType() string {
	if m != nil {
		return m.AttestationType
	}
	return ""
}

func (m *WebauthnCredential) GetTransport() []string {
	if m != nil {
		return m.Transport
	}
	return nil
}

func (m *WebauthnCredential) GetAuthenticator() *WebauthnAuthenticator {
	if m != nil {
		return m.Authenticator
	}
	return nil
}

func (m *WebauthnCredential) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

// WebauthnAuthenticator contains certificate information about a WebAuthn authenticator
type WebauthnAuthenticator struct {
	// The AAGUID of the authenticator. An AAGUID is defined as an array containing the globally unique
	// identifier of the authenticator model being sought.
	Aaguid []byte `protobuf:"bytes,1,opt,name=aaguid,proto3" json:"aaguid,omitempty"`
	// SignCount -Upon a new login operation, the Relying Party compares the stored signature counter value
	// with the new signCount value returned in the assertion’s authenticator data. If this new
	// signCount value is less than or equal to the stored value, a cloned authenticator may
	// exist, or the authenticator may be malfunctioning.
	SignCount uint32 `protobuf:"varint,2,opt,name=sign_count,json=signCount,proto3" json:"sign_count,omitempty"`
	// Attachment - This is a signal that the authenticator may be cloned, i.e. at least two copies of the
	// credential private key may exist and are being used in parallel. Relying Parties should incorporate
	// this information into their risk scoring. Whether the Relying Party updates the stored signature
	// counter value in this case, or not, or fails the authentication ceremony or not, is Relying Party-specific.
	Attachment string `protobuf:"bytes,3,opt,name=attachment,proto3" json:"attachment,omitempty"`
}

func (m *WebauthnAuthenticator) Reset()         { *m = WebauthnAuthenticator{} }
func (m *WebauthnAuthenticator) String() string { return proto.CompactTextString(m) }
func (*WebauthnAuthenticator) ProtoMessage()    {}
func (*WebauthnAuthenticator) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e8ca7c12592694f, []int{1}
}
func (m *WebauthnAuthenticator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WebauthnAuthenticator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WebauthnAuthenticator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WebauthnAuthenticator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebauthnAuthenticator.Merge(m, src)
}
func (m *WebauthnAuthenticator) XXX_Size() int {
	return m.Size()
}
func (m *WebauthnAuthenticator) XXX_DiscardUnknown() {
	xxx_messageInfo_WebauthnAuthenticator.DiscardUnknown(m)
}

var xxx_messageInfo_WebauthnAuthenticator proto.InternalMessageInfo

func (m *WebauthnAuthenticator) GetAaguid() []byte {
	if m != nil {
		return m.Aaguid
	}
	return nil
}

func (m *WebauthnAuthenticator) GetSignCount() uint32 {
	if m != nil {
		return m.SignCount
	}
	return 0
}

func (m *WebauthnAuthenticator) GetAttachment() string {
	if m != nil {
		return m.Attachment
	}
	return ""
}

func init() {
	proto.RegisterType((*WebauthnCredential)(nil), "sonrhq.core.service.WebauthnCredential")
	proto.RegisterType((*WebauthnAuthenticator)(nil), "sonrhq.core.service.WebauthnAuthenticator")
}

func init() { proto.RegisterFile("core/service/credential.proto", fileDescriptor_3e8ca7c12592694f) }

var fileDescriptor_3e8ca7c12592694f = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x3b, 0xed, 0xbd, 0x85, 0xcc, 0xbd, 0xbd, 0x57, 0x46, 0x94, 0x2c, 0xec, 0x10, 0xba,
	0x90, 0xe8, 0x22, 0x01, 0x7d, 0x01, 0xb5, 0x4b, 0x37, 0x12, 0x04, 0xc1, 0x4d, 0x99, 0x4c, 0x0f,
	0xcd, 0x60, 0x3a, 0x13, 0x27, 0x27, 0x62, 0xde, 0xc2, 0xc7, 0x72, 0xd9, 0xa5, 0x4b, 0x69, 0xdf,
	0x43, 0x64, 0xd2, 0x5a, 0x23, 0x74, 0x99, 0x8f, 0x73, 0x4e, 0xfe, 0xf9, 0x7e, 0x3a, 0x94, 0xc6,
	0x42, 0x5c, 0x82, 0x7d, 0x52, 0x12, 0x62, 0x69, 0x61, 0x0a, 0x1a, 0x95, 0xc8, 0xa3, 0xc2, 0x1a,
	0x34, 0x6c, 0xbf, 0x34, 0xda, 0x66, 0x8f, 0x91, 0x9b, 0x8a, 0x36, 0x53, 0xa3, 0x0f, 0x42, 0xd9,
	0x1d, 0xa4, 0xa2, 0xc2, 0x4c, 0x8f, 0xb7, 0x1b, 0xec, 0x1f, 0xed, 0xaa, 0xa9, 0x4f, 0x02, 0x12,
	0xfe, 0x4d, 0xba, 0x6a, 0xca, 0x86, 0x94, 0x16, 0x55, 0x9a, 0x2b, 0x39, 0x79, 0x80, 0xda, 0xef,
	0x36, 0xdc, 0x5b, 0x93, 0x6b, 0xa8, 0xd9, 0x09, 0xdd, 0x13, 0x88, 0x50, 0xa2, 0x40, 0x65, 0xf4,
	0x04, 0xeb, 0x02, 0xfc, 0x5e, 0x40, 0x42, 0x2f, 0xf9, 0xdf, 0xe2, 0xb7, 0x75, 0x01, 0xec, 0x88,
	0x7a, 0x68, 0x85, 0x2e, 0x0b, 0x63, 0xd1, 0xff, 0x15, 0xf4, 0x42, 0x2f, 0xf9, 0x06, 0xec, 0x86,
	0x0e, 0x5c, 0x14, 0x97, 0x42, 0x0a, 0x34, 0xd6, 0xff, 0x1d, 0x90, 0xf0, 0xcf, 0xd9, 0x69, 0xb4,
	0x23, 0x7b, 0xf4, 0x95, 0xfb, 0xb2, 0xbd, 0x91, 0xfc, 0x3c, 0xc0, 0x38, 0xa5, 0xd2, 0x68, 0xb4,
	0x26, 0xcf, 0xc1, 0xfa, 0xfd, 0x26, 0x54, 0x8b, 0x8c, 0x34, 0x3d, 0xd8, 0x79, 0x87, 0x1d, 0xd2,
	0xbe, 0x10, 0xb3, 0x6a, 0xab, 0x61, 0xf3, 0xe5, 0x54, 0x94, 0x6a, 0xa6, 0x27, 0xd2, 0x54, 0x1a,
	0x1b, 0x15, 0x83, 0xc4, 0x73, 0x64, 0xec, 0x80, 0xfb, 0x9f, 0x40, 0x14, 0x32, 0x9b, 0x83, 0xc6,
	0x8d, 0x84, 0x16, 0xb9, 0xba, 0x78, 0x5d, 0x72, 0xb2, 0x58, 0x72, 0xf2, 0xbe, 0xe4, 0xe4, 0x65,
	0xc5, 0x3b, 0x8b, 0x15, 0xef, 0xbc, 0xad, 0x78, 0xe7, 0xfe, 0x78, 0xa6, 0x30, 0xab, 0xd2, 0x48,
	0x9a, 0x79, 0xbc, 0x7e, 0x6e, 0xdc, 0x14, 0xfa, 0xbc, 0xad, 0xd4, 0x79, 0x2d, 0xd3, 0x7e, 0x53,
	0xe7, 0xf9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x40, 0xdc, 0x27, 0x63, 0xef, 0x01, 0x00, 0x00,
}

func (m *WebauthnCredential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WebauthnCredential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WebauthnCredential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x32
	}
	if m.Authenticator != nil {
		{
			size, err := m.Authenticator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCredential(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Transport) > 0 {
		for iNdEx := len(m.Transport) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Transport[iNdEx])
			copy(dAtA[i:], m.Transport[iNdEx])
			i = encodeVarintCredential(dAtA, i, uint64(len(m.Transport[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.AttestationType) > 0 {
		i -= len(m.AttestationType)
		copy(dAtA[i:], m.AttestationType)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.AttestationType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WebauthnAuthenticator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WebauthnAuthenticator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WebauthnAuthenticator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attachment) > 0 {
		i -= len(m.Attachment)
		copy(dAtA[i:], m.Attachment)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.Attachment)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SignCount != 0 {
		i = encodeVarintCredential(dAtA, i, uint64(m.SignCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Aaguid) > 0 {
		i -= len(m.Aaguid)
		copy(dAtA[i:], m.Aaguid)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.Aaguid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCredential(dAtA []byte, offset int, v uint64) int {
	offset -= sovCredential(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WebauthnCredential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	l = len(m.AttestationType)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	if len(m.Transport) > 0 {
		for _, s := range m.Transport {
			l = len(s)
			n += 1 + l + sovCredential(uint64(l))
		}
	}
	if m.Authenticator != nil {
		l = m.Authenticator.Size()
		n += 1 + l + sovCredential(uint64(l))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	return n
}

func (m *WebauthnAuthenticator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Aaguid)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	if m.SignCount != 0 {
		n += 1 + sovCredential(uint64(m.SignCount))
	}
	l = len(m.Attachment)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	return n
}

func sovCredential(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCredential(x uint64) (n int) {
	return sovCredential(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WebauthnCredential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCredential
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
			return fmt.Errorf("proto: WebauthnCredential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WebauthnCredential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = append(m.PublicKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PublicKey == nil {
				m.PublicKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttestationType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttestationType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transport", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transport = append(m.Transport, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authenticator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authenticator == nil {
				m.Authenticator = &WebauthnAuthenticator{}
			}
			if err := m.Authenticator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCredential(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCredential
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
func (m *WebauthnAuthenticator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCredential
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
			return fmt.Errorf("proto: WebauthnAuthenticator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WebauthnAuthenticator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Aaguid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Aaguid = append(m.Aaguid[:0], dAtA[iNdEx:postIndex]...)
			if m.Aaguid == nil {
				m.Aaguid = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignCount", wireType)
			}
			m.SignCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignCount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attachment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attachment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCredential(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCredential
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
func skipCredential(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCredential
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
					return 0, ErrIntOverflowCredential
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
					return 0, ErrIntOverflowCredential
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
				return 0, ErrInvalidLengthCredential
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCredential
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCredential
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCredential        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCredential          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCredential = fmt.Errorf("proto: unexpected end of group")
)
