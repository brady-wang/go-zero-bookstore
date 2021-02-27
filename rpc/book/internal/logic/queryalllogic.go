package logic

import (
	"context"

	"bookstore/rpc/book/book"
	"bookstore/rpc/book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type QueryAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllLogic {
	return &QueryAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAllLogic) QueryAll(in *book.QueryAllRequest) (*book.QueryAllResponse, error) {
	return &book.QueryAllResponse{},nil
}
