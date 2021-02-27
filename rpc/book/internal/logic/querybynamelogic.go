package logic

import (
	"context"

	"bookstore/rpc/book/book"
	"bookstore/rpc/book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type QueryByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryByNameLogic {
	return &QueryByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryByNameLogic) QueryByName(in *book.QueryByNameRequest) (*book.QueryByNameResponse, error) {
	// 手动代码开始
	bookInfo, err := l.svcCtx.Model.FindByName(in.Name)
	if err != nil {
		return nil, err
	}

	return &book.QueryByNameResponse{
		Id:    bookInfo.Id,
		Name:  bookInfo.Name,
		Price: bookInfo.Price,
	}, nil
	// 手动代码结束
}
