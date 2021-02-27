package logic

import (
	"bookstore/api/internal/types"
	"bookstore/rpc/book/book"
	"context"

	"bookstore/api/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

type AddBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddBookLogic {
	return AddBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddBookLogic) AddBook(req types.AddBookReq) (*types.AddBookResp, error) {
	// 手动代码开始
	resp, err :=l.svcCtx.Book.AddBook(l.ctx,&book.AddRequest{
		Name:  req.Name,
		Price: req.Price,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddBookResp{
		Ok: resp.Ok,
	}, nil
	// 手动代码结束
}
