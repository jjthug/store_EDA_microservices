// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: basketspb/api.proto

package basketspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BasketService_StartBasket_FullMethodName    = "/basketspb.BasketService/StartBasket"
	BasketService_CancelBasket_FullMethodName   = "/basketspb.BasketService/CancelBasket"
	BasketService_CheckoutBasket_FullMethodName = "/basketspb.BasketService/CheckoutBasket"
	BasketService_AddItem_FullMethodName        = "/basketspb.BasketService/AddItem"
	BasketService_RemoveItem_FullMethodName     = "/basketspb.BasketService/RemoveItem"
	BasketService_GetBasket_FullMethodName      = "/basketspb.BasketService/GetBasket"
)

// BasketServiceClient is the client API for BasketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BasketServiceClient interface {
	StartBasket(ctx context.Context, in *StartBasketRequest, opts ...grpc.CallOption) (*StartBasketResponse, error)
	CancelBasket(ctx context.Context, in *CancelBasketRequest, opts ...grpc.CallOption) (*CancelBasketResponse, error)
	CheckoutBasket(ctx context.Context, in *CheckoutBasketRequest, opts ...grpc.CallOption) (*CheckoutBasketResponse, error)
	AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error)
	RemoveItem(ctx context.Context, in *RemoveItemRequest, opts ...grpc.CallOption) (*RemoveItemResponse, error)
	GetBasket(ctx context.Context, in *GetBasketRequest, opts ...grpc.CallOption) (*GetBasketResponse, error)
}

type basketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBasketServiceClient(cc grpc.ClientConnInterface) BasketServiceClient {
	return &basketServiceClient{cc}
}

func (c *basketServiceClient) StartBasket(ctx context.Context, in *StartBasketRequest, opts ...grpc.CallOption) (*StartBasketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartBasketResponse)
	err := c.cc.Invoke(ctx, BasketService_StartBasket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) CancelBasket(ctx context.Context, in *CancelBasketRequest, opts ...grpc.CallOption) (*CancelBasketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelBasketResponse)
	err := c.cc.Invoke(ctx, BasketService_CancelBasket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) CheckoutBasket(ctx context.Context, in *CheckoutBasketRequest, opts ...grpc.CallOption) (*CheckoutBasketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckoutBasketResponse)
	err := c.cc.Invoke(ctx, BasketService_CheckoutBasket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddItemResponse)
	err := c.cc.Invoke(ctx, BasketService_AddItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) RemoveItem(ctx context.Context, in *RemoveItemRequest, opts ...grpc.CallOption) (*RemoveItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveItemResponse)
	err := c.cc.Invoke(ctx, BasketService_RemoveItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basketServiceClient) GetBasket(ctx context.Context, in *GetBasketRequest, opts ...grpc.CallOption) (*GetBasketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBasketResponse)
	err := c.cc.Invoke(ctx, BasketService_GetBasket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasketServiceServer is the server API for BasketService service.
// All implementations must embed UnimplementedBasketServiceServer
// for forward compatibility.
type BasketServiceServer interface {
	StartBasket(context.Context, *StartBasketRequest) (*StartBasketResponse, error)
	CancelBasket(context.Context, *CancelBasketRequest) (*CancelBasketResponse, error)
	CheckoutBasket(context.Context, *CheckoutBasketRequest) (*CheckoutBasketResponse, error)
	AddItem(context.Context, *AddItemRequest) (*AddItemResponse, error)
	RemoveItem(context.Context, *RemoveItemRequest) (*RemoveItemResponse, error)
	GetBasket(context.Context, *GetBasketRequest) (*GetBasketResponse, error)
	mustEmbedUnimplementedBasketServiceServer()
}

// UnimplementedBasketServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBasketServiceServer struct{}

func (UnimplementedBasketServiceServer) StartBasket(context.Context, *StartBasketRequest) (*StartBasketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartBasket not implemented")
}
func (UnimplementedBasketServiceServer) CancelBasket(context.Context, *CancelBasketRequest) (*CancelBasketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBasket not implemented")
}
func (UnimplementedBasketServiceServer) CheckoutBasket(context.Context, *CheckoutBasketRequest) (*CheckoutBasketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckoutBasket not implemented")
}
func (UnimplementedBasketServiceServer) AddItem(context.Context, *AddItemRequest) (*AddItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (UnimplementedBasketServiceServer) RemoveItem(context.Context, *RemoveItemRequest) (*RemoveItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveItem not implemented")
}
func (UnimplementedBasketServiceServer) GetBasket(context.Context, *GetBasketRequest) (*GetBasketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBasket not implemented")
}
func (UnimplementedBasketServiceServer) mustEmbedUnimplementedBasketServiceServer() {}
func (UnimplementedBasketServiceServer) testEmbeddedByValue()                       {}

// UnsafeBasketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BasketServiceServer will
// result in compilation errors.
type UnsafeBasketServiceServer interface {
	mustEmbedUnimplementedBasketServiceServer()
}

func RegisterBasketServiceServer(s grpc.ServiceRegistrar, srv BasketServiceServer) {
	// If the following call pancis, it indicates UnimplementedBasketServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BasketService_ServiceDesc, srv)
}

func _BasketService_StartBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartBasketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).StartBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BasketService_StartBasket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).StartBasket(ctx, req.(*StartBasketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_CancelBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelBasketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).CancelBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BasketService_CancelBasket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).CancelBasket(ctx, req.(*CancelBasketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_CheckoutBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckoutBasketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).CheckoutBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BasketService_CheckoutBasket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).CheckoutBasket(ctx, req.(*CheckoutBasketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BasketService_AddItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).AddItem(ctx, req.(*AddItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_RemoveItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).RemoveItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BasketService_RemoveItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).RemoveItem(ctx, req.(*RemoveItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasketService_GetBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBasketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasketServiceServer).GetBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BasketService_GetBasket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasketServiceServer).GetBasket(ctx, req.(*GetBasketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BasketService_ServiceDesc is the grpc.ServiceDesc for BasketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BasketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "basketspb.BasketService",
	HandlerType: (*BasketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartBasket",
			Handler:    _BasketService_StartBasket_Handler,
		},
		{
			MethodName: "CancelBasket",
			Handler:    _BasketService_CancelBasket_Handler,
		},
		{
			MethodName: "CheckoutBasket",
			Handler:    _BasketService_CheckoutBasket_Handler,
		},
		{
			MethodName: "AddItem",
			Handler:    _BasketService_AddItem_Handler,
		},
		{
			MethodName: "RemoveItem",
			Handler:    _BasketService_RemoveItem_Handler,
		},
		{
			MethodName: "GetBasket",
			Handler:    _BasketService_GetBasket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "basketspb/api.proto",
}