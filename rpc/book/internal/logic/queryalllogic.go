package logic

import (
	"bookstore/rpc/book/book"
	"bookstore/rpc/book/internal/svc"
	"context"
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
	// 手动代码开始
	bookList, err := l.svcCtx.Model.FindAll()
	if err != nil {
		return nil, err
	}

	l.Logger.Error(bookList)
	var item book.QueryByNameResponse
	var result []*book.QueryByNameResponse
	for _,v := range bookList{
		item  = book.QueryByNameResponse{
			Id:    v.Id,
			Name:  v.Name,
			Price: v.Price,
		}
		result = append(result,&item)
	}

	// 手动代码结束
	return &book.QueryAllResponse{List: result},nil
}
