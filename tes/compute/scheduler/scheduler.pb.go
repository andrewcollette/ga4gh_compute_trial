// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scheduler.proto

/*
Package scheduler is a generated protocol buffer package.

It is generated from these files:
	scheduler.proto

It has these top-level messages:
	Resources
	Node
	GetNodeRequest
	ListNodesRequest
	ListNodesResponse
	PutNodeResponse
	DeleteNodeResponse
*/
package scheduler

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NodeState int32

const (
	NodeState_UNINITIALIZED NodeState = 0
	NodeState_ALIVE         NodeState = 1
	NodeState_DEAD          NodeState = 2
	NodeState_GONE          NodeState = 3
	NodeState_INITIALIZING  NodeState = 4
	NodeState_DRAIN         NodeState = 5
)

var NodeState_name = map[int32]string{
	0: "UNINITIALIZED",
	1: "ALIVE",
	2: "DEAD",
	3: "GONE",
	4: "INITIALIZING",
	5: "DRAIN",
}
var NodeState_value = map[string]int32{
	"UNINITIALIZED": 0,
	"ALIVE":         1,
	"DEAD":          2,
	"GONE":          3,
	"INITIALIZING":  4,
	"DRAIN":         5,
}

func (x NodeState) String() string {
	return proto.EnumName(NodeState_name, int32(x))
}
func (NodeState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Resources struct {
	Cpus uint32 `protobuf:"varint,1,opt,name=cpus" json:"cpus,omitempty"`
	// In GB
	RamGb float64 `protobuf:"fixed64,2,opt,name=ram_gb,json=ramGb" json:"ram_gb,omitempty"`
	// In GB
	DiskGb float64 `protobuf:"fixed64,3,opt,name=disk_gb,json=diskGb" json:"disk_gb,omitempty"`
}

func (m *Resources) Reset()                    { *m = Resources{} }
func (m *Resources) String() string            { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()               {}
func (*Resources) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Resources) GetCpus() uint32 {
	if m != nil {
		return m.Cpus
	}
	return 0
}

func (m *Resources) GetRamGb() float64 {
	if m != nil {
		return m.RamGb
	}
	return 0
}

func (m *Resources) GetDiskGb() float64 {
	if m != nil {
		return m.DiskGb
	}
	return 0
}

type Node struct {
	Id          string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Resources   *Resources `protobuf:"bytes,5,opt,name=resources" json:"resources,omitempty"`
	Available   *Resources `protobuf:"bytes,6,opt,name=available" json:"available,omitempty"`
	State       NodeState  `protobuf:"varint,8,opt,name=state,enum=scheduler.NodeState" json:"state,omitempty"`
	Preemptible bool       `protobuf:"varint,9,opt,name=preemptible" json:"preemptible,omitempty"`
	Zone        string     `protobuf:"bytes,11,opt,name=zone" json:"zone,omitempty"`
	Hostname    string     `protobuf:"bytes,13,opt,name=hostname" json:"hostname,omitempty"`
	// Timestamp version of the record in the database. Used to prevent write conflicts and as the last ping time.
	Version  int64             `protobuf:"varint,14,opt,name=version" json:"version,omitempty"`
	Metadata map[string]string `protobuf:"bytes,15,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	TaskIds  []string          `protobuf:"bytes,16,rep,name=task_ids,json=taskIds" json:"task_ids,omitempty"`
	LastPing int64             `protobuf:"varint,17,opt,name=last_ping,json=lastPing" json:"last_ping,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Node) GetAvailable() *Resources {
	if m != nil {
		return m.Available
	}
	return nil
}

func (m *Node) GetState() NodeState {
	if m != nil {
		return m.State
	}
	return NodeState_UNINITIALIZED
}

func (m *Node) GetPreemptible() bool {
	if m != nil {
		return m.Preemptible
	}
	return false
}

func (m *Node) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *Node) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Node) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Node) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Node) GetTaskIds() []string {
	if m != nil {
		return m.TaskIds
	}
	return nil
}

func (m *Node) GetLastPing() int64 {
	if m != nil {
		return m.LastPing
	}
	return 0
}

type GetNodeRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetNodeRequest) Reset()                    { *m = GetNodeRequest{} }
func (m *GetNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNodeRequest) ProtoMessage()               {}
func (*GetNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetNodeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListNodesRequest struct {
}

func (m *ListNodesRequest) Reset()                    { *m = ListNodesRequest{} }
func (m *ListNodesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListNodesRequest) ProtoMessage()               {}
func (*ListNodesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ListNodesResponse struct {
	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *ListNodesResponse) Reset()                    { *m = ListNodesResponse{} }
func (m *ListNodesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListNodesResponse) ProtoMessage()               {}
func (*ListNodesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListNodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type PutNodeResponse struct {
}

func (m *PutNodeResponse) Reset()                    { *m = PutNodeResponse{} }
func (m *PutNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*PutNodeResponse) ProtoMessage()               {}
func (*PutNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type DeleteNodeResponse struct {
}

func (m *DeleteNodeResponse) Reset()                    { *m = DeleteNodeResponse{} }
func (m *DeleteNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeResponse) ProtoMessage()               {}
func (*DeleteNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*Resources)(nil), "scheduler.Resources")
	proto.RegisterType((*Node)(nil), "scheduler.Node")
	proto.RegisterType((*GetNodeRequest)(nil), "scheduler.GetNodeRequest")
	proto.RegisterType((*ListNodesRequest)(nil), "scheduler.ListNodesRequest")
	proto.RegisterType((*ListNodesResponse)(nil), "scheduler.ListNodesResponse")
	proto.RegisterType((*PutNodeResponse)(nil), "scheduler.PutNodeResponse")
	proto.RegisterType((*DeleteNodeResponse)(nil), "scheduler.DeleteNodeResponse")
	proto.RegisterEnum("scheduler.NodeState", NodeState_name, NodeState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SchedulerService service

type SchedulerServiceClient interface {
	PutNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*PutNodeResponse, error)
	DeleteNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*DeleteNodeResponse, error)
	ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error)
	GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*Node, error)
}

type schedulerServiceClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerServiceClient(cc *grpc.ClientConn) SchedulerServiceClient {
	return &schedulerServiceClient{cc}
}

func (c *schedulerServiceClient) PutNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*PutNodeResponse, error) {
	out := new(PutNodeResponse)
	err := grpc.Invoke(ctx, "/scheduler.SchedulerService/PutNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) DeleteNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*DeleteNodeResponse, error) {
	out := new(DeleteNodeResponse)
	err := grpc.Invoke(ctx, "/scheduler.SchedulerService/DeleteNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error) {
	out := new(ListNodesResponse)
	err := grpc.Invoke(ctx, "/scheduler.SchedulerService/ListNodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := grpc.Invoke(ctx, "/scheduler.SchedulerService/GetNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SchedulerService service

type SchedulerServiceServer interface {
	PutNode(context.Context, *Node) (*PutNodeResponse, error)
	DeleteNode(context.Context, *Node) (*DeleteNodeResponse, error)
	ListNodes(context.Context, *ListNodesRequest) (*ListNodesResponse, error)
	GetNode(context.Context, *GetNodeRequest) (*Node, error)
}

func RegisterSchedulerServiceServer(s *grpc.Server, srv SchedulerServiceServer) {
	s.RegisterService(&_SchedulerService_serviceDesc, srv)
}

func _SchedulerService_PutNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).PutNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerService/PutNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).PutNode(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_DeleteNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).DeleteNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerService/DeleteNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).DeleteNode(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerService/ListNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).ListNodes(ctx, req.(*ListNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerService/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).GetNode(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SchedulerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scheduler.SchedulerService",
	HandlerType: (*SchedulerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutNode",
			Handler:    _SchedulerService_PutNode_Handler,
		},
		{
			MethodName: "DeleteNode",
			Handler:    _SchedulerService_DeleteNode_Handler,
		},
		{
			MethodName: "ListNodes",
			Handler:    _SchedulerService_ListNodes_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _SchedulerService_GetNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheduler.proto",
}

func init() { proto.RegisterFile("scheduler.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x6a, 0xdb, 0x4a,
	0x10, 0x8e, 0xfc, 0xaf, 0xf1, 0xb1, 0x2d, 0x0f, 0x39, 0xe7, 0x28, 0x4a, 0x0a, 0x42, 0x50, 0x10,
	0xb9, 0x88, 0xa9, 0x7b, 0x93, 0xa6, 0x50, 0x30, 0xd8, 0x18, 0x41, 0xe2, 0x84, 0x4d, 0x5b, 0x68,
	0x6e, 0xc2, 0xda, 0x5a, 0x9c, 0x25, 0xb2, 0xa4, 0x6a, 0xd7, 0x86, 0xb4, 0xf4, 0xa6, 0xaf, 0xd0,
	0x77, 0xe9, 0x23, 0xf4, 0x05, 0xfa, 0x0a, 0x7d, 0x90, 0xb2, 0xb2, 0xa2, 0x28, 0x4a, 0xe9, 0xdd,
	0xce, 0x37, 0xdf, 0xcc, 0x7c, 0x33, 0x9f, 0x10, 0xf4, 0xc4, 0xe2, 0x86, 0xf9, 0xeb, 0x80, 0x25,
	0x47, 0x71, 0x12, 0xc9, 0x08, 0xf5, 0x1c, 0xb0, 0x0e, 0x96, 0x51, 0xb4, 0x0c, 0xd8, 0x80, 0xc6,
	0x7c, 0x40, 0xc3, 0x30, 0x92, 0x54, 0xf2, 0x28, 0x14, 0x5b, 0xa2, 0x73, 0x0e, 0x3a, 0x61, 0x22,
	0x5a, 0x27, 0x0b, 0x26, 0x10, 0xa1, 0xb6, 0x88, 0xd7, 0xc2, 0xd4, 0x6c, 0xcd, 0xed, 0x90, 0xf4,
	0x8d, 0xff, 0x42, 0x23, 0xa1, 0xab, 0xeb, 0xe5, 0xdc, 0xac, 0xd8, 0x9a, 0xab, 0x91, 0x7a, 0x42,
	0x57, 0xd3, 0x39, 0xfe, 0x0f, 0x4d, 0x9f, 0x8b, 0x5b, 0x85, 0x57, 0x53, 0xbc, 0xa1, 0xc2, 0xe9,
	0xdc, 0xf9, 0x51, 0x85, 0xda, 0x2c, 0xf2, 0x19, 0x76, 0xa1, 0xc2, 0xfd, 0xb4, 0x95, 0x4e, 0x2a,
	0xdc, 0xc7, 0x21, 0xe8, 0xc9, 0xfd, 0x24, 0xb3, 0x6e, 0x6b, 0x6e, 0x7b, 0xb8, 0x7b, 0xf4, 0xa0,
	0x3b, 0x57, 0x41, 0x1e, 0x68, 0xaa, 0x86, 0x6e, 0x28, 0x0f, 0xe8, 0x3c, 0x60, 0x66, 0xe3, 0x6f,
	0x35, 0x39, 0x0d, 0x0f, 0xa1, 0x2e, 0x24, 0x95, 0xcc, 0x6c, 0xd9, 0x9a, 0xdb, 0x7d, 0xc4, 0x57,
	0xba, 0x2e, 0x55, 0x8e, 0x6c, 0x29, 0x68, 0x43, 0x3b, 0x4e, 0x18, 0x5b, 0xc5, 0x92, 0xab, 0x09,
	0xba, 0xad, 0xb9, 0x2d, 0x52, 0x84, 0xd4, 0x49, 0x3e, 0x45, 0x21, 0x33, 0xdb, 0xe9, 0x1e, 0xe9,
	0x1b, 0x2d, 0x68, 0xdd, 0x44, 0x42, 0x86, 0x74, 0xc5, 0xcc, 0x4e, 0x8a, 0xe7, 0x31, 0x9a, 0xd0,
	0xdc, 0xb0, 0x44, 0xf0, 0x28, 0x34, 0xbb, 0xb6, 0xe6, 0x56, 0xc9, 0x7d, 0x88, 0xaf, 0xa0, 0xb5,
	0x62, 0x92, 0xfa, 0x54, 0x52, 0xb3, 0x67, 0x57, 0xdd, 0xf6, 0xf0, 0x59, 0x49, 0xda, 0xd1, 0x59,
	0x96, 0x9f, 0x84, 0x32, 0xb9, 0x23, 0x39, 0x1d, 0xf7, 0xa0, 0x25, 0xa9, 0xb8, 0xbd, 0xe6, 0xbe,
	0x30, 0x0d, 0xbb, 0xea, 0xea, 0xa4, 0xa9, 0x62, 0xcf, 0x17, 0xb8, 0x0f, 0x7a, 0x40, 0x85, 0xbc,
	0x8e, 0x79, 0xb8, 0x34, 0xfb, 0xe9, 0xc4, 0x96, 0x02, 0x2e, 0x78, 0xb8, 0xb4, 0x5e, 0x43, 0xe7,
	0x51, 0x4b, 0x34, 0xa0, 0x7a, 0xcb, 0xee, 0x32, 0x53, 0xd4, 0x13, 0x77, 0xa1, 0xbe, 0xa1, 0xc1,
	0x9a, 0xa5, 0xee, 0xea, 0x64, 0x1b, 0x9c, 0x54, 0x8e, 0x35, 0xc7, 0x86, 0xee, 0x94, 0x49, 0xa5,
	0x8b, 0xb0, 0x8f, 0x6b, 0x26, 0x64, 0xd9, 0x51, 0x07, 0xc1, 0x38, 0xe5, 0x22, 0xa5, 0x88, 0x8c,
	0xe3, 0x9c, 0x40, 0xbf, 0x80, 0x89, 0x38, 0x0a, 0x05, 0xc3, 0xe7, 0x50, 0x0f, 0x15, 0x60, 0x6a,
	0xe9, 0xde, 0xbd, 0xd2, 0xde, 0x64, 0x9b, 0x75, 0xfa, 0xd0, 0xbb, 0x58, 0x67, 0x13, 0xb7, 0x95,
	0xce, 0x2e, 0xe0, 0x98, 0x05, 0x4c, 0xb2, 0x22, 0x7a, 0x78, 0x05, 0x7a, 0x6e, 0x25, 0xf6, 0xa1,
	0xf3, 0x6e, 0xe6, 0xcd, 0xbc, 0xb7, 0xde, 0xe8, 0xd4, 0xbb, 0x9a, 0x8c, 0x8d, 0x1d, 0xd4, 0xa1,
	0x3e, 0x3a, 0xf5, 0xde, 0x4f, 0x0c, 0x0d, 0x5b, 0x50, 0x1b, 0x4f, 0x46, 0x63, 0xa3, 0xa2, 0x5e,
	0xd3, 0xf3, 0xd9, 0xc4, 0xa8, 0xa2, 0x01, 0xff, 0xe4, 0x7c, 0x6f, 0x36, 0x35, 0x6a, 0xaa, 0x60,
	0x4c, 0x46, 0xde, 0xcc, 0xa8, 0x0f, 0xbf, 0x57, 0xc0, 0xb8, 0xbc, 0x97, 0x77, 0xc9, 0x92, 0x0d,
	0x5f, 0x30, 0x3c, 0x86, 0x66, 0xa6, 0x0c, 0xcb, 0xe2, 0x2d, 0xab, 0x00, 0x94, 0xe5, 0xef, 0xe0,
	0x1b, 0x80, 0x87, 0x05, 0x9e, 0x16, 0x17, 0x3f, 0x81, 0xa7, 0x8b, 0x3a, 0x3b, 0xf8, 0x01, 0xf4,
	0xfc, 0x9e, 0xb8, 0x5f, 0x60, 0x97, 0x2f, 0x6f, 0x1d, 0xfc, 0x39, 0x99, 0x75, 0xea, 0x7f, 0xfd,
	0xf9, 0xeb, 0x5b, 0xa5, 0x8d, 0xfa, 0x60, 0xf3, 0x62, 0x90, 0x9e, 0x1b, 0xcf, 0xa0, 0x99, 0x19,
	0x8c, 0x7b, 0x85, 0xda, 0xc7, 0xa6, 0x5b, 0x65, 0xc9, 0xce, 0x7f, 0x69, 0x27, 0x03, 0xbb, 0x79,
	0xa7, 0xc1, 0x67, 0xee, 0x7f, 0x99, 0x37, 0xd2, 0x1f, 0xca, 0xcb, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xdb, 0x64, 0x56, 0xb7, 0x8c, 0x04, 0x00, 0x00,
}
