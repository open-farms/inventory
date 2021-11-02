// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EquipmentServiceClient is the client API for EquipmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EquipmentServiceClient interface {
	CreateEquipment(ctx context.Context, in *CreateEquipmentRequest, opts ...grpc.CallOption) (*CreateEquipmentResponse, error)
	UpdateEquipment(ctx context.Context, in *UpdateEquipmentRequest, opts ...grpc.CallOption) (*UpdateEquipmentResponse, error)
	DeleteEquipment(ctx context.Context, in *DeleteEquipmentRequest, opts ...grpc.CallOption) (*DeleteEquipmentResponse, error)
	GetEquipment(ctx context.Context, in *GetEquipmentRequest, opts ...grpc.CallOption) (*GetEquipmentResponse, error)
	ListEquipment(ctx context.Context, in *ListEquipmentRequest, opts ...grpc.CallOption) (*ListEquipmentResponse, error)
}

type equipmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEquipmentServiceClient(cc grpc.ClientConnInterface) EquipmentServiceClient {
	return &equipmentServiceClient{cc}
}

func (c *equipmentServiceClient) CreateEquipment(ctx context.Context, in *CreateEquipmentRequest, opts ...grpc.CallOption) (*CreateEquipmentResponse, error) {
	out := new(CreateEquipmentResponse)
	err := c.cc.Invoke(ctx, "/api.equipment.service.v1.EquipmentService/CreateEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *equipmentServiceClient) UpdateEquipment(ctx context.Context, in *UpdateEquipmentRequest, opts ...grpc.CallOption) (*UpdateEquipmentResponse, error) {
	out := new(UpdateEquipmentResponse)
	err := c.cc.Invoke(ctx, "/api.equipment.service.v1.EquipmentService/UpdateEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *equipmentServiceClient) DeleteEquipment(ctx context.Context, in *DeleteEquipmentRequest, opts ...grpc.CallOption) (*DeleteEquipmentResponse, error) {
	out := new(DeleteEquipmentResponse)
	err := c.cc.Invoke(ctx, "/api.equipment.service.v1.EquipmentService/DeleteEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *equipmentServiceClient) GetEquipment(ctx context.Context, in *GetEquipmentRequest, opts ...grpc.CallOption) (*GetEquipmentResponse, error) {
	out := new(GetEquipmentResponse)
	err := c.cc.Invoke(ctx, "/api.equipment.service.v1.EquipmentService/GetEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *equipmentServiceClient) ListEquipment(ctx context.Context, in *ListEquipmentRequest, opts ...grpc.CallOption) (*ListEquipmentResponse, error) {
	out := new(ListEquipmentResponse)
	err := c.cc.Invoke(ctx, "/api.equipment.service.v1.EquipmentService/ListEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EquipmentServiceServer is the server API for EquipmentService service.
// All implementations must embed UnimplementedEquipmentServiceServer
// for forward compatibility
type EquipmentServiceServer interface {
	CreateEquipment(context.Context, *CreateEquipmentRequest) (*CreateEquipmentResponse, error)
	UpdateEquipment(context.Context, *UpdateEquipmentRequest) (*UpdateEquipmentResponse, error)
	DeleteEquipment(context.Context, *DeleteEquipmentRequest) (*DeleteEquipmentResponse, error)
	GetEquipment(context.Context, *GetEquipmentRequest) (*GetEquipmentResponse, error)
	ListEquipment(context.Context, *ListEquipmentRequest) (*ListEquipmentResponse, error)
	mustEmbedUnimplementedEquipmentServiceServer()
}

// UnimplementedEquipmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEquipmentServiceServer struct {
}

func (UnimplementedEquipmentServiceServer) CreateEquipment(context.Context, *CreateEquipmentRequest) (*CreateEquipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEquipment not implemented")
}
func (UnimplementedEquipmentServiceServer) UpdateEquipment(context.Context, *UpdateEquipmentRequest) (*UpdateEquipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEquipment not implemented")
}
func (UnimplementedEquipmentServiceServer) DeleteEquipment(context.Context, *DeleteEquipmentRequest) (*DeleteEquipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEquipment not implemented")
}
func (UnimplementedEquipmentServiceServer) GetEquipment(context.Context, *GetEquipmentRequest) (*GetEquipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEquipment not implemented")
}
func (UnimplementedEquipmentServiceServer) ListEquipment(context.Context, *ListEquipmentRequest) (*ListEquipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEquipment not implemented")
}
func (UnimplementedEquipmentServiceServer) mustEmbedUnimplementedEquipmentServiceServer() {}

// UnsafeEquipmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EquipmentServiceServer will
// result in compilation errors.
type UnsafeEquipmentServiceServer interface {
	mustEmbedUnimplementedEquipmentServiceServer()
}

func RegisterEquipmentServiceServer(s grpc.ServiceRegistrar, srv EquipmentServiceServer) {
	s.RegisterService(&EquipmentService_ServiceDesc, srv)
}

func _EquipmentService_CreateEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEquipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EquipmentServiceServer).CreateEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.equipment.service.v1.EquipmentService/CreateEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EquipmentServiceServer).CreateEquipment(ctx, req.(*CreateEquipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EquipmentService_UpdateEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEquipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EquipmentServiceServer).UpdateEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.equipment.service.v1.EquipmentService/UpdateEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EquipmentServiceServer).UpdateEquipment(ctx, req.(*UpdateEquipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EquipmentService_DeleteEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEquipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EquipmentServiceServer).DeleteEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.equipment.service.v1.EquipmentService/DeleteEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EquipmentServiceServer).DeleteEquipment(ctx, req.(*DeleteEquipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EquipmentService_GetEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEquipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EquipmentServiceServer).GetEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.equipment.service.v1.EquipmentService/GetEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EquipmentServiceServer).GetEquipment(ctx, req.(*GetEquipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EquipmentService_ListEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEquipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EquipmentServiceServer).ListEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.equipment.service.v1.EquipmentService/ListEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EquipmentServiceServer).ListEquipment(ctx, req.(*ListEquipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EquipmentService_ServiceDesc is the grpc.ServiceDesc for EquipmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EquipmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.equipment.service.v1.EquipmentService",
	HandlerType: (*EquipmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEquipment",
			Handler:    _EquipmentService_CreateEquipment_Handler,
		},
		{
			MethodName: "UpdateEquipment",
			Handler:    _EquipmentService_UpdateEquipment_Handler,
		},
		{
			MethodName: "DeleteEquipment",
			Handler:    _EquipmentService_DeleteEquipment_Handler,
		},
		{
			MethodName: "GetEquipment",
			Handler:    _EquipmentService_GetEquipment_Handler,
		},
		{
			MethodName: "ListEquipment",
			Handler:    _EquipmentService_ListEquipment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/equipment/service/v1/equipment.proto",
}
