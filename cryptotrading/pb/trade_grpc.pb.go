// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: pb/trade.proto

package pb

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
	TradeService_ExecuteTrade_FullMethodName = "/pb.TradeService/ExecuteTrade"
)

// TradeServiceClient is the client API for TradeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TradeServiceClient interface {
	ExecuteTrade(ctx context.Context, in *TradeRequest, opts ...grpc.CallOption) (*TradeResponse, error)
}

type tradeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTradeServiceClient(cc grpc.ClientConnInterface) TradeServiceClient {
	return &tradeServiceClient{cc}
}

func (c *tradeServiceClient) ExecuteTrade(ctx context.Context, in *TradeRequest, opts ...grpc.CallOption) (*TradeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TradeResponse)
	err := c.cc.Invoke(ctx, TradeService_ExecuteTrade_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TradeServiceServer is the server API for TradeService service.
// All implementations must embed UnimplementedTradeServiceServer
// for forward compatibility
type TradeServiceServer interface {
	ExecuteTrade(context.Context, *TradeRequest) (*TradeResponse, error)
	mustEmbedUnimplementedTradeServiceServer()
}

// UnimplementedTradeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTradeServiceServer struct {
}

func (UnimplementedTradeServiceServer) ExecuteTrade(context.Context, *TradeRequest) (*TradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteTrade not implemented")
}
func (UnimplementedTradeServiceServer) mustEmbedUnimplementedTradeServiceServer() {}

// UnsafeTradeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TradeServiceServer will
// result in compilation errors.
type UnsafeTradeServiceServer interface {
	mustEmbedUnimplementedTradeServiceServer()
}

func RegisterTradeServiceServer(s grpc.ServiceRegistrar, srv TradeServiceServer) {
	s.RegisterService(&TradeService_ServiceDesc, srv)
}

func _TradeService_ExecuteTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServiceServer).ExecuteTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeService_ExecuteTrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServiceServer).ExecuteTrade(ctx, req.(*TradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TradeService_ServiceDesc is the grpc.ServiceDesc for TradeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TradeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TradeService",
	HandlerType: (*TradeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteTrade",
			Handler:    _TradeService_ExecuteTrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/trade.proto",
}
