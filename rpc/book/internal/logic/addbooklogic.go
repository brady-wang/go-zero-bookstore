package logic

import (
	"bookstore/rpc/model"
	"context"

	"bookstore/rpc/book/book"
	"bookstore/rpc/book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBookLogic {
	return &AddBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddBookLogic) AddBook(in *book.AddRequest) (*book.AddResponse, error) {
	// 手动代码开始
	_, err := l.svcCtx.Model.Insert(model.Book{
		Name:  in.Name,
		Price: in.Price,
	})
	if err != nil {
		return nil, err
	}

	return &book.AddResponse{
		Ok: true,
	}, nil
	// 手动代码结束
}
