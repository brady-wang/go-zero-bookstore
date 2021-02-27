package logic

import (
	"bookstore/api/internal/types"
	"bookstore/rpc/book/book"
	"context"

	"bookstore/api/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

type QueryByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryByNameLogic {
	return QueryByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryByNameLogic) QueryByName(req types.QueryByNameReq) (*types.QueryByNameResponse, error) {
	// 手动代码开始
	resp, err :=l.svcCtx.Book.QueryByName(l.ctx,&book.QueryByNameRequest{Name: req.Name})
	if err != nil {
		return nil, err
	}

	return &types.QueryByNameResponse{
		Id: resp.GetId(),
		Name: resp.GetName(),
		Price:resp.GetPrice(),
	}, nil
	// 手动代码结束

}
