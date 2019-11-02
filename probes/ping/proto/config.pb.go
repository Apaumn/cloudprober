// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/probes/ping/proto/config.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Next tag: 1
type ProbeConf struct {
	// Packets per probe
	PacketsPerProbe *int32 `protobuf:"varint,6,opt,name=packets_per_probe,json=packetsPerProbe,def=2" json:"packets_per_probe,omitempty"`
	// How long to wait between two packets to the same target
	PacketsIntervalMsec *int32 `protobuf:"varint,7,opt,name=packets_interval_msec,json=packetsIntervalMsec,def=25" json:"packets_interval_msec,omitempty"`
	// Resolve targets after these many probes
	ResolveTargetsInterval *int32 `protobuf:"varint,9,opt,name=resolve_targets_interval,json=resolveTargetsInterval,def=5" json:"resolve_targets_interval,omitempty"`
	// Ping payload size in bytes. It cannot be smaller than 8, number of bytes
	// required for the nanoseconds timestamp.
	PayloadSize *int32 `protobuf:"varint,10,opt,name=payload_size,json=payloadSize,def=56" json:"payload_size,omitempty"`
	// Use datagram socket for ICMP.
	// This option enables unprivileged pings (that is, you don't require root
	// privilege to send ICMP packets). Note that most of the Linux distributions
	// don't allow unprivileged pings by default. To enable unprivileged pings on
	// some Linux distributions, you may need to run the following command:
	//     sudo sysctl -w net.ipv4.ping_group_range="0 <large valid group id>"
	// net.ipv4.ping_group_range system setting takes two integers that specify
	// the group id range that is allowed to execute the unprivileged pings. Note
	// that the same setting (with ipv4 in the path) applies to IPv6 as well.
	UseDatagramSocket *bool `protobuf:"varint,12,opt,name=use_datagram_socket,json=useDatagramSocket,def=1" json:"use_datagram_socket,omitempty"`
	// Disable integrity checks. To detect data courruption in the network, we
	// craft the outgoing ICMP packet payload in a certain format and verify that
	// the reply payload matches the same format.
	DisableIntegrityCheck *bool    `protobuf:"varint,13,opt,name=disable_integrity_check,json=disableIntegrityCheck,def=0" json:"disable_integrity_check,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ProbeConf) Reset()         { *m = ProbeConf{} }
func (m *ProbeConf) String() string { return proto.CompactTextString(m) }
func (*ProbeConf) ProtoMessage()    {}
func (*ProbeConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_07f3181ec59da86e, []int{0}
}

func (m *ProbeConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeConf.Unmarshal(m, b)
}
func (m *ProbeConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeConf.Marshal(b, m, deterministic)
}
func (m *ProbeConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeConf.Merge(m, src)
}
func (m *ProbeConf) XXX_Size() int {
	return xxx_messageInfo_ProbeConf.Size(m)
}
func (m *ProbeConf) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeConf.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeConf proto.InternalMessageInfo

const Default_ProbeConf_PacketsPerProbe int32 = 2
const Default_ProbeConf_PacketsIntervalMsec int32 = 25
const Default_ProbeConf_ResolveTargetsInterval int32 = 5
const Default_ProbeConf_PayloadSize int32 = 56
const Default_ProbeConf_UseDatagramSocket bool = true
const Default_ProbeConf_DisableIntegrityCheck bool = false

func (m *ProbeConf) GetPacketsPerProbe() int32 {
	if m != nil && m.PacketsPerProbe != nil {
		return *m.PacketsPerProbe
	}
	return Default_ProbeConf_PacketsPerProbe
}

func (m *ProbeConf) GetPacketsIntervalMsec() int32 {
	if m != nil && m.PacketsIntervalMsec != nil {
		return *m.PacketsIntervalMsec
	}
	return Default_ProbeConf_PacketsIntervalMsec
}

func (m *ProbeConf) GetResolveTargetsInterval() int32 {
	if m != nil && m.ResolveTargetsInterval != nil {
		return *m.ResolveTargetsInterval
	}
	return Default_ProbeConf_ResolveTargetsInterval
}

func (m *ProbeConf) GetPayloadSize() int32 {
	if m != nil && m.PayloadSize != nil {
		return *m.PayloadSize
	}
	return Default_ProbeConf_PayloadSize
}

func (m *ProbeConf) GetUseDatagramSocket() bool {
	if m != nil && m.UseDatagramSocket != nil {
		return *m.UseDatagramSocket
	}
	return Default_ProbeConf_UseDatagramSocket
}

func (m *ProbeConf) GetDisableIntegrityCheck() bool {
	if m != nil && m.DisableIntegrityCheck != nil {
		return *m.DisableIntegrityCheck
	}
	return Default_ProbeConf_DisableIntegrityCheck
}

func init() {
	proto.RegisterType((*ProbeConf)(nil), "cloudprober.probes.ping.ProbeConf")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/probes/ping/proto/config.proto", fileDescriptor_07f3181ec59da86e)
}

var fileDescriptor_07f3181ec59da86e = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xce, 0x41, 0x6b, 0xc2, 0x30,
	0x14, 0x07, 0x70, 0x94, 0xb9, 0xcd, 0xcc, 0x31, 0x8c, 0x38, 0x73, 0x94, 0xc1, 0xc0, 0xcb, 0x5a,
	0x90, 0xe9, 0xc1, 0x6d, 0x27, 0x77, 0xf1, 0x30, 0x90, 0xb8, 0x7b, 0x88, 0xe9, 0x33, 0x06, 0x63,
	0x53, 0x92, 0x54, 0xd0, 0xaf, 0xb8, 0x2f, 0x35, 0x9a, 0xa6, 0xe0, 0xa9, 0x7d, 0xfd, 0xff, 0xfe,
	0x7d, 0x0f, 0x7d, 0x4a, 0xe5, 0xf7, 0xe5, 0x36, 0x11, 0xe6, 0x98, 0x4a, 0x63, 0xa4, 0x86, 0x54,
	0x68, 0x53, 0x66, 0x85, 0x35, 0x5b, 0xb0, 0x69, 0x78, 0xb8, 0xb4, 0x50, 0xb9, 0xac, 0xde, 0xbd,
	0x49, 0x85, 0xc9, 0x77, 0x4a, 0x26, 0x61, 0xc0, 0xa3, 0x2b, 0x9b, 0xd4, 0x36, 0xa9, 0xec, 0xcb,
	0x5f, 0x1b, 0x75, 0xd7, 0xd5, 0xbc, 0x34, 0xf9, 0x0e, 0xbf, 0xa1, 0x7e, 0xc1, 0xc5, 0x01, 0xbc,
	0x63, 0x05, 0x58, 0x16, 0x20, 0xb9, 0x1d, 0xb7, 0x26, 0x9d, 0x45, 0x6b, 0x4a, 0x9f, 0x62, 0xb6,
	0x06, 0x1b, 0x2a, 0x78, 0x8e, 0x86, 0x0d, 0x57, 0xb9, 0x07, 0x7b, 0xe2, 0x9a, 0x1d, 0x1d, 0x08,
	0x72, 0x17, 0x2a, 0xed, 0xe9, 0x8c, 0x0e, 0x22, 0x58, 0xc5, 0xfc, 0xc7, 0x81, 0xc0, 0x1f, 0x88,
	0x58, 0x70, 0x46, 0x9f, 0x80, 0x79, 0x6e, 0xe5, 0x75, 0x9f, 0x74, 0xeb, 0x6d, 0x33, 0xfa, 0x1c,
	0xc9, 0x6f, 0x2d, 0x9a, 0x1f, 0xe0, 0x57, 0xd4, 0x2b, 0xf8, 0x59, 0x1b, 0x9e, 0x31, 0xa7, 0x2e,
	0x40, 0x50, 0xbd, 0x6b, 0x36, 0xa7, 0x0f, 0xf1, 0xfb, 0x46, 0x5d, 0x00, 0xbf, 0xa3, 0x41, 0xe9,
	0x80, 0x65, 0xdc, 0x73, 0x69, 0xf9, 0x91, 0x39, 0x53, 0xdd, 0x41, 0x7a, 0xe3, 0xd6, 0xe4, 0x7e,
	0x71, 0xe3, 0x6d, 0x09, 0xb4, 0x5f, 0x3a, 0xf8, 0x8e, 0xf9, 0x26, 0xc4, 0xf8, 0x0b, 0x8d, 0x32,
	0xe5, 0xf8, 0x56, 0x43, 0xb8, 0x48, 0x5a, 0xe5, 0xcf, 0x4c, 0xec, 0x41, 0x1c, 0xc8, 0x63, 0x68,
	0x76, 0x76, 0x5c, 0x3b, 0xa0, 0xc3, 0xa8, 0x56, 0x0d, 0x5a, 0x56, 0xe6, 0x3f, 0x00, 0x00, 0xff,
	0xff, 0xe4, 0x5d, 0x49, 0x99, 0xa5, 0x01, 0x00, 0x00,
}