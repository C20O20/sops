// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/websecurityscanner/v1alpha/finding.proto

package websecurityscanner // import "google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1alpha"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Types of Findings.
type Finding_FindingType int32

const (
	// The invalid finding type.
	Finding_FINDING_TYPE_UNSPECIFIED Finding_FindingType = 0
	// A page that was served over HTTPS also resources over HTTP. A
	// man-in-the-middle attacker could tamper with the HTTP resource and gain
	// full access to the website that loads the resource or to monitor the
	// actions taken by the user.
	Finding_MIXED_CONTENT Finding_FindingType = 1
	// The version of an included library is known to contain a security issue.
	// The scanner checks the version of library in use against a known list of
	// vulnerable libraries. False positives are possible if the version
	// detection fails or if the library has been manually patched.
	Finding_OUTDATED_LIBRARY Finding_FindingType = 2
	// This type of vulnerability occurs when the value of a request parameter
	// is reflected at the beginning of the response, for example, in requests
	// using JSONP. Under certain circumstances, an attacker may be able to
	// supply an alphanumeric-only Flash file in the vulnerable parameter
	// causing the browser to execute the Flash file as if it originated on the
	// vulnerable server.
	Finding_ROSETTA_FLASH Finding_FindingType = 5
	// A cross-site scripting (XSS) bug is found via JavaScript callback. For
	// detailed explanations on XSS, see
	// https://www.google.com/about/appsecurity/learning/xss/.
	Finding_XSS_CALLBACK Finding_FindingType = 3
	// A potential cross-site scripting (XSS) bug due to JavaScript breakage.
	// In some circumstances, the application under test might modify the test
	// string before it is parsed by the browser. When the browser attempts to
	// runs this modified test string, it will likely break and throw a
	// JavaScript execution error, thus an injection issue is occurring.
	// However, it may not be exploitable. Manual verification is needed to see
	// if the test string modifications can be evaded and confirm that the issue
	// is in fact an XSS vulnerability. For detailed explanations on XSS, see
	// https://www.google.com/about/appsecurity/learning/xss/.
	Finding_XSS_ERROR Finding_FindingType = 4
	// An application appears to be transmitting a password field in clear text.
	// An attacker can eavesdrop network traffic and sniff the password field.
	Finding_CLEAR_TEXT_PASSWORD Finding_FindingType = 6
)

var Finding_FindingType_name = map[int32]string{
	0: "FINDING_TYPE_UNSPECIFIED",
	1: "MIXED_CONTENT",
	2: "OUTDATED_LIBRARY",
	5: "ROSETTA_FLASH",
	3: "XSS_CALLBACK",
	4: "XSS_ERROR",
	6: "CLEAR_TEXT_PASSWORD",
}

var Finding_FindingType_value = map[string]int32{
	"FINDING_TYPE_UNSPECIFIED": 0,
	"MIXED_CONTENT":            1,
	"OUTDATED_LIBRARY":         2,
	"ROSETTA_FLASH":            5,
	"XSS_CALLBACK":             3,
	"XSS_ERROR":                4,
	"CLEAR_TEXT_PASSWORD":      6,
}

func (x Finding_FindingType) String() string {
	return proto.EnumName(Finding_FindingType_name, int32(x))
}

func (Finding_FindingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_811918cc7b338327, []int{0, 0}
}

// A Finding resource represents a vulnerability instance identified during a
// ScanRun.
type Finding struct {
	// Output only.
	// The resource name of the Finding. The name follows the format of
	// 'projects/{projectId}/scanConfigs/{scanConfigId}/scanruns/{scanRunId}/findings/{findingId}'.
	// The finding IDs are generated by the system.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only.
	// The type of the Finding.
	FindingType Finding_FindingType `protobuf:"varint,2,opt,name=finding_type,json=findingType,proto3,enum=google.cloud.websecurityscanner.v1alpha.Finding_FindingType" json:"finding_type,omitempty"`
	// Output only.
	// The http method of the request that triggered the vulnerability, in
	// uppercase.
	HttpMethod string `protobuf:"bytes,3,opt,name=http_method,json=httpMethod,proto3" json:"http_method,omitempty"`
	// Output only.
	// The URL produced by the server-side fuzzer and used in the request that
	// triggered the vulnerability.
	FuzzedUrl string `protobuf:"bytes,4,opt,name=fuzzed_url,json=fuzzedUrl,proto3" json:"fuzzed_url,omitempty"`
	// Output only.
	// The body of the request that triggered the vulnerability.
	Body string `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	// Output only.
	// The description of the vulnerability.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// Output only.
	// The URL containing human-readable payload that user can leverage to
	// reproduce the vulnerability.
	ReproductionUrl string `protobuf:"bytes,7,opt,name=reproduction_url,json=reproductionUrl,proto3" json:"reproduction_url,omitempty"`
	// Output only.
	// If the vulnerability was originated from nested IFrame, the immediate
	// parent IFrame is reported.
	FrameUrl string `protobuf:"bytes,8,opt,name=frame_url,json=frameUrl,proto3" json:"frame_url,omitempty"`
	// Output only.
	// The URL where the browser lands when the vulnerability is detected.
	FinalUrl string `protobuf:"bytes,9,opt,name=final_url,json=finalUrl,proto3" json:"final_url,omitempty"`
	// Output only.
	// The tracking ID uniquely identifies a vulnerability instance across
	// multiple ScanRuns.
	TrackingId string `protobuf:"bytes,10,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"`
	// Output only.
	// An addon containing information about outdated libraries.
	OutdatedLibrary *OutdatedLibrary `protobuf:"bytes,11,opt,name=outdated_library,json=outdatedLibrary,proto3" json:"outdated_library,omitempty"`
	// Output only.
	// An addon containing detailed information regarding any resource causing the
	// vulnerability such as JavaScript sources, image, audio files, etc.
	ViolatingResource *ViolatingResource `protobuf:"bytes,12,opt,name=violating_resource,json=violatingResource,proto3" json:"violating_resource,omitempty"`
	// Output only.
	// An addon containing information about request parameters which were found
	// to be vulnerable.
	VulnerableParameters *VulnerableParameters `protobuf:"bytes,13,opt,name=vulnerable_parameters,json=vulnerableParameters,proto3" json:"vulnerable_parameters,omitempty"`
	// Output only.
	// An addon containing information reported for an XSS, if any.
	Xss                  *Xss     `protobuf:"bytes,14,opt,name=xss,proto3" json:"xss,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Finding) Reset()         { *m = Finding{} }
func (m *Finding) String() string { return proto.CompactTextString(m) }
func (*Finding) ProtoMessage()    {}
func (*Finding) Descriptor() ([]byte, []int) {
	return fileDescriptor_811918cc7b338327, []int{0}
}
func (m *Finding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Finding.Unmarshal(m, b)
}
func (m *Finding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Finding.Marshal(b, m, deterministic)
}
func (m *Finding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Finding.Merge(m, src)
}
func (m *Finding) XXX_Size() int {
	return xxx_messageInfo_Finding.Size(m)
}
func (m *Finding) XXX_DiscardUnknown() {
	xxx_messageInfo_Finding.DiscardUnknown(m)
}

var xxx_messageInfo_Finding proto.InternalMessageInfo

func (m *Finding) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Finding) GetFindingType() Finding_FindingType {
	if m != nil {
		return m.FindingType
	}
	return Finding_FINDING_TYPE_UNSPECIFIED
}

func (m *Finding) GetHttpMethod() string {
	if m != nil {
		return m.HttpMethod
	}
	return ""
}

func (m *Finding) GetFuzzedUrl() string {
	if m != nil {
		return m.FuzzedUrl
	}
	return ""
}

func (m *Finding) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Finding) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Finding) GetReproductionUrl() string {
	if m != nil {
		return m.ReproductionUrl
	}
	return ""
}

func (m *Finding) GetFrameUrl() string {
	if m != nil {
		return m.FrameUrl
	}
	return ""
}

func (m *Finding) GetFinalUrl() string {
	if m != nil {
		return m.FinalUrl
	}
	return ""
}

func (m *Finding) GetTrackingId() string {
	if m != nil {
		return m.TrackingId
	}
	return ""
}

func (m *Finding) GetOutdatedLibrary() *OutdatedLibrary {
	if m != nil {
		return m.OutdatedLibrary
	}
	return nil
}

func (m *Finding) GetViolatingResource() *ViolatingResource {
	if m != nil {
		return m.ViolatingResource
	}
	return nil
}

func (m *Finding) GetVulnerableParameters() *VulnerableParameters {
	if m != nil {
		return m.VulnerableParameters
	}
	return nil
}

func (m *Finding) GetXss() *Xss {
	if m != nil {
		return m.Xss
	}
	return nil
}

func init() {
	proto.RegisterType((*Finding)(nil), "google.cloud.websecurityscanner.v1alpha.Finding")
	proto.RegisterEnum("google.cloud.websecurityscanner.v1alpha.Finding_FindingType", Finding_FindingType_name, Finding_FindingType_value)
}

func init() {
	proto.RegisterFile("google/cloud/websecurityscanner/v1alpha/finding.proto", fileDescriptor_811918cc7b338327)
}

var fileDescriptor_811918cc7b338327 = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0xfd, 0x65, 0xff, 0xeb, 0x74, 0x5b, 0xe6, 0xdf, 0x10, 0xd1, 0x18, 0x5a, 0xb5, 0x17, 0x86,
	0x40, 0x89, 0x18, 0x42, 0x42, 0x0c, 0x90, 0xd2, 0x36, 0x85, 0x88, 0xae, 0xad, 0x9c, 0x0c, 0x3a,
	0x5e, 0x2c, 0x37, 0x71, 0xb3, 0x88, 0x34, 0x8e, 0x9c, 0xa4, 0xd0, 0x7d, 0x12, 0x5e, 0xe1, 0x93,
	0x22, 0x3b, 0xe9, 0x28, 0x0c, 0x89, 0xf2, 0x54, 0xfb, 0x9c, 0x7b, 0xce, 0xb9, 0xf7, 0xaa, 0x31,
	0x78, 0x16, 0x32, 0x16, 0xc6, 0xd4, 0xf4, 0x63, 0x56, 0x04, 0xe6, 0x67, 0x3a, 0xca, 0xa8, 0x5f,
	0xf0, 0x28, 0x9f, 0x65, 0x3e, 0x49, 0x12, 0xca, 0xcd, 0xe9, 0x13, 0x12, 0xa7, 0x57, 0xc4, 0x1c,
	0x47, 0x49, 0x10, 0x25, 0xa1, 0x91, 0x72, 0x96, 0x33, 0xf8, 0xa0, 0x94, 0x19, 0x52, 0x66, 0xdc,
	0x96, 0x19, 0x95, 0xec, 0xe0, 0xb0, 0xf2, 0x27, 0x69, 0x64, 0x92, 0x24, 0x61, 0x39, 0xc9, 0x23,
	0x96, 0x64, 0xa5, 0xcd, 0xc1, 0xd9, 0x3f, 0xa6, 0x63, 0x12, 0x04, 0x2c, 0x29, 0xc5, 0xc7, 0xdf,
	0x37, 0xc1, 0x66, 0xa7, 0xc4, 0x21, 0x04, 0x6b, 0x09, 0x99, 0x50, 0x5d, 0x69, 0x28, 0x27, 0x35,
	0x24, 0xcf, 0x10, 0x83, 0xfa, 0x5c, 0x96, 0xcf, 0x52, 0xaa, 0xaf, 0x34, 0x94, 0x93, 0x9d, 0xd3,
	0x97, 0xc6, 0x92, 0xad, 0x1b, 0x95, 0xf7, 0xfc, 0xd7, 0x9b, 0xa5, 0x14, 0xa9, 0xe3, 0x9f, 0x17,
	0x78, 0x04, 0xd4, 0xab, 0x3c, 0x4f, 0xf1, 0x84, 0xe6, 0x57, 0x2c, 0xd0, 0x57, 0x65, 0x36, 0x10,
	0xd0, 0xb9, 0x44, 0xe0, 0x7d, 0x00, 0xc6, 0xc5, 0xf5, 0x35, 0x0d, 0x70, 0xc1, 0x63, 0x7d, 0x4d,
	0xf2, 0xb5, 0x12, 0xb9, 0xe0, 0xb1, 0x68, 0x7a, 0xc4, 0x82, 0x99, 0xbe, 0x5e, 0x36, 0x2d, 0xce,
	0xb0, 0x01, 0xd4, 0x80, 0x66, 0x3e, 0x8f, 0x52, 0xb1, 0x27, 0x7d, 0x43, 0x52, 0x8b, 0x10, 0x7c,
	0x08, 0x34, 0x4e, 0x53, 0xce, 0x82, 0xc2, 0x17, 0x77, 0x69, 0xbd, 0x29, 0xcb, 0x76, 0x17, 0x71,
	0x11, 0x70, 0x0f, 0xd4, 0xc6, 0x9c, 0x4c, 0xa8, 0xac, 0xd9, 0x92, 0x35, 0x5b, 0x12, 0x98, 0x93,
	0x51, 0x42, 0x62, 0x49, 0xd6, 0x2a, 0x52, 0x00, 0x82, 0x3c, 0x02, 0x6a, 0xce, 0x89, 0xff, 0x49,
	0x2c, 0x2f, 0x0a, 0x74, 0x50, 0x8e, 0x36, 0x87, 0x9c, 0x00, 0xfa, 0x40, 0x63, 0x45, 0x1e, 0x90,
	0x9c, 0x06, 0x38, 0x8e, 0x46, 0x9c, 0xf0, 0x99, 0xae, 0x36, 0x94, 0x13, 0xf5, 0xf4, 0xf9, 0xd2,
	0x0b, 0xee, 0x57, 0x06, 0xdd, 0x52, 0x8f, 0x76, 0xd9, 0xaf, 0x00, 0x8c, 0x00, 0x9c, 0x46, 0x2c,
	0x26, 0xb9, 0x68, 0x83, 0xd3, 0x8c, 0x15, 0xdc, 0xa7, 0x7a, 0x5d, 0xc6, 0xbc, 0x58, 0x3a, 0xe6,
	0xfd, 0xdc, 0x02, 0x55, 0x0e, 0x68, 0x6f, 0xfa, 0x3b, 0x04, 0x39, 0xb8, 0x33, 0x2d, 0xe2, 0x84,
	0x72, 0x32, 0x8a, 0x29, 0x4e, 0x89, 0xd8, 0x52, 0x4e, 0x79, 0xa6, 0x6f, 0xcb, 0xb4, 0x57, 0xcb,
	0xa7, 0xdd, 0xb8, 0x0c, 0x6e, 0x4c, 0xd0, 0xfe, 0xf4, 0x0f, 0x28, 0x7c, 0x0d, 0x56, 0xbf, 0x64,
	0x99, 0xbe, 0x23, 0x13, 0x1e, 0x2f, 0x9d, 0x30, 0xcc, 0x32, 0x24, 0x84, 0xc7, 0xdf, 0x14, 0xa0,
	0x2e, 0xfc, 0x39, 0xe1, 0x21, 0xd0, 0x3b, 0x4e, 0xaf, 0xed, 0xf4, 0xde, 0x60, 0xef, 0x72, 0x60,
	0xe3, 0x8b, 0x9e, 0x3b, 0xb0, 0x5b, 0x4e, 0xc7, 0xb1, 0xdb, 0xda, 0x7f, 0x70, 0x0f, 0x6c, 0x9f,
	0x3b, 0x43, 0xbb, 0x8d, 0x5b, 0xfd, 0x9e, 0x67, 0xf7, 0x3c, 0x4d, 0x81, 0xfb, 0x40, 0xeb, 0x5f,
	0x78, 0x6d, 0xcb, 0xb3, 0xdb, 0xb8, 0xeb, 0x34, 0x91, 0x85, 0x2e, 0xb5, 0x15, 0x51, 0x88, 0xfa,
	0xae, 0xed, 0x79, 0x16, 0xee, 0x74, 0x2d, 0xf7, 0xad, 0xb6, 0x0e, 0x35, 0x50, 0x1f, 0xba, 0x2e,
	0x6e, 0x59, 0xdd, 0x6e, 0xd3, 0x6a, 0xbd, 0xd3, 0x56, 0xe1, 0x36, 0xa8, 0x09, 0xc4, 0x46, 0xa8,
	0x8f, 0xb4, 0x35, 0x78, 0x17, 0xfc, 0xdf, 0xea, 0xda, 0x16, 0xc2, 0x9e, 0x3d, 0xf4, 0xf0, 0xc0,
	0x72, 0xdd, 0x0f, 0x7d, 0xd4, 0xd6, 0x36, 0x9a, 0x5f, 0x15, 0xf0, 0xc8, 0x67, 0x93, 0x65, 0x87,
	0x6b, 0xd6, 0xab, 0x81, 0x06, 0xe2, 0x13, 0x1f, 0x28, 0x1f, 0x2f, 0x2b, 0x61, 0xc8, 0x62, 0x92,
	0x84, 0x06, 0xe3, 0xa1, 0x19, 0xd2, 0x44, 0x3e, 0x00, 0x66, 0x49, 0x91, 0x34, 0xca, 0xfe, 0xfa,
	0x80, 0x9c, 0xdd, 0xa6, 0x46, 0x1b, 0xd2, 0xe5, 0xe9, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a,
	0x7d, 0x4d, 0xc7, 0x03, 0x05, 0x00, 0x00,
}
