// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: proto/bank/service.proto

package bank

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

// BankHandlerClient is the client API for BankHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankHandlerClient interface {
	GetCurrentBalance(ctx context.Context, in *CurrentBalanceRequest, opts ...grpc.CallOption) (*CurrentBalanceResponse, error)
}

type bankHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewBankHandlerClient(cc grpc.ClientConnInterface) BankHandlerClient {
	return &bankHandlerClient{cc}
}

func (c *bankHandlerClient) GetCurrentBalance(ctx context.Context, in *CurrentBalanceRequest, opts ...grpc.CallOption) (*CurrentBalanceResponse, error) {
	out := new(CurrentBalanceResponse)
	err := c.cc.Invoke(ctx, "/bank.BankHandler/GetCurrentBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankHandlerServer is the server API for BankHandler service.
// All implementations must embed UnimplementedBankHandlerServer
// for forward compatibility
type BankHandlerServer interface {
	GetCurrentBalance(context.Context, *CurrentBalanceRequest) (*CurrentBalanceResponse, error)
	mustEmbedUnimplementedBankHandlerServer()
}

// UnimplementedBankHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedBankHandlerServer struct {
}

func (UnimplementedBankHandlerServer) GetCurrentBalance(context.Context, *CurrentBalanceRequest) (*CurrentBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentBalance not implemented")
}
func (UnimplementedBankHandlerServer) mustEmbedUnimplementedBankHandlerServer() {}

// UnsafeBankHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankHandlerServer will
// result in compilation errors.
type UnsafeBankHandlerServer interface {
	mustEmbedUnimplementedBankHandlerServer()
}

func RegisterBankHandlerServer(s grpc.ServiceRegistrar, srv BankHandlerServer) {
	s.RegisterService(&BankHandler_ServiceDesc, srv)
}

func _BankHandler_GetCurrentBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankHandlerServer).GetCurrentBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank.BankHandler/GetCurrentBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankHandlerServer).GetCurrentBalance(ctx, req.(*CurrentBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BankHandler_ServiceDesc is the grpc.ServiceDesc for BankHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bank.BankHandler",
	HandlerType: (*BankHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentBalance",
			Handler:    _BankHandler_GetCurrentBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/bank/service.proto",
}
