// Code generated by protoc-gen-go. DO NOT EDIT.
// source: certificate.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	certificate.proto
	claim.proto
	fee.proto
	metadata.proto
	signature.proto
	source.proto
	stream.proto

It has these top-level messages:
	Certificate
	Claim
	Fee
	Metadata
	Signature
	Source
	Stream
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KeyType int32

const (
	KeyType_UNKNOWN_PUBLIC_KEY_TYPE KeyType = 0
	KeyType_NIST256p                KeyType = 1
	KeyType_NIST384p                KeyType = 2
	KeyType_SECP256k1               KeyType = 3
)

var KeyType_name = map[int32]string{
	0: "UNKNOWN_PUBLIC_KEY_TYPE",
	1: "NIST256p",
	2: "NIST384p",
	3: "SECP256k1",
}
var KeyType_value = map[string]int32{
	"UNKNOWN_PUBLIC_KEY_TYPE": 0,
	"NIST256p":                1,
	"NIST384p":                2,
	"SECP256k1":               3,
}

func (x KeyType) Enum() *KeyType {
	p := new(KeyType)
	*p = x
	return p
}
func (x KeyType) String() string {
	return proto.EnumName(KeyType_name, int32(x))
}
func (x *KeyType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(KeyType_value, data, "KeyType")
	if err != nil {
		return err
	}
	*x = KeyType(value)
	return nil
}
func (KeyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Certificate_Version int32

const (
	Certificate_UNKNOWN_VERSION Certificate_Version = 0
	Certificate__0_0_1          Certificate_Version = 1
)

var Certificate_Version_name = map[int32]string{
	0: "UNKNOWN_VERSION",
	1: "_0_0_1",
}
var Certificate_Version_value = map[string]int32{
	"UNKNOWN_VERSION": 0,
	"_0_0_1":          1,
}

func (x Certificate_Version) Enum() *Certificate_Version {
	p := new(Certificate_Version)
	*p = x
	return p
}
func (x Certificate_Version) String() string {
	return proto.EnumName(Certificate_Version_name, int32(x))
}
func (x *Certificate_Version) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Certificate_Version_value, data, "Certificate_Version")
	if err != nil {
		return err
	}
	*x = Certificate_Version(value)
	return nil
}
func (Certificate_Version) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Certificate struct {
	Version          *Certificate_Version `protobuf:"varint,1,req,name=version,enum=pb.Certificate_Version" json:"version,omitempty"`
	KeyType          *KeyType             `protobuf:"varint,2,req,name=keyType,enum=pb.KeyType" json:"keyType,omitempty"`
	PublicKey        []byte               `protobuf:"bytes,4,req,name=publicKey" json:"publicKey,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *Certificate) Reset()                    { *m = Certificate{} }
func (m *Certificate) String() string            { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()               {}
func (*Certificate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Certificate) GetVersion() Certificate_Version {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return Certificate_UNKNOWN_VERSION
}

func (m *Certificate) GetKeyType() KeyType {
	if m != nil && m.KeyType != nil {
		return *m.KeyType
	}
	return KeyType_UNKNOWN_PUBLIC_KEY_TYPE
}

func (m *Certificate) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func init() {
	proto.RegisterType((*Certificate)(nil), "pb.Certificate")
	proto.RegisterEnum("pb.KeyType", KeyType_name, KeyType_value)
	proto.RegisterEnum("pb.Certificate_Version", Certificate_Version_name, Certificate_Version_value)
}

func init() { proto.RegisterFile("certificate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4e, 0x2d, 0x2a,
	0xc9, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a,
	0x48, 0x52, 0xda, 0xc8, 0xc8, 0xc5, 0xed, 0x8c, 0x90, 0x11, 0x32, 0xe4, 0x62, 0x2f, 0x4b, 0x2d,
	0x2a, 0xce, 0xcc, 0xcf, 0x93, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x33, 0x12, 0xd7, 0x2b, 0x48, 0xd2,
	0x43, 0x52, 0xa1, 0x17, 0x06, 0x91, 0x0e, 0x82, 0xa9, 0x13, 0x52, 0xe5, 0x62, 0xcf, 0x4e, 0xad,
	0x0c, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x02, 0x6b, 0xe1, 0x06, 0x69, 0xf1, 0x86, 0x08, 0x05, 0xc1,
	0xe4, 0x84, 0x64, 0xb8, 0x38, 0x0b, 0x4a, 0x93, 0x72, 0x32, 0x93, 0xbd, 0x53, 0x2b, 0x25, 0x58,
	0x14, 0x98, 0x34, 0x78, 0x82, 0x10, 0x02, 0x4a, 0x5a, 0x5c, 0xec, 0x50, 0x83, 0x85, 0x84, 0xb9,
	0xf8, 0x43, 0xfd, 0xbc, 0xfd, 0xfc, 0xc3, 0xfd, 0xe2, 0xc3, 0x5c, 0x83, 0x82, 0x3d, 0xfd, 0xfd,
	0x04, 0x18, 0x84, 0xb8, 0xb8, 0xd8, 0xe2, 0x0d, 0xe2, 0x0d, 0xe2, 0x0d, 0x05, 0x18, 0xb5, 0x02,
	0xb9, 0xd8, 0xa1, 0xa6, 0x0b, 0x49, 0x73, 0x89, 0xc3, 0xd4, 0x06, 0x84, 0x3a, 0xf9, 0x78, 0x3a,
	0xc7, 0x7b, 0xbb, 0x46, 0xc6, 0x87, 0x44, 0x06, 0xb8, 0x0a, 0x30, 0x08, 0xf1, 0x70, 0x71, 0xf8,
	0x79, 0x06, 0x87, 0x18, 0x99, 0x9a, 0x15, 0x08, 0x30, 0xc2, 0x78, 0xc6, 0x16, 0x26, 0x05, 0x02,
	0x4c, 0x42, 0xbc, 0x5c, 0x9c, 0xc1, 0xae, 0xce, 0x01, 0x46, 0xa6, 0x66, 0xd9, 0x86, 0x02, 0xcc,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xd0, 0xba, 0x16, 0x1e, 0x01, 0x00, 0x00,
}
