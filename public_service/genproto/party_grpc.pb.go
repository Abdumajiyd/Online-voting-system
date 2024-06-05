// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: protos/party.proto

package public

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	PartyService_Create_FullMethodName  = "/protos.PartyService/Create"
	PartyService_Update_FullMethodName  = "/protos.PartyService/Update"
	PartyService_Delete_FullMethodName  = "/protos.PartyService/Delete"
	PartyService_GetById_FullMethodName = "/protos.PartyService/GetById"
	PartyService_GetAll_FullMethodName  = "/protos.PartyService/GetAll"
)

// PartyServiceClient is the client API for PartyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartyServiceClient interface {
	Create(ctx context.Context, in *PartyCreate, opts ...grpc.CallOption) (*Void, error)
	Update(ctx context.Context, in *PartyUpdate, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *PartyDelete, opts ...grpc.CallOption) (*Void, error)
	GetById(ctx context.Context, in *PartyById, opts ...grpc.CallOption) (*Party, error)
	GetAll(ctx context.Context, in *GetAllPartyRequest, opts ...grpc.CallOption) (*GetAllPartyResponse, error)
}

type partyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartyServiceClient(cc grpc.ClientConnInterface) PartyServiceClient {
	return &partyServiceClient{cc}
}

func (c *partyServiceClient) Create(ctx context.Context, in *PartyCreate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, PartyService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) Update(ctx context.Context, in *PartyUpdate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, PartyService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) Delete(ctx context.Context, in *PartyDelete, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, PartyService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) GetById(ctx context.Context, in *PartyById, opts ...grpc.CallOption) (*Party, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Party)
	err := c.cc.Invoke(ctx, PartyService_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) GetAll(ctx context.Context, in *GetAllPartyRequest, opts ...grpc.CallOption) (*GetAllPartyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllPartyResponse)
	err := c.cc.Invoke(ctx, PartyService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartyServiceServer is the server API for PartyService service.
// All implementations must embed UnimplementedPartyServiceServer
// for forward compatibility
type PartyServiceServer interface {
	Create(context.Context, *PartyCreate) (*Void, error)
	Update(context.Context, *PartyUpdate) (*Void, error)
	Delete(context.Context, *PartyDelete) (*Void, error)
	GetById(context.Context, *PartyById) (*Party, error)
	GetAll(context.Context, *GetAllPartyRequest) (*GetAllPartyResponse, error)
	mustEmbedUnimplementedPartyServiceServer()
}

// UnimplementedPartyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPartyServiceServer struct {
}

func (UnimplementedPartyServiceServer) Create(context.Context, *PartyCreate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPartyServiceServer) Update(context.Context, *PartyUpdate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPartyServiceServer) Delete(context.Context, *PartyDelete) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPartyServiceServer) GetById(context.Context, *PartyById) (*Party, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedPartyServiceServer) GetAll(context.Context, *GetAllPartyRequest) (*GetAllPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedPartyServiceServer) mustEmbedUnimplementedPartyServiceServer() {}

// UnsafePartyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartyServiceServer will
// result in compilation errors.
type UnsafePartyServiceServer interface {
	mustEmbedUnimplementedPartyServiceServer()
}

func RegisterPartyServiceServer(s grpc.ServiceRegistrar, srv PartyServiceServer) {
	s.RegisterService(&PartyService_ServiceDesc, srv)
}

func _PartyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartyCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).Create(ctx, req.(*PartyCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartyUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).Update(ctx, req.(*PartyUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartyDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).Delete(ctx, req.(*PartyDelete))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartyById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).GetById(ctx, req.(*PartyById))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartyService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).GetAll(ctx, req.(*GetAllPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PartyService_ServiceDesc is the grpc.ServiceDesc for PartyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.PartyService",
	HandlerType: (*PartyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PartyService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PartyService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PartyService_Delete_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _PartyService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PartyService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/party.proto",
}