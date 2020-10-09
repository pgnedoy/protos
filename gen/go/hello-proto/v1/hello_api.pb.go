// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello-proto/v1/hello_api.proto

package helloprotov1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_150f973d63063ff6, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Greeting             string               `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	GreetTime            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=greet_time,json=greetTime,proto3" json:"greet_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_150f973d63063ff6, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func (m *HelloResponse) GetGreetTime() *timestamp.Timestamp {
	if m != nil {
		return m.GreetTime
	}
	return nil
}

type CreateUserRequest struct {
	AuthId               string   `protobuf:"bytes,100,opt,name=auth_id,json=authId,proto3" json:"auth_id,omitempty"`
	AuthType             int32    `protobuf:"varint,2,opt,name=auth_type,json=authType,proto3" json:"auth_type,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Country              string   `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	Interest             string   `protobuf:"bytes,6,opt,name=interest,proto3" json:"interest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_150f973d63063ff6, []int{2}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetAuthId() string {
	if m != nil {
		return m.AuthId
	}
	return ""
}

func (m *CreateUserRequest) GetAuthType() int32 {
	if m != nil {
		return m.AuthType
	}
	return 0
}

func (m *CreateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserRequest) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *CreateUserRequest) GetInterest() string {
	if m != nil {
		return m.Interest
	}
	return ""
}

type CreateUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_150f973d63063ff6, []int{3}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloproto.v1.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "helloproto.v1.HelloResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "helloproto.v1.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "helloproto.v1.CreateUserResponse")
}

func init() { proto.RegisterFile("hello-proto/v1/hello_api.proto", fileDescriptor_150f973d63063ff6) }

var fileDescriptor_150f973d63063ff6 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x6e, 0xda, 0x40,
	0x10, 0x95, 0x29, 0x18, 0x18, 0xa0, 0x12, 0xdb, 0x43, 0x2d, 0x83, 0x5a, 0xea, 0x4b, 0x51, 0xa5,
	0xda, 0x82, 0x9e, 0x1a, 0x29, 0x87, 0x24, 0x97, 0x70, 0x4b, 0x2c, 0xb8, 0xe4, 0x82, 0x96, 0x30,
	0x18, 0x4b, 0xb0, 0xeb, 0x78, 0xd7, 0x48, 0x5c, 0xf3, 0x0b, 0xfc, 0x4a, 0xfe, 0x24, 0xbf, 0x90,
	0x0f, 0x89, 0x3c, 0xb6, 0x81, 0x90, 0x28, 0xb7, 0x79, 0x33, 0x6f, 0x66, 0xde, 0x9b, 0x81, 0x1f,
	0x4b, 0x5c, 0xad, 0xe4, 0xdf, 0x28, 0x96, 0x5a, 0x7a, 0x9b, 0x81, 0x47, 0x70, 0xca, 0xa3, 0xd0,
	0xa5, 0x14, 0x6b, 0x51, 0x82, 0x62, 0x77, 0x33, 0xb0, 0xbb, 0x81, 0x94, 0xc1, 0x0a, 0x3d, 0x1e,
	0x85, 0x1e, 0x17, 0x42, 0x6a, 0xae, 0x43, 0x29, 0x54, 0x46, 0xb6, 0x7f, 0xe6, 0x55, 0x42, 0xb3,
	0x64, 0xe1, 0xe9, 0x70, 0x8d, 0x4a, 0xf3, 0x75, 0x94, 0x13, 0xec, 0x93, 0x6d, 0x7a, 0x1b, 0x61,
	0xde, 0xec, 0x38, 0xd0, 0xbc, 0x4e, 0xab, 0x3e, 0x3e, 0x24, 0xa8, 0x34, 0x63, 0x50, 0x16, 0x7c,
	0x8d, 0x96, 0xd1, 0x33, 0xfa, 0x75, 0x9f, 0x62, 0x67, 0x01, 0xad, 0x9c, 0xa3, 0x22, 0x29, 0x14,
	0x32, 0x1b, 0x6a, 0x41, 0x8c, 0xa8, 0x43, 0x11, 0xe4, 0xc4, 0x3d, 0x66, 0xff, 0x01, 0x28, 0x9e,
	0xa6, 0x2a, 0xac, 0x52, 0xcf, 0xe8, 0x37, 0x86, 0xb6, 0x9b, 0x49, 0x74, 0x0b, 0x89, 0xee, 0xb8,
	0x90, 0xe8, 0xd7, 0x89, 0x9d, 0x62, 0x67, 0x67, 0x40, 0xfb, 0x2a, 0x46, 0xae, 0x71, 0xa2, 0x30,
	0x2e, 0x14, 0x7d, 0x87, 0x2a, 0x4f, 0xf4, 0x72, 0x1a, 0xce, 0xad, 0x39, 0xed, 0x32, 0x53, 0x38,
	0x9a, 0xb3, 0x0e, 0xd4, 0xa9, 0x90, 0xda, 0xa1, 0x45, 0x15, 0xbf, 0x96, 0x26, 0xc6, 0xdb, 0x08,
	0xf7, 0x3e, 0xbe, 0x1c, 0x7c, 0x30, 0x0b, 0xaa, 0xf7, 0x32, 0x11, 0x3a, 0xde, 0x5a, 0x15, 0x4a,
	0x17, 0x30, 0x35, 0x14, 0x0a, 0x8d, 0x31, 0x2a, 0x6d, 0x99, 0x99, 0xa1, 0x02, 0x3b, 0xe7, 0xc0,
	0x8e, 0x45, 0xe5, 0x27, 0xf8, 0x0d, 0xe5, 0x44, 0x61, 0x4c, 0xf6, 0x1b, 0xc3, 0x6f, 0xee, 0x9b,
	0x87, 0xb9, 0x44, 0x25, 0xc2, 0xf0, 0xc9, 0x80, 0x1a, 0x5d, 0xef, 0xe2, 0x66, 0xc4, 0x6e, 0x01,
	0x0e, 0xb3, 0x58, 0xef, 0xa4, 0xeb, 0x9d, 0x77, 0xfb, 0xd7, 0x27, 0x8c, 0x5c, 0xc8, 0x04, 0x2a,
	0x34, 0x9e, 0x75, 0x4e, 0xb8, 0xc7, 0x6f, 0xb5, 0xbb, 0x1f, 0x17, 0xb3, 0x19, 0x4e, 0xfb, 0xf1,
	0xf9, 0x65, 0x57, 0x6a, 0x38, 0xa6, 0x47, 0xcf, 0x38, 0x33, 0xfe, 0x5c, 0x7e, 0xbd, 0x6b, 0x1e,
	0x3a, 0x36, 0x83, 0x99, 0x49, 0xc1, 0xbf, 0xd7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x67, 0x2a, 0xea,
	0xdd, 0xba, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloAPIClient is the client API for HelloAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloAPIClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloAPIClient struct {
	cc *grpc.ClientConn
}

func NewHelloAPIClient(cc *grpc.ClientConn) HelloAPIClient {
	return &helloAPIClient{cc}
}

func (c *helloAPIClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/helloproto.v1.HelloAPI/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloAPIClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/helloproto.v1.HelloAPI/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloAPIServer is the server API for HelloAPI service.
type HelloAPIServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterHelloAPIServer(s *grpc.Server, srv HelloAPIServer) {
	s.RegisterService(&_HelloAPI_serviceDesc, srv)
}

func _HelloAPI_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloAPIServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloproto.v1.HelloAPI/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloAPIServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloAPI_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloAPIServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloproto.v1.HelloAPI/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloAPIServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloproto.v1.HelloAPI",
	HandlerType: (*HelloAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _HelloAPI_CreateUser_Handler,
		},
		{
			MethodName: "Hello",
			Handler:    _HelloAPI_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello-proto/v1/hello_api.proto",
}
