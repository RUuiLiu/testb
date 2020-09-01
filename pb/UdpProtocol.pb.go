// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.12.3
// source: UdpProtocol.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CommandCode int32

const (
	CommandCode_NONE           CommandCode = 1
	CommandCode_ACK            CommandCode = 2  //确认
	CommandCode_CONNECT        CommandCode = 3  //连接请求
	CommandCode_ACKCONNECT     CommandCode = 4  //连接请求确认
	CommandCode_DISCONNECT     CommandCode = 5  //断线请求
	CommandCode_ACKDISCONNECT  CommandCode = 6  //断线请求确认
	CommandCode_PING           CommandCode = 7  //ping值
	CommandCode_SENDRELIABLE   CommandCode = 8  //可靠消息
	CommandCode_SENDUNRELIABLE CommandCode = 9  //不可靠消息
	CommandCode_SENDFRAGMENT   CommandCode = 10 //片段消息
	CommandCode_COUNT          CommandCode = 11
)

// Enum value maps for CommandCode.
var (
	CommandCode_name = map[int32]string{
		1:  "NONE",
		2:  "ACK",
		3:  "CONNECT",
		4:  "ACKCONNECT",
		5:  "DISCONNECT",
		6:  "ACKDISCONNECT",
		7:  "PING",
		8:  "SENDRELIABLE",
		9:  "SENDUNRELIABLE",
		10: "SENDFRAGMENT",
		11: "COUNT",
	}
	CommandCode_value = map[string]int32{
		"NONE":           1,
		"ACK":            2,
		"CONNECT":        3,
		"ACKCONNECT":     4,
		"DISCONNECT":     5,
		"ACKDISCONNECT":  6,
		"PING":           7,
		"SENDRELIABLE":   8,
		"SENDUNRELIABLE": 9,
		"SENDFRAGMENT":   10,
		"COUNT":          11,
	}
)

func (x CommandCode) Enum() *CommandCode {
	p := new(CommandCode)
	*p = x
	return p
}

func (x CommandCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandCode) Descriptor() protoreflect.EnumDescriptor {
	return file_UdpProtocol_proto_enumTypes[0].Descriptor()
}

func (CommandCode) Type() protoreflect.EnumType {
	return &file_UdpProtocol_proto_enumTypes[0]
}

func (x CommandCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CommandCode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CommandCode(num)
	return nil
}

// Deprecated: Use CommandCode.Descriptor instead.
func (CommandCode) EnumDescriptor() ([]byte, []int) {
	return file_UdpProtocol_proto_rawDescGZIP(), []int{0}
}

type UdpHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerID   *int32 `protobuf:"varint,1,req,name=peerID" json:"peerID,omitempty"`
	SentTime *int64 `protobuf:"varint,2,req,name=sentTime" json:"sentTime,omitempty"`
	Uno      *int32 `protobuf:"varint,3,req,name=uno" json:"uno,omitempty"`
}

func (x *UdpHeader) Reset() {
	*x = UdpHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UdpProtocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UdpHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UdpHeader) ProtoMessage() {}

func (x *UdpHeader) ProtoReflect() protoreflect.Message {
	mi := &file_UdpProtocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UdpHeader.ProtoReflect.Descriptor instead.
func (*UdpHeader) Descriptor() ([]byte, []int) {
	return file_UdpProtocol_proto_rawDescGZIP(), []int{0}
}

func (x *UdpHeader) GetPeerID() int32 {
	if x != nil && x.PeerID != nil {
		return *x.PeerID
	}
	return 0
}

func (x *UdpHeader) GetSentTime() int64 {
	if x != nil && x.SentTime != nil {
		return *x.SentTime
	}
	return 0
}

func (x *UdpHeader) GetUno() int32 {
	if x != nil && x.Uno != nil {
		return *x.Uno
	}
	return 0
}

type CommandHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandCode *CommandCode `protobuf:"varint,1,req,name=commandCode,enum=pb.CommandCode" json:"commandCode,omitempty"`
	Flag        *int32       `protobuf:"varint,2,req,name=flag" json:"flag,omitempty"`
}

func (x *CommandHeader) Reset() {
	*x = CommandHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UdpProtocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandHeader) ProtoMessage() {}

func (x *CommandHeader) ProtoReflect() protoreflect.Message {
	mi := &file_UdpProtocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandHeader.ProtoReflect.Descriptor instead.
func (*CommandHeader) Descriptor() ([]byte, []int) {
	return file_UdpProtocol_proto_rawDescGZIP(), []int{1}
}

func (x *CommandHeader) GetCommandCode() CommandCode {
	if x != nil && x.CommandCode != nil {
		return *x.CommandCode
	}
	return CommandCode_NONE
}

func (x *CommandHeader) GetFlag() int32 {
	if x != nil && x.Flag != nil {
		return *x.Flag
	}
	return 0
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn         *int32  `protobuf:"varint,1,req,name=sn" json:"sn,omitempty"`
	ExpectedSn []int32 `protobuf:"varint,2,rep,name=expectedSn" json:"expectedSn,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UdpProtocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_UdpProtocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_UdpProtocol_proto_rawDescGZIP(), []int{2}
}

func (x *Ack) GetSn() int32 {
	if x != nil && x.Sn != nil {
		return *x.Sn
	}
	return 0
}

func (x *Ack) GetExpectedSn() []int32 {
	if x != nil {
		return x.ExpectedSn
	}
	return nil
}

type Connect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientPeerID *int32 `protobuf:"varint,1,req,name=clientPeerID" json:"clientPeerID,omitempty"`
	Mtu          *int32 `protobuf:"varint,2,req,name=mtu" json:"mtu,omitempty"`
	ConnectID    *int32 `protobuf:"varint,3,req,name=connectID" json:"connectID,omitempty"`
	Data         []byte `protobuf:"bytes,4,opt,name=data" json:"data,omitempty"`
}

func (x *Connect) Reset() {
	*x = Connect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UdpProtocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connect) ProtoMessage() {}

func (x *Connect) ProtoReflect() protoreflect.Message {
	mi := &file_UdpProtocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connect.ProtoReflect.Descriptor instead.
func (*Connect) Descriptor() ([]byte, []int) {
	return file_UdpProtocol_proto_rawDescGZIP(), []int{3}
}

func (x *Connect) GetClientPeerID() int32 {
	if x != nil && x.ClientPeerID != nil {
		return *x.ClientPeerID
	}
	return 0
}

func (x *Connect) GetMtu() int32 {
	if x != nil && x.Mtu != nil {
		return *x.Mtu
	}
	return 0
}

func (x *Connect) GetConnectID() int32 {
	if x != nil && x.ConnectID != nil {
		return *x.ConnectID
	}
	return 0
}

func (x *Connect) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type AckConnect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerID    *int32  `protobuf:"varint,1,req,name=peerID" json:"peerID,omitempty"`
	Mtu       *int32  `protobuf:"varint,2,req,name=mtu" json:"mtu,omitempty"`
	ConnectID []int32 `protobuf:"varint,3,rep,name=connectID" json:"connectID,omitempty"`
}

func (x *AckConnect) Reset() {
	*x = AckConnect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UdpProtocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckConnect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckConnect) ProtoMessage() {}

func (x *AckConnect) ProtoReflect() protoreflect.Message {
	mi := &file_UdpProtocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckConnect.ProtoReflect.Descriptor instead.
func (*AckConnect) Descriptor() ([]byte, []int) {
	return file_UdpProtocol_proto_rawDescGZIP(), []int{4}
}

func (x *AckConnect) GetPeerID() int32 {
	if x != nil && x.PeerID != nil {
		return *x.PeerID
	}
	return 0
}

func (x *AckConnect) GetMtu() int32 {
	if x != nil && x.Mtu != nil {
		return *x.Mtu
	}
	return 0
}

func (x *AckConnect) GetConnectID() []int32 {
	if x != nil {
		return x.ConnectID
	}
	return nil
}

type Disconnect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason *int32 `protobuf:"varint,1,opt,name=reason" json:"reason,omitempty"`
}

func (x *Disconnect) Reset() {
	*x = Disconnect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UdpProtocol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Disconnect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Disconnect) ProtoMessage() {}

func (x *Disconnect) ProtoReflect() protoreflect.Message {
	mi := &file_UdpProtocol_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Disconnect.ProtoReflect.Descriptor instead.
func (*Disconnect) Descriptor() ([]byte, []int) {
	return file_UdpProtocol_proto_rawDescGZIP(), []int{5}
}

func (x *Disconnect) GetReason() int32 {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return 0
}

type Send struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn         *int32 `protobuf:"varint,1,req,name=sn" json:"sn,omitempty"`
	DataLength *int32 `protobuf:"varint,2,req,name=dataLength" json:"dataLength,omitempty"`
}

func (x *Send) Reset() {
	*x = Send{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UdpProtocol_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Send) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Send) ProtoMessage() {}

func (x *Send) ProtoReflect() protoreflect.Message {
	mi := &file_UdpProtocol_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Send.ProtoReflect.Descriptor instead.
func (*Send) Descriptor() ([]byte, []int) {
	return file_UdpProtocol_proto_rawDescGZIP(), []int{6}
}

func (x *Send) GetSn() int32 {
	if x != nil && x.Sn != nil {
		return *x.Sn
	}
	return 0
}

func (x *Send) GetDataLength() int32 {
	if x != nil && x.DataLength != nil {
		return *x.DataLength
	}
	return 0
}

type SendMultiFragment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalLength *int64 `protobuf:"varint,1,req,name=totalLength" json:"totalLength,omitempty"`
	FragCount   *int32 `protobuf:"varint,2,req,name=fragCount" json:"fragCount,omitempty"`
	FragSn      *int32 `protobuf:"varint,3,req,name=fragSn" json:"fragSn,omitempty"`
}

func (x *SendMultiFragment) Reset() {
	*x = SendMultiFragment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UdpProtocol_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMultiFragment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMultiFragment) ProtoMessage() {}

func (x *SendMultiFragment) ProtoReflect() protoreflect.Message {
	mi := &file_UdpProtocol_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMultiFragment.ProtoReflect.Descriptor instead.
func (*SendMultiFragment) Descriptor() ([]byte, []int) {
	return file_UdpProtocol_proto_rawDescGZIP(), []int{7}
}

func (x *SendMultiFragment) GetTotalLength() int64 {
	if x != nil && x.TotalLength != nil {
		return *x.TotalLength
	}
	return 0
}

func (x *SendMultiFragment) GetFragCount() int32 {
	if x != nil && x.FragCount != nil {
		return *x.FragCount
	}
	return 0
}

func (x *SendMultiFragment) GetFragSn() int32 {
	if x != nil && x.FragSn != nil {
		return *x.FragSn
	}
	return 0
}

var File_UdpProtocol_proto protoreflect.FileDescriptor

var file_UdpProtocol_proto_rawDesc = []byte{
	0x0a, 0x11, 0x55, 0x64, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x51, 0x0a, 0x09, 0x55, 0x64, 0x70, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03, 0x52, 0x08,
	0x73, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x6e, 0x6f, 0x18,
	0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x75, 0x6e, 0x6f, 0x22, 0x56, 0x0a, 0x0d, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x04, 0x66, 0x6c,
	0x61, 0x67, 0x22, 0x35, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x65,
	0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x6e, 0x22, 0x71, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x65,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x50, 0x65, 0x65, 0x72, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x74, 0x75, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x74, 0x75, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x54, 0x0a, 0x0a,
	0x41, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x74, 0x75, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52,
	0x03, 0x6d, 0x74, 0x75, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x49, 0x44, 0x22, 0x24, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x36, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x02, 0x73, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x22, 0x6b, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x46, 0x72, 0x61,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x61, 0x67, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x09, 0x66, 0x72, 0x61, 0x67,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72, 0x61, 0x67, 0x53, 0x6e, 0x18,
	0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x66, 0x72, 0x61, 0x67, 0x53, 0x6e, 0x2a, 0xad, 0x01,
	0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x43, 0x4b, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x03, 0x12, 0x0e, 0x0a,
	0x0a, 0x41, 0x43, 0x4b, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x04, 0x12, 0x0e, 0x0a,
	0x0a, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x05, 0x12, 0x11, 0x0a,
	0x0d, 0x41, 0x43, 0x4b, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x06,
	0x12, 0x08, 0x0a, 0x04, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45,
	0x4e, 0x44, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x08, 0x12, 0x12, 0x0a, 0x0e,
	0x53, 0x45, 0x4e, 0x44, 0x55, 0x4e, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x09,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x4e, 0x44, 0x46, 0x52, 0x41, 0x47, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x0b,
}

var (
	file_UdpProtocol_proto_rawDescOnce sync.Once
	file_UdpProtocol_proto_rawDescData = file_UdpProtocol_proto_rawDesc
)

func file_UdpProtocol_proto_rawDescGZIP() []byte {
	file_UdpProtocol_proto_rawDescOnce.Do(func() {
		file_UdpProtocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_UdpProtocol_proto_rawDescData)
	})
	return file_UdpProtocol_proto_rawDescData
}

var file_UdpProtocol_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_UdpProtocol_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_UdpProtocol_proto_goTypes = []interface{}{
	(CommandCode)(0),          // 0: pb.CommandCode
	(*UdpHeader)(nil),         // 1: pb.UdpHeader
	(*CommandHeader)(nil),     // 2: pb.CommandHeader
	(*Ack)(nil),               // 3: pb.Ack
	(*Connect)(nil),           // 4: pb.Connect
	(*AckConnect)(nil),        // 5: pb.AckConnect
	(*Disconnect)(nil),        // 6: pb.Disconnect
	(*Send)(nil),              // 7: pb.Send
	(*SendMultiFragment)(nil), // 8: pb.SendMultiFragment
}
var file_UdpProtocol_proto_depIdxs = []int32{
	0, // 0: pb.CommandHeader.commandCode:type_name -> pb.CommandCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_UdpProtocol_proto_init() }
func file_UdpProtocol_proto_init() {
	if File_UdpProtocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_UdpProtocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UdpHeader); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_UdpProtocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandHeader); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_UdpProtocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_UdpProtocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connect); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_UdpProtocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckConnect); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_UdpProtocol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Disconnect); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_UdpProtocol_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Send); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_UdpProtocol_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMultiFragment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_UdpProtocol_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UdpProtocol_proto_goTypes,
		DependencyIndexes: file_UdpProtocol_proto_depIdxs,
		EnumInfos:         file_UdpProtocol_proto_enumTypes,
		MessageInfos:      file_UdpProtocol_proto_msgTypes,
	}.Build()
	File_UdpProtocol_proto = out.File
	file_UdpProtocol_proto_rawDesc = nil
	file_UdpProtocol_proto_goTypes = nil
	file_UdpProtocol_proto_depIdxs = nil
}