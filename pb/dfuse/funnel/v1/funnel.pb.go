// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/funnel/v1/funnel.proto

package pbfunnel

import (
	context "context"
	fmt "fmt"
	eos "github.com/dfuse-io/dfuse-eosio/pb/dfuse/codecs/eos"
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

type StreamBlockRequest struct {
	FromBlockNum         int64    `protobuf:"varint,2,opt,name=fromBlockNum,proto3" json:"fromBlockNum,omitempty"`
	IrreversibleOnly     bool     `protobuf:"varint,3,opt,name=irreversibleOnly,proto3" json:"irreversibleOnly,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamBlockRequest) Reset()         { *m = StreamBlockRequest{} }
func (m *StreamBlockRequest) String() string { return proto.CompactTextString(m) }
func (*StreamBlockRequest) ProtoMessage()    {}
func (*StreamBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3f872ed66cead6b, []int{0}
}

func (m *StreamBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamBlockRequest.Unmarshal(m, b)
}
func (m *StreamBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamBlockRequest.Marshal(b, m, deterministic)
}
func (m *StreamBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamBlockRequest.Merge(m, src)
}
func (m *StreamBlockRequest) XXX_Size() int {
	return xxx_messageInfo_StreamBlockRequest.Size(m)
}
func (m *StreamBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamBlockRequest proto.InternalMessageInfo

func (m *StreamBlockRequest) GetFromBlockNum() int64 {
	if m != nil {
		return m.FromBlockNum
	}
	return 0
}

func (m *StreamBlockRequest) GetIrreversibleOnly() bool {
	if m != nil {
		return m.IrreversibleOnly
	}
	return false
}

type StreamBlockResponse struct {
	Undo                 bool       `protobuf:"varint,1,opt,name=undo,proto3" json:"undo,omitempty"`
	Block                *eos.Block `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StreamBlockResponse) Reset()         { *m = StreamBlockResponse{} }
func (m *StreamBlockResponse) String() string { return proto.CompactTextString(m) }
func (*StreamBlockResponse) ProtoMessage()    {}
func (*StreamBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3f872ed66cead6b, []int{1}
}

func (m *StreamBlockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamBlockResponse.Unmarshal(m, b)
}
func (m *StreamBlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamBlockResponse.Marshal(b, m, deterministic)
}
func (m *StreamBlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamBlockResponse.Merge(m, src)
}
func (m *StreamBlockResponse) XXX_Size() int {
	return xxx_messageInfo_StreamBlockResponse.Size(m)
}
func (m *StreamBlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamBlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamBlockResponse proto.InternalMessageInfo

func (m *StreamBlockResponse) GetUndo() bool {
	if m != nil {
		return m.Undo
	}
	return false
}

func (m *StreamBlockResponse) GetBlock() *eos.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamBlockRequest)(nil), "dfuse.funnel.v1.StreamBlockRequest")
	proto.RegisterType((*StreamBlockResponse)(nil), "dfuse.funnel.v1.StreamBlockResponse")
}

func init() { proto.RegisterFile("dfuse/funnel/v1/funnel.proto", fileDescriptor_d3f872ed66cead6b) }

var fileDescriptor_d3f872ed66cead6b = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xa6, 0x4e, 0x87, 0xc4, 0x81, 0xf2, 0x3c, 0x58, 0x8a, 0x87, 0x51, 0x3d, 0x0c, 0x61, 0x89,
	0x9b, 0xc7, 0xe1, 0x65, 0x07, 0x8f, 0x0a, 0xf5, 0x22, 0x7a, 0x32, 0xed, 0xab, 0x16, 0xdb, 0xbc,
	0x9a, 0x34, 0x05, 0xff, 0xbd, 0x2c, 0xaf, 0x07, 0xb7, 0x81, 0x87, 0xc0, 0xc7, 0x97, 0xef, 0xe5,
	0xfb, 0xbe, 0x17, 0x71, 0x59, 0x94, 0xde, 0xa1, 0x2a, 0xbd, 0x31, 0x58, 0xab, 0x7e, 0x31, 0x20,
	0xd9, 0x5a, 0xea, 0x08, 0x4e, 0xc3, 0xad, 0x1c, 0xb8, 0x7e, 0x91, 0x24, 0x2c, 0xcf, 0xa9, 0xc0,
	0xdc, 0x29, 0xa4, 0x70, 0x58, 0x9c, 0x16, 0x02, 0x9e, 0x3b, 0x8b, 0xef, 0xcd, 0xba, 0xa6, 0xfc,
	0x2b, 0xc3, 0x6f, 0x8f, 0xae, 0x83, 0x54, 0x4c, 0x4a, 0x4b, 0xcc, 0x3d, 0xfa, 0x26, 0x3e, 0x98,
	0x46, 0xb3, 0x51, 0xb6, 0xc5, 0xc1, 0x8d, 0x38, 0xab, 0xac, 0xc5, 0x1e, 0xad, 0xab, 0x74, 0x8d,
	0x4f, 0xa6, 0xfe, 0x89, 0x47, 0xd3, 0x68, 0x76, 0x9c, 0xed, 0xf1, 0xe9, 0x8b, 0x38, 0xdf, 0x72,
	0x71, 0x2d, 0x19, 0x87, 0x00, 0xe2, 0xd0, 0x9b, 0x82, 0xe2, 0x28, 0x8c, 0x05, 0x0c, 0x73, 0x71,
	0xa4, 0x37, 0xa2, 0xe0, 0x79, 0xb2, 0xbc, 0x90, 0xdc, 0x86, 0xc3, 0xcb, 0x4d, 0x70, 0x7e, 0x83,
	0x55, 0x4b, 0x14, 0xe3, 0x87, 0x50, 0x14, 0xde, 0xc4, 0xe4, 0x8f, 0x87, 0x83, 0x2b, 0xb9, 0xb3,
	0x07, 0xb9, 0x5f, 0x34, 0xb9, 0xfe, 0x5f, 0xc4, 0x39, 0x6f, 0xa3, 0xf5, 0xfd, 0xeb, 0xea, 0xa3,
	0xea, 0x3e, 0xbd, 0x96, 0x39, 0x35, 0x2a, 0xcc, 0xcc, 0x2b, 0x1a, 0x00, 0x92, 0xab, 0x48, 0xb5,
	0x5a, 0xed, 0x7c, 0xcb, 0xaa, 0xd5, 0x8c, 0xf5, 0x38, 0x2c, 0xfb, 0xee, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x68, 0x0f, 0x02, 0xbb, 0xb9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FunnelClient is the client API for Funnel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FunnelClient interface {
	StreamBlocks(ctx context.Context, in *StreamBlockRequest, opts ...grpc.CallOption) (Funnel_StreamBlocksClient, error)
}

type funnelClient struct {
	cc *grpc.ClientConn
}

func NewFunnelClient(cc *grpc.ClientConn) FunnelClient {
	return &funnelClient{cc}
}

func (c *funnelClient) StreamBlocks(ctx context.Context, in *StreamBlockRequest, opts ...grpc.CallOption) (Funnel_StreamBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Funnel_serviceDesc.Streams[0], "/dfuse.funnel.v1.Funnel/StreamBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &funnelStreamBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Funnel_StreamBlocksClient interface {
	Recv() (*StreamBlockResponse, error)
	grpc.ClientStream
}

type funnelStreamBlocksClient struct {
	grpc.ClientStream
}

func (x *funnelStreamBlocksClient) Recv() (*StreamBlockResponse, error) {
	m := new(StreamBlockResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FunnelServer is the server API for Funnel service.
type FunnelServer interface {
	StreamBlocks(*StreamBlockRequest, Funnel_StreamBlocksServer) error
}

// UnimplementedFunnelServer can be embedded to have forward compatible implementations.
type UnimplementedFunnelServer struct {
}

func (*UnimplementedFunnelServer) StreamBlocks(req *StreamBlockRequest, srv Funnel_StreamBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamBlocks not implemented")
}

func RegisterFunnelServer(s *grpc.Server, srv FunnelServer) {
	s.RegisterService(&_Funnel_serviceDesc, srv)
}

func _Funnel_StreamBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamBlockRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FunnelServer).StreamBlocks(m, &funnelStreamBlocksServer{stream})
}

type Funnel_StreamBlocksServer interface {
	Send(*StreamBlockResponse) error
	grpc.ServerStream
}

type funnelStreamBlocksServer struct {
	grpc.ServerStream
}

func (x *funnelStreamBlocksServer) Send(m *StreamBlockResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Funnel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.funnel.v1.Funnel",
	HandlerType: (*FunnelServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamBlocks",
			Handler:       _Funnel_StreamBlocks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dfuse/funnel/v1/funnel.proto",
}
