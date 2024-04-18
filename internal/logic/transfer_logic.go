package logic

import (
	"context"

	"abao/abao"
	"abao/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransferLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTransferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransferLogic {
	return &TransferLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TransferLogic) Transfer(in *abao.TransferRequest) (*abao.TransferResponse, error) {
	// todo: add your logic here and delete this line

	return &abao.TransferResponse{}, nil
}