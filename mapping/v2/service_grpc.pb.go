// Copyright © 2021 The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: packetbroker/api/mapping/v2/service.proto

package mappingpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GatewayVisibilityManager_GetDefaultVisibility_FullMethodName     = "/org.packetbroker.mapping.v2.GatewayVisibilityManager/GetDefaultVisibility"
	GatewayVisibilityManager_SetDefaultVisibility_FullMethodName     = "/org.packetbroker.mapping.v2.GatewayVisibilityManager/SetDefaultVisibility"
	GatewayVisibilityManager_GetHomeNetworkVisibility_FullMethodName = "/org.packetbroker.mapping.v2.GatewayVisibilityManager/GetHomeNetworkVisibility"
	GatewayVisibilityManager_SetHomeNetworkVisibility_FullMethodName = "/org.packetbroker.mapping.v2.GatewayVisibilityManager/SetHomeNetworkVisibility"
)

// GatewayVisibilityManagerClient is the client API for GatewayVisibilityManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayVisibilityManagerClient interface {
	// Get the Default Gateway Visibility.
	GetDefaultVisibility(ctx context.Context, in *GetDefaultGatewayVisibilityRequest, opts ...grpc.CallOption) (*GetGatewayVisibilityResponse, error)
	// Set the Default Gateway Visibility.
	SetDefaultVisibility(ctx context.Context, in *SetGatewayVisibilityRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get the Gateway Visibility with the Home Network.
	GetHomeNetworkVisibility(ctx context.Context, in *GetHomeNetworkGatewayVisibilityRequest, opts ...grpc.CallOption) (*GetGatewayVisibilityResponse, error)
	// Set the Gateway Visibility with the Home Network.
	SetHomeNetworkVisibility(ctx context.Context, in *SetGatewayVisibilityRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type gatewayVisibilityManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayVisibilityManagerClient(cc grpc.ClientConnInterface) GatewayVisibilityManagerClient {
	return &gatewayVisibilityManagerClient{cc}
}

func (c *gatewayVisibilityManagerClient) GetDefaultVisibility(ctx context.Context, in *GetDefaultGatewayVisibilityRequest, opts ...grpc.CallOption) (*GetGatewayVisibilityResponse, error) {
	out := new(GetGatewayVisibilityResponse)
	err := c.cc.Invoke(ctx, GatewayVisibilityManager_GetDefaultVisibility_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayVisibilityManagerClient) SetDefaultVisibility(ctx context.Context, in *SetGatewayVisibilityRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GatewayVisibilityManager_SetDefaultVisibility_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayVisibilityManagerClient) GetHomeNetworkVisibility(ctx context.Context, in *GetHomeNetworkGatewayVisibilityRequest, opts ...grpc.CallOption) (*GetGatewayVisibilityResponse, error) {
	out := new(GetGatewayVisibilityResponse)
	err := c.cc.Invoke(ctx, GatewayVisibilityManager_GetHomeNetworkVisibility_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayVisibilityManagerClient) SetHomeNetworkVisibility(ctx context.Context, in *SetGatewayVisibilityRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GatewayVisibilityManager_SetHomeNetworkVisibility_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayVisibilityManagerServer is the server API for GatewayVisibilityManager service.
// All implementations must embed UnimplementedGatewayVisibilityManagerServer
// for forward compatibility
type GatewayVisibilityManagerServer interface {
	// Get the Default Gateway Visibility.
	GetDefaultVisibility(context.Context, *GetDefaultGatewayVisibilityRequest) (*GetGatewayVisibilityResponse, error)
	// Set the Default Gateway Visibility.
	SetDefaultVisibility(context.Context, *SetGatewayVisibilityRequest) (*emptypb.Empty, error)
	// Get the Gateway Visibility with the Home Network.
	GetHomeNetworkVisibility(context.Context, *GetHomeNetworkGatewayVisibilityRequest) (*GetGatewayVisibilityResponse, error)
	// Set the Gateway Visibility with the Home Network.
	SetHomeNetworkVisibility(context.Context, *SetGatewayVisibilityRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedGatewayVisibilityManagerServer()
}

// UnimplementedGatewayVisibilityManagerServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayVisibilityManagerServer struct {
}

func (UnimplementedGatewayVisibilityManagerServer) GetDefaultVisibility(context.Context, *GetDefaultGatewayVisibilityRequest) (*GetGatewayVisibilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultVisibility not implemented")
}
func (UnimplementedGatewayVisibilityManagerServer) SetDefaultVisibility(context.Context, *SetGatewayVisibilityRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDefaultVisibility not implemented")
}
func (UnimplementedGatewayVisibilityManagerServer) GetHomeNetworkVisibility(context.Context, *GetHomeNetworkGatewayVisibilityRequest) (*GetGatewayVisibilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHomeNetworkVisibility not implemented")
}
func (UnimplementedGatewayVisibilityManagerServer) SetHomeNetworkVisibility(context.Context, *SetGatewayVisibilityRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHomeNetworkVisibility not implemented")
}
func (UnimplementedGatewayVisibilityManagerServer) mustEmbedUnimplementedGatewayVisibilityManagerServer() {
}

// UnsafeGatewayVisibilityManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayVisibilityManagerServer will
// result in compilation errors.
type UnsafeGatewayVisibilityManagerServer interface {
	mustEmbedUnimplementedGatewayVisibilityManagerServer()
}

func RegisterGatewayVisibilityManagerServer(s grpc.ServiceRegistrar, srv GatewayVisibilityManagerServer) {
	s.RegisterService(&GatewayVisibilityManager_ServiceDesc, srv)
}

func _GatewayVisibilityManager_GetDefaultVisibility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefaultGatewayVisibilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayVisibilityManagerServer).GetDefaultVisibility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayVisibilityManager_GetDefaultVisibility_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayVisibilityManagerServer).GetDefaultVisibility(ctx, req.(*GetDefaultGatewayVisibilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayVisibilityManager_SetDefaultVisibility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGatewayVisibilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayVisibilityManagerServer).SetDefaultVisibility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayVisibilityManager_SetDefaultVisibility_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayVisibilityManagerServer).SetDefaultVisibility(ctx, req.(*SetGatewayVisibilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayVisibilityManager_GetHomeNetworkVisibility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHomeNetworkGatewayVisibilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayVisibilityManagerServer).GetHomeNetworkVisibility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayVisibilityManager_GetHomeNetworkVisibility_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayVisibilityManagerServer).GetHomeNetworkVisibility(ctx, req.(*GetHomeNetworkGatewayVisibilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayVisibilityManager_SetHomeNetworkVisibility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGatewayVisibilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayVisibilityManagerServer).SetHomeNetworkVisibility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayVisibilityManager_SetHomeNetworkVisibility_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayVisibilityManagerServer).SetHomeNetworkVisibility(ctx, req.(*SetGatewayVisibilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayVisibilityManager_ServiceDesc is the grpc.ServiceDesc for GatewayVisibilityManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayVisibilityManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "org.packetbroker.mapping.v2.GatewayVisibilityManager",
	HandlerType: (*GatewayVisibilityManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDefaultVisibility",
			Handler:    _GatewayVisibilityManager_GetDefaultVisibility_Handler,
		},
		{
			MethodName: "SetDefaultVisibility",
			Handler:    _GatewayVisibilityManager_SetDefaultVisibility_Handler,
		},
		{
			MethodName: "GetHomeNetworkVisibility",
			Handler:    _GatewayVisibilityManager_GetHomeNetworkVisibility_Handler,
		},
		{
			MethodName: "SetHomeNetworkVisibility",
			Handler:    _GatewayVisibilityManager_SetHomeNetworkVisibility_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "packetbroker/api/mapping/v2/service.proto",
}

const (
	Mapper_UpdateGateway_FullMethodName = "/org.packetbroker.mapping.v2.Mapper/UpdateGateway"
)

// MapperClient is the client API for Mapper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MapperClient interface {
	// Update a gateway.
	UpdateGateway(ctx context.Context, in *UpdateGatewayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type mapperClient struct {
	cc grpc.ClientConnInterface
}

func NewMapperClient(cc grpc.ClientConnInterface) MapperClient {
	return &mapperClient{cc}
}

func (c *mapperClient) UpdateGateway(ctx context.Context, in *UpdateGatewayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Mapper_UpdateGateway_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MapperServer is the server API for Mapper service.
// All implementations must embed UnimplementedMapperServer
// for forward compatibility
type MapperServer interface {
	// Update a gateway.
	UpdateGateway(context.Context, *UpdateGatewayRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMapperServer()
}

// UnimplementedMapperServer must be embedded to have forward compatible implementations.
type UnimplementedMapperServer struct {
}

func (UnimplementedMapperServer) UpdateGateway(context.Context, *UpdateGatewayRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGateway not implemented")
}
func (UnimplementedMapperServer) mustEmbedUnimplementedMapperServer() {}

// UnsafeMapperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MapperServer will
// result in compilation errors.
type UnsafeMapperServer interface {
	mustEmbedUnimplementedMapperServer()
}

func RegisterMapperServer(s grpc.ServiceRegistrar, srv MapperServer) {
	s.RegisterService(&Mapper_ServiceDesc, srv)
}

func _Mapper_UpdateGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapperServer).UpdateGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mapper_UpdateGateway_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapperServer).UpdateGateway(ctx, req.(*UpdateGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mapper_ServiceDesc is the grpc.ServiceDesc for Mapper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mapper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "org.packetbroker.mapping.v2.Mapper",
	HandlerType: (*MapperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateGateway",
			Handler:    _Mapper_UpdateGateway_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "packetbroker/api/mapping/v2/service.proto",
}
