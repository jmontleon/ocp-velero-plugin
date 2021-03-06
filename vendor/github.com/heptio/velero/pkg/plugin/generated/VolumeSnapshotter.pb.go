// Code generated by protoc-gen-go. DO NOT EDIT.
// source: VolumeSnapshotter.proto

package generated

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateVolumeRequest struct {
	Plugin     string `protobuf:"bytes,1,opt,name=plugin" json:"plugin,omitempty"`
	SnapshotID string `protobuf:"bytes,2,opt,name=snapshotID" json:"snapshotID,omitempty"`
	VolumeType string `protobuf:"bytes,3,opt,name=volumeType" json:"volumeType,omitempty"`
	VolumeAZ   string `protobuf:"bytes,4,opt,name=volumeAZ" json:"volumeAZ,omitempty"`
	Iops       int64  `protobuf:"varint,5,opt,name=iops" json:"iops,omitempty"`
}

func (m *CreateVolumeRequest) Reset()                    { *m = CreateVolumeRequest{} }
func (m *CreateVolumeRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateVolumeRequest) ProtoMessage()               {}
func (*CreateVolumeRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *CreateVolumeRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *CreateVolumeRequest) GetSnapshotID() string {
	if m != nil {
		return m.SnapshotID
	}
	return ""
}

func (m *CreateVolumeRequest) GetVolumeType() string {
	if m != nil {
		return m.VolumeType
	}
	return ""
}

func (m *CreateVolumeRequest) GetVolumeAZ() string {
	if m != nil {
		return m.VolumeAZ
	}
	return ""
}

func (m *CreateVolumeRequest) GetIops() int64 {
	if m != nil {
		return m.Iops
	}
	return 0
}

type CreateVolumeResponse struct {
	VolumeID string `protobuf:"bytes,1,opt,name=volumeID" json:"volumeID,omitempty"`
}

func (m *CreateVolumeResponse) Reset()                    { *m = CreateVolumeResponse{} }
func (m *CreateVolumeResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateVolumeResponse) ProtoMessage()               {}
func (*CreateVolumeResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *CreateVolumeResponse) GetVolumeID() string {
	if m != nil {
		return m.VolumeID
	}
	return ""
}

type GetVolumeInfoRequest struct {
	Plugin   string `protobuf:"bytes,1,opt,name=plugin" json:"plugin,omitempty"`
	VolumeID string `protobuf:"bytes,2,opt,name=volumeID" json:"volumeID,omitempty"`
	VolumeAZ string `protobuf:"bytes,3,opt,name=volumeAZ" json:"volumeAZ,omitempty"`
}

func (m *GetVolumeInfoRequest) Reset()                    { *m = GetVolumeInfoRequest{} }
func (m *GetVolumeInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetVolumeInfoRequest) ProtoMessage()               {}
func (*GetVolumeInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *GetVolumeInfoRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *GetVolumeInfoRequest) GetVolumeID() string {
	if m != nil {
		return m.VolumeID
	}
	return ""
}

func (m *GetVolumeInfoRequest) GetVolumeAZ() string {
	if m != nil {
		return m.VolumeAZ
	}
	return ""
}

type GetVolumeInfoResponse struct {
	VolumeType string `protobuf:"bytes,1,opt,name=volumeType" json:"volumeType,omitempty"`
	Iops       int64  `protobuf:"varint,2,opt,name=iops" json:"iops,omitempty"`
}

func (m *GetVolumeInfoResponse) Reset()                    { *m = GetVolumeInfoResponse{} }
func (m *GetVolumeInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetVolumeInfoResponse) ProtoMessage()               {}
func (*GetVolumeInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *GetVolumeInfoResponse) GetVolumeType() string {
	if m != nil {
		return m.VolumeType
	}
	return ""
}

func (m *GetVolumeInfoResponse) GetIops() int64 {
	if m != nil {
		return m.Iops
	}
	return 0
}

type CreateSnapshotRequest struct {
	Plugin   string            `protobuf:"bytes,1,opt,name=plugin" json:"plugin,omitempty"`
	VolumeID string            `protobuf:"bytes,2,opt,name=volumeID" json:"volumeID,omitempty"`
	VolumeAZ string            `protobuf:"bytes,3,opt,name=volumeAZ" json:"volumeAZ,omitempty"`
	Tags     map[string]string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CreateSnapshotRequest) Reset()                    { *m = CreateSnapshotRequest{} }
func (m *CreateSnapshotRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateSnapshotRequest) ProtoMessage()               {}
func (*CreateSnapshotRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *CreateSnapshotRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *CreateSnapshotRequest) GetVolumeID() string {
	if m != nil {
		return m.VolumeID
	}
	return ""
}

func (m *CreateSnapshotRequest) GetVolumeAZ() string {
	if m != nil {
		return m.VolumeAZ
	}
	return ""
}

func (m *CreateSnapshotRequest) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type CreateSnapshotResponse struct {
	SnapshotID string `protobuf:"bytes,1,opt,name=snapshotID" json:"snapshotID,omitempty"`
}

func (m *CreateSnapshotResponse) Reset()                    { *m = CreateSnapshotResponse{} }
func (m *CreateSnapshotResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateSnapshotResponse) ProtoMessage()               {}
func (*CreateSnapshotResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *CreateSnapshotResponse) GetSnapshotID() string {
	if m != nil {
		return m.SnapshotID
	}
	return ""
}

type DeleteSnapshotRequest struct {
	Plugin     string `protobuf:"bytes,1,opt,name=plugin" json:"plugin,omitempty"`
	SnapshotID string `protobuf:"bytes,2,opt,name=snapshotID" json:"snapshotID,omitempty"`
}

func (m *DeleteSnapshotRequest) Reset()                    { *m = DeleteSnapshotRequest{} }
func (m *DeleteSnapshotRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteSnapshotRequest) ProtoMessage()               {}
func (*DeleteSnapshotRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *DeleteSnapshotRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *DeleteSnapshotRequest) GetSnapshotID() string {
	if m != nil {
		return m.SnapshotID
	}
	return ""
}

type GetVolumeIDRequest struct {
	Plugin           string `protobuf:"bytes,1,opt,name=plugin" json:"plugin,omitempty"`
	PersistentVolume []byte `protobuf:"bytes,2,opt,name=persistentVolume,proto3" json:"persistentVolume,omitempty"`
}

func (m *GetVolumeIDRequest) Reset()                    { *m = GetVolumeIDRequest{} }
func (m *GetVolumeIDRequest) String() string            { return proto.CompactTextString(m) }
func (*GetVolumeIDRequest) ProtoMessage()               {}
func (*GetVolumeIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *GetVolumeIDRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *GetVolumeIDRequest) GetPersistentVolume() []byte {
	if m != nil {
		return m.PersistentVolume
	}
	return nil
}

type GetVolumeIDResponse struct {
	VolumeID string `protobuf:"bytes,1,opt,name=volumeID" json:"volumeID,omitempty"`
}

func (m *GetVolumeIDResponse) Reset()                    { *m = GetVolumeIDResponse{} }
func (m *GetVolumeIDResponse) String() string            { return proto.CompactTextString(m) }
func (*GetVolumeIDResponse) ProtoMessage()               {}
func (*GetVolumeIDResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *GetVolumeIDResponse) GetVolumeID() string {
	if m != nil {
		return m.VolumeID
	}
	return ""
}

type SetVolumeIDRequest struct {
	Plugin           string `protobuf:"bytes,1,opt,name=plugin" json:"plugin,omitempty"`
	PersistentVolume []byte `protobuf:"bytes,2,opt,name=persistentVolume,proto3" json:"persistentVolume,omitempty"`
	VolumeID         string `protobuf:"bytes,3,opt,name=volumeID" json:"volumeID,omitempty"`
}

func (m *SetVolumeIDRequest) Reset()                    { *m = SetVolumeIDRequest{} }
func (m *SetVolumeIDRequest) String() string            { return proto.CompactTextString(m) }
func (*SetVolumeIDRequest) ProtoMessage()               {}
func (*SetVolumeIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

func (m *SetVolumeIDRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *SetVolumeIDRequest) GetPersistentVolume() []byte {
	if m != nil {
		return m.PersistentVolume
	}
	return nil
}

func (m *SetVolumeIDRequest) GetVolumeID() string {
	if m != nil {
		return m.VolumeID
	}
	return ""
}

type SetVolumeIDResponse struct {
	PersistentVolume []byte `protobuf:"bytes,1,opt,name=persistentVolume,proto3" json:"persistentVolume,omitempty"`
}

func (m *SetVolumeIDResponse) Reset()                    { *m = SetVolumeIDResponse{} }
func (m *SetVolumeIDResponse) String() string            { return proto.CompactTextString(m) }
func (*SetVolumeIDResponse) ProtoMessage()               {}
func (*SetVolumeIDResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{10} }

func (m *SetVolumeIDResponse) GetPersistentVolume() []byte {
	if m != nil {
		return m.PersistentVolume
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateVolumeRequest)(nil), "generated.CreateVolumeRequest")
	proto.RegisterType((*CreateVolumeResponse)(nil), "generated.CreateVolumeResponse")
	proto.RegisterType((*GetVolumeInfoRequest)(nil), "generated.GetVolumeInfoRequest")
	proto.RegisterType((*GetVolumeInfoResponse)(nil), "generated.GetVolumeInfoResponse")
	proto.RegisterType((*CreateSnapshotRequest)(nil), "generated.CreateSnapshotRequest")
	proto.RegisterType((*CreateSnapshotResponse)(nil), "generated.CreateSnapshotResponse")
	proto.RegisterType((*DeleteSnapshotRequest)(nil), "generated.DeleteSnapshotRequest")
	proto.RegisterType((*GetVolumeIDRequest)(nil), "generated.GetVolumeIDRequest")
	proto.RegisterType((*GetVolumeIDResponse)(nil), "generated.GetVolumeIDResponse")
	proto.RegisterType((*SetVolumeIDRequest)(nil), "generated.SetVolumeIDRequest")
	proto.RegisterType((*SetVolumeIDResponse)(nil), "generated.SetVolumeIDResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VolumeSnapshotter service

type VolumeSnapshotterClient interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error)
	CreateVolumeFromSnapshot(ctx context.Context, in *CreateVolumeRequest, opts ...grpc.CallOption) (*CreateVolumeResponse, error)
	GetVolumeInfo(ctx context.Context, in *GetVolumeInfoRequest, opts ...grpc.CallOption) (*GetVolumeInfoResponse, error)
	CreateSnapshot(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*CreateSnapshotResponse, error)
	DeleteSnapshot(ctx context.Context, in *DeleteSnapshotRequest, opts ...grpc.CallOption) (*Empty, error)
	GetVolumeID(ctx context.Context, in *GetVolumeIDRequest, opts ...grpc.CallOption) (*GetVolumeIDResponse, error)
	SetVolumeID(ctx context.Context, in *SetVolumeIDRequest, opts ...grpc.CallOption) (*SetVolumeIDResponse, error)
}

type volumeSnapshotterClient struct {
	cc *grpc.ClientConn
}

func NewVolumeSnapshotterClient(cc *grpc.ClientConn) VolumeSnapshotterClient {
	return &volumeSnapshotterClient{cc}
}

func (c *volumeSnapshotterClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/generated.VolumeSnapshotter/Init", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeSnapshotterClient) CreateVolumeFromSnapshot(ctx context.Context, in *CreateVolumeRequest, opts ...grpc.CallOption) (*CreateVolumeResponse, error) {
	out := new(CreateVolumeResponse)
	err := grpc.Invoke(ctx, "/generated.VolumeSnapshotter/CreateVolumeFromSnapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeSnapshotterClient) GetVolumeInfo(ctx context.Context, in *GetVolumeInfoRequest, opts ...grpc.CallOption) (*GetVolumeInfoResponse, error) {
	out := new(GetVolumeInfoResponse)
	err := grpc.Invoke(ctx, "/generated.VolumeSnapshotter/GetVolumeInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeSnapshotterClient) CreateSnapshot(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*CreateSnapshotResponse, error) {
	out := new(CreateSnapshotResponse)
	err := grpc.Invoke(ctx, "/generated.VolumeSnapshotter/CreateSnapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeSnapshotterClient) DeleteSnapshot(ctx context.Context, in *DeleteSnapshotRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/generated.VolumeSnapshotter/DeleteSnapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeSnapshotterClient) GetVolumeID(ctx context.Context, in *GetVolumeIDRequest, opts ...grpc.CallOption) (*GetVolumeIDResponse, error) {
	out := new(GetVolumeIDResponse)
	err := grpc.Invoke(ctx, "/generated.VolumeSnapshotter/GetVolumeID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeSnapshotterClient) SetVolumeID(ctx context.Context, in *SetVolumeIDRequest, opts ...grpc.CallOption) (*SetVolumeIDResponse, error) {
	out := new(SetVolumeIDResponse)
	err := grpc.Invoke(ctx, "/generated.VolumeSnapshotter/SetVolumeID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VolumeSnapshotter service

type VolumeSnapshotterServer interface {
	Init(context.Context, *InitRequest) (*Empty, error)
	CreateVolumeFromSnapshot(context.Context, *CreateVolumeRequest) (*CreateVolumeResponse, error)
	GetVolumeInfo(context.Context, *GetVolumeInfoRequest) (*GetVolumeInfoResponse, error)
	CreateSnapshot(context.Context, *CreateSnapshotRequest) (*CreateSnapshotResponse, error)
	DeleteSnapshot(context.Context, *DeleteSnapshotRequest) (*Empty, error)
	GetVolumeID(context.Context, *GetVolumeIDRequest) (*GetVolumeIDResponse, error)
	SetVolumeID(context.Context, *SetVolumeIDRequest) (*SetVolumeIDResponse, error)
}

func RegisterVolumeSnapshotterServer(s *grpc.Server, srv VolumeSnapshotterServer) {
	s.RegisterService(&_VolumeSnapshotter_serviceDesc, srv)
}

func _VolumeSnapshotter_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeSnapshotterServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.VolumeSnapshotter/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeSnapshotterServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeSnapshotter_CreateVolumeFromSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeSnapshotterServer).CreateVolumeFromSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.VolumeSnapshotter/CreateVolumeFromSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeSnapshotterServer).CreateVolumeFromSnapshot(ctx, req.(*CreateVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeSnapshotter_GetVolumeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVolumeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeSnapshotterServer).GetVolumeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.VolumeSnapshotter/GetVolumeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeSnapshotterServer).GetVolumeInfo(ctx, req.(*GetVolumeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeSnapshotter_CreateSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeSnapshotterServer).CreateSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.VolumeSnapshotter/CreateSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeSnapshotterServer).CreateSnapshot(ctx, req.(*CreateSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeSnapshotter_DeleteSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeSnapshotterServer).DeleteSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.VolumeSnapshotter/DeleteSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeSnapshotterServer).DeleteSnapshot(ctx, req.(*DeleteSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeSnapshotter_GetVolumeID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVolumeIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeSnapshotterServer).GetVolumeID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.VolumeSnapshotter/GetVolumeID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeSnapshotterServer).GetVolumeID(ctx, req.(*GetVolumeIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeSnapshotter_SetVolumeID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetVolumeIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeSnapshotterServer).SetVolumeID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.VolumeSnapshotter/SetVolumeID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeSnapshotterServer).SetVolumeID(ctx, req.(*SetVolumeIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VolumeSnapshotter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generated.VolumeSnapshotter",
	HandlerType: (*VolumeSnapshotterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _VolumeSnapshotter_Init_Handler,
		},
		{
			MethodName: "CreateVolumeFromSnapshot",
			Handler:    _VolumeSnapshotter_CreateVolumeFromSnapshot_Handler,
		},
		{
			MethodName: "GetVolumeInfo",
			Handler:    _VolumeSnapshotter_GetVolumeInfo_Handler,
		},
		{
			MethodName: "CreateSnapshot",
			Handler:    _VolumeSnapshotter_CreateSnapshot_Handler,
		},
		{
			MethodName: "DeleteSnapshot",
			Handler:    _VolumeSnapshotter_DeleteSnapshot_Handler,
		},
		{
			MethodName: "GetVolumeID",
			Handler:    _VolumeSnapshotter_GetVolumeID_Handler,
		},
		{
			MethodName: "SetVolumeID",
			Handler:    _VolumeSnapshotter_SetVolumeID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "VolumeSnapshotter.proto",
}

func init() { proto.RegisterFile("VolumeSnapshotter.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0xc6, 0x6e, 0x45, 0x26, 0xa5, 0x0a, 0x93, 0x1f, 0x2c, 0x4b, 0x04, 0xe3, 0x53, 0xd4,
	0x43, 0x04, 0xe1, 0x40, 0xc5, 0x01, 0xa9, 0xc2, 0x05, 0x45, 0x54, 0x42, 0xb2, 0x0b, 0x42, 0x70,
	0x32, 0xea, 0x34, 0x8d, 0x48, 0x6c, 0xe3, 0xdd, 0x54, 0xca, 0xc3, 0xf0, 0x0c, 0xbc, 0x12, 0x8f,
	0x82, 0x62, 0x6f, 0x92, 0xdd, 0x64, 0x53, 0xf7, 0xd2, 0x9b, 0x67, 0x66, 0xe7, 0x9b, 0x6f, 0x66,
	0xbf, 0x59, 0xc3, 0xd3, 0xaf, 0xe9, 0x74, 0x3e, 0xa3, 0x28, 0x89, 0x33, 0x7e, 0x93, 0x0a, 0x41,
	0xf9, 0x20, 0xcb, 0x53, 0x91, 0x62, 0x7d, 0x4c, 0x09, 0xe5, 0xb1, 0xa0, 0x2b, 0xf7, 0x28, 0xba,
	0x89, 0x73, 0xba, 0x2a, 0x03, 0xfe, 0x1f, 0x06, 0xad, 0xf7, 0x39, 0xc5, 0x82, 0xca, 0xd4, 0x90,
	0x7e, 0xcf, 0x89, 0x0b, 0xec, 0xc2, 0x61, 0x36, 0x9d, 0x8f, 0x27, 0x89, 0xc3, 0x3c, 0xd6, 0xaf,
	0x87, 0xd2, 0xc2, 0x1e, 0x00, 0x97, 0xe8, 0xa3, 0xc0, 0xa9, 0x15, 0x31, 0xc5, 0xb3, 0x8c, 0xdf,
	0x16, 0x40, 0x97, 0x8b, 0x8c, 0x1c, 0xab, 0x8c, 0x6f, 0x3c, 0xe8, 0xc2, 0xa3, 0xd2, 0x3a, 0xfb,
	0xee, 0xd8, 0x45, 0x74, 0x6d, 0x23, 0x82, 0x3d, 0x49, 0x33, 0xee, 0x1c, 0x78, 0xac, 0x6f, 0x85,
	0xc5, 0xb7, 0x3f, 0x84, 0xb6, 0x4e, 0x8f, 0x67, 0x69, 0xc2, 0x15, 0x9c, 0x51, 0x20, 0x19, 0xae,
	0x6d, 0xff, 0x1a, 0xda, 0x1f, 0x49, 0x94, 0x09, 0xa3, 0xe4, 0x3a, 0xad, 0xea, 0x49, 0xc5, 0xaa,
	0xe9, 0x58, 0x1a, 0x5f, 0x4b, 0xe7, 0xeb, 0x7f, 0x82, 0xce, 0x56, 0x1d, 0x49, 0x4e, 0x1f, 0x02,
	0xdb, 0x19, 0xc2, 0xaa, 0xd1, 0x9a, 0xd2, 0xe8, 0x3f, 0x06, 0x9d, 0xb2, 0xd3, 0xd5, 0xed, 0x3d,
	0x10, 0x6d, 0x7c, 0x07, 0xb6, 0x88, 0xc7, 0xdc, 0xb1, 0x3d, 0xab, 0xdf, 0x18, 0x9e, 0x0c, 0xd6,
	0xd2, 0x18, 0x18, 0xeb, 0x0f, 0x2e, 0xe3, 0x31, 0x3f, 0x4f, 0x44, 0xbe, 0x08, 0x8b, 0x3c, 0xf7,
	0x0d, 0xd4, 0xd7, 0x2e, 0x6c, 0x82, 0xf5, 0x8b, 0x16, 0x92, 0xd9, 0xf2, 0x13, 0xdb, 0x70, 0x70,
	0x1b, 0x4f, 0xe7, 0x24, 0x39, 0x95, 0xc6, 0xdb, 0xda, 0x29, 0xf3, 0x4f, 0xa1, 0xbb, 0x5d, 0x61,
	0x33, 0x30, 0x45, 0x55, 0x6c, 0x5b, 0x55, 0xfe, 0x67, 0xe8, 0x04, 0x34, 0xa5, 0xfb, 0xcf, 0xa6,
	0x42, 0xa6, 0xfe, 0x37, 0xc0, 0xcd, 0xd5, 0x05, 0x55, 0x68, 0x27, 0xd0, 0xcc, 0x28, 0xe7, 0x13,
	0x2e, 0x28, 0x91, 0x49, 0x05, 0xe6, 0x51, 0xb8, 0xe3, 0xf7, 0x5f, 0x41, 0x4b, 0x43, 0xbe, 0x87,
	0x5e, 0x05, 0x60, 0xf4, 0x20, 0x64, 0xb4, 0xaa, 0xd6, 0x56, 0xd5, 0x33, 0x68, 0x45, 0x06, 0xa2,
	0x26, 0x78, 0x66, 0x86, 0x1f, 0xfe, 0xb5, 0xe1, 0xc9, 0xce, 0x8b, 0x83, 0x2f, 0xc1, 0x1e, 0x25,
	0x13, 0x81, 0x5d, 0x45, 0x59, 0x4b, 0x87, 0x6c, 0xcc, 0x6d, 0x2a, 0xfe, 0xf3, 0x59, 0x26, 0x16,
	0xf8, 0x03, 0x1c, 0x75, 0xc9, 0x3f, 0xe4, 0xe9, 0x6c, 0x05, 0x88, 0xbd, 0x1d, 0x7d, 0x6a, 0x0f,
	0x95, 0xfb, 0x7c, 0x6f, 0x5c, 0x36, 0x14, 0xc2, 0x63, 0x6d, 0x4b, 0x51, 0xcd, 0x30, 0xbd, 0x13,
	0xae, 0xb7, 0xff, 0x80, 0xc4, 0xfc, 0x02, 0xc7, 0xba, 0x92, 0xd1, 0xab, 0x5a, 0x23, 0xf7, 0xc5,
	0x1d, 0x27, 0x24, 0x6c, 0x00, 0xc7, 0xba, 0xcc, 0x35, 0x58, 0xe3, 0x06, 0x18, 0xa6, 0x79, 0x01,
	0x0d, 0x45, 0x81, 0xf8, 0xcc, 0xd8, 0xcd, 0x4a, 0x66, 0x6e, 0x6f, 0x5f, 0x58, 0x72, 0xba, 0x80,
	0x46, 0xb4, 0x07, 0x2d, 0xba, 0x1b, 0xcd, 0xa0, 0xae, 0x9f, 0x87, 0xc5, 0x5f, 0xe7, 0xf5, 0xff,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0x03, 0xe3, 0x22, 0xa9, 0x06, 0x00, 0x00,
}
