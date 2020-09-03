// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/xxdawn/micro/examples/booking/srv/profile/proto/profile.proto

/*
Package profile is a generated protocol buffer package.

It is generated from these files:
	github.com/xxdawn/micro/examples/booking/srv/profile/proto/profile.proto

It has these top-level messages:
	Request
	Result
	Hotel
	Address
	Image
*/
package profile

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	HotelIds []string `protobuf:"bytes,1,rep,name=hotelIds" json:"hotelIds,omitempty"`
	Locale   string   `protobuf:"bytes,2,opt,name=locale" json:"locale,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetHotelIds() []string {
	if m != nil {
		return m.HotelIds
	}
	return nil
}

func (m *Request) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

type Result struct {
	Hotels []*Hotel `protobuf:"bytes,1,rep,name=hotels" json:"hotels,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetHotels() []*Hotel {
	if m != nil {
		return m.Hotels
	}
	return nil
}

type Hotel struct {
	Id          string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	PhoneNumber string   `protobuf:"bytes,3,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Address     *Address `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	Images      []*Image `protobuf:"bytes,6,rep,name=images" json:"images,omitempty"`
}

func (m *Hotel) Reset()                    { *m = Hotel{} }
func (m *Hotel) String() string            { return proto.CompactTextString(m) }
func (*Hotel) ProtoMessage()               {}
func (*Hotel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Hotel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Hotel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Hotel) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Hotel) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Hotel) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Hotel) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

type Address struct {
	StreetNumber string `protobuf:"bytes,1,opt,name=streetNumber" json:"streetNumber,omitempty"`
	StreetName   string `protobuf:"bytes,2,opt,name=streetName" json:"streetName,omitempty"`
	City         string `protobuf:"bytes,3,opt,name=city" json:"city,omitempty"`
	State        string `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	Country      string `protobuf:"bytes,5,opt,name=country" json:"country,omitempty"`
	PostalCode   string `protobuf:"bytes,6,opt,name=postalCode" json:"postalCode,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Address) GetStreetNumber() string {
	if m != nil {
		return m.StreetNumber
	}
	return ""
}

func (m *Address) GetStreetName() string {
	if m != nil {
		return m.StreetName
	}
	return ""
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Address) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

type Image struct {
	Url     string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Default bool   `protobuf:"varint,2,opt,name=default" json:"default,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Image) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Image) GetDefault() bool {
	if m != nil {
		return m.Default
	}
	return false
}

func init() {
	proto.RegisterType((*Request)(nil), "profile.Request")
	proto.RegisterType((*Result)(nil), "profile.Result")
	proto.RegisterType((*Hotel)(nil), "profile.Hotel")
	proto.RegisterType((*Address)(nil), "profile.Address")
	proto.RegisterType((*Image)(nil), "profile.Image")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Profile service

type ProfileClient interface {
	GetProfiles(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type profileClient struct {
	cc *grpc.ClientConn
}

func NewProfileClient(cc *grpc.ClientConn) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) GetProfiles(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/profile.Profile/GetProfiles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Profile service

type ProfileServer interface {
	GetProfiles(context.Context, *Request) (*Result, error)
}

func RegisterProfileServer(s *grpc.Server, srv ProfileServer) {
	s.RegisterService(&_Profile_serviceDesc, srv)
}

func _Profile_GetProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/GetProfiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfiles(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Profile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profile.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfiles",
			Handler:    _Profile_GetProfiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/xxdawn/micro/examples/booking/srv/profile/proto/profile.proto",
}

func init() {
	proto.RegisterFile("github.com/xxdawn/micro/examples/booking/srv/profile/proto/profile.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0x95, 0xdd, 0x6e, 0xd2, 0x9d, 0x45, 0xa5, 0x1a, 0x21, 0x64, 0xf5, 0x80, 0x56, 0x39,
	0xa0, 0x15, 0x87, 0x4d, 0xb5, 0x3d, 0x22, 0x0e, 0x15, 0x07, 0xe8, 0x05, 0x21, 0xbf, 0x81, 0x13,
	0x4f, 0x77, 0x2d, 0x9c, 0x38, 0xd8, 0x0e, 0xa2, 0x8f, 0xc5, 0x33, 0xf0, 0x62, 0xc8, 0x8e, 0xd3,
	0x0d, 0x3d, 0x79, 0xfe, 0x6f, 0xc6, 0x9e, 0xf9, 0xad, 0x81, 0xfb, 0xa3, 0xf2, 0xa7, 0xa1, 0xde,
	0x37, 0xa6, 0xad, 0x5a, 0xd5, 0x58, 0x53, 0xd1, 0x6f, 0xd1, 0xf6, 0x9a, 0x5c, 0x55, 0x1b, 0xf3,
	0x43, 0x75, 0xc7, 0xca, 0xd9, 0x5f, 0x55, 0x6f, 0xcd, 0xa3, 0xd2, 0x14, 0x4e, 0x6f, 0x26, 0xb5,
	0x8f, 0x0a, 0x8b, 0x24, 0xcb, 0x4f, 0x50, 0x70, 0xfa, 0x39, 0x90, 0xf3, 0x78, 0x03, 0x97, 0x27,
	0xe3, 0x49, 0x3f, 0x48, 0xc7, 0xb2, 0xed, 0x72, 0xb7, 0xe6, 0xcf, 0x1a, 0xdf, 0x42, 0xae, 0x4d,
	0x23, 0x34, 0xb1, 0xc5, 0x36, 0xdb, 0xad, 0x79, 0x52, 0xe5, 0x2d, 0xe4, 0x9c, 0xdc, 0xa0, 0x3d,
	0xbe, 0x87, 0x3c, 0x56, 0x8f, 0x77, 0x37, 0x87, 0xab, 0xfd, 0xd4, 0xf1, 0x6b, 0xc0, 0x3c, 0x65,
	0xcb, 0xbf, 0x19, 0xac, 0x22, 0xc1, 0x2b, 0x58, 0x28, 0xc9, 0xb2, 0xf8, 0xde, 0x42, 0x49, 0x44,
	0xb8, 0xe8, 0x44, 0x3b, 0x75, 0x88, 0x31, 0x6e, 0x61, 0xd3, 0x9f, 0x4c, 0x47, 0xdf, 0x86, 0xb6,
	0x26, 0xcb, 0x96, 0x31, 0x35, 0x47, 0xa1, 0x42, 0x92, 0x6b, 0xac, 0xea, 0xbd, 0x32, 0x1d, 0xbb,
	0x18, 0x2b, 0x66, 0x08, 0x3f, 0x40, 0x21, 0xa4, 0xb4, 0xe4, 0x1c, 0x5b, 0x6d, 0xb3, 0xdd, 0xe6,
	0x70, 0xfd, 0x3c, 0xda, 0xfd, 0xc8, 0xf9, 0x54, 0x10, 0x5c, 0xa8, 0x56, 0x1c, 0xc9, 0xb1, 0xfc,
	0x85, 0x8b, 0x87, 0x80, 0x79, 0xca, 0x96, 0x7f, 0x32, 0x28, 0xd2, 0x65, 0x2c, 0xe1, 0x95, 0xf3,
	0x96, 0xc8, 0xa7, 0x21, 0x47, 0x47, 0xff, 0x31, 0x7c, 0x07, 0x90, 0xf4, 0xd9, 0xe1, 0x8c, 0x04,
	0xef, 0x8d, 0xf2, 0x4f, 0xc9, 0x60, 0x8c, 0xf1, 0x0d, 0xac, 0x9c, 0x17, 0x9e, 0x92, 0xa7, 0x51,
	0x20, 0x83, 0xa2, 0x31, 0x43, 0xe7, 0xed, 0x53, 0x74, 0xb3, 0xe6, 0x93, 0x0c, 0x3d, 0x7a, 0xe3,
	0xbc, 0xd0, 0x9f, 0x8d, 0x24, 0x96, 0x8f, 0x3d, 0xce, 0xa4, 0xbc, 0x83, 0x55, 0x34, 0x81, 0xd7,
	0xb0, 0x1c, 0xac, 0x4e, 0x73, 0x86, 0x30, 0x3c, 0x2a, 0xe9, 0x51, 0x0c, 0xda, 0xc7, 0xd9, 0x2e,
	0xf9, 0x24, 0x0f, 0x1f, 0xa1, 0xf8, 0x3e, 0xfe, 0x00, 0xde, 0xc2, 0xe6, 0x0b, 0xf9, 0xa4, 0x1c,
	0x9e, 0x7f, 0x31, 0x2d, 0xd0, 0xcd, 0xeb, 0x19, 0x09, 0x3b, 0x51, 0xe7, 0x71, 0xd9, 0xee, 0xfe,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x00, 0x45, 0x96, 0xb1, 0x02, 0x00, 0x00,
}
