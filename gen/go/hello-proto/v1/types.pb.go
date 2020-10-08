// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello-proto/v1/types.proto

package hellogrpcv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type AuthType int32

const (
	AuthType_AUTH_TYPE_INVALID AuthType = 0
	AuthType_AUTH_TYPE_PHONE   AuthType = 1
	AuthType_AUTH_TYPE_SNAP    AuthType = 2
)

var AuthType_name = map[int32]string{
	0: "AUTH_TYPE_INVALID",
	1: "AUTH_TYPE_PHONE",
	2: "AUTH_TYPE_SNAP",
}

var AuthType_value = map[string]int32{
	"AUTH_TYPE_INVALID": 0,
	"AUTH_TYPE_PHONE":   1,
	"AUTH_TYPE_SNAP":    2,
}

func (x AuthType) String() string {
	return proto.EnumName(AuthType_name, int32(x))
}

func (AuthType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b2b3bc71e736bc01, []int{0}
}

type GenderType int32

const (
	GenderType_GENDER_TYPE_INVALID GenderType = 0
	GenderType_GENDER_TYPE_MALE    GenderType = 1
	GenderType_GENDER_TYPE_FEMALE  GenderType = 2
)

var GenderType_name = map[int32]string{
	0: "GENDER_TYPE_INVALID",
	1: "GENDER_TYPE_MALE",
	2: "GENDER_TYPE_FEMALE",
}

var GenderType_value = map[string]int32{
	"GENDER_TYPE_INVALID": 0,
	"GENDER_TYPE_MALE":    1,
	"GENDER_TYPE_FEMALE":  2,
}

func (x GenderType) String() string {
	return proto.EnumName(GenderType_name, int32(x))
}

func (GenderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b2b3bc71e736bc01, []int{1}
}

type DeletionReason int32

const (
	DeletionReason_DELETION_REASON_INVALID           DeletionReason = 0
	DeletionReason_DELETION_REASON_SELF_DELETION     DeletionReason = 1
	DeletionReason_DELETION_REASON_BOT_BEHAVIOUR     DeletionReason = 2
	DeletionReason_DELETION_REASON_NUDITY            DeletionReason = 3
	DeletionReason_DELETION_REASON_SEXUAL_ACTIVITY   DeletionReason = 4
	DeletionReason_DELETION_REASON_ADULT_TOYS        DeletionReason = 5
	DeletionReason_DELETION_REASON_UNDERWEAR         DeletionReason = 6
	DeletionReason_DELETION_REASON_PHYSICAL_VIOLENCE DeletionReason = 7
	DeletionReason_DELETION_REASON_WEAPON_VIOLENCE   DeletionReason = 8
	DeletionReason_DELETION_REASON_WEAPONS           DeletionReason = 9
	DeletionReason_DELETION_REASON_SELF_INJURY       DeletionReason = 10
)

var DeletionReason_name = map[int32]string{
	0:  "DELETION_REASON_INVALID",
	1:  "DELETION_REASON_SELF_DELETION",
	2:  "DELETION_REASON_BOT_BEHAVIOUR",
	3:  "DELETION_REASON_NUDITY",
	4:  "DELETION_REASON_SEXUAL_ACTIVITY",
	5:  "DELETION_REASON_ADULT_TOYS",
	6:  "DELETION_REASON_UNDERWEAR",
	7:  "DELETION_REASON_PHYSICAL_VIOLENCE",
	8:  "DELETION_REASON_WEAPON_VIOLENCE",
	9:  "DELETION_REASON_WEAPONS",
	10: "DELETION_REASON_SELF_INJURY",
}

var DeletionReason_value = map[string]int32{
	"DELETION_REASON_INVALID":           0,
	"DELETION_REASON_SELF_DELETION":     1,
	"DELETION_REASON_BOT_BEHAVIOUR":     2,
	"DELETION_REASON_NUDITY":            3,
	"DELETION_REASON_SEXUAL_ACTIVITY":   4,
	"DELETION_REASON_ADULT_TOYS":        5,
	"DELETION_REASON_UNDERWEAR":         6,
	"DELETION_REASON_PHYSICAL_VIOLENCE": 7,
	"DELETION_REASON_WEAPON_VIOLENCE":   8,
	"DELETION_REASON_WEAPONS":           9,
	"DELETION_REASON_SELF_INJURY":       10,
}

func (x DeletionReason) String() string {
	return proto.EnumName(DeletionReason_name, int32(x))
}

func (DeletionReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b2b3bc71e736bc01, []int{2}
}

// This is a leading comment for a message
type User struct {
	// the uuid of user
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	// unique value for each user
	WinkId   string     `protobuf:"bytes,4,opt,name=wink_id,json=winkId,proto3" json:"wink_id,omitempty"`
	Country  string     `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	Gender   GenderType `protobuf:"varint,6,opt,name=gender,proto3,enum=boilerplates.hellogrpc.v1.GenderType" json:"gender,omitempty"`
	About    string     `protobuf:"bytes,7,opt,name=about,proto3" json:"about,omitempty"`
	AuthType AuthType   `protobuf:"varint,8,opt,name=auth_type,json=authType,proto3,enum=boilerplates.hellogrpc.v1.AuthType" json:"auth_type,omitempty"`
	// birthday of this user
	Birthday       string               `protobuf:"bytes,9,opt,name=birthday,proto3" json:"birthday,omitempty"`
	DeletionReason DeletionReason       `protobuf:"varint,10,opt,name=deletion_reason,json=deletionReason,proto3,enum=boilerplates.hellogrpc.v1.DeletionReason" json:"deletion_reason,omitempty"`
	CreateTime     *timestamp.Timestamp `protobuf:"bytes,11,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime     *timestamp.Timestamp `protobuf:"bytes,12,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	DeleteTime     *timestamp.Timestamp `protobuf:"bytes,13,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// snapchat username
	SnapName             string   `protobuf:"bytes,14,opt,name=snap_name,json=snapName,proto3" json:"snap_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2b3bc71e736bc01, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetWinkId() string {
	if m != nil {
		return m.WinkId
	}
	return ""
}

func (m *User) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *User) GetGender() GenderType {
	if m != nil {
		return m.Gender
	}
	return GenderType_GENDER_TYPE_INVALID
}

func (m *User) GetAbout() string {
	if m != nil {
		return m.About
	}
	return ""
}

func (m *User) GetAuthType() AuthType {
	if m != nil {
		return m.AuthType
	}
	return AuthType_AUTH_TYPE_INVALID
}

func (m *User) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *User) GetDeletionReason() DeletionReason {
	if m != nil {
		return m.DeletionReason
	}
	return DeletionReason_DELETION_REASON_INVALID
}

func (m *User) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *User) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *User) GetDeleteTime() *timestamp.Timestamp {
	if m != nil {
		return m.DeleteTime
	}
	return nil
}

func (m *User) GetSnapName() string {
	if m != nil {
		return m.SnapName
	}
	return ""
}

func init() {
	proto.RegisterEnum("boilerplates.hellogrpc.v1.AuthType", AuthType_name, AuthType_value)
	proto.RegisterEnum("boilerplates.hellogrpc.v1.GenderType", GenderType_name, GenderType_value)
	proto.RegisterEnum("boilerplates.hellogrpc.v1.DeletionReason", DeletionReason_name, DeletionReason_value)
	proto.RegisterType((*User)(nil), "boilerplates.hellogrpc.v1.User")
}

func init() { proto.RegisterFile("hello-proto/v1/types.proto", fileDescriptor_b2b3bc71e736bc01) }

var fileDescriptor_b2b3bc71e736bc01 = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x6f, 0x4f, 0xdb, 0x48,
	0x10, 0xc6, 0xcf, 0x26, 0x7f, 0x27, 0x47, 0xf0, 0x0d, 0x1c, 0x98, 0x20, 0x8e, 0x70, 0x08, 0x29,
	0x45, 0xaa, 0x23, 0xe8, 0xcb, 0xaa, 0x52, 0x0d, 0x59, 0x88, 0x91, 0x6b, 0x47, 0xfe, 0x13, 0x9a,
	0xbe, 0xb1, 0x1c, 0xbc, 0x4d, 0xac, 0x26, 0xb6, 0xe5, 0x38, 0x54, 0xf9, 0x8a, 0xfd, 0x4e, 0x95,
	0x2a, 0xaf, 0x13, 0x02, 0x69, 0x0a, 0xef, 0x76, 0x9e, 0xf9, 0x3d, 0xcf, 0x6e, 0x66, 0x1c, 0xa8,
	0x0d, 0xe9, 0x68, 0x14, 0xbe, 0x8d, 0xe2, 0x30, 0x09, 0x9b, 0x0f, 0xe7, 0xcd, 0x64, 0x16, 0xd1,
	0x89, 0xc4, 0x4a, 0xdc, 0xef, 0x87, 0xfe, 0x88, 0xc6, 0xd1, 0xc8, 0x4d, 0xe8, 0x44, 0x62, 0xe0,
	0x20, 0x8e, 0xee, 0xa5, 0x87, 0xf3, 0xda, 0xd1, 0x20, 0x0c, 0x07, 0x23, 0xda, 0x64, 0x60, 0x7f,
	0xfa, 0xb5, 0x99, 0xf8, 0x63, 0x3a, 0x49, 0xdc, 0x71, 0x94, 0x79, 0xff, 0xff, 0x91, 0x83, 0x9c,
	0x3d, 0xa1, 0x31, 0x56, 0x81, 0xf7, 0x3d, 0x91, 0xab, 0x73, 0x8d, 0xb2, 0xc1, 0xfb, 0x1e, 0x22,
	0xe4, 0x02, 0x77, 0x4c, 0x45, 0x9e, 0x29, 0xec, 0x8c, 0x02, 0x6c, 0xb8, 0x03, 0x2a, 0x6e, 0xd4,
	0xb9, 0x46, 0xde, 0x48, 0x8f, 0xb8, 0x07, 0xc5, 0xef, 0x7e, 0xf0, 0xcd, 0xf1, 0x3d, 0x31, 0xc7,
	0xc0, 0x42, 0x5a, 0x2a, 0x1e, 0x8a, 0x50, 0xbc, 0x0f, 0xa7, 0x41, 0x12, 0xcf, 0xc4, 0x3c, 0x6b,
	0x2c, 0x4a, 0xfc, 0x00, 0x85, 0x01, 0x0d, 0x3c, 0x1a, 0x8b, 0x85, 0x3a, 0xd7, 0xa8, 0x5e, 0x9c,
	0x4a, 0x7f, 0x7c, 0xbe, 0x74, 0xc3, 0x40, 0x6b, 0x16, 0x51, 0x63, 0x6e, 0xc2, 0x1d, 0xc8, 0xbb,
	0xfd, 0x70, 0x9a, 0x88, 0x45, 0x16, 0x9b, 0x15, 0xf8, 0x11, 0xca, 0xee, 0x34, 0x19, 0x3a, 0xe9,
	0x58, 0xc4, 0x12, 0xcb, 0x3d, 0x79, 0x21, 0x57, 0x9e, 0x26, 0x43, 0x96, 0x5a, 0x72, 0xe7, 0x27,
	0xac, 0x41, 0xa9, 0xef, 0xc7, 0xc9, 0xd0, 0x73, 0x67, 0x62, 0x99, 0x45, 0x3f, 0xd6, 0x68, 0xc0,
	0x96, 0x47, 0x47, 0x34, 0xf1, 0xc3, 0xc0, 0x89, 0xa9, 0x3b, 0x09, 0x03, 0x11, 0xd8, 0x1d, 0x6f,
	0x5e, 0xb8, 0xa3, 0x35, 0x77, 0x18, 0xcc, 0x60, 0x54, 0xbd, 0x67, 0x35, 0xbe, 0x87, 0xca, 0x7d,
	0x4c, 0xdd, 0x84, 0x3a, 0xe9, 0x4a, 0xc4, 0x4a, 0x9d, 0x6b, 0x54, 0x2e, 0x6a, 0x52, 0xb6, 0x2f,
	0x69, 0xb1, 0x2f, 0xc9, 0x5a, 0xec, 0xcb, 0x80, 0x0c, 0x4f, 0x85, 0xd4, 0x3c, 0x8d, 0xbc, 0x47,
	0xf3, 0xdf, 0xaf, 0x9b, 0x33, 0x7c, 0x61, 0x66, 0x6f, 0x99, 0x9b, 0x37, 0x5f, 0x37, 0x67, 0x38,
	0x33, 0x1f, 0x40, 0x79, 0x12, 0xb8, 0x91, 0xc3, 0xbe, 0x8d, 0x6a, 0x36, 0xa7, 0x54, 0xd0, 0xdc,
	0x31, 0x3d, 0xbb, 0x85, 0xd2, 0x62, 0xb2, 0xf8, 0x2f, 0xfc, 0x23, 0xdb, 0x56, 0xdb, 0xb1, 0x7a,
	0x1d, 0xe2, 0x28, 0x5a, 0x57, 0x56, 0x95, 0x96, 0xf0, 0x17, 0x6e, 0xc3, 0xd6, 0x52, 0xee, 0xb4,
	0x75, 0x8d, 0x08, 0x1c, 0x22, 0x54, 0x97, 0xa2, 0xa9, 0xc9, 0x1d, 0x81, 0x3f, 0x33, 0x01, 0x96,
	0xdb, 0xc7, 0x3d, 0xd8, 0xbe, 0x21, 0x5a, 0x8b, 0x18, 0xab, 0x79, 0x3b, 0x20, 0x3c, 0x6d, 0x7c,
	0x92, 0xd5, 0x34, 0x70, 0x17, 0xf0, 0xa9, 0x7a, 0x4d, 0x98, 0xce, 0x9f, 0xfd, 0xe4, 0xa1, 0xfa,
	0x7c, 0x2f, 0x78, 0x00, 0x7b, 0x2d, 0xa2, 0x12, 0x4b, 0xd1, 0x35, 0xc7, 0x20, 0xb2, 0xa9, 0x6b,
	0x4f, 0xd2, 0x8f, 0xe1, 0x70, 0xb5, 0x69, 0x12, 0xf5, 0xda, 0x59, 0x88, 0x02, 0xb7, 0x0e, 0xb9,
	0xd4, 0x2d, 0xe7, 0x92, 0xb4, 0xe5, 0xae, 0xa2, 0xdb, 0x86, 0xc0, 0x63, 0x0d, 0x76, 0x57, 0x11,
	0xcd, 0x6e, 0x29, 0x56, 0x4f, 0xd8, 0xc0, 0x13, 0x38, 0xfa, 0xfd, 0x86, 0xcf, 0xb6, 0xac, 0x3a,
	0xf2, 0x95, 0xa5, 0x74, 0x53, 0x28, 0x87, 0xff, 0x41, 0x6d, 0x15, 0x92, 0x5b, 0xb6, 0x6a, 0x39,
	0x96, 0xde, 0x33, 0x85, 0x3c, 0x1e, 0xc2, 0xfe, 0x6a, 0xdf, 0x4e, 0x7f, 0xfd, 0x1d, 0x91, 0x0d,
	0xa1, 0x80, 0xa7, 0x70, 0xbc, 0xda, 0xee, 0xb4, 0x7b, 0xa6, 0x72, 0x25, 0xab, 0x4e, 0x57, 0xd1,
	0x55, 0xa2, 0x5d, 0x11, 0xa1, 0xb8, 0xee, 0x29, 0x77, 0x44, 0xee, 0xe8, 0xda, 0x12, 0x2a, 0xad,
	0x1b, 0x57, 0x06, 0x99, 0x42, 0x19, 0x8f, 0xe0, 0x60, 0xed, 0xb8, 0x14, 0xed, 0xd6, 0x36, 0x7a,
	0x02, 0x5c, 0x6e, 0x7e, 0xa9, 0x3c, 0xfe, 0x47, 0x1e, 0xce, 0xfb, 0x05, 0xf6, 0xb1, 0xbd, 0xfb,
	0x15, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x7f, 0x70, 0xce, 0xdd, 0x04, 0x00, 0x00,
}