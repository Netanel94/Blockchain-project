// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lottery/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgCreateLotteryItem struct {
	Creator      string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Amount       int32    `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Participants []string `protobuf:"bytes,4,rep,name=participants,proto3" json:"participants,omitempty"`
}

func (m *MsgCreateLotteryItem) Reset()         { *m = MsgCreateLotteryItem{} }
func (m *MsgCreateLotteryItem) String() string { return proto.CompactTextString(m) }
func (*MsgCreateLotteryItem) ProtoMessage()    {}
func (*MsgCreateLotteryItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_16a4f365e1c6455d, []int{0}
}
func (m *MsgCreateLotteryItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateLotteryItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateLotteryItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateLotteryItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateLotteryItem.Merge(m, src)
}
func (m *MsgCreateLotteryItem) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateLotteryItem) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateLotteryItem.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateLotteryItem proto.InternalMessageInfo

func (m *MsgCreateLotteryItem) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateLotteryItem) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *MsgCreateLotteryItem) GetParticipants() []string {
	if m != nil {
		return m.Participants
	}
	return nil
}

type MsgCreateLotteryItemResponse struct {
}

func (m *MsgCreateLotteryItemResponse) Reset()         { *m = MsgCreateLotteryItemResponse{} }
func (m *MsgCreateLotteryItemResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateLotteryItemResponse) ProtoMessage()    {}
func (*MsgCreateLotteryItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_16a4f365e1c6455d, []int{1}
}
func (m *MsgCreateLotteryItemResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateLotteryItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateLotteryItemResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateLotteryItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateLotteryItemResponse.Merge(m, src)
}
func (m *MsgCreateLotteryItemResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateLotteryItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateLotteryItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateLotteryItemResponse proto.InternalMessageInfo

type MsgUpdateLotteryItem struct {
	Creator      string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Amount       int32    `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Participants []string `protobuf:"bytes,4,rep,name=participants,proto3" json:"participants,omitempty"`
}

func (m *MsgUpdateLotteryItem) Reset()         { *m = MsgUpdateLotteryItem{} }
func (m *MsgUpdateLotteryItem) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateLotteryItem) ProtoMessage()    {}
func (*MsgUpdateLotteryItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_16a4f365e1c6455d, []int{2}
}
func (m *MsgUpdateLotteryItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateLotteryItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateLotteryItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateLotteryItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateLotteryItem.Merge(m, src)
}
func (m *MsgUpdateLotteryItem) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateLotteryItem) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateLotteryItem.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateLotteryItem proto.InternalMessageInfo

func (m *MsgUpdateLotteryItem) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateLotteryItem) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *MsgUpdateLotteryItem) GetParticipants() []string {
	if m != nil {
		return m.Participants
	}
	return nil
}

type MsgUpdateLotteryItemResponse struct {
}

func (m *MsgUpdateLotteryItemResponse) Reset()         { *m = MsgUpdateLotteryItemResponse{} }
func (m *MsgUpdateLotteryItemResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateLotteryItemResponse) ProtoMessage()    {}
func (*MsgUpdateLotteryItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_16a4f365e1c6455d, []int{3}
}
func (m *MsgUpdateLotteryItemResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateLotteryItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateLotteryItemResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateLotteryItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateLotteryItemResponse.Merge(m, src)
}
func (m *MsgUpdateLotteryItemResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateLotteryItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateLotteryItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateLotteryItemResponse proto.InternalMessageInfo

type MsgDeleteLotteryItem struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgDeleteLotteryItem) Reset()         { *m = MsgDeleteLotteryItem{} }
func (m *MsgDeleteLotteryItem) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteLotteryItem) ProtoMessage()    {}
func (*MsgDeleteLotteryItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_16a4f365e1c6455d, []int{4}
}
func (m *MsgDeleteLotteryItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteLotteryItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteLotteryItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteLotteryItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteLotteryItem.Merge(m, src)
}
func (m *MsgDeleteLotteryItem) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteLotteryItem) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteLotteryItem.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteLotteryItem proto.InternalMessageInfo

func (m *MsgDeleteLotteryItem) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgDeleteLotteryItemResponse struct {
}

func (m *MsgDeleteLotteryItemResponse) Reset()         { *m = MsgDeleteLotteryItemResponse{} }
func (m *MsgDeleteLotteryItemResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteLotteryItemResponse) ProtoMessage()    {}
func (*MsgDeleteLotteryItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_16a4f365e1c6455d, []int{5}
}
func (m *MsgDeleteLotteryItemResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteLotteryItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteLotteryItemResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteLotteryItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteLotteryItemResponse.Merge(m, src)
}
func (m *MsgDeleteLotteryItemResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteLotteryItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteLotteryItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteLotteryItemResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateLotteryItem)(nil), "lottery.lottery.MsgCreateLotteryItem")
	proto.RegisterType((*MsgCreateLotteryItemResponse)(nil), "lottery.lottery.MsgCreateLotteryItemResponse")
	proto.RegisterType((*MsgUpdateLotteryItem)(nil), "lottery.lottery.MsgUpdateLotteryItem")
	proto.RegisterType((*MsgUpdateLotteryItemResponse)(nil), "lottery.lottery.MsgUpdateLotteryItemResponse")
	proto.RegisterType((*MsgDeleteLotteryItem)(nil), "lottery.lottery.MsgDeleteLotteryItem")
	proto.RegisterType((*MsgDeleteLotteryItemResponse)(nil), "lottery.lottery.MsgDeleteLotteryItemResponse")
}

func init() { proto.RegisterFile("lottery/tx.proto", fileDescriptor_16a4f365e1c6455d) }

var fileDescriptor_16a4f365e1c6455d = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xc9, 0x2f, 0x29,
	0x49, 0x2d, 0xaa, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x87, 0x8a,
	0xe8, 0x41, 0x69, 0x29, 0x29, 0x98, 0x12, 0x28, 0x1d, 0x9f, 0x59, 0x92, 0x9a, 0x0b, 0x51, 0xac,
	0x94, 0xc3, 0x25, 0xe2, 0x5b, 0x9c, 0xee, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0xea, 0x03, 0x91, 0xf6,
	0x2c, 0x49, 0xcd, 0x15, 0x92, 0xe0, 0x62, 0x4f, 0x06, 0x09, 0xe6, 0x17, 0x49, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42, 0x62, 0x5c, 0x6c, 0x89, 0xb9, 0xf9, 0xa5, 0x79, 0x25, 0x12,
	0xcc, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x50, 0x9e, 0x90, 0x12, 0x17, 0x4f, 0x41, 0x62, 0x51, 0x49,
	0x66, 0x72, 0x66, 0x41, 0x62, 0x5e, 0x49, 0xb1, 0x04, 0x8b, 0x02, 0xb3, 0x06, 0x67, 0x10, 0x8a,
	0x98, 0x92, 0x1c, 0x97, 0x0c, 0x36, 0xdb, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0xa1,
	0xae, 0x09, 0x2d, 0x48, 0xa1, 0xa3, 0x6b, 0x30, 0x6c, 0x83, 0xbb, 0xc6, 0x00, 0xec, 0x1a, 0x97,
	0xd4, 0x9c, 0x54, 0x22, 0x5d, 0x03, 0x35, 0x11, 0x43, 0x07, 0xcc, 0x44, 0xa3, 0x63, 0x4c, 0x5c,
	0xcc, 0xbe, 0xc5, 0xe9, 0x42, 0x99, 0x5c, 0x82, 0x98, 0x41, 0xae, 0xaa, 0x87, 0x16, 0x71, 0x7a,
	0xd8, 0xc2, 0x4a, 0x4a, 0x97, 0x28, 0x65, 0x30, 0x2b, 0x41, 0x56, 0x61, 0x86, 0x27, 0x56, 0xab,
	0x30, 0x94, 0x61, 0xb7, 0x0a, 0x67, 0x78, 0x81, 0xac, 0xc2, 0x0c, 0x2c, 0xac, 0x56, 0x61, 0x28,
	0xc3, 0x6e, 0x15, 0xce, 0x80, 0x74, 0x32, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x86, 0x28, 0x71, 0x58, 0x62, 0xaf, 0xd0, 0x87, 0xe7, 0x8c, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36,
	0x70, 0x82, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xd9, 0xa0, 0xf4, 0x31, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateLotteryItem(ctx context.Context, in *MsgCreateLotteryItem, opts ...grpc.CallOption) (*MsgCreateLotteryItemResponse, error)
	UpdateLotteryItem(ctx context.Context, in *MsgUpdateLotteryItem, opts ...grpc.CallOption) (*MsgUpdateLotteryItemResponse, error)
	DeleteLotteryItem(ctx context.Context, in *MsgDeleteLotteryItem, opts ...grpc.CallOption) (*MsgDeleteLotteryItemResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateLotteryItem(ctx context.Context, in *MsgCreateLotteryItem, opts ...grpc.CallOption) (*MsgCreateLotteryItemResponse, error) {
	out := new(MsgCreateLotteryItemResponse)
	err := c.cc.Invoke(ctx, "/lottery.lottery.Msg/CreateLotteryItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateLotteryItem(ctx context.Context, in *MsgUpdateLotteryItem, opts ...grpc.CallOption) (*MsgUpdateLotteryItemResponse, error) {
	out := new(MsgUpdateLotteryItemResponse)
	err := c.cc.Invoke(ctx, "/lottery.lottery.Msg/UpdateLotteryItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteLotteryItem(ctx context.Context, in *MsgDeleteLotteryItem, opts ...grpc.CallOption) (*MsgDeleteLotteryItemResponse, error) {
	out := new(MsgDeleteLotteryItemResponse)
	err := c.cc.Invoke(ctx, "/lottery.lottery.Msg/DeleteLotteryItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateLotteryItem(context.Context, *MsgCreateLotteryItem) (*MsgCreateLotteryItemResponse, error)
	UpdateLotteryItem(context.Context, *MsgUpdateLotteryItem) (*MsgUpdateLotteryItemResponse, error)
	DeleteLotteryItem(context.Context, *MsgDeleteLotteryItem) (*MsgDeleteLotteryItemResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateLotteryItem(ctx context.Context, req *MsgCreateLotteryItem) (*MsgCreateLotteryItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLotteryItem not implemented")
}
func (*UnimplementedMsgServer) UpdateLotteryItem(ctx context.Context, req *MsgUpdateLotteryItem) (*MsgUpdateLotteryItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLotteryItem not implemented")
}
func (*UnimplementedMsgServer) DeleteLotteryItem(ctx context.Context, req *MsgDeleteLotteryItem) (*MsgDeleteLotteryItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLotteryItem not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateLotteryItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateLotteryItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateLotteryItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lottery.lottery.Msg/CreateLotteryItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateLotteryItem(ctx, req.(*MsgCreateLotteryItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateLotteryItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateLotteryItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateLotteryItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lottery.lottery.Msg/UpdateLotteryItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateLotteryItem(ctx, req.(*MsgUpdateLotteryItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteLotteryItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteLotteryItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteLotteryItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lottery.lottery.Msg/DeleteLotteryItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteLotteryItem(ctx, req.(*MsgDeleteLotteryItem))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lottery.lottery.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLotteryItem",
			Handler:    _Msg_CreateLotteryItem_Handler,
		},
		{
			MethodName: "UpdateLotteryItem",
			Handler:    _Msg_UpdateLotteryItem_Handler,
		},
		{
			MethodName: "DeleteLotteryItem",
			Handler:    _Msg_DeleteLotteryItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lottery/tx.proto",
}

func (m *MsgCreateLotteryItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateLotteryItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateLotteryItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Participants) > 0 {
		for iNdEx := len(m.Participants) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Participants[iNdEx])
			copy(dAtA[i:], m.Participants[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Participants[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Amount != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateLotteryItemResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateLotteryItemResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateLotteryItemResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateLotteryItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateLotteryItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateLotteryItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Participants) > 0 {
		for iNdEx := len(m.Participants) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Participants[iNdEx])
			copy(dAtA[i:], m.Participants[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Participants[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Amount != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateLotteryItemResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateLotteryItemResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateLotteryItemResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteLotteryItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteLotteryItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteLotteryItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteLotteryItemResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteLotteryItemResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteLotteryItemResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateLotteryItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovTx(uint64(m.Amount))
	}
	if len(m.Participants) > 0 {
		for _, s := range m.Participants {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCreateLotteryItemResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateLotteryItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovTx(uint64(m.Amount))
	}
	if len(m.Participants) > 0 {
		for _, s := range m.Participants {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgUpdateLotteryItemResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteLotteryItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeleteLotteryItemResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateLotteryItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateLotteryItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateLotteryItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participants", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participants = append(m.Participants, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateLotteryItemResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateLotteryItemResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateLotteryItemResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateLotteryItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateLotteryItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateLotteryItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participants", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participants = append(m.Participants, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateLotteryItemResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateLotteryItemResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateLotteryItemResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDeleteLotteryItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeleteLotteryItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteLotteryItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDeleteLotteryItemResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeleteLotteryItemResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteLotteryItemResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
