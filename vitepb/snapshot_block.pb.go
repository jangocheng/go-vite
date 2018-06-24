// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snapshot_block.proto

package vitepb

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

type SnapshotBlock struct {
	PrevHash             []byte            `protobuf:"bytes,1,opt,name=prevHash,proto3" json:"prevHash,omitempty"`
	BlockNum             []byte            `protobuf:"bytes,2,opt,name=blockNum,proto3" json:"blockNum,omitempty"`
	Producer             []byte            `protobuf:"bytes,3,opt,name=producer,proto3" json:"producer,omitempty"`
	Snapshot             map[string][]byte `protobuf:"bytes,4,rep,name=snapshot,proto3" json:"snapshot,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Signature            []byte            `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SnapshotBlock) Reset()         { *m = SnapshotBlock{} }
func (m *SnapshotBlock) String() string { return proto.CompactTextString(m) }
func (*SnapshotBlock) ProtoMessage()    {}
func (*SnapshotBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_snapshot_block_cb79d60975da5157, []int{0}
}
func (m *SnapshotBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotBlock.Unmarshal(m, b)
}
func (m *SnapshotBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotBlock.Marshal(b, m, deterministic)
}
func (dst *SnapshotBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotBlock.Merge(dst, src)
}
func (m *SnapshotBlock) XXX_Size() int {
	return xxx_messageInfo_SnapshotBlock.Size(m)
}
func (m *SnapshotBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotBlock.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotBlock proto.InternalMessageInfo

func (m *SnapshotBlock) GetPrevHash() []byte {
	if m != nil {
		return m.PrevHash
	}
	return nil
}

func (m *SnapshotBlock) GetBlockNum() []byte {
	if m != nil {
		return m.BlockNum
	}
	return nil
}

func (m *SnapshotBlock) GetProducer() []byte {
	if m != nil {
		return m.Producer
	}
	return nil
}

func (m *SnapshotBlock) GetSnapshot() map[string][]byte {
	if m != nil {
		return m.Snapshot
	}
	return nil
}

func (m *SnapshotBlock) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*SnapshotBlock)(nil), "vitepb.SnapshotBlock")
	proto.RegisterMapType((map[string][]byte)(nil), "vitepb.SnapshotBlock.SnapshotEntry")
}

func init() {
	proto.RegisterFile("snapshot_block.proto", fileDescriptor_snapshot_block_cb79d60975da5157)
}

var fileDescriptor_snapshot_block_cb79d60975da5157 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0xce, 0x4b, 0x2c,
	0x28, 0xce, 0xc8, 0x2f, 0x89, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x2b, 0xcb, 0x2c, 0x49, 0x2d, 0x48, 0x52, 0xfa, 0xcf, 0xc8, 0xc5, 0x1b, 0x0c, 0x55,
	0xe0, 0x04, 0x92, 0x17, 0x92, 0xe2, 0xe2, 0x28, 0x28, 0x4a, 0x2d, 0xf3, 0x48, 0x2c, 0xce, 0x90,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0xf3, 0x41, 0x72, 0x60, 0x43, 0xfc, 0x4a, 0x73, 0x25,
	0x98, 0x20, 0x72, 0x30, 0x3e, 0x44, 0x5f, 0x7e, 0x4a, 0x69, 0x72, 0x6a, 0x91, 0x04, 0x33, 0x4c,
	0x1f, 0x84, 0x2f, 0x64, 0xcf, 0xc5, 0x01, 0x73, 0x85, 0x04, 0x8b, 0x02, 0xb3, 0x06, 0xb7, 0x91,
	0xb2, 0x1e, 0xc4, 0x01, 0x7a, 0x28, 0x96, 0xc3, 0x79, 0xae, 0x79, 0x25, 0x45, 0x95, 0x41, 0x70,
	0x4d, 0x42, 0x32, 0x5c, 0x9c, 0xc5, 0x99, 0xe9, 0x79, 0x89, 0x25, 0xa5, 0x45, 0xa9, 0x12, 0xac,
	0x60, 0xd3, 0x11, 0x02, 0x52, 0xd6, 0x08, 0x3f, 0x80, 0x35, 0x0a, 0x09, 0x70, 0x31, 0x67, 0xa7,
	0x56, 0x82, 0x9d, 0xcf, 0x19, 0x04, 0x62, 0x0a, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6,
	0x42, 0x9d, 0x0d, 0xe1, 0x58, 0x31, 0x59, 0x30, 0x26, 0xb1, 0x81, 0x03, 0xc4, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xa7, 0x46, 0x64, 0x4f, 0x28, 0x01, 0x00, 0x00,
}