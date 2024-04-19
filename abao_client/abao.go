// Code generated by goctl. DO NOT EDIT.
// Source: abao.proto

package abao_client

import (
	"context"

	"abao/abao"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Account            = abao.Account
	GetAccountRequest  = abao.GetAccountRequest
	GetAccountResponse = abao.GetAccountResponse
	GetBalanceRequest  = abao.GetBalanceRequest
	GetBalanceResponse = abao.GetBalanceResponse
	TransferRequest    = abao.TransferRequest
	TransferResponse   = abao.TransferResponse

	Abao interface {
		GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
		GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
		Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	}

	defaultAbao struct {
		cli zrpc.Client
	}
)

func NewAbao(cli zrpc.Client) Abao {
	return &defaultAbao{
		cli: cli,
	}
}

func (m *defaultAbao) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	client := abao.NewAbaoClient(m.cli.Conn())
	return client.GetAccount(ctx, in, opts...)
}

func (m *defaultAbao) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	client := abao.NewAbaoClient(m.cli.Conn())
	return client.GetBalance(ctx, in, opts...)
}

func (m *defaultAbao) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	client := abao.NewAbaoClient(m.cli.Conn())
	return client.Transfer(ctx, in, opts...)
}
