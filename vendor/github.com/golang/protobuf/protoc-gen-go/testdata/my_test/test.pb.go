// Code generated by protoc-gen-go. DO NOT EDIT.
// source: my_test/test.proto

package test // import "github.com/golang/protobuf/protoc-gen-go/testdata/my_test"

/*
This package holds interesting messages.
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/protoc-gen-go/testdata/multi"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HatType int32

const (
	// deliberately skipping 0
	HatType_FEDORA HatType = 1
	HatType_FEZ    HatType = 2
)

var HatType_name = map[int32]string{
	1: "FEDORA",
	2: "FEZ",
}

var HatType_value = map[string]int32{
	"FEDORA": 1,
	"FEZ":    2,
}

func (x HatType) Enum() *HatType {
	p := new(HatType)
	*p = x
	return p
}

func (x HatType) String() string {
	return proto.EnumName(HatType_name, int32(x))
}

func (x *HatType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HatType_value, data, "HatType")
	if err != nil {
		return err
	}
	*x = HatType(value)
	return nil
}

func (HatType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{0}
}

// This enum represents days of the week.
type Days int32

const (
	Days_MONDAY  Days = 1
	Days_TUESDAY Days = 2
	Days_LUNDI   Days = 1
)

var Days_name = map[int32]string{
	1: "MONDAY",
	2: "TUESDAY",
	// Duplicate value: 1: "LUNDI",
}

var Days_value = map[string]int32{
	"MONDAY":  1,
	"TUESDAY": 2,
	"LUNDI":   1,
}

func (x Days) Enum() *Days {
	p := new(Days)
	*p = x
	return p
}

func (x Days) String() string {
	return proto.EnumName(Days_name, int32(x))
}

func (x *Days) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Days_value, data, "Days")
	if err != nil {
		return err
	}
	*x = Days(value)
	return nil
}

func (Days) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{1}
}

type Request_Color int32

const (
	Request_RED   Request_Color = 0
	Request_GREEN Request_Color = 1
	Request_BLUE  Request_Color = 2
)

var Request_Color_name = map[int32]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}

var Request_Color_value = map[string]int32{
	"RED":   0,
	"GREEN": 1,
	"BLUE":  2,
}

func (x Request_Color) Enum() *Request_Color {
	p := new(Request_Color)
	*p = x
	return p
}

func (x Request_Color) String() string {
	return proto.EnumName(Request_Color_name, int32(x))
}

func (x *Request_Color) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Request_Color_value, data, "Request_Color")
	if err != nil {
		return err
	}
	*x = Request_Color(value)
	return nil
}

func (Request_Color) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{0, 0}
}

type Reply_Entry_Game int32

const (
	Reply_Entry_FOOTBALL Reply_Entry_Game = 1
	Reply_Entry_TENNIS   Reply_Entry_Game = 2
)

var Reply_Entry_Game_name = map[int32]string{
	1: "FOOTBALL",
	2: "TENNIS",
}

var Reply_Entry_Game_value = map[string]int32{
	"FOOTBALL": 1,
	"TENNIS":   2,
}

func (x Reply_Entry_Game) Enum() *Reply_Entry_Game {
	p := new(Reply_Entry_Game)
	*p = x
	return p
}

func (x Reply_Entry_Game) String() string {
	return proto.EnumName(Reply_Entry_Game_name, int32(x))
}

func (x *Reply_Entry_Game) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Reply_Entry_Game_value, data, "Reply_Entry_Game")
	if err != nil {
		return err
	}
	*x = Reply_Entry_Game(value)
	return nil
}

func (Reply_Entry_Game) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{1, 0, 0}
}

// This is a message that might be sent somewhere.
type Request struct {
	Key []int64 `protobuf:"varint,1,rep,name=key" json:"key,omitempty"`
	//  optional imp.ImportedMessage imported_message = 2;
	Hue *Request_Color `protobuf:"varint,3,opt,name=hue,enum=my.test.Request_Color" json:"hue,omitempty"`
	Hat *HatType       `protobuf:"varint,4,opt,name=hat,enum=my.test.HatType,def=1" json:"hat,omitempty"`
	//  optional imp.ImportedMessage.Owner owner = 6;
	Deadline  *float32           `protobuf:"fixed32,7,opt,name=deadline,def=inf" json:"deadline,omitempty"`
	Somegroup *Request_SomeGroup `protobuf:"group,8,opt,name=SomeGroup,json=somegroup" json:"somegroup,omitempty"`
	// This is a map field. It will generate map[int32]string.
	NameMapping map[int32]string `protobuf:"bytes,14,rep,name=name_mapping,json=nameMapping" json:"name_mapping,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// This is a map field whose value type is a message.
	MsgMapping map[int64]*Reply `protobuf:"bytes,15,rep,name=msg_mapping,json=msgMapping" json:"msg_mapping,omitempty" protobuf_key:"zigzag64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Reset_     *int32           `protobuf:"varint,12,opt,name=reset" json:"reset,omitempty"`
	// This field should not conflict with any getters.
	GetKey_              *string  `protobuf:"bytes,16,opt,name=get_key,json=getKey" json:"get_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

const Default_Request_Hat HatType = HatType_FEDORA

var Default_Request_Deadline float32 = float32(math.Inf(1))

func (m *Request) GetKey() []int64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Request) GetHue() Request_Color {
	if m != nil && m.Hue != nil {
		return *m.Hue
	}
	return Request_RED
}

func (m *Request) GetHat() HatType {
	if m != nil && m.Hat != nil {
		return *m.Hat
	}
	return Default_Request_Hat
}

func (m *Request) GetDeadline() float32 {
	if m != nil && m.Deadline != nil {
		return *m.Deadline
	}
	return Default_Request_Deadline
}

func (m *Request) GetSomegroup() *Request_SomeGroup {
	if m != nil {
		return m.Somegroup
	}
	return nil
}

func (m *Request) GetNameMapping() map[int32]string {
	if m != nil {
		return m.NameMapping
	}
	return nil
}

func (m *Request) GetMsgMapping() map[int64]*Reply {
	if m != nil {
		return m.MsgMapping
	}
	return nil
}

func (m *Request) GetReset_() int32 {
	if m != nil && m.Reset_ != nil {
		return *m.Reset_
	}
	return 0
}

func (m *Request) GetGetKey_() string {
	if m != nil && m.GetKey_ != nil {
		return *m.GetKey_
	}
	return ""
}

type Request_SomeGroup struct {
	GroupField           *int32   `protobuf:"varint,9,opt,name=group_field,json=groupField" json:"group_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_SomeGroup) Reset()         { *m = Request_SomeGroup{} }
func (m *Request_SomeGroup) String() string { return proto.CompactTextString(m) }
func (*Request_SomeGroup) ProtoMessage()    {}
func (*Request_SomeGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{0, 0}
}
func (m *Request_SomeGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_SomeGroup.Unmarshal(m, b)
}
func (m *Request_SomeGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_SomeGroup.Marshal(b, m, deterministic)
}
func (m *Request_SomeGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_SomeGroup.Merge(m, src)
}
func (m *Request_SomeGroup) XXX_Size() int {
	return xxx_messageInfo_Request_SomeGroup.Size(m)
}
func (m *Request_SomeGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_SomeGroup.DiscardUnknown(m)
}

var xxx_messageInfo_Request_SomeGroup proto.InternalMessageInfo

func (m *Request_SomeGroup) GetGroupField() int32 {
	if m != nil && m.GroupField != nil {
		return *m.GroupField
	}
	return 0
}

type Reply struct {
	Found                        []*Reply_Entry `protobuf:"bytes,1,rep,name=found" json:"found,omitempty"`
	CompactKeys                  []int32        `protobuf:"varint,2,rep,packed,name=compact_keys,json=compactKeys" json:"compact_keys,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}       `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{1}
}

var extRange_Reply = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*Reply) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Reply
}
func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetFound() []*Reply_Entry {
	if m != nil {
		return m.Found
	}
	return nil
}

func (m *Reply) GetCompactKeys() []int32 {
	if m != nil {
		return m.CompactKeys
	}
	return nil
}

type Reply_Entry struct {
	KeyThatNeeds_1234Camel_CasIng *int64   `protobuf:"varint,1,req,name=key_that_needs_1234camel_CasIng,json=keyThatNeeds1234camelCasIng" json:"key_that_needs_1234camel_CasIng,omitempty"`
	Value                         *int64   `protobuf:"varint,2,opt,name=value,def=7" json:"value,omitempty"`
	XMyFieldName_2                *int64   `protobuf:"varint,3,opt,name=_my_field_name_2,json=MyFieldName2" json:"_my_field_name_2,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *Reply_Entry) Reset()         { *m = Reply_Entry{} }
func (m *Reply_Entry) String() string { return proto.CompactTextString(m) }
func (*Reply_Entry) ProtoMessage()    {}
func (*Reply_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{1, 0}
}
func (m *Reply_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply_Entry.Unmarshal(m, b)
}
func (m *Reply_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply_Entry.Marshal(b, m, deterministic)
}
func (m *Reply_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply_Entry.Merge(m, src)
}
func (m *Reply_Entry) XXX_Size() int {
	return xxx_messageInfo_Reply_Entry.Size(m)
}
func (m *Reply_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Reply_Entry proto.InternalMessageInfo

const Default_Reply_Entry_Value int64 = 7

func (m *Reply_Entry) GetKeyThatNeeds_1234Camel_CasIng() int64 {
	if m != nil && m.KeyThatNeeds_1234Camel_CasIng != nil {
		return *m.KeyThatNeeds_1234Camel_CasIng
	}
	return 0
}

func (m *Reply_Entry) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return Default_Reply_Entry_Value
}

func (m *Reply_Entry) GetXMyFieldName_2() int64 {
	if m != nil && m.XMyFieldName_2 != nil {
		return *m.XMyFieldName_2
	}
	return 0
}

type OtherBase struct {
	Name                         *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *OtherBase) Reset()         { *m = OtherBase{} }
func (m *OtherBase) String() string { return proto.CompactTextString(m) }
func (*OtherBase) ProtoMessage()    {}
func (*OtherBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{2}
}

var extRange_OtherBase = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*OtherBase) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_OtherBase
}
func (m *OtherBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtherBase.Unmarshal(m, b)
}
func (m *OtherBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtherBase.Marshal(b, m, deterministic)
}
func (m *OtherBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtherBase.Merge(m, src)
}
func (m *OtherBase) XXX_Size() int {
	return xxx_messageInfo_OtherBase.Size(m)
}
func (m *OtherBase) XXX_DiscardUnknown() {
	xxx_messageInfo_OtherBase.DiscardUnknown(m)
}

var xxx_messageInfo_OtherBase proto.InternalMessageInfo

func (m *OtherBase) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type ReplyExtensions struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyExtensions) Reset()         { *m = ReplyExtensions{} }
func (m *ReplyExtensions) String() string { return proto.CompactTextString(m) }
func (*ReplyExtensions) ProtoMessage()    {}
func (*ReplyExtensions) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{3}
}
func (m *ReplyExtensions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyExtensions.Unmarshal(m, b)
}
func (m *ReplyExtensions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyExtensions.Marshal(b, m, deterministic)
}
func (m *ReplyExtensions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyExtensions.Merge(m, src)
}
func (m *ReplyExtensions) XXX_Size() int {
	return xxx_messageInfo_ReplyExtensions.Size(m)
}
func (m *ReplyExtensions) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyExtensions.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyExtensions proto.InternalMessageInfo

var E_ReplyExtensions_Time = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*float64)(nil),
	Field:         101,
	Name:          "my.test.ReplyExtensions.time",
	Tag:           "fixed64,101,opt,name=time",
	Filename:      "my_test/test.proto",
}

var E_ReplyExtensions_Carrot = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*ReplyExtensions)(nil),
	Field:         105,
	Name:          "my.test.ReplyExtensions.carrot",
	Tag:           "bytes,105,opt,name=carrot",
	Filename:      "my_test/test.proto",
}

var E_ReplyExtensions_Donut = &proto.ExtensionDesc{
	ExtendedType:  (*OtherBase)(nil),
	ExtensionType: (*ReplyExtensions)(nil),
	Field:         101,
	Name:          "my.test.ReplyExtensions.donut",
	Tag:           "bytes,101,opt,name=donut",
	Filename:      "my_test/test.proto",
}

type OtherReplyExtensions struct {
	Key                  *int32   `protobuf:"varint,1,opt,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OtherReplyExtensions) Reset()         { *m = OtherReplyExtensions{} }
func (m *OtherReplyExtensions) String() string { return proto.CompactTextString(m) }
func (*OtherReplyExtensions) ProtoMessage()    {}
func (*OtherReplyExtensions) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{4}
}
func (m *OtherReplyExtensions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtherReplyExtensions.Unmarshal(m, b)
}
func (m *OtherReplyExtensions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtherReplyExtensions.Marshal(b, m, deterministic)
}
func (m *OtherReplyExtensions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtherReplyExtensions.Merge(m, src)
}
func (m *OtherReplyExtensions) XXX_Size() int {
	return xxx_messageInfo_OtherReplyExtensions.Size(m)
}
func (m *OtherReplyExtensions) XXX_DiscardUnknown() {
	xxx_messageInfo_OtherReplyExtensions.DiscardUnknown(m)
}

var xxx_messageInfo_OtherReplyExtensions proto.InternalMessageInfo

func (m *OtherReplyExtensions) GetKey() int32 {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return 0
}

type OldReply struct {
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `protobuf_messageset:"1" json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *OldReply) Reset()         { *m = OldReply{} }
func (m *OldReply) String() string { return proto.CompactTextString(m) }
func (*OldReply) ProtoMessage()    {}
func (*OldReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{5}
}

func (m *OldReply) MarshalJSON() ([]byte, error) {
	return proto.MarshalMessageSetJSON(&m.XXX_InternalExtensions)
}
func (m *OldReply) UnmarshalJSON(buf []byte) error {
	return proto.UnmarshalMessageSetJSON(buf, &m.XXX_InternalExtensions)
}

var extRange_OldReply = []proto.ExtensionRange{
	{Start: 100, End: 2147483646},
}

func (*OldReply) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_OldReply
}
func (m *OldReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OldReply.Unmarshal(m, b)
}
func (m *OldReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OldReply.Marshal(b, m, deterministic)
}
func (m *OldReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OldReply.Merge(m, src)
}
func (m *OldReply) XXX_Size() int {
	return xxx_messageInfo_OldReply.Size(m)
}
func (m *OldReply) XXX_DiscardUnknown() {
	xxx_messageInfo_OldReply.DiscardUnknown(m)
}

var xxx_messageInfo_OldReply proto.InternalMessageInfo

type Communique struct {
	MakeMeCry *bool `protobuf:"varint,1,opt,name=make_me_cry,json=makeMeCry" json:"make_me_cry,omitempty"`
	// This is a oneof, called "union".
	//
	// Types that are valid to be assigned to Union:
	//	*Communique_Number
	//	*Communique_Name
	//	*Communique_Data
	//	*Communique_TempC
	//	*Communique_Height
	//	*Communique_Today
	//	*Communique_Maybe
	//	*Communique_Delta_
	//	*Communique_Msg
	//	*Communique_Somegroup
	Union                isCommunique_Union `protobuf_oneof:"union"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Communique) Reset()         { *m = Communique{} }
func (m *Communique) String() string { return proto.CompactTextString(m) }
func (*Communique) ProtoMessage()    {}
func (*Communique) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{6}
}
func (m *Communique) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Communique.Unmarshal(m, b)
}
func (m *Communique) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Communique.Marshal(b, m, deterministic)
}
func (m *Communique) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Communique.Merge(m, src)
}
func (m *Communique) XXX_Size() int {
	return xxx_messageInfo_Communique.Size(m)
}
func (m *Communique) XXX_DiscardUnknown() {
	xxx_messageInfo_Communique.DiscardUnknown(m)
}

var xxx_messageInfo_Communique proto.InternalMessageInfo

func (m *Communique) GetMakeMeCry() bool {
	if m != nil && m.MakeMeCry != nil {
		return *m.MakeMeCry
	}
	return false
}

type isCommunique_Union interface {
	isCommunique_Union()
}

type Communique_Number struct {
	Number int32 `protobuf:"varint,5,opt,name=number,oneof"`
}

type Communique_Name struct {
	Name string `protobuf:"bytes,6,opt,name=name,oneof"`
}

type Communique_Data struct {
	Data []byte `protobuf:"bytes,7,opt,name=data,oneof"`
}

type Communique_TempC struct {
	TempC float64 `protobuf:"fixed64,8,opt,name=temp_c,json=tempC,oneof"`
}

type Communique_Height struct {
	Height float32 `protobuf:"fixed32,9,opt,name=height,oneof"`
}

type Communique_Today struct {
	Today Days `protobuf:"varint,10,opt,name=today,enum=my.test.Days,oneof"`
}

type Communique_Maybe struct {
	Maybe bool `protobuf:"varint,11,opt,name=maybe,oneof"`
}

type Communique_Delta_ struct {
	Delta int32 `protobuf:"zigzag32,12,opt,name=delta,oneof"`
}

type Communique_Msg struct {
	Msg *Reply `protobuf:"bytes,16,opt,name=msg,oneof"`
}

type Communique_Somegroup struct {
	Somegroup *Communique_SomeGroup `protobuf:"group,14,opt,name=SomeGroup,json=somegroup,oneof"`
}

func (*Communique_Number) isCommunique_Union() {}

func (*Communique_Name) isCommunique_Union() {}

func (*Communique_Data) isCommunique_Union() {}

func (*Communique_TempC) isCommunique_Union() {}

func (*Communique_Height) isCommunique_Union() {}

func (*Communique_Today) isCommunique_Union() {}

func (*Communique_Maybe) isCommunique_Union() {}

func (*Communique_Delta_) isCommunique_Union() {}

func (*Communique_Msg) isCommunique_Union() {}

func (*Communique_Somegroup) isCommunique_Union() {}

func (m *Communique) GetUnion() isCommunique_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *Communique) GetNumber() int32 {
	if x, ok := m.GetUnion().(*Communique_Number); ok {
		return x.Number
	}
	return 0
}

func (m *Communique) GetName() string {
	if x, ok := m.GetUnion().(*Communique_Name); ok {
		return x.Name
	}
	return ""
}

func (m *Communique) GetData() []byte {
	if x, ok := m.GetUnion().(*Communique_Data); ok {
		return x.Data
	}
	return nil
}

func (m *Communique) GetTempC() float64 {
	if x, ok := m.GetUnion().(*Communique_TempC); ok {
		return x.TempC
	}
	return 0
}

func (m *Communique) GetHeight() float32 {
	if x, ok := m.GetUnion().(*Communique_Height); ok {
		return x.Height
	}
	return 0
}

func (m *Communique) GetToday() Days {
	if x, ok := m.GetUnion().(*Communique_Today); ok {
		return x.Today
	}
	return Days_MONDAY
}

func (m *Communique) GetMaybe() bool {
	if x, ok := m.GetUnion().(*Communique_Maybe); ok {
		return x.Maybe
	}
	return false
}

func (m *Communique) GetDelta() int32 {
	if x, ok := m.GetUnion().(*Communique_Delta_); ok {
		return x.Delta
	}
	return 0
}

func (m *Communique) GetMsg() *Reply {
	if x, ok := m.GetUnion().(*Communique_Msg); ok {
		return x.Msg
	}
	return nil
}

func (m *Communique) GetSomegroup() *Communique_SomeGroup {
	if x, ok := m.GetUnion().(*Communique_Somegroup); ok {
		return x.Somegroup
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Communique) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Communique_OneofMarshaler, _Communique_OneofUnmarshaler, _Communique_OneofSizer, []interface{}{
		(*Communique_Number)(nil),
		(*Communique_Name)(nil),
		(*Communique_Data)(nil),
		(*Communique_TempC)(nil),
		(*Communique_Height)(nil),
		(*Communique_Today)(nil),
		(*Communique_Maybe)(nil),
		(*Communique_Delta_)(nil),
		(*Communique_Msg)(nil),
		(*Communique_Somegroup)(nil),
	}
}

func _Communique_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Communique)
	// union
	switch x := m.Union.(type) {
	case *Communique_Number:
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Number))
	case *Communique_Name:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Name)
	case *Communique_Data:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Data)
	case *Communique_TempC:
		b.EncodeVarint(8<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.TempC))
	case *Communique_Height:
		b.EncodeVarint(9<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.Height)))
	case *Communique_Today:
		b.EncodeVarint(10<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Today))
	case *Communique_Maybe:
		t := uint64(0)
		if x.Maybe {
			t = 1
		}
		b.EncodeVarint(11<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Communique_Delta_:
		b.EncodeVarint(12<<3 | proto.WireVarint)
		b.EncodeZigzag32(uint64(x.Delta))
	case *Communique_Msg:
		b.EncodeVarint(16<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Msg); err != nil {
			return err
		}
	case *Communique_Somegroup:
		b.EncodeVarint(14<<3 | proto.WireStartGroup)
		if err := b.Marshal(x.Somegroup); err != nil {
			return err
		}
		b.EncodeVarint(14<<3 | proto.WireEndGroup)
	case nil:
	default:
		return fmt.Errorf("Communique.Union has unexpected type %T", x)
	}
	return nil
}

func _Communique_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Communique)
	switch tag {
	case 5: // union.number
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Union = &Communique_Number{int32(x)}
		return true, err
	case 6: // union.name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Union = &Communique_Name{x}
		return true, err
	case 7: // union.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Union = &Communique_Data{x}
		return true, err
	case 8: // union.temp_c
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Union = &Communique_TempC{math.Float64frombits(x)}
		return true, err
	case 9: // union.height
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Union = &Communique_Height{math.Float32frombits(uint32(x))}
		return true, err
	case 10: // union.today
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Union = &Communique_Today{Days(x)}
		return true, err
	case 11: // union.maybe
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Union = &Communique_Maybe{x != 0}
		return true, err
	case 12: // union.delta
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeZigzag32()
		m.Union = &Communique_Delta_{int32(x)}
		return true, err
	case 16: // union.msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Reply)
		err := b.DecodeMessage(msg)
		m.Union = &Communique_Msg{msg}
		return true, err
	case 14: // union.somegroup
		if wire != proto.WireStartGroup {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Communique_SomeGroup)
		err := b.DecodeGroup(msg)
		m.Union = &Communique_Somegroup{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Communique_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Communique)
	// union
	switch x := m.Union.(type) {
	case *Communique_Number:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Number))
	case *Communique_Name:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case *Communique_Data:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Data)))
		n += len(x.Data)
	case *Communique_TempC:
		n += 1 // tag and wire
		n += 8
	case *Communique_Height:
		n += 1 // tag and wire
		n += 4
	case *Communique_Today:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Today))
	case *Communique_Maybe:
		n += 1 // tag and wire
		n += 1
	case *Communique_Delta_:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64((uint32(x.Delta) << 1) ^ uint32((int32(x.Delta) >> 31))))
	case *Communique_Msg:
		s := proto.Size(x.Msg)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Communique_Somegroup:
		n += 1 // tag and wire
		n += proto.Size(x.Somegroup)
		n += 1 // tag and wire
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Communique_SomeGroup struct {
	Member               *string  `protobuf:"bytes,15,opt,name=member" json:"member,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Communique_SomeGroup) Reset()         { *m = Communique_SomeGroup{} }
func (m *Communique_SomeGroup) String() string { return proto.CompactTextString(m) }
func (*Communique_SomeGroup) ProtoMessage()    {}
func (*Communique_SomeGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{6, 0}
}
func (m *Communique_SomeGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Communique_SomeGroup.Unmarshal(m, b)
}
func (m *Communique_SomeGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Communique_SomeGroup.Marshal(b, m, deterministic)
}
func (m *Communique_SomeGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Communique_SomeGroup.Merge(m, src)
}
func (m *Communique_SomeGroup) XXX_Size() int {
	return xxx_messageInfo_Communique_SomeGroup.Size(m)
}
func (m *Communique_SomeGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_Communique_SomeGroup.DiscardUnknown(m)
}

var xxx_messageInfo_Communique_SomeGroup proto.InternalMessageInfo

func (m *Communique_SomeGroup) GetMember() string {
	if m != nil && m.Member != nil {
		return *m.Member
	}
	return ""
}

type Communique_Delta struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Communique_Delta) Reset()         { *m = Communique_Delta{} }
func (m *Communique_Delta) String() string { return proto.CompactTextString(m) }
func (*Communique_Delta) ProtoMessage()    {}
func (*Communique_Delta) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{6, 1}
}
func (m *Communique_Delta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Communique_Delta.Unmarshal(m, b)
}
func (m *Communique_Delta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Communique_Delta.Marshal(b, m, deterministic)
}
func (m *Communique_Delta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Communique_Delta.Merge(m, src)
}
func (m *Communique_Delta) XXX_Size() int {
	return xxx_messageInfo_Communique_Delta.Size(m)
}
func (m *Communique_Delta) XXX_DiscardUnknown() {
	xxx_messageInfo_Communique_Delta.DiscardUnknown(m)
}

var xxx_messageInfo_Communique_Delta proto.InternalMessageInfo

var E_Tag = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*string)(nil),
	Field:         103,
	Name:          "my.test.tag",
	Tag:           "bytes,103,opt,name=tag",
	Filename:      "my_test/test.proto",
}

var E_Donut = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*OtherReplyExtensions)(nil),
	Field:         106,
	Name:          "my.test.donut",
	Tag:           "bytes,106,opt,name=donut",
	Filename:      "my_test/test.proto",
}

func init() {
	proto.RegisterType((*Request)(nil), "my.test.Request")
	proto.RegisterMapType((map[int64]*Reply)(nil), "my.test.Request.MsgMappingEntry")
	proto.RegisterMapType((map[int32]string)(nil), "my.test.Request.NameMappingEntry")
	proto.RegisterType((*Request_SomeGroup)(nil), "my.test.Request.SomeGroup")
	proto.RegisterType((*Reply)(nil), "my.test.Reply")
	proto.RegisterType((*Reply_Entry)(nil), "my.test.Reply.Entry")
	proto.RegisterType((*OtherBase)(nil), "my.test.OtherBase")
	proto.RegisterType((*ReplyExtensions)(nil), "my.test.ReplyExtensions")
	proto.RegisterType((*OtherReplyExtensions)(nil), "my.test.OtherReplyExtensions")
	proto.RegisterType((*OldReply)(nil), "my.test.OldReply")
	proto.RegisterType((*Communique)(nil), "my.test.Communique")
	proto.RegisterType((*Communique_SomeGroup)(nil), "my.test.Communique.SomeGroup")
	proto.RegisterType((*Communique_Delta)(nil), "my.test.Communique.Delta")
	proto.RegisterEnum("my.test.HatType", HatType_name, HatType_value)
	proto.RegisterEnum("my.test.Days", Days_name, Days_value)
	proto.RegisterEnum("my.test.Request_Color", Request_Color_name, Request_Color_value)
	proto.RegisterEnum("my.test.Reply_Entry_Game", Reply_Entry_Game_name, Reply_Entry_Game_value)
	proto.RegisterExtension(E_ReplyExtensions_Time)
	proto.RegisterExtension(E_ReplyExtensions_Carrot)
	proto.RegisterExtension(E_ReplyExtensions_Donut)
	proto.RegisterExtension(E_Tag)
	proto.RegisterExtension(E_Donut)
}

func init() { proto.RegisterFile("my_test/test.proto", fileDescriptor_2c9b60a40d5131b9) }

var fileDescriptor_2c9b60a40d5131b9 = []byte{
	// 1033 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xce, 0xd8, 0x71, 0x7e, 0x4e, 0x42, 0x6b, 0x46, 0x55, 0x6b, 0x05, 0xed, 0xd6, 0x04, 0x8a,
	0x4c, 0xc5, 0xa6, 0xda, 0x80, 0xc4, 0x2a, 0x88, 0xd5, 0x36, 0x3f, 0x6d, 0xaa, 0x6d, 0x12, 0x69,
	0xda, 0x5e, 0xb0, 0x37, 0xd6, 0x34, 0x9e, 0x3a, 0xa6, 0x19, 0x3b, 0x6b, 0x8f, 0x11, 0xbe, 0xeb,
	0x53, 0xc0, 0x6b, 0x70, 0xcf, 0x0b, 0xf1, 0x16, 0x45, 0x33, 0x0e, 0x49, 0xda, 0xa0, 0xbd, 0xb1,
	0x7c, 0xce, 0xf9, 0xce, 0xe7, 0x39, 0x3f, 0xfe, 0x06, 0x30, 0xcf, 0x5c, 0xc1, 0x12, 0x71, 0x22,
	0x1f, 0xad, 0x45, 0x1c, 0x89, 0x08, 0x97, 0x79, 0xd6, 0x92, 0x66, 0x03, 0xf3, 0x74, 0x2e, 0x82,
	0x13, 0xf5, 0x7c, 0x9d, 0x07, 0x9b, 0xff, 0x14, 0xa1, 0x4c, 0xd8, 0xc7, 0x94, 0x25, 0x02, 0x9b,
	0xa0, 0xdf, 0xb3, 0xcc, 0x42, 0xb6, 0xee, 0xe8, 0x44, 0xbe, 0x62, 0x07, 0xf4, 0x59, 0xca, 0x2c,
	0xdd, 0x46, 0xce, 0x4e, 0x7b, 0xbf, 0xb5, 0x24, 0x6a, 0x2d, 0x13, 0x5a, 0xbd, 0x68, 0x1e, 0xc5,
	0x44, 0x42, 0xf0, 0x31, 0xe8, 0x33, 0x2a, 0xac, 0xa2, 0x42, 0x9a, 0x2b, 0xe4, 0x90, 0x8a, 0xeb,
	0x6c, 0xc1, 0x3a, 0xa5, 0xb3, 0x41, 0x7f, 0x42, 0x4e, 0x89, 0x04, 0xe1, 0x43, 0xa8, 0x78, 0x8c,
	0x7a, 0xf3, 0x20, 0x64, 0x56, 0xd9, 0x46, 0x8e, 0xd6, 0xd1, 0x83, 0xf0, 0x8e, 0xac, 0x9c, 0xf8,
	0x0d, 0x54, 0x93, 0x88, 0x33, 0x3f, 0x8e, 0xd2, 0x85, 0x55, 0xb1, 0x91, 0x03, 0xed, 0xc6, 0xd6,
	0xc7, 0xaf, 0x22, 0xce, 0xce, 0x25, 0x82, 0xac, 0xc1, 0xb8, 0x0f, 0xf5, 0x90, 0x72, 0xe6, 0x72,
	0xba, 0x58, 0x04, 0xa1, 0x6f, 0xed, 0xd8, 0xba, 0x53, 0x6b, 0x7f, 0xb9, 0x95, 0x3c, 0xa6, 0x9c,
	0x8d, 0x72, 0xcc, 0x20, 0x14, 0x71, 0x46, 0x6a, 0xe1, 0xda, 0x83, 0x4f, 0xa1, 0xc6, 0x13, 0x7f,
	0x45, 0xb2, 0xab, 0x48, 0xec, 0x2d, 0x92, 0x51, 0xe2, 0x3f, 0xe1, 0x00, 0xbe, 0x72, 0xe0, 0x3d,
	0x30, 0x62, 0x96, 0x30, 0x61, 0xd5, 0x6d, 0xe4, 0x18, 0x24, 0x37, 0xf0, 0x01, 0x94, 0x7d, 0x26,
	0x5c, 0xd9, 0x65, 0xd3, 0x46, 0x4e, 0x95, 0x94, 0x7c, 0x26, 0xde, 0xb3, 0xac, 0xf1, 0x1d, 0x54,
	0x57, 0xf5, 0xe0, 0x43, 0xa8, 0xa9, 0x6a, 0xdc, 0xbb, 0x80, 0xcd, 0x3d, 0xab, 0xaa, 0x18, 0x40,
	0xb9, 0xce, 0xa4, 0xa7, 0xf1, 0x16, 0xcc, 0xe7, 0x05, 0xac, 0x87, 0x27, 0xc1, 0x6a, 0x78, 0x7b,
	0x60, 0xfc, 0x46, 0xe7, 0x29, 0xb3, 0x34, 0xf5, 0xa9, 0xdc, 0xe8, 0x68, 0x6f, 0x50, 0x63, 0x04,
	0xbb, 0xcf, 0xce, 0xbe, 0x99, 0x8e, 0xf3, 0xf4, 0xaf, 0x37, 0xd3, 0x6b, 0xed, 0x9d, 0x8d, 0xf2,
	0x17, 0xf3, 0x6c, 0x83, 0xae, 0x79, 0x04, 0x86, 0xda, 0x04, 0x5c, 0x06, 0x9d, 0x0c, 0xfa, 0x66,
	0x01, 0x57, 0xc1, 0x38, 0x27, 0x83, 0xc1, 0xd8, 0x44, 0xb8, 0x02, 0xc5, 0xee, 0xe5, 0xcd, 0xc0,
	0xd4, 0x9a, 0x7f, 0x6a, 0x60, 0xa8, 0x5c, 0x7c, 0x0c, 0xc6, 0x5d, 0x94, 0x86, 0x9e, 0x5a, 0xb5,
	0x5a, 0x7b, 0xef, 0x29, 0x75, 0x2b, 0xef, 0x66, 0x0e, 0xc1, 0x47, 0x50, 0x9f, 0x46, 0x7c, 0x41,
	0xa7, 0xaa, 0x6d, 0x89, 0xa5, 0xd9, 0xba, 0x63, 0x74, 0x35, 0x13, 0x91, 0xda, 0xd2, 0xff, 0x9e,
	0x65, 0x49, 0xe3, 0x2f, 0x04, 0x46, 0x5e, 0x49, 0x1f, 0x0e, 0xef, 0x59, 0xe6, 0x8a, 0x19, 0x15,
	0x6e, 0xc8, 0x98, 0x97, 0xb8, 0xaf, 0xdb, 0xdf, 0xff, 0x30, 0xa5, 0x9c, 0xcd, 0xdd, 0x1e, 0x4d,
	0x2e, 0x42, 0xdf, 0x42, 0xb6, 0xe6, 0xe8, 0xe4, 0x8b, 0x7b, 0x96, 0x5d, 0xcf, 0xa8, 0x18, 0x4b,
	0xd0, 0x0a, 0x93, 0x43, 0xf0, 0xc1, 0x66, 0xf5, 0x7a, 0x07, 0xfd, 0xb8, 0x2c, 0x18, 0x7f, 0x03,
	0xa6, 0xcb, 0xb3, 0x7c, 0x34, 0xae, 0xda, 0xb5, 0xb6, 0xfa, 0x3f, 0x74, 0x52, 0x1f, 0x65, 0x6a,
	0x3c, 0x72, 0x34, 0xed, 0xa6, 0x0d, 0xc5, 0x73, 0xca, 0x19, 0xae, 0x43, 0xe5, 0x6c, 0x32, 0xb9,
	0xee, 0x9e, 0x5e, 0x5e, 0x9a, 0x08, 0x03, 0x94, 0xae, 0x07, 0xe3, 0xf1, 0xc5, 0x95, 0xa9, 0x1d,
	0x57, 0x2a, 0x9e, 0xf9, 0xf0, 0xf0, 0xf0, 0xa0, 0x35, 0xbf, 0x85, 0xea, 0x44, 0xcc, 0x58, 0xdc,
	0xa5, 0x09, 0xc3, 0x18, 0x8a, 0x92, 0x56, 0x8d, 0xa2, 0x4a, 0xd4, 0xfb, 0x06, 0xf4, 0x6f, 0x04,
	0xbb, 0xaa, 0x4b, 0x83, 0xdf, 0x05, 0x0b, 0x93, 0x20, 0x0a, 0x93, 0x76, 0x13, 0x8a, 0x22, 0xe0,
	0x0c, 0x3f, 0x1b, 0x91, 0xc5, 0x6c, 0xe4, 0x20, 0xa2, 0x62, 0xed, 0x77, 0x50, 0x9a, 0xd2, 0x38,
	0x8e, 0xc4, 0x16, 0x2a, 0x50, 0xe3, 0xb5, 0x9e, 0x7a, 0xd7, 0xec, 0x64, 0x99, 0xd7, 0xee, 0x82,
	0xe1, 0x45, 0x61, 0x2a, 0x30, 0x5e, 0x41, 0x57, 0x87, 0x56, 0x9f, 0xfa, 0x14, 0x49, 0x9e, 0xda,
	0x74, 0x60, 0x4f, 0xe5, 0x3c, 0x0b, 0x6f, 0x2f, 0x6f, 0xd3, 0x82, 0xca, 0x64, 0xee, 0x29, 0x9c,
	0xaa, 0xfe, 0xf1, 0xf1, 0xf1, 0xb1, 0xdc, 0xd1, 0x2a, 0xa8, 0xf9, 0x87, 0x0e, 0xd0, 0x8b, 0x38,
	0x4f, 0xc3, 0xe0, 0x63, 0xca, 0xf0, 0x4b, 0xa8, 0x71, 0x7a, 0xcf, 0x5c, 0xce, 0xdc, 0x69, 0x9c,
	0x53, 0x54, 0x48, 0x55, 0xba, 0x46, 0xac, 0x17, 0x67, 0xd8, 0x82, 0x52, 0x98, 0xf2, 0x5b, 0x16,
	0x5b, 0x86, 0x64, 0x1f, 0x16, 0xc8, 0xd2, 0xc6, 0x7b, 0xcb, 0x46, 0x97, 0x64, 0xa3, 0x87, 0x85,
	0xbc, 0xd5, 0xd2, 0xeb, 0x51, 0x41, 0x95, 0x30, 0xd5, 0xa5, 0x57, 0x5a, 0xf8, 0x00, 0x4a, 0x82,
	0xf1, 0x85, 0x3b, 0x55, 0x72, 0x84, 0x86, 0x05, 0x62, 0x48, 0xbb, 0x27, 0xe9, 0x67, 0x2c, 0xf0,
	0x67, 0x42, 0xfd, 0xa6, 0x9a, 0xa4, 0xcf, 0x6d, 0x7c, 0x04, 0x86, 0x88, 0x3c, 0x9a, 0x59, 0xa0,
	0x34, 0xf1, 0xb3, 0x55, 0x6f, 0xfa, 0x34, 0x4b, 0x14, 0x81, 0x8c, 0xe2, 0x7d, 0x30, 0x38, 0xcd,
	0x6e, 0x99, 0x55, 0x93, 0x27, 0x97, 0x7e, 0x65, 0x4a, 0xbf, 0xc7, 0xe6, 0x82, 0x2a, 0x01, 0xf9,
	0x5c, 0xfa, 0x95, 0x89, 0x9b, 0xa0, 0xf3, 0xc4, 0x57, 0xf2, 0xb1, 0xf5, 0x53, 0x0e, 0x0b, 0x44,
	0x06, 0xf1, 0xcf, 0x9b, 0xfa, 0xb9, 0xa3, 0xf4, 0xf3, 0xc5, 0x0a, 0xb9, 0xee, 0xdd, 0x5a, 0x42,
	0x87, 0x85, 0x0d, 0x11, 0x6d, 0x7c, 0xb5, 0x29, 0x46, 0xfb, 0x50, 0xe2, 0x4c, 0xf5, 0x6f, 0x37,
	0x57, 0xac, 0xdc, 0x6a, 0x94, 0xc1, 0xe8, 0xcb, 0x03, 0x75, 0xcb, 0x60, 0xa4, 0x61, 0x10, 0x85,
	0xc7, 0x2f, 0xa1, 0xbc, 0x94, 0x7b, 0xb9, 0xe6, 0xb9, 0xe0, 0x9b, 0x48, 0x8a, 0xc2, 0xd9, 0xe0,
	0x83, 0xa9, 0x1d, 0xb7, 0xa0, 0x28, 0x4b, 0x97, 0xc1, 0xd1, 0x64, 0xdc, 0x3f, 0xfd, 0xc5, 0x44,
	0xb8, 0x06, 0xe5, 0xeb, 0x9b, 0xc1, 0x95, 0x34, 0x34, 0xa9, 0x1a, 0x97, 0x37, 0xe3, 0xfe, 0x85,
	0x89, 0x1a, 0x9a, 0x89, 0x3a, 0x36, 0xe8, 0x82, 0xfa, 0x5b, 0xfb, 0xea, 0xab, 0x63, 0xc8, 0x50,
	0xa7, 0xf7, 0xdf, 0x4a, 0x3e, 0xc7, 0xfc, 0xaa, 0xba, 0xf3, 0xe2, 0xe9, 0xa2, 0xfe, 0xff, 0x4e,
	0x76, 0xdf, 0x7d, 0x78, 0xeb, 0x07, 0x62, 0x96, 0xde, 0xb6, 0xa6, 0x11, 0x3f, 0xf1, 0xa3, 0x39,
	0x0d, 0xfd, 0x13, 0x75, 0x39, 0xde, 0xa6, 0x77, 0xf9, 0xcb, 0xf4, 0x95, 0xcf, 0xc2, 0x57, 0x7e,
	0xa4, 0x6e, 0x55, 0xb9, 0x0f, 0x27, 0xcb, 0x6b, 0xf6, 0x27, 0xf9, 0xf8, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x12, 0xd5, 0x46, 0x00, 0x75, 0x07, 0x00, 0x00,
}
