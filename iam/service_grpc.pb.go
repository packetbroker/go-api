// Copyright © 2020 The Things Industries B.V.
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
// source: packetbroker/api/iam/v1/service.proto

package iampb

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
	NetworkRegistry_ListNetworks_FullMethodName        = "/org.packetbroker.iam.v1.NetworkRegistry/ListNetworks"
	NetworkRegistry_CreateNetwork_FullMethodName       = "/org.packetbroker.iam.v1.NetworkRegistry/CreateNetwork"
	NetworkRegistry_GetNetwork_FullMethodName          = "/org.packetbroker.iam.v1.NetworkRegistry/GetNetwork"
	NetworkRegistry_UpdateNetwork_FullMethodName       = "/org.packetbroker.iam.v1.NetworkRegistry/UpdateNetwork"
	NetworkRegistry_UpdateNetworkListed_FullMethodName = "/org.packetbroker.iam.v1.NetworkRegistry/UpdateNetworkListed"
	NetworkRegistry_DeleteNetwork_FullMethodName       = "/org.packetbroker.iam.v1.NetworkRegistry/DeleteNetwork"
)

// NetworkRegistryClient is the client API for NetworkRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkRegistryClient interface {
	// List networks.
	ListNetworks(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error)
	// Create a network.
	CreateNetwork(ctx context.Context, in *CreateNetworkRequest, opts ...grpc.CallOption) (*CreateNetworkResponse, error)
	// Get a network.
	GetNetwork(ctx context.Context, in *NetworkRequest, opts ...grpc.CallOption) (*GetNetworkResponse, error)
	// Update a network.
	UpdateNetwork(ctx context.Context, in *UpdateNetworkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Update the listed state of a network.
	UpdateNetworkListed(ctx context.Context, in *UpdateNetworkListedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete a network.
	DeleteNetwork(ctx context.Context, in *NetworkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type networkRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkRegistryClient(cc grpc.ClientConnInterface) NetworkRegistryClient {
	return &networkRegistryClient{cc}
}

func (c *networkRegistryClient) ListNetworks(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error) {
	out := new(ListNetworksResponse)
	err := c.cc.Invoke(ctx, NetworkRegistry_ListNetworks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkRegistryClient) CreateNetwork(ctx context.Context, in *CreateNetworkRequest, opts ...grpc.CallOption) (*CreateNetworkResponse, error) {
	out := new(CreateNetworkResponse)
	err := c.cc.Invoke(ctx, NetworkRegistry_CreateNetwork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkRegistryClient) GetNetwork(ctx context.Context, in *NetworkRequest, opts ...grpc.CallOption) (*GetNetworkResponse, error) {
	out := new(GetNetworkResponse)
	err := c.cc.Invoke(ctx, NetworkRegistry_GetNetwork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkRegistryClient) UpdateNetwork(ctx context.Context, in *UpdateNetworkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NetworkRegistry_UpdateNetwork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkRegistryClient) UpdateNetworkListed(ctx context.Context, in *UpdateNetworkListedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NetworkRegistry_UpdateNetworkListed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkRegistryClient) DeleteNetwork(ctx context.Context, in *NetworkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NetworkRegistry_DeleteNetwork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkRegistryServer is the server API for NetworkRegistry service.
// All implementations must embed UnimplementedNetworkRegistryServer
// for forward compatibility
type NetworkRegistryServer interface {
	// List networks.
	ListNetworks(context.Context, *ListNetworksRequest) (*ListNetworksResponse, error)
	// Create a network.
	CreateNetwork(context.Context, *CreateNetworkRequest) (*CreateNetworkResponse, error)
	// Get a network.
	GetNetwork(context.Context, *NetworkRequest) (*GetNetworkResponse, error)
	// Update a network.
	UpdateNetwork(context.Context, *UpdateNetworkRequest) (*emptypb.Empty, error)
	// Update the listed state of a network.
	UpdateNetworkListed(context.Context, *UpdateNetworkListedRequest) (*emptypb.Empty, error)
	// Delete a network.
	DeleteNetwork(context.Context, *NetworkRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedNetworkRegistryServer()
}

// UnimplementedNetworkRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkRegistryServer struct {
}

func (UnimplementedNetworkRegistryServer) ListNetworks(context.Context, *ListNetworksRequest) (*ListNetworksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNetworks not implemented")
}
func (UnimplementedNetworkRegistryServer) CreateNetwork(context.Context, *CreateNetworkRequest) (*CreateNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNetwork not implemented")
}
func (UnimplementedNetworkRegistryServer) GetNetwork(context.Context, *NetworkRequest) (*GetNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetwork not implemented")
}
func (UnimplementedNetworkRegistryServer) UpdateNetwork(context.Context, *UpdateNetworkRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNetwork not implemented")
}
func (UnimplementedNetworkRegistryServer) UpdateNetworkListed(context.Context, *UpdateNetworkListedRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNetworkListed not implemented")
}
func (UnimplementedNetworkRegistryServer) DeleteNetwork(context.Context, *NetworkRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNetwork not implemented")
}
func (UnimplementedNetworkRegistryServer) mustEmbedUnimplementedNetworkRegistryServer() {}

// UnsafeNetworkRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkRegistryServer will
// result in compilation errors.
type UnsafeNetworkRegistryServer interface {
	mustEmbedUnimplementedNetworkRegistryServer()
}

func RegisterNetworkRegistryServer(s grpc.ServiceRegistrar, srv NetworkRegistryServer) {
	s.RegisterService(&NetworkRegistry_ServiceDesc, srv)
}

func _NetworkRegistry_ListNetworks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkRegistryServer).ListNetworks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkRegistry_ListNetworks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkRegistryServer).ListNetworks(ctx, req.(*ListNetworksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkRegistry_CreateNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkRegistryServer).CreateNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkRegistry_CreateNetwork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkRegistryServer).CreateNetwork(ctx, req.(*CreateNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkRegistry_GetNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkRegistryServer).GetNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkRegistry_GetNetwork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkRegistryServer).GetNetwork(ctx, req.(*NetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkRegistry_UpdateNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkRegistryServer).UpdateNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkRegistry_UpdateNetwork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkRegistryServer).UpdateNetwork(ctx, req.(*UpdateNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkRegistry_UpdateNetworkListed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNetworkListedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkRegistryServer).UpdateNetworkListed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkRegistry_UpdateNetworkListed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkRegistryServer).UpdateNetworkListed(ctx, req.(*UpdateNetworkListedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkRegistry_DeleteNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkRegistryServer).DeleteNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkRegistry_DeleteNetwork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkRegistryServer).DeleteNetwork(ctx, req.(*NetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkRegistry_ServiceDesc is the grpc.ServiceDesc for NetworkRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "org.packetbroker.iam.v1.NetworkRegistry",
	HandlerType: (*NetworkRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNetworks",
			Handler:    _NetworkRegistry_ListNetworks_Handler,
		},
		{
			MethodName: "CreateNetwork",
			Handler:    _NetworkRegistry_CreateNetwork_Handler,
		},
		{
			MethodName: "GetNetwork",
			Handler:    _NetworkRegistry_GetNetwork_Handler,
		},
		{
			MethodName: "UpdateNetwork",
			Handler:    _NetworkRegistry_UpdateNetwork_Handler,
		},
		{
			MethodName: "UpdateNetworkListed",
			Handler:    _NetworkRegistry_UpdateNetworkListed_Handler,
		},
		{
			MethodName: "DeleteNetwork",
			Handler:    _NetworkRegistry_DeleteNetwork_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "packetbroker/api/iam/v1/service.proto",
}

const (
	TenantRegistry_ListTenants_FullMethodName        = "/org.packetbroker.iam.v1.TenantRegistry/ListTenants"
	TenantRegistry_CreateTenant_FullMethodName       = "/org.packetbroker.iam.v1.TenantRegistry/CreateTenant"
	TenantRegistry_GetTenant_FullMethodName          = "/org.packetbroker.iam.v1.TenantRegistry/GetTenant"
	TenantRegistry_UpdateTenant_FullMethodName       = "/org.packetbroker.iam.v1.TenantRegistry/UpdateTenant"
	TenantRegistry_UpdateTenantListed_FullMethodName = "/org.packetbroker.iam.v1.TenantRegistry/UpdateTenantListed"
	TenantRegistry_DeleteTenant_FullMethodName       = "/org.packetbroker.iam.v1.TenantRegistry/DeleteTenant"
)

// TenantRegistryClient is the client API for TenantRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantRegistryClient interface {
	// List a network's tenants.
	ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error)
	// Create a tenant.
	CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*CreateTenantResponse, error)
	// Get a tenant.
	GetTenant(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*GetTenantResponse, error)
	// Update a tenant.
	UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Update the listed state of a tenant.
	UpdateTenantListed(ctx context.Context, in *UpdateTenantListedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete a tenant.
	DeleteTenant(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type tenantRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantRegistryClient(cc grpc.ClientConnInterface) TenantRegistryClient {
	return &tenantRegistryClient{cc}
}

func (c *tenantRegistryClient) ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error) {
	out := new(ListTenantsResponse)
	err := c.cc.Invoke(ctx, TenantRegistry_ListTenants_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantRegistryClient) CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*CreateTenantResponse, error) {
	out := new(CreateTenantResponse)
	err := c.cc.Invoke(ctx, TenantRegistry_CreateTenant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantRegistryClient) GetTenant(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*GetTenantResponse, error) {
	out := new(GetTenantResponse)
	err := c.cc.Invoke(ctx, TenantRegistry_GetTenant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantRegistryClient) UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TenantRegistry_UpdateTenant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantRegistryClient) UpdateTenantListed(ctx context.Context, in *UpdateTenantListedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TenantRegistry_UpdateTenantListed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantRegistryClient) DeleteTenant(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TenantRegistry_DeleteTenant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantRegistryServer is the server API for TenantRegistry service.
// All implementations must embed UnimplementedTenantRegistryServer
// for forward compatibility
type TenantRegistryServer interface {
	// List a network's tenants.
	ListTenants(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error)
	// Create a tenant.
	CreateTenant(context.Context, *CreateTenantRequest) (*CreateTenantResponse, error)
	// Get a tenant.
	GetTenant(context.Context, *TenantRequest) (*GetTenantResponse, error)
	// Update a tenant.
	UpdateTenant(context.Context, *UpdateTenantRequest) (*emptypb.Empty, error)
	// Update the listed state of a tenant.
	UpdateTenantListed(context.Context, *UpdateTenantListedRequest) (*emptypb.Empty, error)
	// Delete a tenant.
	DeleteTenant(context.Context, *TenantRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTenantRegistryServer()
}

// UnimplementedTenantRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedTenantRegistryServer struct {
}

func (UnimplementedTenantRegistryServer) ListTenants(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTenants not implemented")
}
func (UnimplementedTenantRegistryServer) CreateTenant(context.Context, *CreateTenantRequest) (*CreateTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenant not implemented")
}
func (UnimplementedTenantRegistryServer) GetTenant(context.Context, *TenantRequest) (*GetTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenant not implemented")
}
func (UnimplementedTenantRegistryServer) UpdateTenant(context.Context, *UpdateTenantRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTenant not implemented")
}
func (UnimplementedTenantRegistryServer) UpdateTenantListed(context.Context, *UpdateTenantListedRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTenantListed not implemented")
}
func (UnimplementedTenantRegistryServer) DeleteTenant(context.Context, *TenantRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTenant not implemented")
}
func (UnimplementedTenantRegistryServer) mustEmbedUnimplementedTenantRegistryServer() {}

// UnsafeTenantRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantRegistryServer will
// result in compilation errors.
type UnsafeTenantRegistryServer interface {
	mustEmbedUnimplementedTenantRegistryServer()
}

func RegisterTenantRegistryServer(s grpc.ServiceRegistrar, srv TenantRegistryServer) {
	s.RegisterService(&TenantRegistry_ServiceDesc, srv)
}

func _TenantRegistry_ListTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantRegistryServer).ListTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantRegistry_ListTenants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantRegistryServer).ListTenants(ctx, req.(*ListTenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantRegistry_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantRegistryServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantRegistry_CreateTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantRegistryServer).CreateTenant(ctx, req.(*CreateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantRegistry_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantRegistryServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantRegistry_GetTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantRegistryServer).GetTenant(ctx, req.(*TenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantRegistry_UpdateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantRegistryServer).UpdateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantRegistry_UpdateTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantRegistryServer).UpdateTenant(ctx, req.(*UpdateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantRegistry_UpdateTenantListed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTenantListedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantRegistryServer).UpdateTenantListed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantRegistry_UpdateTenantListed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantRegistryServer).UpdateTenantListed(ctx, req.(*UpdateTenantListedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantRegistry_DeleteTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantRegistryServer).DeleteTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantRegistry_DeleteTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantRegistryServer).DeleteTenant(ctx, req.(*TenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TenantRegistry_ServiceDesc is the grpc.ServiceDesc for TenantRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TenantRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "org.packetbroker.iam.v1.TenantRegistry",
	HandlerType: (*TenantRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTenants",
			Handler:    _TenantRegistry_ListTenants_Handler,
		},
		{
			MethodName: "CreateTenant",
			Handler:    _TenantRegistry_CreateTenant_Handler,
		},
		{
			MethodName: "GetTenant",
			Handler:    _TenantRegistry_GetTenant_Handler,
		},
		{
			MethodName: "UpdateTenant",
			Handler:    _TenantRegistry_UpdateTenant_Handler,
		},
		{
			MethodName: "UpdateTenantListed",
			Handler:    _TenantRegistry_UpdateTenantListed_Handler,
		},
		{
			MethodName: "DeleteTenant",
			Handler:    _TenantRegistry_DeleteTenant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "packetbroker/api/iam/v1/service.proto",
}

const (
	APIKeyVault_ListAPIKeys_FullMethodName  = "/org.packetbroker.iam.v1.APIKeyVault/ListAPIKeys"
	APIKeyVault_CreateAPIKey_FullMethodName = "/org.packetbroker.iam.v1.APIKeyVault/CreateAPIKey"
	APIKeyVault_DeleteAPIKey_FullMethodName = "/org.packetbroker.iam.v1.APIKeyVault/DeleteAPIKey"
)

// APIKeyVaultClient is the client API for APIKeyVault service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIKeyVaultClient interface {
	// List API keys.
	ListAPIKeys(ctx context.Context, in *ListAPIKeysRequest, opts ...grpc.CallOption) (*ListAPIKeysResponse, error)
	// Create an API key.
	CreateAPIKey(ctx context.Context, in *CreateAPIKeyRequest, opts ...grpc.CallOption) (*CreateAPIKeyResponse, error)
	// Delete an API key.
	DeleteAPIKey(ctx context.Context, in *APIKeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type aPIKeyVaultClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIKeyVaultClient(cc grpc.ClientConnInterface) APIKeyVaultClient {
	return &aPIKeyVaultClient{cc}
}

func (c *aPIKeyVaultClient) ListAPIKeys(ctx context.Context, in *ListAPIKeysRequest, opts ...grpc.CallOption) (*ListAPIKeysResponse, error) {
	out := new(ListAPIKeysResponse)
	err := c.cc.Invoke(ctx, APIKeyVault_ListAPIKeys_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIKeyVaultClient) CreateAPIKey(ctx context.Context, in *CreateAPIKeyRequest, opts ...grpc.CallOption) (*CreateAPIKeyResponse, error) {
	out := new(CreateAPIKeyResponse)
	err := c.cc.Invoke(ctx, APIKeyVault_CreateAPIKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIKeyVaultClient) DeleteAPIKey(ctx context.Context, in *APIKeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, APIKeyVault_DeleteAPIKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIKeyVaultServer is the server API for APIKeyVault service.
// All implementations must embed UnimplementedAPIKeyVaultServer
// for forward compatibility
type APIKeyVaultServer interface {
	// List API keys.
	ListAPIKeys(context.Context, *ListAPIKeysRequest) (*ListAPIKeysResponse, error)
	// Create an API key.
	CreateAPIKey(context.Context, *CreateAPIKeyRequest) (*CreateAPIKeyResponse, error)
	// Delete an API key.
	DeleteAPIKey(context.Context, *APIKeyRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAPIKeyVaultServer()
}

// UnimplementedAPIKeyVaultServer must be embedded to have forward compatible implementations.
type UnimplementedAPIKeyVaultServer struct {
}

func (UnimplementedAPIKeyVaultServer) ListAPIKeys(context.Context, *ListAPIKeysRequest) (*ListAPIKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAPIKeys not implemented")
}
func (UnimplementedAPIKeyVaultServer) CreateAPIKey(context.Context, *CreateAPIKeyRequest) (*CreateAPIKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAPIKey not implemented")
}
func (UnimplementedAPIKeyVaultServer) DeleteAPIKey(context.Context, *APIKeyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAPIKey not implemented")
}
func (UnimplementedAPIKeyVaultServer) mustEmbedUnimplementedAPIKeyVaultServer() {}

// UnsafeAPIKeyVaultServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIKeyVaultServer will
// result in compilation errors.
type UnsafeAPIKeyVaultServer interface {
	mustEmbedUnimplementedAPIKeyVaultServer()
}

func RegisterAPIKeyVaultServer(s grpc.ServiceRegistrar, srv APIKeyVaultServer) {
	s.RegisterService(&APIKeyVault_ServiceDesc, srv)
}

func _APIKeyVault_ListAPIKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAPIKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyVaultServer).ListAPIKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKeyVault_ListAPIKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyVaultServer).ListAPIKeys(ctx, req.(*ListAPIKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIKeyVault_CreateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyVaultServer).CreateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKeyVault_CreateAPIKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyVaultServer).CreateAPIKey(ctx, req.(*CreateAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIKeyVault_DeleteAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyVaultServer).DeleteAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKeyVault_DeleteAPIKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyVaultServer).DeleteAPIKey(ctx, req.(*APIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// APIKeyVault_ServiceDesc is the grpc.ServiceDesc for APIKeyVault service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var APIKeyVault_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "org.packetbroker.iam.v1.APIKeyVault",
	HandlerType: (*APIKeyVaultServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAPIKeys",
			Handler:    _APIKeyVault_ListAPIKeys_Handler,
		},
		{
			MethodName: "CreateAPIKey",
			Handler:    _APIKeyVault_CreateAPIKey_Handler,
		},
		{
			MethodName: "DeleteAPIKey",
			Handler:    _APIKeyVault_DeleteAPIKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "packetbroker/api/iam/v1/service.proto",
}
