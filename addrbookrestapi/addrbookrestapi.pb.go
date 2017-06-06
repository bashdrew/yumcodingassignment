// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addrbookrestapi.proto

/*
Package addrbookrestapi is a generated protocol buffer package.

It is generated from these files:
	addrbookrestapi.proto

It has these top-level messages:
	PersonRequest
	PersonReply
	PeopleReply
*/
package addrbookrestapi

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

// The request message containing the user's id.
type PersonRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *PersonRequest) Reset()                    { *m = PersonRequest{} }
func (m *PersonRequest) String() string            { return proto.CompactTextString(m) }
func (*PersonRequest) ProtoMessage()               {}
func (*PersonRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PersonRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// The request message containing the user's id.
type PersonReply struct {
	Id        int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Firstname string `protobuf:"bytes,2,opt,name=firstname" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,3,opt,name=lastname" json:"lastname,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Phoneno   string `protobuf:"bytes,5,opt,name=phoneno" json:"phoneno,omitempty"`
}

func (m *PersonReply) Reset()                    { *m = PersonReply{} }
func (m *PersonReply) String() string            { return proto.CompactTextString(m) }
func (*PersonReply) ProtoMessage()               {}
func (*PersonReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PersonReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PersonReply) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *PersonReply) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *PersonReply) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *PersonReply) GetPhoneno() string {
	if m != nil {
		return m.Phoneno
	}
	return ""
}

// The response message containing the list of persons
type PeopleReply struct {
	People []*PersonReply `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *PeopleReply) Reset()                    { *m = PeopleReply{} }
func (m *PeopleReply) String() string            { return proto.CompactTextString(m) }
func (*PeopleReply) ProtoMessage()               {}
func (*PeopleReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PeopleReply) GetPeople() []*PersonReply {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*PersonRequest)(nil), "addrbookrestapi.PersonRequest")
	proto.RegisterType((*PersonReply)(nil), "addrbookrestapi.PersonReply")
	proto.RegisterType((*PeopleReply)(nil), "addrbookrestapi.PeopleReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AddrBookRestAPI service

type AddrBookRestAPIClient interface {
	// Gets a person details
	GetPerson(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*PersonReply, error)
	// Creates a person details
	PostPerson(ctx context.Context, in *PersonReply, opts ...grpc.CallOption) (*PersonReply, error)
	// Updates a person details
	PutPerson(ctx context.Context, in *PersonReply, opts ...grpc.CallOption) (*PersonReply, error)
	// Deletes a person details
	DeletePerson(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*PersonReply, error)
}

type addrBookRestAPIClient struct {
	cc *grpc.ClientConn
}

func NewAddrBookRestAPIClient(cc *grpc.ClientConn) AddrBookRestAPIClient {
	return &addrBookRestAPIClient{cc}
}

func (c *addrBookRestAPIClient) GetPerson(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*PersonReply, error) {
	out := new(PersonReply)
	err := grpc.Invoke(ctx, "/addrbookrestapi.AddrBookRestAPI/GetPerson", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addrBookRestAPIClient) PostPerson(ctx context.Context, in *PersonReply, opts ...grpc.CallOption) (*PersonReply, error) {
	out := new(PersonReply)
	err := grpc.Invoke(ctx, "/addrbookrestapi.AddrBookRestAPI/PostPerson", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addrBookRestAPIClient) PutPerson(ctx context.Context, in *PersonReply, opts ...grpc.CallOption) (*PersonReply, error) {
	out := new(PersonReply)
	err := grpc.Invoke(ctx, "/addrbookrestapi.AddrBookRestAPI/PutPerson", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addrBookRestAPIClient) DeletePerson(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*PersonReply, error) {
	out := new(PersonReply)
	err := grpc.Invoke(ctx, "/addrbookrestapi.AddrBookRestAPI/DeletePerson", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AddrBookRestAPI service

type AddrBookRestAPIServer interface {
	// Gets a person details
	GetPerson(context.Context, *PersonRequest) (*PersonReply, error)
	// Creates a person details
	PostPerson(context.Context, *PersonReply) (*PersonReply, error)
	// Updates a person details
	PutPerson(context.Context, *PersonReply) (*PersonReply, error)
	// Deletes a person details
	DeletePerson(context.Context, *PersonRequest) (*PersonReply, error)
}

func RegisterAddrBookRestAPIServer(s *grpc.Server, srv AddrBookRestAPIServer) {
	s.RegisterService(&_AddrBookRestAPI_serviceDesc, srv)
}

func _AddrBookRestAPI_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddrBookRestAPIServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addrbookrestapi.AddrBookRestAPI/GetPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddrBookRestAPIServer).GetPerson(ctx, req.(*PersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddrBookRestAPI_PostPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonReply)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddrBookRestAPIServer).PostPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addrbookrestapi.AddrBookRestAPI/PostPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddrBookRestAPIServer).PostPerson(ctx, req.(*PersonReply))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddrBookRestAPI_PutPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonReply)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddrBookRestAPIServer).PutPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addrbookrestapi.AddrBookRestAPI/PutPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddrBookRestAPIServer).PutPerson(ctx, req.(*PersonReply))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddrBookRestAPI_DeletePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddrBookRestAPIServer).DeletePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addrbookrestapi.AddrBookRestAPI/DeletePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddrBookRestAPIServer).DeletePerson(ctx, req.(*PersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddrBookRestAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "addrbookrestapi.AddrBookRestAPI",
	HandlerType: (*AddrBookRestAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPerson",
			Handler:    _AddrBookRestAPI_GetPerson_Handler,
		},
		{
			MethodName: "PostPerson",
			Handler:    _AddrBookRestAPI_PostPerson_Handler,
		},
		{
			MethodName: "PutPerson",
			Handler:    _AddrBookRestAPI_PutPerson_Handler,
		},
		{
			MethodName: "DeletePerson",
			Handler:    _AddrBookRestAPI_DeletePerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "addrbookrestapi.proto",
}

func init() { proto.RegisterFile("addrbookrestapi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x52, 0xdd, 0x4a, 0x02, 0x41,
	0x14, 0x6e, 0xd6, 0xb4, 0xf6, 0x58, 0x09, 0x83, 0xd1, 0x20, 0x52, 0xb2, 0x57, 0x5e, 0x09, 0x59,
	0x2f, 0xa0, 0x05, 0x61, 0x41, 0x0c, 0xfb, 0x06, 0x6b, 0x73, 0xaa, 0xc1, 0x75, 0xcf, 0x34, 0x33,
	0x5e, 0xf4, 0x04, 0x3d, 0x59, 0xef, 0x15, 0xce, 0xba, 0x15, 0x0b, 0xd9, 0x8d, 0x97, 0xdf, 0x0f,
	0xdf, 0xf9, 0xce, 0xe1, 0xc0, 0x69, 0xa6, 0x94, 0x9d, 0x13, 0x2d, 0x2c, 0x3a, 0x9f, 0x19, 0x3d,
	0x32, 0x96, 0x3c, 0xf1, 0x4e, 0x8d, 0x4e, 0x2e, 0xe0, 0x58, 0xa2, 0x75, 0x54, 0xa4, 0xf8, 0xb6,
	0x42, 0xe7, 0xf9, 0x09, 0x44, 0x5a, 0x09, 0x36, 0x60, 0xc3, 0x46, 0x1a, 0x69, 0x95, 0x7c, 0x30,
	0x68, 0x57, 0x0e, 0x93, 0xbf, 0xd7, 0x75, 0xde, 0x87, 0xf8, 0x59, 0x5b, 0xe7, 0x8b, 0x6c, 0x89,
	0x22, 0x1a, 0xb0, 0x61, 0x9c, 0xfe, 0x10, 0xbc, 0x07, 0x87, 0x79, 0xb6, 0x11, 0x1b, 0x41, 0xfc,
	0xc6, 0xbc, 0x0b, 0x4d, 0x5c, 0x66, 0x3a, 0x17, 0xfb, 0x41, 0x28, 0x01, 0x17, 0x70, 0x60, 0x5e,
	0xa9, 0xc0, 0x82, 0x44, 0x33, 0xf0, 0x15, 0x4c, 0x6e, 0xd6, 0x45, 0xc8, 0xe4, 0x58, 0x16, 0xb9,
	0x86, 0x96, 0x09, 0x50, 0xb0, 0x41, 0x63, 0xd8, 0x1e, 0xf7, 0x47, 0xf5, 0x95, 0x7f, 0xd5, 0x4e,
	0x37, 0xde, 0xf1, 0x67, 0x04, 0x9d, 0x89, 0x52, 0x76, 0x4a, 0xb4, 0x48, 0xd1, 0xf9, 0x89, 0x9c,
	0xf1, 0x07, 0x88, 0xef, 0xd0, 0x97, 0x6e, 0x7e, 0xfe, 0x67, 0x4c, 0xb8, 0x4f, 0x6f, 0xeb, 0x98,
	0x64, 0x8f, 0xdf, 0x03, 0x48, 0x72, 0x55, 0xda, 0x56, 0xf7, 0xbf, 0x59, 0x33, 0x88, 0xe5, 0x6a,
	0x37, 0x51, 0x8f, 0x70, 0x74, 0x8b, 0x39, 0x7a, 0xdc, 0xcd, 0x9a, 0xd3, 0x4b, 0x38, 0xd3, 0x34,
	0x7a, 0xb1, 0xe6, 0xa9, 0x6e, 0x9c, 0x76, 0x6b, 0xf7, 0x95, 0xeb, 0xcf, 0x93, 0x6c, 0xde, 0x0a,
	0x2f, 0x78, 0xf5, 0x15, 0x00, 0x00, 0xff, 0xff, 0x72, 0x36, 0x71, 0x6f, 0x9b, 0x02, 0x00, 0x00,
}
