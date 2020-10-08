// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello-proto/v1/hello_api.proto

package hellogrpcv1

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
	AuthId               string   `protobuf:"bytes,1,opt,name=auth_id,json=authId,proto3" json:"auth_id,omitempty"`
	AuthType             AuthType `protobuf:"varint,2,opt,name=auth_type,json=authType,proto3,enum=boilerplates.hellogrpc.v1.AuthType" json:"auth_type,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Country              string   `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
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

func (m *CreateUserRequest) GetAuthType() AuthType {
	if m != nil {
		return m.AuthType
	}
	return AuthType_AUTH_TYPE_INVALID
}

func (m *CreateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserRequest) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *CreateUserRequest) GetCountry() string {
	if m != nil {
		return m.Country
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
	proto.RegisterType((*HelloRequest)(nil), "boilerplates.hellogrpc.v1.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "boilerplates.hellogrpc.v1.HelloResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "boilerplates.hellogrpc.v1.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "boilerplates.hellogrpc.v1.CreateUserResponse")
}

func init() { proto.RegisterFile("hello-proto/v1/hello_api.proto", fileDescriptor_150f973d63063ff6) }

var fileDescriptor_150f973d63063ff6 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0x36, 0x49, 0x93, 0x31, 0x45, 0x74, 0x2f, 0x18, 0x0b, 0xd1, 0xc8, 0x1c, 0xb0,
	0x10, 0x5d, 0x2b, 0xee, 0x09, 0x4e, 0x14, 0x2e, 0xe4, 0x86, 0xac, 0x72, 0xe1, 0x12, 0x6d, 0xda,
	0xa9, 0xb3, 0x92, 0xed, 0x5d, 0x76, 0xd7, 0x91, 0x72, 0xe5, 0x15, 0x78, 0x0f, 0x5e, 0x86, 0x57,
	0xe8, 0x83, 0x20, 0x4f, 0xd6, 0x69, 0x05, 0x22, 0xf4, 0x36, 0xbf, 0xe7, 0x9f, 0x99, 0x6f, 0xc6,
	0x0b, 0x2f, 0x56, 0x58, 0x55, 0xea, 0x4c, 0x1b, 0xe5, 0x54, 0xb6, 0x9e, 0x65, 0x24, 0x17, 0x42,
	0x4b, 0x4e, 0x9f, 0xd8, 0xb3, 0xa5, 0x92, 0x15, 0x1a, 0x5d, 0x09, 0x87, 0x96, 0x53, 0xb6, 0x34,
	0xfa, 0x8a, 0xaf, 0x67, 0xf1, 0xf3, 0x52, 0xa9, 0xb2, 0xc2, 0x4c, 0x68, 0x99, 0x89, 0xa6, 0x51,
	0x4e, 0x38, 0xa9, 0x1a, 0xbb, 0x2d, 0x8c, 0x4f, 0x7d, 0x96, 0xd4, 0xb2, 0xbd, 0xc9, 0x9c, 0xac,
	0xd1, 0x3a, 0x51, 0x6b, 0x6f, 0x88, 0xff, 0x98, 0xec, 0x36, 0x1a, 0x7d, 0x71, 0x92, 0xc0, 0xa3,
	0x4f, 0x5d, 0xb6, 0xc0, 0x6f, 0x2d, 0x5a, 0xc7, 0x18, 0x0c, 0x1a, 0x51, 0x63, 0x14, 0x4c, 0x83,
	0x74, 0x52, 0x50, 0x9c, 0xdc, 0xc0, 0xb1, 0xf7, 0x58, 0xad, 0x1a, 0x8b, 0x2c, 0x86, 0x71, 0x69,
	0x10, 0x9d, 0x6c, 0x4a, 0x6f, 0xdc, 0x69, 0xf6, 0x16, 0x80, 0xe2, 0x45, 0x47, 0x11, 0x1d, 0x4c,
	0x83, 0x34, 0xcc, 0x63, 0xbe, 0x45, 0xe4, 0x3d, 0x22, 0xbf, 0xec, 0x11, 0x8b, 0x09, 0xb9, 0x3b,
	0x9d, 0xfc, 0x0c, 0xe0, 0xe4, 0xa3, 0x41, 0xe1, 0xf0, 0x8b, 0x45, 0xd3, 0x13, 0x3d, 0x85, 0x23,
	0xd1, 0xba, 0xd5, 0x42, 0x5e, 0xfb, 0x59, 0xa3, 0x4e, 0xce, 0xaf, 0xd9, 0x7b, 0x98, 0x50, 0xa2,
	0x5b, 0x87, 0x06, 0x3d, 0xce, 0x5f, 0xf2, 0x7f, 0x1e, 0x91, 0x5f, 0xb4, 0x6e, 0x75, 0xb9, 0xd1,
	0x58, 0x8c, 0x85, 0x8f, 0x76, 0xcb, 0x1e, 0xde, 0x2d, 0xcb, 0x9e, 0xc0, 0xa1, 0x28, 0x31, 0x1a,
	0x4c, 0x83, 0x74, 0x58, 0x74, 0x21, 0x8b, 0xe0, 0xe8, 0x4a, 0xb5, 0x8d, 0x33, 0x9b, 0x68, 0x48,
	0xc6, 0x5e, 0x26, 0x73, 0x60, 0xf7, 0x79, 0xfd, 0x75, 0xce, 0x61, 0xd0, 0x5a, 0x34, 0x44, 0x1b,
	0xe6, 0xa7, 0x7b, 0x90, 0xa8, 0x8c, 0xcc, 0xf9, 0x6d, 0x00, 0x63, 0x3a, 0xf2, 0xc5, 0xe7, 0x39,
	0x93, 0x00, 0x77, 0x7d, 0xd9, 0x9b, 0x3d, 0x1d, 0xfe, 0x3a, 0x57, 0x7c, 0xf6, 0x40, 0xb7, 0x87,
	0xad, 0x61, 0x48, 0x63, 0xd9, 0xab, 0x3d, 0x75, 0xf7, 0x5f, 0x48, 0x9c, 0xfe, 0xdf, 0xb8, 0xed,
	0x9d, 0x9c, 0x7c, 0xff, 0x75, 0xfb, 0xe3, 0x20, 0x4c, 0x46, 0x19, 0xfd, 0xe3, 0x77, 0xc1, 0xeb,
	0x0f, 0xc7, 0x5f, 0xc3, 0x5d, 0xc1, 0x7a, 0xb6, 0x1c, 0xd1, 0x83, 0x38, 0xff, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0x43, 0xfb, 0xe2, 0x53, 0x1c, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/boilerplates.hellogrpc.v1.HelloAPI/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloAPIClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/boilerplates.hellogrpc.v1.HelloAPI/Hello", in, out, opts...)
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
		FullMethod: "/boilerplates.hellogrpc.v1.HelloAPI/CreateUser",
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
		FullMethod: "/boilerplates.hellogrpc.v1.HelloAPI/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloAPIServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "boilerplates.hellogrpc.v1.HelloAPI",
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
