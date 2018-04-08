// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/examples/booking/api/hotel/proto/hotel.proto

/*
Package hotel is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/examples/booking/api/hotel/proto/hotel.proto

It has these top-level messages:
	Request
	Response
*/
package hotel

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import profile "github.com/micro/examples/booking/srv/profile/proto"
import rate "github.com/micro/examples/booking/srv/rate/proto"

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

type Request struct {
	InDate  string `protobuf:"bytes,1,opt,name=inDate" json:"inDate,omitempty"`
	OutDate string `protobuf:"bytes,2,opt,name=outDate" json:"outDate,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *Request) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

type Response struct {
	Hotels    []*profile.Hotel `protobuf:"bytes,1,rep,name=hotels" json:"hotels,omitempty"`
	RatePlans []*rate.RatePlan `protobuf:"bytes,2,rep,name=ratePlans" json:"ratePlans,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetHotels() []*profile.Hotel {
	if m != nil {
		return m.Hotels
	}
	return nil
}

func (m *Response) GetRatePlans() []*rate.RatePlan {
	if m != nil {
		return m.RatePlans
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "hotel.Request")
	proto.RegisterType((*Response)(nil), "hotel.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hotel service

type HotelClient interface {
	Rates(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type hotelClient struct {
	cc *grpc.ClientConn
}

func NewHotelClient(cc *grpc.ClientConn) HotelClient {
	return &hotelClient{cc}
}

func (c *hotelClient) Rates(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/hotel.Hotel/Rates", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hotel service

type HotelServer interface {
	Rates(context.Context, *Request) (*Response, error)
}

func RegisterHotelServer(s *grpc.Server, srv HotelServer) {
	s.RegisterService(&_Hotel_serviceDesc, srv)
}

func _Hotel_Rates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServer).Rates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hotel.Hotel/Rates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServer).Rates(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hotel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hotel.Hotel",
	HandlerType: (*HotelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Rates",
			Handler:    _Hotel_Rates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/micro/examples/booking/api/hotel/proto/hotel.proto",
}

func init() {
	proto.RegisterFile("github.com/micro/examples/booking/api/hotel/proto/hotel.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x35, 0x95, 0xa4, 0x76, 0x84, 0x0a, 0x7b, 0x90, 0x90, 0x53, 0xc9, 0x41, 0x8a, 0x48, 0x16,
	0xda, 0x63, 0xf1, 0x20, 0x78, 0xf0, 0x28, 0xfb, 0x05, 0x6e, 0xca, 0xd8, 0x2c, 0x26, 0x99, 0x98,
	0xdd, 0x88, 0x9f, 0x2f, 0x3b, 0xbb, 0xd1, 0xa3, 0xbd, 0xbd, 0x37, 0x6f, 0xde, 0xdb, 0x37, 0x0b,
	0x8f, 0x27, 0xe3, 0x9a, 0xa9, 0xae, 0x8e, 0xd4, 0xc9, 0xce, 0x1c, 0x47, 0x92, 0xf8, 0xad, 0xbb,
	0xa1, 0x45, 0x2b, 0x6b, 0xa2, 0x0f, 0xd3, 0x9f, 0xa4, 0x1e, 0x8c, 0x6c, 0xc8, 0x61, 0x2b, 0x87,
	0x91, 0x1c, 0x05, 0x5c, 0x31, 0x16, 0x29, 0x93, 0xe2, 0xe9, 0xff, 0x14, 0x3b, 0x7e, 0x79, 0xff,
	0xbb, 0x69, 0x31, 0xe6, 0x44, 0x16, 0x92, 0x8a, 0xc3, 0x79, 0x11, 0xa3, 0x76, 0xb3, 0xdf, 0xc3,
	0x60, 0x2e, 0x0f, 0xb0, 0x54, 0xf8, 0x39, 0xa1, 0x75, 0xe2, 0x16, 0x32, 0xd3, 0x3f, 0x6b, 0x87,
	0x79, 0xb2, 0x49, 0xb6, 0x2b, 0x15, 0x99, 0xc8, 0x61, 0x49, 0x93, 0x63, 0x61, 0xc1, 0xc2, 0x4c,
	0xcb, 0x37, 0xb8, 0x52, 0x68, 0x07, 0xea, 0x2d, 0x8a, 0x3b, 0xc8, 0xf8, 0x22, 0x9b, 0x27, 0x9b,
	0xcb, 0xed, 0xf5, 0x6e, 0x5d, 0xcd, 0x2d, 0x5f, 0xfc, 0x58, 0x45, 0x55, 0x3c, 0xc0, 0xca, 0x3f,
	0xff, 0xda, 0xea, 0xde, 0xe6, 0x8b, 0xb8, 0xca, 0x85, 0x54, 0x1c, 0xab, 0xbf, 0x85, 0xdd, 0x1e,
	0x52, 0xb6, 0x8b, 0x7b, 0x48, 0xbd, 0x6e, 0xc5, 0xba, 0x0a, 0xbf, 0x18, 0x5b, 0x17, 0x37, 0xbf,
	0x3c, 0x14, 0x29, 0x2f, 0xea, 0x8c, 0x4f, 0xdb, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x11, 0xa3,
	0x5e, 0xbd, 0xa2, 0x01, 0x00, 0x00,
}
