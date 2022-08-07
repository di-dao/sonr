// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common/link.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

type LinkKind int32

const (
	// None should be specified when no link is provided for the given schema type
	LinkKind_None   LinkKind = 0
	LinkKind_Object LinkKind = 1
	LinkKind_Schema LinkKind = 2
	LinkKind_Bucket LinkKind = 3
)

var LinkKind_name = map[int32]string{
	0: "None",
	1: "Object",
	2: "Schema",
	3: "Bucket",
}

var LinkKind_value = map[string]int32{
	"None":   0,
	"Object": 1,
	"Schema": 2,
	"Bucket": 3,
}

func (x LinkKind) String() string {
	return proto.EnumName(LinkKind_name, int32(x))
}

func (LinkKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_471b35814ac8730b, []int{0}
}

func init() {
	proto.RegisterEnum("sonrio.sonr.common.LinkKind", LinkKind_name, LinkKind_value)
}

func init() { proto.RegisterFile("common/link.proto", fileDescriptor_471b35814ac8730b) }

var fileDescriptor_471b35814ac8730b = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0xcf, 0xc9, 0xcc, 0xcb, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2a,
	0xce, 0xcf, 0x2b, 0xca, 0xcc, 0xd7, 0x03, 0x51, 0x7a, 0x10, 0x69, 0x2d, 0x0b, 0x2e, 0x0e, 0x9f,
	0xcc, 0xbc, 0x6c, 0xef, 0xcc, 0xbc, 0x14, 0x21, 0x0e, 0x2e, 0x16, 0xbf, 0xfc, 0xbc, 0x54, 0x01,
	0x06, 0x21, 0x2e, 0x2e, 0x36, 0xff, 0xa4, 0xac, 0xd4, 0xe4, 0x12, 0x01, 0x46, 0x10, 0x3b, 0x38,
	0x39, 0x23, 0x35, 0x37, 0x51, 0x80, 0x09, 0xc4, 0x76, 0x2a, 0x4d, 0xce, 0x4e, 0x2d, 0x11, 0x60,
	0x76, 0x72, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27,
	0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xb5, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0x90, 0x15, 0xfa, 0x20, 0xbb, 0x74, 0x33, 0xf3, 0xc1, 0xb4, 0x7e, 0x85,
	0x3e, 0xd4, 0x51, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x67, 0x19, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x0f, 0x2f, 0x0f, 0xf7, 0xab, 0x00, 0x00, 0x00,
}
