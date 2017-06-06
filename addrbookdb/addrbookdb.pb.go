// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addrbookdb.proto

/*
Package addrbookdb is a generated protocol buffer package.

It is generated from these files:
	addrbookdb.proto

It has these top-level messages:
	PersonRequestDB
	PersonReplyDB
	PeopleReplyDB
*/
package addrbookdb

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
type PersonRequestDB struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *PersonRequestDB) Reset()                    { *m = PersonRequestDB{} }
func (m *PersonRequestDB) String() string            { return proto.CompactTextString(m) }
func (*PersonRequestDB) ProtoMessage()               {}
func (*PersonRequestDB) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PersonRequestDB) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// The request message containing the user's id.
type PersonReplyDB struct {
	Id        int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Firstname string `protobuf:"bytes,2,opt,name=firstname" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,3,opt,name=lastname" json:"lastname,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Phoneno   string `protobuf:"bytes,5,opt,name=phoneno" json:"phoneno,omitempty"`
}

func (m *PersonReplyDB) Reset()                    { *m = PersonReplyDB{} }
func (m *PersonReplyDB) String() string            { return proto.CompactTextString(m) }
func (*PersonReplyDB) ProtoMessage()               {}
func (*PersonReplyDB) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PersonReplyDB) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PersonReplyDB) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *PersonReplyDB) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *PersonReplyDB) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *PersonReplyDB) GetPhoneno() string {
	if m != nil {
		return m.Phoneno
	}
	return ""
}

// The response message containing the list of persons
type PeopleReplyDB struct {
	People []*PersonReplyDB `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *PeopleReplyDB) Reset()                    { *m = PeopleReplyDB{} }
func (m *PeopleReplyDB) String() string            { return proto.CompactTextString(m) }
func (*PeopleReplyDB) ProtoMessage()               {}
func (*PeopleReplyDB) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PeopleReplyDB) GetPeople() []*PersonReplyDB {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*PersonRequestDB)(nil), "addrbookdb.PersonRequestDB")
	proto.RegisterType((*PersonReplyDB)(nil), "addrbookdb.PersonReplyDB")
	proto.RegisterType((*PeopleReplyDB)(nil), "addrbookdb.PeopleReplyDB")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AddrBookDB service

type AddrBookDBClient interface {
	// Gets a person details
	ReadPersonDB(ctx context.Context, in *PersonRequestDB, opts ...grpc.CallOption) (*PersonReplyDB, error)
	// Creats a person details
	CreatePersonDB(ctx context.Context, in *PersonReplyDB, opts ...grpc.CallOption) (*PersonReplyDB, error)
	// Update a person details
	UpdatePersonDB(ctx context.Context, in *PersonReplyDB, opts ...grpc.CallOption) (*PersonReplyDB, error)
	// Delete a person details
	DeletePersonDB(ctx context.Context, in *PersonRequestDB, opts ...grpc.CallOption) (*PersonReplyDB, error)
}

type addrBookDBClient struct {
	cc *grpc.ClientConn
}

func NewAddrBookDBClient(cc *grpc.ClientConn) AddrBookDBClient {
	return &addrBookDBClient{cc}
}

func (c *addrBookDBClient) ReadPersonDB(ctx context.Context, in *PersonRequestDB, opts ...grpc.CallOption) (*PersonReplyDB, error) {
	out := new(PersonReplyDB)
	err := grpc.Invoke(ctx, "/addrbookdb.AddrBookDB/ReadPersonDB", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addrBookDBClient) CreatePersonDB(ctx context.Context, in *PersonReplyDB, opts ...grpc.CallOption) (*PersonReplyDB, error) {
	out := new(PersonReplyDB)
	err := grpc.Invoke(ctx, "/addrbookdb.AddrBookDB/CreatePersonDB", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addrBookDBClient) UpdatePersonDB(ctx context.Context, in *PersonReplyDB, opts ...grpc.CallOption) (*PersonReplyDB, error) {
	out := new(PersonReplyDB)
	err := grpc.Invoke(ctx, "/addrbookdb.AddrBookDB/UpdatePersonDB", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addrBookDBClient) DeletePersonDB(ctx context.Context, in *PersonRequestDB, opts ...grpc.CallOption) (*PersonReplyDB, error) {
	out := new(PersonReplyDB)
	err := grpc.Invoke(ctx, "/addrbookdb.AddrBookDB/DeletePersonDB", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AddrBookDB service

type AddrBookDBServer interface {
	// Gets a person details
	ReadPersonDB(context.Context, *PersonRequestDB) (*PersonReplyDB, error)
	// Creats a person details
	CreatePersonDB(context.Context, *PersonReplyDB) (*PersonReplyDB, error)
	// Update a person details
	UpdatePersonDB(context.Context, *PersonReplyDB) (*PersonReplyDB, error)
	// Delete a person details
	DeletePersonDB(context.Context, *PersonRequestDB) (*PersonReplyDB, error)
}

func RegisterAddrBookDBServer(s *grpc.Server, srv AddrBookDBServer) {
	s.RegisterService(&_AddrBookDB_serviceDesc, srv)
}

func _AddrBookDB_ReadPersonDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonRequestDB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddrBookDBServer).ReadPersonDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addrbookdb.AddrBookDB/ReadPersonDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddrBookDBServer).ReadPersonDB(ctx, req.(*PersonRequestDB))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddrBookDB_CreatePersonDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonReplyDB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddrBookDBServer).CreatePersonDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addrbookdb.AddrBookDB/CreatePersonDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddrBookDBServer).CreatePersonDB(ctx, req.(*PersonReplyDB))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddrBookDB_UpdatePersonDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonReplyDB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddrBookDBServer).UpdatePersonDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addrbookdb.AddrBookDB/UpdatePersonDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddrBookDBServer).UpdatePersonDB(ctx, req.(*PersonReplyDB))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddrBookDB_DeletePersonDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonRequestDB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddrBookDBServer).DeletePersonDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addrbookdb.AddrBookDB/DeletePersonDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddrBookDBServer).DeletePersonDB(ctx, req.(*PersonRequestDB))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddrBookDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "addrbookdb.AddrBookDB",
	HandlerType: (*AddrBookDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadPersonDB",
			Handler:    _AddrBookDB_ReadPersonDB_Handler,
		},
		{
			MethodName: "CreatePersonDB",
			Handler:    _AddrBookDB_CreatePersonDB_Handler,
		},
		{
			MethodName: "UpdatePersonDB",
			Handler:    _AddrBookDB_UpdatePersonDB_Handler,
		},
		{
			MethodName: "DeletePersonDB",
			Handler:    _AddrBookDB_DeletePersonDB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "addrbookdb.proto",
}

func init() { proto.RegisterFile("addrbookdb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0x5d, 0x6a, 0xab, 0x1d, 0x95, 0x9a, 0x8d, 0x87, 0xb5, 0x7a, 0x40, 0x2e, 0x72, 0x22,
	0xb1, 0x3e, 0x81, 0x2b, 0x87, 0xc6, 0x13, 0x21, 0xf1, 0x01, 0x16, 0x77, 0x54, 0x52, 0xca, 0xac,
	0x0b, 0x1e, 0x7c, 0x04, 0x5f, 0xc8, 0xe7, 0x33, 0xa5, 0x50, 0x30, 0xda, 0x98, 0x18, 0x8f, 0xf3,
	0xcf, 0x97, 0x2f, 0x33, 0x93, 0x81, 0x63, 0xa5, 0xb5, 0x4d, 0x89, 0x16, 0x3a, 0x0d, 0x8d, 0xa5,
	0x8a, 0x38, 0x74, 0x89, 0x7f, 0x01, 0x93, 0x18, 0x6d, 0x49, 0x45, 0x82, 0x2f, 0xaf, 0x58, 0x56,
	0x91, 0xe4, 0x2e, 0x38, 0x99, 0x16, 0xcc, 0x63, 0xc1, 0x20, 0x71, 0x32, 0xed, 0xbf, 0x33, 0x38,
	0x6a, 0x19, 0x93, 0xbf, 0x7d, 0x27, 0xf8, 0x39, 0x8c, 0x1f, 0x33, 0x5b, 0x56, 0x85, 0x5a, 0xa2,
	0x70, 0x3c, 0x16, 0x8c, 0x93, 0x2e, 0xe0, 0x53, 0xd8, 0xcf, 0x55, 0xd3, 0x1c, 0xd4, 0xcd, 0x4d,
	0xcd, 0x4f, 0x60, 0x88, 0x4b, 0x95, 0xe5, 0x62, 0xb7, 0x6e, 0xac, 0x0b, 0x2e, 0x60, 0xcf, 0x3c,
	0x53, 0x81, 0x05, 0x89, 0x61, 0x9d, 0xb7, 0xa5, 0x2f, 0x57, 0xa3, 0x90, 0xc9, 0xb1, 0x1d, 0xe5,
	0x0a, 0x46, 0xa6, 0x0e, 0x04, 0xf3, 0x06, 0xc1, 0xc1, 0xec, 0x34, 0xec, 0xad, 0xfb, 0x65, 0xea,
	0xa4, 0x01, 0x67, 0x1f, 0x0e, 0xc0, 0x8d, 0xd6, 0x56, 0x12, 0x2d, 0x22, 0xc9, 0xe7, 0x70, 0x98,
	0xa0, 0xd2, 0x6b, 0x36, 0x92, 0xfc, 0xec, 0x27, 0x43, 0x73, 0x9b, 0xe9, 0x76, 0xbd, 0xbf, 0xc3,
	0xe7, 0xe0, 0xde, 0x5a, 0x54, 0x15, 0x6e, 0x5c, 0xdb, 0xf1, 0x5f, 0x4d, 0xf7, 0x46, 0xff, 0x87,
	0xe9, 0x0e, 0xdc, 0x08, 0x73, 0xec, 0x99, 0xfe, 0xbc, 0x9f, 0xbc, 0x04, 0x9e, 0x51, 0xf8, 0x64,
	0xcd, 0x43, 0x8f, 0x92, 0x93, 0xee, 0x96, 0xf1, 0xea, 0xbd, 0x62, 0x96, 0x8e, 0xea, 0x3f, 0xbb,
	0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xf6, 0xc6, 0x6c, 0x7b, 0x02, 0x00, 0x00,
}
