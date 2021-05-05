// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: heimdall/clerk/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryRecordParams is request type for the Query/Record RPC method
type QueryRecordParams struct {
	RecordId uint64 `protobuf:"varint,1,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
}

func (m *QueryRecordParams) Reset()         { *m = QueryRecordParams{} }
func (m *QueryRecordParams) String() string { return proto.CompactTextString(m) }
func (*QueryRecordParams) ProtoMessage()    {}
func (*QueryRecordParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aed47df6f39147d, []int{0}
}
func (m *QueryRecordParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRecordParams.Unmarshal(m, b)
}
func (m *QueryRecordParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRecordParams.Marshal(b, m, deterministic)
}
func (m *QueryRecordParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRecordParams.Merge(m, src)
}
func (m *QueryRecordParams) XXX_Size() int {
	return xxx_messageInfo_QueryRecordParams.Size(m)
}
func (m *QueryRecordParams) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRecordParams.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRecordParams proto.InternalMessageInfo

func (m *QueryRecordParams) GetRecordId() uint64 {
	if m != nil {
		return m.RecordId
	}
	return 0
}

// QueryRecordResponse is response type for the Query/Record RPC method
type QueryRecordResponse struct {
	EventRecord *EventRecord `protobuf:"bytes,1,opt,name=event_record,json=eventRecord,proto3" json:"event_record,omitempty"`
}

func (m *QueryRecordResponse) Reset()         { *m = QueryRecordResponse{} }
func (m *QueryRecordResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRecordResponse) ProtoMessage()    {}
func (*QueryRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aed47df6f39147d, []int{1}
}
func (m *QueryRecordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRecordResponse.Unmarshal(m, b)
}
func (m *QueryRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRecordResponse.Marshal(b, m, deterministic)
}
func (m *QueryRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRecordResponse.Merge(m, src)
}
func (m *QueryRecordResponse) XXX_Size() int {
	return xxx_messageInfo_QueryRecordResponse.Size(m)
}
func (m *QueryRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRecordResponse proto.InternalMessageInfo

func (m *QueryRecordResponse) GetEventRecord() *EventRecord {
	if m != nil {
		return m.EventRecord
	}
	return nil
}

// IsOldTx request and response messages
type QueryIsOldTxRequest struct {
	TxHash   string `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	LogIndex uint64 `protobuf:"varint,2,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
}

func (m *QueryIsOldTxRequest) Reset()         { *m = QueryIsOldTxRequest{} }
func (m *QueryIsOldTxRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIsOldTxRequest) ProtoMessage()    {}
func (*QueryIsOldTxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aed47df6f39147d, []int{2}
}
func (m *QueryIsOldTxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryIsOldTxRequest.Unmarshal(m, b)
}
func (m *QueryIsOldTxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryIsOldTxRequest.Marshal(b, m, deterministic)
}
func (m *QueryIsOldTxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIsOldTxRequest.Merge(m, src)
}
func (m *QueryIsOldTxRequest) XXX_Size() int {
	return xxx_messageInfo_QueryIsOldTxRequest.Size(m)
}
func (m *QueryIsOldTxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIsOldTxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIsOldTxRequest proto.InternalMessageInfo

func (m *QueryIsOldTxRequest) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *QueryIsOldTxRequest) GetLogIndex() uint64 {
	if m != nil {
		return m.LogIndex
	}
	return 0
}

type QueryIsOldTxResponse struct {
	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *QueryIsOldTxResponse) Reset()         { *m = QueryIsOldTxResponse{} }
func (m *QueryIsOldTxResponse) String() string { return proto.CompactTextString(m) }
func (*QueryIsOldTxResponse) ProtoMessage()    {}
func (*QueryIsOldTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aed47df6f39147d, []int{3}
}
func (m *QueryIsOldTxResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryIsOldTxResponse.Unmarshal(m, b)
}
func (m *QueryIsOldTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryIsOldTxResponse.Marshal(b, m, deterministic)
}
func (m *QueryIsOldTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIsOldTxResponse.Merge(m, src)
}
func (m *QueryIsOldTxResponse) XXX_Size() int {
	return xxx_messageInfo_QueryIsOldTxResponse.Size(m)
}
func (m *QueryIsOldTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIsOldTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIsOldTxResponse proto.InternalMessageInfo

func (m *QueryIsOldTxResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

// QueryRecordListRequest
type QueryRecordListRequest struct {
	Page     uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit    uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	FromId   uint64 `protobuf:"varint,3,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	FromTime uint64 `protobuf:"varint,4,opt,name=from_time,json=fromTime,proto3" json:"from_time,omitempty"`
	ToTime   uint64 `protobuf:"varint,5,opt,name=to_time,json=toTime,proto3" json:"to_time,omitempty"`
}

func (m *QueryRecordListRequest) Reset()         { *m = QueryRecordListRequest{} }
func (m *QueryRecordListRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRecordListRequest) ProtoMessage()    {}
func (*QueryRecordListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aed47df6f39147d, []int{4}
}
func (m *QueryRecordListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRecordListRequest.Unmarshal(m, b)
}
func (m *QueryRecordListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRecordListRequest.Marshal(b, m, deterministic)
}
func (m *QueryRecordListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRecordListRequest.Merge(m, src)
}
func (m *QueryRecordListRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRecordListRequest.Size(m)
}
func (m *QueryRecordListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRecordListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRecordListRequest proto.InternalMessageInfo

func (m *QueryRecordListRequest) GetPage() uint64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *QueryRecordListRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *QueryRecordListRequest) GetFromId() uint64 {
	if m != nil {
		return m.FromId
	}
	return 0
}

func (m *QueryRecordListRequest) GetFromTime() uint64 {
	if m != nil {
		return m.FromTime
	}
	return 0
}

func (m *QueryRecordListRequest) GetToTime() uint64 {
	if m != nil {
		return m.ToTime
	}
	return 0
}

type QueryRecordListResponse struct {
	EventRecords []*EventRecord `protobuf:"bytes,1,rep,name=event_records,json=eventRecords,proto3" json:"event_records,omitempty"`
}

func (m *QueryRecordListResponse) Reset()         { *m = QueryRecordListResponse{} }
func (m *QueryRecordListResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRecordListResponse) ProtoMessage()    {}
func (*QueryRecordListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aed47df6f39147d, []int{5}
}
func (m *QueryRecordListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRecordListResponse.Unmarshal(m, b)
}
func (m *QueryRecordListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRecordListResponse.Marshal(b, m, deterministic)
}
func (m *QueryRecordListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRecordListResponse.Merge(m, src)
}
func (m *QueryRecordListResponse) XXX_Size() int {
	return xxx_messageInfo_QueryRecordListResponse.Size(m)
}
func (m *QueryRecordListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRecordListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRecordListResponse proto.InternalMessageInfo

func (m *QueryRecordListResponse) GetEventRecords() []*EventRecord {
	if m != nil {
		return m.EventRecords
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRecordParams)(nil), "heimdall.clerk.v1beta1.QueryRecordParams")
	proto.RegisterType((*QueryRecordResponse)(nil), "heimdall.clerk.v1beta1.QueryRecordResponse")
	proto.RegisterType((*QueryIsOldTxRequest)(nil), "heimdall.clerk.v1beta1.QueryIsOldTxRequest")
	proto.RegisterType((*QueryIsOldTxResponse)(nil), "heimdall.clerk.v1beta1.QueryIsOldTxResponse")
	proto.RegisterType((*QueryRecordListRequest)(nil), "heimdall.clerk.v1beta1.QueryRecordListRequest")
	proto.RegisterType((*QueryRecordListResponse)(nil), "heimdall.clerk.v1beta1.QueryRecordListResponse")
}

func init() {
	proto.RegisterFile("heimdall/clerk/v1beta1/query.proto", fileDescriptor_7aed47df6f39147d)
}

var fileDescriptor_7aed47df6f39147d = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xb5, 0x9b, 0xc4, 0x6d, 0xa7, 0x65, 0xc1, 0x10, 0xa5, 0x91, 0x41, 0x06, 0xcc, 0x82, 0x57,
	0xe5, 0xa1, 0xe1, 0x0f, 0x40, 0xa0, 0x44, 0x20, 0x01, 0x56, 0x57, 0x48, 0x28, 0x9a, 0xd8, 0x83,
	0x3d, 0xaa, 0xed, 0x71, 0x3d, 0x93, 0xe2, 0x0a, 0xb1, 0xe1, 0x0b, 0x40, 0x45, 0xe2, 0x97, 0x58,
	0xa1, 0x4a, 0x6c, 0x58, 0xa2, 0x84, 0x0f, 0x41, 0xf3, 0x68, 0x9b, 0x8a, 0xa6, 0xf2, 0x6e, 0xee,
	0x9d, 0x73, 0xcf, 0x3d, 0xe7, 0x68, 0x6c, 0xe0, 0xa7, 0x84, 0xe6, 0x31, 0xce, 0x32, 0x14, 0x65,
	0xa4, 0xda, 0x43, 0x07, 0x3b, 0x13, 0x22, 0xf0, 0x0e, 0xda, 0x9f, 0x92, 0xea, 0x30, 0x28, 0x2b,
	0x26, 0x18, 0xec, 0x9d, 0x60, 0x02, 0x85, 0x09, 0x0c, 0xc6, 0xbd, 0x91, 0x30, 0x96, 0x64, 0x04,
	0xe1, 0x92, 0x22, 0x5c, 0x14, 0x4c, 0x60, 0x41, 0x59, 0xc1, 0xf5, 0x94, 0xbb, 0x8c, 0x59, 0x73,
	0x68, 0x4c, 0x37, 0x61, 0x09, 0x53, 0x47, 0x24, 0x4f, 0xba, 0xeb, 0x3f, 0x02, 0x57, 0xdf, 0xc8,
	0xf5, 0x21, 0x89, 0x58, 0x15, 0xbf, 0xc6, 0x15, 0xce, 0x39, 0xbc, 0x0e, 0xd6, 0x2b, 0x55, 0x8f,
	0x69, 0xdc, 0xb7, 0x6f, 0xd9, 0xf7, 0xda, 0xe1, 0x9a, 0x6e, 0x8c, 0x62, 0xff, 0x1d, 0xb8, 0xb6,
	0x30, 0x11, 0x12, 0x5e, 0xb2, 0x82, 0x13, 0xf8, 0x1c, 0x6c, 0x92, 0x03, 0x52, 0x88, 0xb1, 0x06,
	0xaa, 0xb1, 0x8d, 0xc1, 0x9d, 0xe0, 0x62, 0x3f, 0xc1, 0x33, 0x89, 0x35, 0x14, 0x1b, 0xe4, 0xac,
	0xf0, 0x5f, 0x18, 0xfa, 0x11, 0x7f, 0x95, 0xc5, 0xbb, 0x75, 0x48, 0xf6, 0xa7, 0x84, 0x0b, 0xb8,
	0x05, 0x56, 0x45, 0x3d, 0x4e, 0x31, 0x4f, 0x15, 0xf3, 0x7a, 0xe8, 0x88, 0x7a, 0x88, 0x79, 0x2a,
	0xb5, 0x66, 0x2c, 0x19, 0xd3, 0x22, 0x26, 0x75, 0x7f, 0x45, 0x6b, 0xcd, 0x58, 0x32, 0x92, 0xb5,
	0x1f, 0x80, 0xee, 0x79, 0x32, 0x23, 0xb6, 0x07, 0x1c, 0x2e, 0xb0, 0x98, 0x72, 0x45, 0xb6, 0x16,
	0x9a, 0xca, 0xff, 0x6a, 0x83, 0xde, 0x82, 0xb9, 0x97, 0x94, 0x8b, 0x13, 0x01, 0x10, 0xb4, 0x4b,
	0x9c, 0x10, 0x13, 0x87, 0x3a, 0xc3, 0x2e, 0xe8, 0x64, 0x34, 0xa7, 0xc2, 0xec, 0xd5, 0x85, 0x94,
	0xfa, 0xbe, 0x62, 0xb9, 0xcc, 0xae, 0xa5, 0xfa, 0x8e, 0x2c, 0x47, 0xb1, 0x94, 0xaa, 0x2e, 0x04,
	0xcd, 0x49, 0xbf, 0xad, 0xa5, 0xca, 0xc6, 0x2e, 0xcd, 0x89, 0x32, 0xc8, 0xf4, 0x55, 0x47, 0x4f,
	0x09, 0x26, 0x2f, 0xfc, 0x08, 0x6c, 0xfd, 0x27, 0xc9, 0xd8, 0x18, 0x82, 0x2b, 0x8b, 0x99, 0x4b,
	0x37, 0xad, 0xa6, 0xa1, 0x6f, 0x2e, 0x84, 0xce, 0x07, 0x3f, 0x5b, 0xa0, 0xa3, 0xb6, 0xc0, 0x23,
	0x1b, 0xac, 0x9a, 0x2e, 0x0c, 0x96, 0x11, 0x5d, 0x9c, 0x91, 0x8b, 0x1a, 0xe3, 0xb5, 0x01, 0xff,
	0xee, 0xe7, 0x5f, 0x7f, 0x8f, 0x56, 0x6e, 0xc3, 0x9b, 0x68, 0xc9, 0x03, 0x36, 0xc6, 0xe0, 0x37,
	0x1b, 0x38, 0x7a, 0x1e, 0xde, 0x6f, 0xb0, 0x44, 0xbf, 0x63, 0xf7, 0x61, 0x03, 0xe8, 0xa9, 0x96,
	0x81, 0xd2, 0xb2, 0x0d, 0x1f, 0x5c, 0xae, 0x05, 0x7d, 0x3c, 0xfd, 0x34, 0x3e, 0xc1, 0xef, 0xb6,
	0xf9, 0x7c, 0xcc, 0x03, 0x7b, 0x2a, 0x07, 0xe0, 0xe5, 0x6b, 0xcf, 0x3f, 0x6c, 0x77, 0xbb, 0x19,
	0xb8, 0x69, 0x60, 0x94, 0xb3, 0x2c, 0x16, 0xf5, 0x93, 0xe1, 0x8f, 0x99, 0x67, 0x1d, 0xcf, 0x3c,
	0xeb, 0xcf, 0xcc, 0xb3, 0xbf, 0xcc, 0x3d, 0xeb, 0x78, 0xee, 0x59, 0xbf, 0xe7, 0x9e, 0xf5, 0x36,
	0x48, 0xa8, 0x48, 0xa7, 0x93, 0x20, 0x62, 0x39, 0xca, 0xb1, 0xa0, 0x51, 0x41, 0xc4, 0x07, 0x56,
	0xed, 0x9d, 0x31, 0xd6, 0x86, 0x53, 0x1c, 0x96, 0x84, 0x4f, 0x1c, 0xf5, 0xa3, 0x78, 0xfc, 0x2f,
	0x00, 0x00, 0xff, 0xff, 0xa7, 0x7a, 0xdf, 0xbc, 0xbe, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Records
	Records(ctx context.Context, in *QueryRecordListRequest, opts ...grpc.CallOption) (*QueryRecordListResponse, error)
	// Record queries the record that match by record id.
	Record(ctx context.Context, in *QueryRecordParams, opts ...grpc.CallOption) (*QueryRecordResponse, error)
	QueryIsOldTxClerk(ctx context.Context, in *QueryIsOldTxRequest, opts ...grpc.CallOption) (*QueryIsOldTxResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Records(ctx context.Context, in *QueryRecordListRequest, opts ...grpc.CallOption) (*QueryRecordListResponse, error) {
	out := new(QueryRecordListResponse)
	err := c.cc.Invoke(ctx, "/heimdall.clerk.v1beta1.Query/Records", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Record(ctx context.Context, in *QueryRecordParams, opts ...grpc.CallOption) (*QueryRecordResponse, error) {
	out := new(QueryRecordResponse)
	err := c.cc.Invoke(ctx, "/heimdall.clerk.v1beta1.Query/Record", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryIsOldTxClerk(ctx context.Context, in *QueryIsOldTxRequest, opts ...grpc.CallOption) (*QueryIsOldTxResponse, error) {
	out := new(QueryIsOldTxResponse)
	err := c.cc.Invoke(ctx, "/heimdall.clerk.v1beta1.Query/QueryIsOldTxClerk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Records
	Records(context.Context, *QueryRecordListRequest) (*QueryRecordListResponse, error)
	// Record queries the record that match by record id.
	Record(context.Context, *QueryRecordParams) (*QueryRecordResponse, error)
	QueryIsOldTxClerk(context.Context, *QueryIsOldTxRequest) (*QueryIsOldTxResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Records(ctx context.Context, req *QueryRecordListRequest) (*QueryRecordListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Records not implemented")
}
func (*UnimplementedQueryServer) Record(ctx context.Context, req *QueryRecordParams) (*QueryRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Record not implemented")
}
func (*UnimplementedQueryServer) QueryIsOldTxClerk(ctx context.Context, req *QueryIsOldTxRequest) (*QueryIsOldTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryIsOldTxClerk not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Records_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRecordListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Records(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall.clerk.v1beta1.Query/Records",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Records(ctx, req.(*QueryRecordListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Record_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRecordParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Record(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall.clerk.v1beta1.Query/Record",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Record(ctx, req.(*QueryRecordParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryIsOldTxClerk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIsOldTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryIsOldTxClerk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall.clerk.v1beta1.Query/QueryIsOldTxClerk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryIsOldTxClerk(ctx, req.(*QueryIsOldTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heimdall.clerk.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Records",
			Handler:    _Query_Records_Handler,
		},
		{
			MethodName: "Record",
			Handler:    _Query_Record_Handler,
		},
		{
			MethodName: "QueryIsOldTxClerk",
			Handler:    _Query_QueryIsOldTxClerk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heimdall/clerk/v1beta1/query.proto",
}

func (m *QueryRecordParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RecordId != 0 {
		n += 1 + sovQuery(uint64(m.RecordId))
	}
	return n
}

func (m *QueryRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EventRecord != nil {
		l = m.EventRecord.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryIsOldTxRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.LogIndex != 0 {
		n += 1 + sovQuery(uint64(m.LogIndex))
	}
	return n
}

func (m *QueryIsOldTxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status {
		n += 2
	}
	return n
}

func (m *QueryRecordListRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Page != 0 {
		n += 1 + sovQuery(uint64(m.Page))
	}
	if m.Limit != 0 {
		n += 1 + sovQuery(uint64(m.Limit))
	}
	if m.FromId != 0 {
		n += 1 + sovQuery(uint64(m.FromId))
	}
	if m.FromTime != 0 {
		n += 1 + sovQuery(uint64(m.FromTime))
	}
	if m.ToTime != 0 {
		n += 1 + sovQuery(uint64(m.ToTime))
	}
	return n
}

func (m *QueryRecordListResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.EventRecords) > 0 {
		for _, e := range m.EventRecords {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
