// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planet/blog/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgSendIbcPost struct {
	Creator          string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Port             string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	ChannelID        string `protobuf:"bytes,3,opt,name=channelID,proto3" json:"channelID,omitempty"`
	TimeoutTimestamp uint64 `protobuf:"varint,4,opt,name=timeoutTimestamp,proto3" json:"timeoutTimestamp,omitempty"`
	Title            string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Content          string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *MsgSendIbcPost) Reset()         { *m = MsgSendIbcPost{} }
func (m *MsgSendIbcPost) String() string { return proto.CompactTextString(m) }
func (*MsgSendIbcPost) ProtoMessage()    {}
func (*MsgSendIbcPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_32d5736426604c20, []int{0}
}
func (m *MsgSendIbcPost) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendIbcPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendIbcPost.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendIbcPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendIbcPost.Merge(m, src)
}
func (m *MsgSendIbcPost) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendIbcPost) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendIbcPost.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendIbcPost proto.InternalMessageInfo

func (m *MsgSendIbcPost) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSendIbcPost) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *MsgSendIbcPost) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *MsgSendIbcPost) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

func (m *MsgSendIbcPost) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MsgSendIbcPost) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type MsgSendIbcPostResponse struct {
}

func (m *MsgSendIbcPostResponse) Reset()         { *m = MsgSendIbcPostResponse{} }
func (m *MsgSendIbcPostResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendIbcPostResponse) ProtoMessage()    {}
func (*MsgSendIbcPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32d5736426604c20, []int{1}
}
func (m *MsgSendIbcPostResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendIbcPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendIbcPostResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendIbcPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendIbcPostResponse.Merge(m, src)
}
func (m *MsgSendIbcPostResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendIbcPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendIbcPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendIbcPostResponse proto.InternalMessageInfo

type MsgSendUpdatePost struct {
	PostID           string `protobuf:"bytes,5,opt,name=postID,proto3" json:"postID,omitempty"`
	Title            string `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	Content          string `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	Creator          string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Port             string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	ChannelID        string `protobuf:"bytes,3,opt,name=channelID,proto3" json:"channelID,omitempty"`
	TimeoutTimestamp uint64 `protobuf:"varint,4,opt,name=timeoutTimestamp,proto3" json:"timeoutTimestamp,omitempty"`
}

func (m *MsgSendUpdatePost) Reset()         { *m = MsgSendUpdatePost{} }
func (m *MsgSendUpdatePost) String() string { return proto.CompactTextString(m) }
func (*MsgSendUpdatePost) ProtoMessage()    {}
func (*MsgSendUpdatePost) Descriptor() ([]byte, []int) {
	return fileDescriptor_32d5736426604c20, []int{2}
}
func (m *MsgSendUpdatePost) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendUpdatePost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendUpdatePost.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendUpdatePost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendUpdatePost.Merge(m, src)
}
func (m *MsgSendUpdatePost) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendUpdatePost) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendUpdatePost.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendUpdatePost proto.InternalMessageInfo

func (m *MsgSendUpdatePost) GetPostID() string {
	if m != nil {
		return m.PostID
	}
	return ""
}

func (m *MsgSendUpdatePost) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MsgSendUpdatePost) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *MsgSendUpdatePost) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSendUpdatePost) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *MsgSendUpdatePost) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *MsgSendUpdatePost) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

type MsgSendUpdatePostResponse struct {
}

func (m *MsgSendUpdatePostResponse) Reset()         { *m = MsgSendUpdatePostResponse{} }
func (m *MsgSendUpdatePostResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendUpdatePostResponse) ProtoMessage()    {}
func (*MsgSendUpdatePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32d5736426604c20, []int{3}
}
func (m *MsgSendUpdatePostResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendUpdatePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendUpdatePostResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendUpdatePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendUpdatePostResponse.Merge(m, src)
}
func (m *MsgSendUpdatePostResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendUpdatePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendUpdatePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendUpdatePostResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSendIbcPost)(nil), "planet.blog.MsgSendIbcPost")
	proto.RegisterType((*MsgSendIbcPostResponse)(nil), "planet.blog.MsgSendIbcPostResponse")
	proto.RegisterType((*MsgSendUpdatePost)(nil), "planet.blog.MsgSendUpdatePost")
	proto.RegisterType((*MsgSendUpdatePostResponse)(nil), "planet.blog.MsgSendUpdatePostResponse")
}

func init() { proto.RegisterFile("planet/blog/tx.proto", fileDescriptor_32d5736426604c20) }

var fileDescriptor_32d5736426604c20 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x52, 0xcd, 0x4a, 0x42, 0x41,
	0x18, 0x75, 0x52, 0xaf, 0xf8, 0x09, 0x52, 0x93, 0xc8, 0xa4, 0x31, 0xc8, 0x0d, 0x42, 0x82, 0xae,
	0x50, 0x6f, 0x10, 0x6e, 0x5c, 0x48, 0x61, 0x05, 0xd1, 0xee, 0xaa, 0x83, 0x09, 0xd7, 0x99, 0xe1,
	0xce, 0x17, 0xd8, 0x5b, 0xf4, 0x28, 0x2d, 0x7b, 0x84, 0x96, 0x42, 0x9b, 0x96, 0xa1, 0x2f, 0x12,
	0x8d, 0x73, 0xd1, 0x9b, 0xe5, 0xb6, 0xdd, 0x9c, 0x73, 0x98, 0xef, 0x7c, 0x3f, 0x07, 0x2a, 0x3a,
	0x0a, 0xa5, 0xc0, 0x56, 0x3f, 0x52, 0xa3, 0x16, 0x4e, 0x03, 0x1d, 0x2b, 0x54, 0xb4, 0xb4, 0x64,
	0x83, 0x6f, 0xd6, 0x7f, 0x25, 0x50, 0xee, 0x9a, 0xd1, 0xb5, 0x90, 0xc3, 0x4e, 0x7f, 0x70, 0xa5,
	0x0c, 0x52, 0x06, 0x85, 0x41, 0x2c, 0x42, 0x54, 0x31, 0x23, 0x0d, 0xd2, 0x2c, 0xf6, 0x12, 0x48,
	0x29, 0xe4, 0xb4, 0x8a, 0x91, 0xed, 0x58, 0xda, 0xbe, 0xe9, 0x21, 0x14, 0x07, 0x0f, 0xa1, 0x94,
	0x22, 0xea, 0xb4, 0x59, 0xd6, 0x0a, 0x2b, 0x82, 0x9e, 0xc0, 0x2e, 0x8e, 0x27, 0x42, 0x3d, 0xe2,
	0xcd, 0x78, 0x22, 0x0c, 0x86, 0x13, 0xcd, 0x72, 0x0d, 0xd2, 0xcc, 0xf5, 0x36, 0x78, 0x5a, 0x81,
	0x3c, 0x8e, 0x31, 0x12, 0x2c, 0x6f, 0xab, 0x2c, 0x81, 0xed, 0x46, 0x49, 0x14, 0x12, 0x99, 0xe7,
	0xba, 0x59, 0x42, 0x9f, 0x41, 0x35, 0xdd, 0x79, 0x4f, 0x18, 0xad, 0xa4, 0x11, 0xfe, 0x3b, 0x81,
	0x3d, 0x27, 0xdd, 0xea, 0x61, 0x88, 0xc2, 0xce, 0x55, 0x05, 0x4f, 0x2b, 0x83, 0x9d, 0xb6, 0x33,
	0x70, 0x68, 0xe5, 0xeb, 0xfd, 0xe1, 0x5b, 0x48, 0xf9, 0xfe, 0xd7, 0x7e, 0xfc, 0x3a, 0x1c, 0x6c,
	0x0c, 0x95, 0x8c, 0x7c, 0xf6, 0x42, 0x20, 0xdb, 0x35, 0x23, 0x7a, 0x09, 0xa5, 0xf5, 0x5b, 0xd6,
	0x83, 0xb5, 0x63, 0x07, 0xe9, 0x75, 0xd5, 0x8e, 0xb6, 0x88, 0x49, 0x61, 0x7a, 0x07, 0xe5, 0x1f,
	0x7b, 0xe4, 0xbf, 0x7d, 0x5b, 0xe9, 0xb5, 0xe3, 0xed, 0x7a, 0x52, 0xf9, 0xe2, 0xf4, 0x6d, 0xce,
	0xc9, 0x6c, 0xce, 0xc9, 0xe7, 0x9c, 0x93, 0xe7, 0x05, 0xcf, 0xcc, 0x16, 0x3c, 0xf3, 0xb1, 0xe0,
	0x99, 0xfb, 0x7d, 0x97, 0xdb, 0xa9, 0x4b, 0xee, 0x93, 0x16, 0xa6, 0xef, 0xd9, 0xf4, 0x9e, 0x7f,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x50, 0x09, 0x04, 0x6a, 0xd5, 0x02, 0x00, 0x00,
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
	SendIbcPost(ctx context.Context, in *MsgSendIbcPost, opts ...grpc.CallOption) (*MsgSendIbcPostResponse, error)
	SendUpdatePost(ctx context.Context, in *MsgSendUpdatePost, opts ...grpc.CallOption) (*MsgSendUpdatePostResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SendIbcPost(ctx context.Context, in *MsgSendIbcPost, opts ...grpc.CallOption) (*MsgSendIbcPostResponse, error) {
	out := new(MsgSendIbcPostResponse)
	err := c.cc.Invoke(ctx, "/planet.blog.Msg/SendIbcPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendUpdatePost(ctx context.Context, in *MsgSendUpdatePost, opts ...grpc.CallOption) (*MsgSendUpdatePostResponse, error) {
	out := new(MsgSendUpdatePostResponse)
	err := c.cc.Invoke(ctx, "/planet.blog.Msg/SendUpdatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SendIbcPost(context.Context, *MsgSendIbcPost) (*MsgSendIbcPostResponse, error)
	SendUpdatePost(context.Context, *MsgSendUpdatePost) (*MsgSendUpdatePostResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SendIbcPost(ctx context.Context, req *MsgSendIbcPost) (*MsgSendIbcPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendIbcPost not implemented")
}
func (*UnimplementedMsgServer) SendUpdatePost(ctx context.Context, req *MsgSendUpdatePost) (*MsgSendUpdatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUpdatePost not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SendIbcPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendIbcPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendIbcPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/planet.blog.Msg/SendIbcPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendIbcPost(ctx, req.(*MsgSendIbcPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendUpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendUpdatePost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendUpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/planet.blog.Msg/SendUpdatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendUpdatePost(ctx, req.(*MsgSendUpdatePost))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "planet.blog.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendIbcPost",
			Handler:    _Msg_SendIbcPost_Handler,
		},
		{
			MethodName: "SendUpdatePost",
			Handler:    _Msg_SendUpdatePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "planet/blog/tx.proto",
}

func (m *MsgSendIbcPost) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendIbcPost) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendIbcPost) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0x12
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

func (m *MsgSendIbcPostResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendIbcPostResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendIbcPostResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSendUpdatePost) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendUpdatePost) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendUpdatePost) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.PostID) > 0 {
		i -= len(m.PostID)
		copy(dAtA[i:], m.PostID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PostID)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0x12
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

func (m *MsgSendUpdatePostResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendUpdatePostResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendUpdatePostResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgSendIbcPost) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSendIbcPostResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSendUpdatePost) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	l = len(m.PostID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSendUpdatePostResponse) Size() (n int) {
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
func (m *MsgSendIbcPost) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSendIbcPost: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendIbcPost: illegal tag %d (wire type %d)", fieldNum, wire)
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
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
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
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
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
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
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
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
			m.Content = string(dAtA[iNdEx:postIndex])
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
func (m *MsgSendIbcPostResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSendIbcPostResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendIbcPostResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgSendUpdatePost) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSendUpdatePost: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendUpdatePost: illegal tag %d (wire type %d)", fieldNum, wire)
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
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
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
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
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
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
			m.PostID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
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
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
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
			m.Content = string(dAtA[iNdEx:postIndex])
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
func (m *MsgSendUpdatePostResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSendUpdatePostResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendUpdatePostResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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