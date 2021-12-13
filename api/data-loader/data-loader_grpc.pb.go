// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package data_loader

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

// DataLoaderClient is the client API for DataLoader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataLoaderClient interface {
	SaveClient(ctx context.Context, in *ClientReq, opts ...grpc.CallOption) (*ClientResp, error)
	SaveComponent(ctx context.Context, in *ComponentReq, opts ...grpc.CallOption) (*ComponentResp, error)
	SaveEngine(ctx context.Context, in *EngineReq, opts ...grpc.CallOption) (*EngineResp, error)
	SaveOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderResp, error)
	SaveProvider(ctx context.Context, in *ProviderReq, opts ...grpc.CallOption) (*ProviderResp, error)
	SaveRequirement(ctx context.Context, in *RequirementReq, opts ...grpc.CallOption) (*RequirementResp, error)
	SaveSupply(ctx context.Context, in *SupplyReq, opts ...grpc.CallOption) (*SupplyResp, error)
}

type dataLoaderClient struct {
	cc grpc.ClientConnInterface
}

func NewDataLoaderClient(cc grpc.ClientConnInterface) DataLoaderClient {
	return &dataLoaderClient{cc}
}

func (c *dataLoaderClient) SaveClient(ctx context.Context, in *ClientReq, opts ...grpc.CallOption) (*ClientResp, error) {
	out := new(ClientResp)
	err := c.cc.Invoke(ctx, "/DataLoader/SaveClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLoaderClient) SaveComponent(ctx context.Context, in *ComponentReq, opts ...grpc.CallOption) (*ComponentResp, error) {
	out := new(ComponentResp)
	err := c.cc.Invoke(ctx, "/DataLoader/SaveComponent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLoaderClient) SaveEngine(ctx context.Context, in *EngineReq, opts ...grpc.CallOption) (*EngineResp, error) {
	out := new(EngineResp)
	err := c.cc.Invoke(ctx, "/DataLoader/SaveEngine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLoaderClient) SaveOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderResp, error) {
	out := new(OrderResp)
	err := c.cc.Invoke(ctx, "/DataLoader/SaveOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLoaderClient) SaveProvider(ctx context.Context, in *ProviderReq, opts ...grpc.CallOption) (*ProviderResp, error) {
	out := new(ProviderResp)
	err := c.cc.Invoke(ctx, "/DataLoader/SaveProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLoaderClient) SaveRequirement(ctx context.Context, in *RequirementReq, opts ...grpc.CallOption) (*RequirementResp, error) {
	out := new(RequirementResp)
	err := c.cc.Invoke(ctx, "/DataLoader/SaveRequirement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataLoaderClient) SaveSupply(ctx context.Context, in *SupplyReq, opts ...grpc.CallOption) (*SupplyResp, error) {
	out := new(SupplyResp)
	err := c.cc.Invoke(ctx, "/DataLoader/SaveSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataLoaderServer is the server API for DataLoader service.
// All implementations must embed UnimplementedDataLoaderServer
// for forward compatibility
type DataLoaderServer interface {
	SaveClient(context.Context, *ClientReq) (*ClientResp, error)
	SaveComponent(context.Context, *ComponentReq) (*ComponentResp, error)
	SaveEngine(context.Context, *EngineReq) (*EngineResp, error)
	SaveOrder(context.Context, *OrderReq) (*OrderResp, error)
	SaveProvider(context.Context, *ProviderReq) (*ProviderResp, error)
	SaveRequirement(context.Context, *RequirementReq) (*RequirementResp, error)
	SaveSupply(context.Context, *SupplyReq) (*SupplyResp, error)
	mustEmbedUnimplementedDataLoaderServer()
}

// UnimplementedDataLoaderServer must be embedded to have forward compatible implementations.
type UnimplementedDataLoaderServer struct {
}

func (UnimplementedDataLoaderServer) SaveClient(context.Context, *ClientReq) (*ClientResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveClient not implemented")
}
func (UnimplementedDataLoaderServer) SaveComponent(context.Context, *ComponentReq) (*ComponentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveComponent not implemented")
}
func (UnimplementedDataLoaderServer) SaveEngine(context.Context, *EngineReq) (*EngineResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveEngine not implemented")
}
func (UnimplementedDataLoaderServer) SaveOrder(context.Context, *OrderReq) (*OrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveOrder not implemented")
}
func (UnimplementedDataLoaderServer) SaveProvider(context.Context, *ProviderReq) (*ProviderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveProvider not implemented")
}
func (UnimplementedDataLoaderServer) SaveRequirement(context.Context, *RequirementReq) (*RequirementResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveRequirement not implemented")
}
func (UnimplementedDataLoaderServer) SaveSupply(context.Context, *SupplyReq) (*SupplyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveSupply not implemented")
}
func (UnimplementedDataLoaderServer) mustEmbedUnimplementedDataLoaderServer() {}

// UnsafeDataLoaderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataLoaderServer will
// result in compilation errors.
type UnsafeDataLoaderServer interface {
	mustEmbedUnimplementedDataLoaderServer()
}

func RegisterDataLoaderServer(s grpc.ServiceRegistrar, srv DataLoaderServer) {
	s.RegisterService(&DataLoader_ServiceDesc, srv)
}

func _DataLoader_SaveClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLoaderServer).SaveClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataLoader/SaveClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLoaderServer).SaveClient(ctx, req.(*ClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLoader_SaveComponent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComponentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLoaderServer).SaveComponent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataLoader/SaveComponent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLoaderServer).SaveComponent(ctx, req.(*ComponentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLoader_SaveEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EngineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLoaderServer).SaveEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataLoader/SaveEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLoaderServer).SaveEngine(ctx, req.(*EngineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLoader_SaveOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLoaderServer).SaveOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataLoader/SaveOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLoaderServer).SaveOrder(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLoader_SaveProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLoaderServer).SaveProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataLoader/SaveProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLoaderServer).SaveProvider(ctx, req.(*ProviderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLoader_SaveRequirement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequirementReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLoaderServer).SaveRequirement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataLoader/SaveRequirement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLoaderServer).SaveRequirement(ctx, req.(*RequirementReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataLoader_SaveSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLoaderServer).SaveSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataLoader/SaveSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLoaderServer).SaveSupply(ctx, req.(*SupplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DataLoader_ServiceDesc is the grpc.ServiceDesc for DataLoader service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataLoader_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DataLoader",
	HandlerType: (*DataLoaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveClient",
			Handler:    _DataLoader_SaveClient_Handler,
		},
		{
			MethodName: "SaveComponent",
			Handler:    _DataLoader_SaveComponent_Handler,
		},
		{
			MethodName: "SaveEngine",
			Handler:    _DataLoader_SaveEngine_Handler,
		},
		{
			MethodName: "SaveOrder",
			Handler:    _DataLoader_SaveOrder_Handler,
		},
		{
			MethodName: "SaveProvider",
			Handler:    _DataLoader_SaveProvider_Handler,
		},
		{
			MethodName: "SaveRequirement",
			Handler:    _DataLoader_SaveRequirement_Handler,
		},
		{
			MethodName: "SaveSupply",
			Handler:    _DataLoader_SaveSupply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/data-loader/data-loader.proto",
}
