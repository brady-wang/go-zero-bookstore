package logic

import (
	"context"

	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type QueryAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryAllLogic {
	return QueryAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAllLogic) QueryAll(req types.QueryAllReq) (*types.QueryAllResp, error) {
	// todo: add your logic here and delete this line

	return &types.QueryAllResp{}, nil
}
