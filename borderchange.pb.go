// Code generated by protoc-gen-go. DO NOT EDIT.
// source: borderchange.proto

package game_border

import (
	context "context"
	fmt "fmt"
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

type GetBorderRequest struct {
	Style                string   `protobuf:"bytes,1,opt,name=style,proto3" json:"style,omitempty"`
	Width                string   `protobuf:"bytes,2,opt,name=width,proto3" json:"width,omitempty"`
	ShapeColorRed        float64  `protobuf:"fixed64,3,opt,name=shape_color_red,json=shapeColorRed,proto3" json:"shape_color_red,omitempty"`
	ShapeColorGreen      float64  `protobuf:"fixed64,4,opt,name=shape_color_green,json=shapeColorGreen,proto3" json:"shape_color_green,omitempty"`
	ShapeColorBlue       float64  `protobuf:"fixed64,5,opt,name=shape_color_blue,json=shapeColorBlue,proto3" json:"shape_color_blue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBorderRequest) Reset()         { *m = GetBorderRequest{} }
func (m *GetBorderRequest) String() string { return proto.CompactTextString(m) }
func (*GetBorderRequest) ProtoMessage()    {}
func (*GetBorderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65e06fa8348ad9be, []int{0}
}

func (m *GetBorderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBorderRequest.Unmarshal(m, b)
}
func (m *GetBorderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBorderRequest.Marshal(b, m, deterministic)
}
func (m *GetBorderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBorderRequest.Merge(m, src)
}
func (m *GetBorderRequest) XXX_Size() int {
	return xxx_messageInfo_GetBorderRequest.Size(m)
}
func (m *GetBorderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBorderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBorderRequest proto.InternalMessageInfo

func (m *GetBorderRequest) GetStyle() string {
	if m != nil {
		return m.Style
	}
	return ""
}

func (m *GetBorderRequest) GetWidth() string {
	if m != nil {
		return m.Width
	}
	return ""
}

func (m *GetBorderRequest) GetShapeColorRed() float64 {
	if m != nil {
		return m.ShapeColorRed
	}
	return 0
}

func (m *GetBorderRequest) GetShapeColorGreen() float64 {
	if m != nil {
		return m.ShapeColorGreen
	}
	return 0
}

func (m *GetBorderRequest) GetShapeColorBlue() float64 {
	if m != nil {
		return m.ShapeColorBlue
	}
	return 0
}

type GetBorderResponse struct {
	Style                string   `protobuf:"bytes,1,opt,name=style,proto3" json:"style,omitempty"`
	Width                string   `protobuf:"bytes,2,opt,name=width,proto3" json:"width,omitempty"`
	BorderColorRed       float64  `protobuf:"fixed64,3,opt,name=border_color_red,json=borderColorRed,proto3" json:"border_color_red,omitempty"`
	BorderColorGreen     float64  `protobuf:"fixed64,4,opt,name=border_color_green,json=borderColorGreen,proto3" json:"border_color_green,omitempty"`
	BorderColorBlue      float64  `protobuf:"fixed64,5,opt,name=border_color_blue,json=borderColorBlue,proto3" json:"border_color_blue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBorderResponse) Reset()         { *m = GetBorderResponse{} }
func (m *GetBorderResponse) String() string { return proto.CompactTextString(m) }
func (*GetBorderResponse) ProtoMessage()    {}
func (*GetBorderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65e06fa8348ad9be, []int{1}
}

func (m *GetBorderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBorderResponse.Unmarshal(m, b)
}
func (m *GetBorderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBorderResponse.Marshal(b, m, deterministic)
}
func (m *GetBorderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBorderResponse.Merge(m, src)
}
func (m *GetBorderResponse) XXX_Size() int {
	return xxx_messageInfo_GetBorderResponse.Size(m)
}
func (m *GetBorderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBorderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBorderResponse proto.InternalMessageInfo

func (m *GetBorderResponse) GetStyle() string {
	if m != nil {
		return m.Style
	}
	return ""
}

func (m *GetBorderResponse) GetWidth() string {
	if m != nil {
		return m.Width
	}
	return ""
}

func (m *GetBorderResponse) GetBorderColorRed() float64 {
	if m != nil {
		return m.BorderColorRed
	}
	return 0
}

func (m *GetBorderResponse) GetBorderColorGreen() float64 {
	if m != nil {
		return m.BorderColorGreen
	}
	return 0
}

func (m *GetBorderResponse) GetBorderColorBlue() float64 {
	if m != nil {
		return m.BorderColorBlue
	}
	return 0
}

func init() {
	proto.RegisterType((*GetBorderRequest)(nil), "m.borderchange.v1.GetBorderRequest")
	proto.RegisterType((*GetBorderResponse)(nil), "m.borderchange.v1.GetBorderResponse")
}

func init() { proto.RegisterFile("borderchange.proto", fileDescriptor_65e06fa8348ad9be) }

var fileDescriptor_65e06fa8348ad9be = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x80, 0x89, 0x3a, 0x61, 0x4f, 0xb6, 0xb6, 0xc1, 0x43, 0xf0, 0x34, 0xa6, 0x48, 0x19, 0x52,
	0x50, 0xff, 0x41, 0x3d, 0xf4, 0xde, 0x83, 0x07, 0x2f, 0xa5, 0x5d, 0x1e, 0xad, 0xd0, 0x36, 0x35,
	0x49, 0x15, 0xff, 0x9b, 0x37, 0xff, 0x98, 0x24, 0xc1, 0x2d, 0xb3, 0x22, 0x78, 0xcc, 0xc7, 0x77,
	0x78, 0xdf, 0x7b, 0x01, 0x5a, 0x09, 0xc9, 0x51, 0x6e, 0x9b, 0xb2, 0xaf, 0x31, 0x19, 0xa4, 0xd0,
	0x82, 0x46, 0x5d, 0x72, 0x40, 0x5f, 0x6f, 0xd7, 0x1f, 0x04, 0xc2, 0x0c, 0x75, 0x6a, 0x71, 0x8e,
	0x2f, 0x23, 0x2a, 0x4d, 0xcf, 0x61, 0xa6, 0xf4, 0x7b, 0x8b, 0x8c, 0xac, 0x48, 0x3c, 0xcf, 0xdd,
	0xc3, 0xd0, 0xb7, 0x67, 0xae, 0x1b, 0x76, 0xe4, 0xa8, 0x7d, 0xd0, 0x6b, 0x08, 0x54, 0x53, 0x0e,
	0x58, 0x6c, 0x45, 0x2b, 0x64, 0x21, 0x91, 0xb3, 0xe3, 0x15, 0x89, 0x49, 0xbe, 0xb0, 0xf8, 0xc1,
	0xd0, 0x1c, 0x39, 0xdd, 0x40, 0xe4, 0x7b, 0xb5, 0x44, 0xec, 0xd9, 0x89, 0x35, 0x83, 0xbd, 0x99,
	0x19, 0x4c, 0x63, 0x08, 0x7d, 0xb7, 0x6a, 0x47, 0x64, 0x33, 0xab, 0x2e, 0xf7, 0x6a, 0xda, 0x8e,
	0xb8, 0xfe, 0x24, 0x10, 0x79, 0xe3, 0xab, 0x41, 0xf4, 0x0a, 0xff, 0x35, 0x7f, 0x0c, 0xa1, 0xdb,
	0xc9, 0x24, 0x60, 0xe9, 0xf8, 0xae, 0xe0, 0xe6, 0x7b, 0xa7, 0xbf, 0x24, 0x84, 0x9e, 0xeb, 0x1a,
	0x36, 0x10, 0x1d, 0xd8, 0x5e, 0x44, 0xe0, 0xc9, 0xa6, 0xe2, 0x8e, 0x03, 0x64, 0x65, 0x87, 0xae,
	0x82, 0x3e, 0xc2, 0x7c, 0x97, 0x44, 0x2f, 0x93, 0xc9, 0xcd, 0x92, 0x9f, 0xf7, 0xba, 0xb8, 0xfa,
	0x5b, 0x72, 0x5b, 0x49, 0x17, 0x4f, 0x67, 0x75, 0xd9, 0x61, 0xe1, 0xcc, 0xea, 0xd4, 0xfe, 0x89,
	0xfb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xb1, 0xe6, 0x6e, 0x29, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameBorderClient is the client API for GameBorder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameBorderClient interface {
	GetBorder(ctx context.Context, in *GetBorderRequest, opts ...grpc.CallOption) (*GetBorderResponse, error)
}

type gameBorderClient struct {
	cc *grpc.ClientConn
}

func NewGameBorderClient(cc *grpc.ClientConn) GameBorderClient {
	return &gameBorderClient{cc}
}

func (c *gameBorderClient) GetBorder(ctx context.Context, in *GetBorderRequest, opts ...grpc.CallOption) (*GetBorderResponse, error) {
	out := new(GetBorderResponse)
	err := c.cc.Invoke(ctx, "/m.borderchange.v1.GameBorder/GetBorder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameBorderServer is the server API for GameBorder service.
type GameBorderServer interface {
	GetBorder(context.Context, *GetBorderRequest) (*GetBorderResponse, error)
}

// UnimplementedGameBorderServer can be embedded to have forward compatible implementations.
type UnimplementedGameBorderServer struct {
}

func (*UnimplementedGameBorderServer) GetBorder(ctx context.Context, req *GetBorderRequest) (*GetBorderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBorder not implemented")
}

func RegisterGameBorderServer(s *grpc.Server, srv GameBorderServer) {
	s.RegisterService(&_GameBorder_serviceDesc, srv)
}

func _GameBorder_GetBorder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBorderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameBorderServer).GetBorder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/m.borderchange.v1.GameBorder/GetBorder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameBorderServer).GetBorder(ctx, req.(*GetBorderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameBorder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "m.borderchange.v1.GameBorder",
	HandlerType: (*GameBorderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBorder",
			Handler:    _GameBorder_GetBorder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "borderchange.proto",
}
