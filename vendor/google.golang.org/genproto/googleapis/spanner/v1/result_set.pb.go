// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/spanner/v1/result_set.proto

package spanner // import "google.golang.org/genproto/googleapis/spanner/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"
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

// Results from [Read][google.spanner.v1.Spanner.Read] or
// [ExecuteSql][google.spanner.v1.Spanner.ExecuteSql].
type ResultSet struct {
	// Metadata about the result set, such as row type information.
	Metadata *ResultSetMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Each element in `rows` is a row whose format is defined by
	// [metadata.row_type][google.spanner.v1.ResultSetMetadata.row_type]. The ith element
	// in each row matches the ith field in
	// [metadata.row_type][google.spanner.v1.ResultSetMetadata.row_type]. Elements are
	// encoded based on type as described
	// [here][google.spanner.v1.TypeCode].
	Rows []*_struct.ListValue `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	// Query plan and execution statistics for the query that produced this
	// result set. These can be requested by setting
	// [ExecuteSqlRequest.query_mode][google.spanner.v1.ExecuteSqlRequest.query_mode].
	Stats                *ResultSetStats `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ResultSet) Reset()         { *m = ResultSet{} }
func (m *ResultSet) String() string { return proto.CompactTextString(m) }
func (*ResultSet) ProtoMessage()    {}
func (*ResultSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8c1400ddbf24213, []int{0}
}
func (m *ResultSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultSet.Unmarshal(m, b)
}
func (m *ResultSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultSet.Marshal(b, m, deterministic)
}
func (m *ResultSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultSet.Merge(m, src)
}
func (m *ResultSet) XXX_Size() int {
	return xxx_messageInfo_ResultSet.Size(m)
}
func (m *ResultSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultSet.DiscardUnknown(m)
}

var xxx_messageInfo_ResultSet proto.InternalMessageInfo

func (m *ResultSet) GetMetadata() *ResultSetMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ResultSet) GetRows() []*_struct.ListValue {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *ResultSet) GetStats() *ResultSetStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// Partial results from a streaming read or SQL query. Streaming reads and
// SQL queries better tolerate large result sets, large rows, and large
// values, but are a little trickier to consume.
type PartialResultSet struct {
	// Metadata about the result set, such as row type information.
	// Only present in the first response.
	Metadata *ResultSetMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// A streamed result set consists of a stream of values, which might
	// be split into many `PartialResultSet` messages to accommodate
	// large rows and/or large values. Every N complete values defines a
	// row, where N is equal to the number of entries in
	// [metadata.row_type.fields][google.spanner.v1.StructType.fields].
	//
	// Most values are encoded based on type as described
	// [here][google.spanner.v1.TypeCode].
	//
	// It is possible that the last value in values is "chunked",
	// meaning that the rest of the value is sent in subsequent
	// `PartialResultSet`(s). This is denoted by the [chunked_value][google.spanner.v1.PartialResultSet.chunked_value]
	// field. Two or more chunked values can be merged to form a
	// complete value as follows:
	//
	//   * `bool/number/null`: cannot be chunked
	//   * `string`: concatenate the strings
	//   * `list`: concatenate the lists. If the last element in a list is a
	//     `string`, `list`, or `object`, merge it with the first element in
	//     the next list by applying these rules recursively.
	//   * `object`: concatenate the (field name, field value) pairs. If a
	//     field name is duplicated, then apply these rules recursively
	//     to merge the field values.
	//
	// Some examples of merging:
	//
	//     # Strings are concatenated.
	//     "foo", "bar" => "foobar"
	//
	//     # Lists of non-strings are concatenated.
	//     [2, 3], [4] => [2, 3, 4]
	//
	//     # Lists are concatenated, but the last and first elements are merged
	//     # because they are strings.
	//     ["a", "b"], ["c", "d"] => ["a", "bc", "d"]
	//
	//     # Lists are concatenated, but the last and first elements are merged
	//     # because they are lists. Recursively, the last and first elements
	//     # of the inner lists are merged because they are strings.
	//     ["a", ["b", "c"]], [["d"], "e"] => ["a", ["b", "cd"], "e"]
	//
	//     # Non-overlapping object fields are combined.
	//     {"a": "1"}, {"b": "2"} => {"a": "1", "b": 2"}
	//
	//     # Overlapping object fields are merged.
	//     {"a": "1"}, {"a": "2"} => {"a": "12"}
	//
	//     # Examples of merging objects containing lists of strings.
	//     {"a": ["1"]}, {"a": ["2"]} => {"a": ["12"]}
	//
	// For a more complete example, suppose a streaming SQL query is
	// yielding a result set whose rows contain a single string
	// field. The following `PartialResultSet`s might be yielded:
	//
	//     {
	//       "metadata": { ... }
	//       "values": ["Hello", "W"]
	//       "chunked_value": true
	//       "resume_token": "Af65..."
	//     }
	//     {
	//       "values": ["orl"]
	//       "chunked_value": true
	//       "resume_token": "Bqp2..."
	//     }
	//     {
	//       "values": ["d"]
	//       "resume_token": "Zx1B..."
	//     }
	//
	// This sequence of `PartialResultSet`s encodes two rows, one
	// containing the field value `"Hello"`, and a second containing the
	// field value `"World" = "W" + "orl" + "d"`.
	Values []*_struct.Value `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	// If true, then the final value in [values][google.spanner.v1.PartialResultSet.values] is chunked, and must
	// be combined with more values from subsequent `PartialResultSet`s
	// to obtain a complete field value.
	ChunkedValue bool `protobuf:"varint,3,opt,name=chunked_value,json=chunkedValue,proto3" json:"chunked_value,omitempty"`
	// Streaming calls might be interrupted for a variety of reasons, such
	// as TCP connection loss. If this occurs, the stream of results can
	// be resumed by re-sending the original request and including
	// `resume_token`. Note that executing any other transaction in the
	// same session invalidates the token.
	ResumeToken []byte `protobuf:"bytes,4,opt,name=resume_token,json=resumeToken,proto3" json:"resume_token,omitempty"`
	// Query plan and execution statistics for the query that produced this
	// streaming result set. These can be requested by setting
	// [ExecuteSqlRequest.query_mode][google.spanner.v1.ExecuteSqlRequest.query_mode] and are sent
	// only once with the last response in the stream.
	Stats                *ResultSetStats `protobuf:"bytes,5,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PartialResultSet) Reset()         { *m = PartialResultSet{} }
func (m *PartialResultSet) String() string { return proto.CompactTextString(m) }
func (*PartialResultSet) ProtoMessage()    {}
func (*PartialResultSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8c1400ddbf24213, []int{1}
}
func (m *PartialResultSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartialResultSet.Unmarshal(m, b)
}
func (m *PartialResultSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartialResultSet.Marshal(b, m, deterministic)
}
func (m *PartialResultSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartialResultSet.Merge(m, src)
}
func (m *PartialResultSet) XXX_Size() int {
	return xxx_messageInfo_PartialResultSet.Size(m)
}
func (m *PartialResultSet) XXX_DiscardUnknown() {
	xxx_messageInfo_PartialResultSet.DiscardUnknown(m)
}

var xxx_messageInfo_PartialResultSet proto.InternalMessageInfo

func (m *PartialResultSet) GetMetadata() *ResultSetMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PartialResultSet) GetValues() []*_struct.Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *PartialResultSet) GetChunkedValue() bool {
	if m != nil {
		return m.ChunkedValue
	}
	return false
}

func (m *PartialResultSet) GetResumeToken() []byte {
	if m != nil {
		return m.ResumeToken
	}
	return nil
}

func (m *PartialResultSet) GetStats() *ResultSetStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// Metadata about a [ResultSet][google.spanner.v1.ResultSet] or [PartialResultSet][google.spanner.v1.PartialResultSet].
type ResultSetMetadata struct {
	// Indicates the field names and types for the rows in the result
	// set.  For example, a SQL query like `"SELECT UserId, UserName FROM
	// Users"` could return a `row_type` value like:
	//
	//     "fields": [
	//       { "name": "UserId", "type": { "code": "INT64" } },
	//       { "name": "UserName", "type": { "code": "STRING" } },
	//     ]
	RowType *StructType `protobuf:"bytes,1,opt,name=row_type,json=rowType,proto3" json:"row_type,omitempty"`
	// If the read or SQL query began a transaction as a side-effect, the
	// information about the new transaction is yielded here.
	Transaction          *Transaction `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ResultSetMetadata) Reset()         { *m = ResultSetMetadata{} }
func (m *ResultSetMetadata) String() string { return proto.CompactTextString(m) }
func (*ResultSetMetadata) ProtoMessage()    {}
func (*ResultSetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8c1400ddbf24213, []int{2}
}
func (m *ResultSetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultSetMetadata.Unmarshal(m, b)
}
func (m *ResultSetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultSetMetadata.Marshal(b, m, deterministic)
}
func (m *ResultSetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultSetMetadata.Merge(m, src)
}
func (m *ResultSetMetadata) XXX_Size() int {
	return xxx_messageInfo_ResultSetMetadata.Size(m)
}
func (m *ResultSetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultSetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ResultSetMetadata proto.InternalMessageInfo

func (m *ResultSetMetadata) GetRowType() *StructType {
	if m != nil {
		return m.RowType
	}
	return nil
}

func (m *ResultSetMetadata) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

// Additional statistics about a [ResultSet][google.spanner.v1.ResultSet] or [PartialResultSet][google.spanner.v1.PartialResultSet].
type ResultSetStats struct {
	// [QueryPlan][google.spanner.v1.QueryPlan] for the query associated with this result.
	QueryPlan *QueryPlan `protobuf:"bytes,1,opt,name=query_plan,json=queryPlan,proto3" json:"query_plan,omitempty"`
	// Aggregated statistics from the execution of the query. Only present when
	// the query is profiled. For example, a query could return the statistics as
	// follows:
	//
	//     {
	//       "rows_returned": "3",
	//       "elapsed_time": "1.22 secs",
	//       "cpu_time": "1.19 secs"
	//     }
	QueryStats           *_struct.Struct `protobuf:"bytes,2,opt,name=query_stats,json=queryStats,proto3" json:"query_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ResultSetStats) Reset()         { *m = ResultSetStats{} }
func (m *ResultSetStats) String() string { return proto.CompactTextString(m) }
func (*ResultSetStats) ProtoMessage()    {}
func (*ResultSetStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8c1400ddbf24213, []int{3}
}
func (m *ResultSetStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultSetStats.Unmarshal(m, b)
}
func (m *ResultSetStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultSetStats.Marshal(b, m, deterministic)
}
func (m *ResultSetStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultSetStats.Merge(m, src)
}
func (m *ResultSetStats) XXX_Size() int {
	return xxx_messageInfo_ResultSetStats.Size(m)
}
func (m *ResultSetStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultSetStats.DiscardUnknown(m)
}

var xxx_messageInfo_ResultSetStats proto.InternalMessageInfo

func (m *ResultSetStats) GetQueryPlan() *QueryPlan {
	if m != nil {
		return m.QueryPlan
	}
	return nil
}

func (m *ResultSetStats) GetQueryStats() *_struct.Struct {
	if m != nil {
		return m.QueryStats
	}
	return nil
}

func init() {
	proto.RegisterType((*ResultSet)(nil), "google.spanner.v1.ResultSet")
	proto.RegisterType((*PartialResultSet)(nil), "google.spanner.v1.PartialResultSet")
	proto.RegisterType((*ResultSetMetadata)(nil), "google.spanner.v1.ResultSetMetadata")
	proto.RegisterType((*ResultSetStats)(nil), "google.spanner.v1.ResultSetStats")
}

func init() { proto.RegisterFile("google/spanner/v1/result_set.proto", fileDescriptor_e8c1400ddbf24213) }

var fileDescriptor_e8c1400ddbf24213 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcd, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0xe5, 0xf4, 0x87, 0xd4, 0x13, 0x10, 0xb5, 0x04, 0x1d, 0x45, 0x05, 0xa5, 0x29, 0x8b,
	0xac, 0x3c, 0x4a, 0x59, 0x10, 0xa9, 0x9b, 0xaa, 0x2c, 0xd8, 0x80, 0x14, 0x9c, 0x28, 0x0b, 0x14,
	0x69, 0xe4, 0x26, 0x66, 0x18, 0x75, 0x62, 0x4f, 0x6d, 0x4f, 0xa2, 0x3c, 0x00, 0x62, 0xc9, 0x9e,
	0x47, 0xe0, 0x01, 0x78, 0x08, 0x9e, 0x88, 0x25, 0xf2, 0xcf, 0x24, 0x81, 0x19, 0x21, 0x21, 0x75,
	0xe7, 0xf8, 0x7e, 0xe7, 0x9e, 0x7b, 0x3c, 0x37, 0xb0, 0x9b, 0x08, 0x91, 0x64, 0x2c, 0x52, 0x39,
	0xe5, 0x9c, 0xc9, 0x68, 0xd9, 0x8f, 0x24, 0x53, 0x45, 0xa6, 0x63, 0xc5, 0x34, 0xce, 0xa5, 0xd0,
	0x02, 0x1d, 0x3b, 0x06, 0x7b, 0x06, 0x2f, 0xfb, 0xed, 0x53, 0x2f, 0xa3, 0x79, 0x1a, 0x51, 0xce,
	0x85, 0xa6, 0x3a, 0x15, 0x5c, 0x39, 0xc1, 0xa6, 0x6a, 0x7f, 0xdd, 0x14, 0x1f, 0x23, 0xa5, 0x65,
	0x31, 0xf3, 0xed, 0xda, 0x35, 0x96, 0x77, 0x05, 0x93, 0xeb, 0x38, 0xcf, 0x28, 0xf7, 0xcc, 0x79,
	0x95, 0xd1, 0x92, 0x72, 0x45, 0x67, 0xc6, 0xe7, 0x2f, 0x9b, 0x5d, 0x68, 0x9d, 0x33, 0x57, 0xed,
	0xfe, 0x00, 0xf0, 0x88, 0xd8, 0x28, 0x23, 0xa6, 0xd1, 0x15, 0x6c, 0x2e, 0x98, 0xa6, 0x73, 0xaa,
	0x69, 0x08, 0x3a, 0xa0, 0x17, 0x5c, 0xbc, 0xc0, 0x95, 0x58, 0x78, 0xc3, 0xbf, 0xf3, 0x2c, 0xd9,
	0xa8, 0x10, 0x86, 0xfb, 0x52, 0xac, 0x54, 0xd8, 0xe8, 0xec, 0xf5, 0x82, 0x8b, 0x76, 0xa9, 0x2e,
	0x33, 0xe2, 0xb7, 0xa9, 0xd2, 0x13, 0x9a, 0x15, 0x8c, 0x58, 0x0e, 0xbd, 0x82, 0x07, 0x4a, 0x53,
	0xad, 0xc2, 0x3d, 0x6b, 0x77, 0xf6, 0x2f, 0xbb, 0x91, 0x01, 0x89, 0xe3, 0xbb, 0x9f, 0x1b, 0xf0,
	0xf1, 0x90, 0x4a, 0x9d, 0xd2, 0xec, 0x7e, 0xe7, 0x3f, 0x5c, 0x9a, 0xf1, 0xca, 0x04, 0x4f, 0x2b,
	0x09, 0xdc, 0xf4, 0x9e, 0x42, 0xe7, 0xf0, 0xe1, 0xec, 0x53, 0xc1, 0x6f, 0xd9, 0x3c, 0xb6, 0x37,
	0x36, 0x47, 0x93, 0xb4, 0xfc, 0xa5, 0x85, 0xd1, 0x19, 0x6c, 0x99, 0x75, 0x59, 0xb0, 0x58, 0x8b,
	0x5b, 0xc6, 0xc3, 0xfd, 0x0e, 0xe8, 0xb5, 0x48, 0xe0, 0xee, 0xc6, 0xe6, 0x6a, 0xfb, 0x0e, 0x07,
	0xff, 0xf9, 0x0e, 0x5f, 0x01, 0x3c, 0xae, 0x04, 0x42, 0x03, 0xd8, 0x94, 0x62, 0x15, 0x9b, 0x0f,
	0xed, 0x1f, 0xe2, 0x59, 0x4d, 0xc7, 0x91, 0x5d, 0xb8, 0xf1, 0x3a, 0x67, 0xe4, 0x81, 0x14, 0x2b,
	0x73, 0x40, 0x57, 0x30, 0xd8, 0xd9, 0xa1, 0xb0, 0x61, 0xc5, 0xcf, 0x6b, 0xc4, 0xe3, 0x2d, 0x45,
	0x76, 0x25, 0xdd, 0x2f, 0x00, 0x3e, 0xfa, 0x73, 0x56, 0x74, 0x09, 0xe1, 0x76, 0x79, 0xfd, 0x40,
	0xa7, 0x35, 0x3d, 0xdf, 0x1b, 0x68, 0x98, 0x51, 0x4e, 0x8e, 0xee, 0xca, 0x23, 0x1a, 0xc0, 0xc0,
	0x89, 0xdd, 0x03, 0xb9, 0x89, 0x4e, 0x2a, 0xdf, 0xc5, 0x85, 0x21, 0xce, 0xc8, 0xda, 0x5e, 0x7f,
	0x03, 0xf0, 0xc9, 0x4c, 0x2c, 0xaa, 0x46, 0xd7, 0xdb, 0x01, 0x87, 0x46, 0x3f, 0x04, 0x1f, 0x06,
	0x1e, 0x4a, 0x44, 0x46, 0x79, 0x82, 0x85, 0x4c, 0xa2, 0x84, 0x71, 0xdb, 0x3d, 0x72, 0x25, 0x9a,
	0xa7, 0x6a, 0xe7, 0x5f, 0x74, 0xe9, 0x8f, 0xbf, 0x00, 0xf8, 0xde, 0x38, 0x79, 0xe3, 0xd4, 0xaf,
	0x33, 0x51, 0xcc, 0xf1, 0xc8, 0x1b, 0x4d, 0xfa, 0x3f, 0xcb, 0xca, 0xd4, 0x56, 0xa6, 0xbe, 0x32,
	0x9d, 0xf4, 0x6f, 0x0e, 0x6d, 0xef, 0x97, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x06, 0x67, 0x6c,
	0xee, 0x5c, 0x04, 0x00, 0x00,
}
