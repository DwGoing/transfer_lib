// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: transfer_lib.proto

package transfer_lib

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

const (
	TransferLib_GetAccount_FullMethodName     = "/transfer_lib.TransferLib/GetAccount"
	TransferLib_GetBalance_FullMethodName     = "/transfer_lib.TransferLib/GetBalance"
	TransferLib_Transfer_FullMethodName       = "/transfer_lib.TransferLib/Transfer"
	TransferLib_GetTranscation_FullMethodName = "/transfer_lib.TransferLib/GetTranscation"
)

// TransferLibClient is the client API for TransferLib service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransferLibClient interface {
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	GetTranscation(ctx context.Context, in *GetTranscationRequest, opts ...grpc.CallOption) (*GetTranscationResponse, error)
}

type transferLibClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferLibClient(cc grpc.ClientConnInterface) TransferLibClient {
	return &transferLibClient{cc}
}

func (c *transferLibClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, TransferLib_GetAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferLibClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, TransferLib_GetBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferLibClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, TransferLib_Transfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferLibClient) GetTranscation(ctx context.Context, in *GetTranscationRequest, opts ...grpc.CallOption) (*GetTranscationResponse, error) {
	out := new(GetTranscationResponse)
	err := c.cc.Invoke(ctx, TransferLib_GetTranscation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransferLibServer is the server API for TransferLib service.
// All implementations must embed UnimplementedTransferLibServer
// for forward compatibility
type TransferLibServer interface {
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	Transfer(context.Context, *TransferRequest) (*TransferResponse, error)
	GetTranscation(context.Context, *GetTranscationRequest) (*GetTranscationResponse, error)
	mustEmbedUnimplementedTransferLibServer()
}

// UnimplementedTransferLibServer must be embedded to have forward compatible implementations.
type UnimplementedTransferLibServer struct {
}

func (UnimplementedTransferLibServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedTransferLibServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedTransferLibServer) Transfer(context.Context, *TransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedTransferLibServer) GetTranscation(context.Context, *GetTranscationRequest) (*GetTranscationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTranscation not implemented")
}
func (UnimplementedTransferLibServer) mustEmbedUnimplementedTransferLibServer() {}

// UnsafeTransferLibServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransferLibServer will
// result in compilation errors.
type UnsafeTransferLibServer interface {
	mustEmbedUnimplementedTransferLibServer()
}

func RegisterTransferLibServer(s grpc.ServiceRegistrar, srv TransferLibServer) {
	s.RegisterService(&TransferLib_ServiceDesc, srv)
}

func _TransferLib_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferLibServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferLib_GetAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferLibServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferLib_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferLibServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferLib_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferLibServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferLib_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferLibServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferLib_Transfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferLibServer).Transfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferLib_GetTranscation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTranscationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferLibServer).GetTranscation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferLib_GetTranscation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferLibServer).GetTranscation(ctx, req.(*GetTranscationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransferLib_ServiceDesc is the grpc.ServiceDesc for TransferLib service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransferLib_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transfer_lib.TransferLib",
	HandlerType: (*TransferLibServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccount",
			Handler:    _TransferLib_GetAccount_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _TransferLib_GetBalance_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _TransferLib_Transfer_Handler,
		},
		{
			MethodName: "GetTranscation",
			Handler:    _TransferLib_GetTranscation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transfer_lib.proto",
}