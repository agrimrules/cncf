// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/message.proto

package message

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Message struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f3a5e1293a7bcd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CreateMessageReq struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMessageReq) Reset()         { *m = CreateMessageReq{} }
func (m *CreateMessageReq) String() string { return proto.CompactTextString(m) }
func (*CreateMessageReq) ProtoMessage()    {}
func (*CreateMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f3a5e1293a7bcd, []int{1}
}

func (m *CreateMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessageReq.Unmarshal(m, b)
}
func (m *CreateMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessageReq.Marshal(b, m, deterministic)
}
func (m *CreateMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessageReq.Merge(m, src)
}
func (m *CreateMessageReq) XXX_Size() int {
	return xxx_messageInfo_CreateMessageReq.Size(m)
}
func (m *CreateMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessageReq proto.InternalMessageInfo

func (m *CreateMessageReq) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type CreateMessageRes struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMessageRes) Reset()         { *m = CreateMessageRes{} }
func (m *CreateMessageRes) String() string { return proto.CompactTextString(m) }
func (*CreateMessageRes) ProtoMessage()    {}
func (*CreateMessageRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f3a5e1293a7bcd, []int{2}
}

func (m *CreateMessageRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessageRes.Unmarshal(m, b)
}
func (m *CreateMessageRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessageRes.Marshal(b, m, deterministic)
}
func (m *CreateMessageRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessageRes.Merge(m, src)
}
func (m *CreateMessageRes) XXX_Size() int {
	return xxx_messageInfo_CreateMessageRes.Size(m)
}
func (m *CreateMessageRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessageRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessageRes proto.InternalMessageInfo

func (m *CreateMessageRes) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type GetMessageReq struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMessageReq) Reset()         { *m = GetMessageReq{} }
func (m *GetMessageReq) String() string { return proto.CompactTextString(m) }
func (*GetMessageReq) ProtoMessage()    {}
func (*GetMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f3a5e1293a7bcd, []int{3}
}

func (m *GetMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessageReq.Unmarshal(m, b)
}
func (m *GetMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessageReq.Marshal(b, m, deterministic)
}
func (m *GetMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessageReq.Merge(m, src)
}
func (m *GetMessageReq) XXX_Size() int {
	return xxx_messageInfo_GetMessageReq.Size(m)
}
func (m *GetMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessageReq proto.InternalMessageInfo

func (m *GetMessageReq) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type GetMessageRes struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMessageRes) Reset()         { *m = GetMessageRes{} }
func (m *GetMessageRes) String() string { return proto.CompactTextString(m) }
func (*GetMessageRes) ProtoMessage()    {}
func (*GetMessageRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f3a5e1293a7bcd, []int{4}
}

func (m *GetMessageRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessageRes.Unmarshal(m, b)
}
func (m *GetMessageRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessageRes.Marshal(b, m, deterministic)
}
func (m *GetMessageRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessageRes.Merge(m, src)
}
func (m *GetMessageRes) XXX_Size() int {
	return xxx_messageInfo_GetMessageRes.Size(m)
}
func (m *GetMessageRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessageRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessageRes proto.InternalMessageInfo

func (m *GetMessageRes) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type ListMessageRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMessageRequest) Reset()         { *m = ListMessageRequest{} }
func (m *ListMessageRequest) String() string { return proto.CompactTextString(m) }
func (*ListMessageRequest) ProtoMessage()    {}
func (*ListMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f3a5e1293a7bcd, []int{5}
}

func (m *ListMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMessageRequest.Unmarshal(m, b)
}
func (m *ListMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMessageRequest.Marshal(b, m, deterministic)
}
func (m *ListMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMessageRequest.Merge(m, src)
}
func (m *ListMessageRequest) XXX_Size() int {
	return xxx_messageInfo_ListMessageRequest.Size(m)
}
func (m *ListMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMessageRequest proto.InternalMessageInfo

type ListMessageResponse struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMessageResponse) Reset()         { *m = ListMessageResponse{} }
func (m *ListMessageResponse) String() string { return proto.CompactTextString(m) }
func (*ListMessageResponse) ProtoMessage()    {}
func (*ListMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f3a5e1293a7bcd, []int{6}
}

func (m *ListMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMessageResponse.Unmarshal(m, b)
}
func (m *ListMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMessageResponse.Marshal(b, m, deterministic)
}
func (m *ListMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMessageResponse.Merge(m, src)
}
func (m *ListMessageResponse) XXX_Size() int {
	return xxx_messageInfo_ListMessageResponse.Size(m)
}
func (m *ListMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMessageResponse proto.InternalMessageInfo

func (m *ListMessageResponse) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "message.Message")
	proto.RegisterType((*CreateMessageReq)(nil), "message.CreateMessageReq")
	proto.RegisterType((*CreateMessageRes)(nil), "message.CreateMessageRes")
	proto.RegisterType((*GetMessageReq)(nil), "message.GetMessageReq")
	proto.RegisterType((*GetMessageRes)(nil), "message.GetMessageRes")
	proto.RegisterType((*ListMessageRequest)(nil), "message.ListMessageRequest")
	proto.RegisterType((*ListMessageResponse)(nil), "message.ListMessageResponse")
}

func init() { proto.RegisterFile("proto/message.proto", fileDescriptor_33f3a5e1293a7bcd) }

var fileDescriptor_33f3a5e1293a7bcd = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x03, 0xf3, 0x84, 0xd8, 0xa1, 0x5c,
	0x25, 0x73, 0x2e, 0x76, 0x5f, 0x08, 0x53, 0x48, 0x88, 0x8b, 0xa5, 0xb4, 0x38, 0xb5, 0x48, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x92, 0xe0, 0x82, 0xa9, 0x94, 0x60, 0x02, 0x0b,
	0xc3, 0x35, 0xda, 0x71, 0x09, 0x38, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0x42, 0xb5, 0x07, 0xa5, 0x16,
	0x0a, 0x69, 0x21, 0x54, 0x83, 0x0c, 0xe1, 0x36, 0x12, 0xd0, 0x83, 0x59, 0x0b, 0x53, 0x85, 0x47,
	0x7f, 0x31, 0x49, 0xfa, 0x95, 0xb9, 0x78, 0xdd, 0x53, 0x4b, 0x90, 0x2c, 0xc7, 0xe2, 0x7c, 0x25,
	0x6b, 0x54, 0x45, 0xa4, 0xd9, 0x20, 0xc2, 0x25, 0xe4, 0x93, 0x59, 0x8c, 0x64, 0x45, 0x69, 0x6a,
	0x71, 0x89, 0x92, 0x23, 0x97, 0x30, 0x8a, 0x68, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x29, 0x06,
	0x1b, 0x3d, 0x66, 0xe4, 0xe2, 0x83, 0x0a, 0x06, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xb9,
	0x72, 0xf1, 0xa2, 0x84, 0x86, 0x90, 0x24, 0x5c, 0x3b, 0x7a, 0x28, 0x4b, 0xe1, 0x94, 0x2a, 0x16,
	0xb2, 0xe3, 0xe2, 0x42, 0xf8, 0x57, 0x48, 0x0c, 0xae, 0x10, 0x25, 0xa4, 0xa4, 0xb0, 0x8b, 0x17,
	0x1b, 0x30, 0x0a, 0x79, 0x71, 0x71, 0x23, 0x79, 0x4e, 0x48, 0x1a, 0xae, 0x10, 0x33, 0x20, 0xa4,
	0x64, 0xb0, 0x4b, 0x42, 0xc2, 0xc3, 0x80, 0x31, 0x89, 0x0d, 0x9c, 0xd2, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x5c, 0x81, 0xe6, 0x61, 0x80, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageServiceClient interface {
	CreateMessage(ctx context.Context, in *CreateMessageReq, opts ...grpc.CallOption) (*CreateMessageRes, error)
	GetMessage(ctx context.Context, in *GetMessageReq, opts ...grpc.CallOption) (MessageService_GetMessageClient, error)
	ListMessage(ctx context.Context, in *ListMessageRequest, opts ...grpc.CallOption) (MessageService_ListMessageClient, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) CreateMessage(ctx context.Context, in *CreateMessageReq, opts ...grpc.CallOption) (*CreateMessageRes, error) {
	out := new(CreateMessageRes)
	err := c.cc.Invoke(ctx, "/message.MessageService/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetMessage(ctx context.Context, in *GetMessageReq, opts ...grpc.CallOption) (MessageService_GetMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MessageService_serviceDesc.Streams[0], "/message.MessageService/GetMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceGetMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageService_GetMessageClient interface {
	Recv() (*GetMessageRes, error)
	grpc.ClientStream
}

type messageServiceGetMessageClient struct {
	grpc.ClientStream
}

func (x *messageServiceGetMessageClient) Recv() (*GetMessageRes, error) {
	m := new(GetMessageRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageServiceClient) ListMessage(ctx context.Context, in *ListMessageRequest, opts ...grpc.CallOption) (MessageService_ListMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MessageService_serviceDesc.Streams[1], "/message.MessageService/ListMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceListMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageService_ListMessageClient interface {
	Recv() (*ListMessageResponse, error)
	grpc.ClientStream
}

type messageServiceListMessageClient struct {
	grpc.ClientStream
}

func (x *messageServiceListMessageClient) Recv() (*ListMessageResponse, error) {
	m := new(ListMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageServiceServer is the server API for MessageService service.
type MessageServiceServer interface {
	CreateMessage(context.Context, *CreateMessageReq) (*CreateMessageRes, error)
	GetMessage(*GetMessageReq, MessageService_GetMessageServer) error
	ListMessage(*ListMessageRequest, MessageService_ListMessageServer) error
}

// UnimplementedMessageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (*UnimplementedMessageServiceServer) CreateMessage(ctx context.Context, req *CreateMessageReq) (*CreateMessageRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (*UnimplementedMessageServiceServer) GetMessage(req *GetMessageReq, srv MessageService_GetMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (*UnimplementedMessageServiceServer) ListMessage(req *ListMessageRequest, srv MessageService_ListMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method ListMessage not implemented")
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).CreateMessage(ctx, req.(*CreateMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetMessageReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageServiceServer).GetMessage(m, &messageServiceGetMessageServer{stream})
}

type MessageService_GetMessageServer interface {
	Send(*GetMessageRes) error
	grpc.ServerStream
}

type messageServiceGetMessageServer struct {
	grpc.ServerStream
}

func (x *messageServiceGetMessageServer) Send(m *GetMessageRes) error {
	return x.ServerStream.SendMsg(m)
}

func _MessageService_ListMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListMessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageServiceServer).ListMessage(m, &messageServiceListMessageServer{stream})
}

type MessageService_ListMessageServer interface {
	Send(*ListMessageResponse) error
	grpc.ServerStream
}

type messageServiceListMessageServer struct {
	grpc.ServerStream
}

func (x *messageServiceListMessageServer) Send(m *ListMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _MessageService_CreateMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMessage",
			Handler:       _MessageService_GetMessage_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListMessage",
			Handler:       _MessageService_ListMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/message.proto",
}
