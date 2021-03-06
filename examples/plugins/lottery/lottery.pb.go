// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/examples/plugins/lottery/lottery.proto

/*
Package lottery is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/examples/plugins/lottery/lottery.proto

It has these top-level messages:
	LotteryInit
*/
package lottery

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type LotteryInit struct {
	Winner *types.Address `protobuf:"bytes,2,opt,name=winner" json:"winner,omitempty"`
}

func (m *LotteryInit) Reset()                    { *m = LotteryInit{} }
func (m *LotteryInit) String() string            { return proto.CompactTextString(m) }
func (*LotteryInit) ProtoMessage()               {}
func (*LotteryInit) Descriptor() ([]byte, []int) { return fileDescriptorLottery, []int{0} }

func (m *LotteryInit) GetWinner() *types.Address {
	if m != nil {
		return m.Winner
	}
	return nil
}

func init() {
	proto.RegisterType((*LotteryInit)(nil), "LotteryInit")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/examples/plugins/lottery/lottery.proto", fileDescriptorLottery)
}

var fileDescriptorLottery = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xc9, 0xcf, 0xcf, 0xcd, 0x4b, 0x2d, 0x29, 0xcf,
	0x2f, 0xca, 0xd6, 0x4f, 0xcf, 0xd7, 0x05, 0x71, 0xf5, 0x53, 0x2b, 0x12, 0x73, 0x0b, 0x72, 0x52,
	0x8b, 0xf5, 0x0b, 0x72, 0x4a, 0xd3, 0x33, 0xf3, 0x8a, 0xf5, 0x73, 0xf2, 0x4b, 0x4a, 0x52, 0x8b,
	0x2a, 0x61, 0xb4, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x94, 0x01, 0x01, 0x63, 0x4a, 0x2a, 0x0b,
	0x52, 0x8b, 0x21, 0x24, 0x44, 0x87, 0x92, 0x3e, 0x17, 0xb7, 0x0f, 0xc4, 0x08, 0xcf, 0xbc, 0xcc,
	0x12, 0x21, 0x05, 0x2e, 0xb6, 0xf2, 0xcc, 0xbc, 0xbc, 0xd4, 0x22, 0x09, 0x26, 0x05, 0x46, 0x0d,
	0x6e, 0x23, 0x0e, 0x3d, 0xc7, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0xe2, 0x20, 0xa8, 0x78, 0x12, 0x1b,
	0x58, 0x9f, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x42, 0x4d, 0xcd, 0xb2, 0x00, 0x00, 0x00,
}
