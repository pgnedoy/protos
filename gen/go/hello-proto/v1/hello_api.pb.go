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

func init() {
	proto.RegisterType((*HelloRequest)(nil), "boilerplates.hellogrpc.v1.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "boilerplates.hellogrpc.v1.HelloResponse")
}

func init() { proto.RegisterFile("hello-proto/v1/hello_api.proto", fileDescriptor_150f973d63063ff6) }

var fileDescriptor_150f973d63063ff6 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x95, 0x0a, 0xaa, 0xd6, 0xa1, 0x03, 0x9e, 0x4a, 0x84, 0xa0, 0xf2, 0x42, 0x84, 0x84,
	0xad, 0x94, 0x09, 0x36, 0x98, 0x60, 0x43, 0x11, 0x13, 0x4b, 0xe5, 0xa0, 0xab, 0xb1, 0xe4, 0xd8,
	0x26, 0x76, 0x23, 0xb1, 0xf2, 0x0a, 0x3c, 0x1a, 0xaf, 0xc0, 0x83, 0xa0, 0x5c, 0xd2, 0x8a, 0x05,
	0x75, 0xfb, 0x7f, 0xdf, 0x7f, 0xfe, 0xbf, 0x23, 0x67, 0x6f, 0x60, 0x8c, 0xbb, 0xf2, 0x8d, 0x8b,
	0x4e, 0xb4, 0x85, 0x40, 0xbb, 0x92, 0x5e, 0x73, 0x7c, 0xa2, 0x27, 0x95, 0xd3, 0x06, 0x1a, 0x6f,
	0x64, 0x84, 0xc0, 0x71, 0xaa, 0x1a, 0xff, 0xca, 0xdb, 0x22, 0x3b, 0x55, 0xce, 0x29, 0x03, 0x42,
	0x7a, 0x2d, 0xa4, 0xb5, 0x2e, 0xca, 0xa8, 0x9d, 0x0d, 0xfd, 0x62, 0x76, 0x3e, 0x4c, 0xd1, 0x55,
	0x9b, 0xb5, 0x88, 0xba, 0x86, 0x10, 0x65, 0xed, 0xfb, 0x00, 0x63, 0xe4, 0xe8, 0xa1, 0xfb, 0xae,
	0x84, 0xf7, 0x0d, 0x84, 0x48, 0x29, 0x39, 0xb0, 0xb2, 0x86, 0x79, 0xb2, 0x48, 0xf2, 0x69, 0x89,
	0x9a, 0xad, 0xc9, 0x6c, 0xc8, 0x04, 0xef, 0x6c, 0x00, 0x9a, 0x91, 0x89, 0x6a, 0x00, 0xa2, 0xb6,
	0x6a, 0x08, 0xee, 0x3c, 0xbd, 0x21, 0x04, 0xf5, 0xaa, 0x6b, 0x9a, 0x8f, 0x16, 0x49, 0x9e, 0x2e,
	0x33, 0xde, 0x63, 0xf0, 0x2d, 0x06, 0x7f, 0xde, 0x62, 0x94, 0x53, 0x4c, 0x77, 0x7e, 0xf9, 0x41,
	0x26, 0xd8, 0x73, 0xf7, 0xf4, 0x48, 0x6b, 0x72, 0x88, 0x9a, 0x5e, 0xf0, 0x7f, 0x6f, 0xe7, 0x7f,
	0xc9, 0xb3, 0x7c, 0x7f, 0xb0, 0xc7, 0x67, 0xc7, 0x9f, 0xdf, 0x3f, 0x5f, 0xa3, 0x94, 0x8d, 0x05,
	0x76, 0xdf, 0x26, 0x97, 0xf7, 0xb3, 0x97, 0x74, 0xb7, 0xd0, 0x16, 0xd5, 0x18, 0x41, 0xaf, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x03, 0xde, 0xe4, 0x98, 0x01, 0x00, 0x00,
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
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloAPIClient struct {
	cc *grpc.ClientConn
}

func NewHelloAPIClient(cc *grpc.ClientConn) HelloAPIClient {
	return &helloAPIClient{cc}
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
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterHelloAPIServer(s *grpc.Server, srv HelloAPIServer) {
	s.RegisterService(&_HelloAPI_serviceDesc, srv)
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
			MethodName: "Hello",
			Handler:    _HelloAPI_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello-proto/v1/hello_api.proto",
}