package logic

import (
	"bookstore/rpc/book/book"
	"context"
	"fmt"

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
	// 手动代码开始
	resp, err :=l.svcCtx.Book.QueryAll(l.ctx,&book.QueryAllRequest{

	})
	logx.Info(resp)
	if err != nil {
		return nil, err
	}

	var list []*types.QueryByNameResponse
	for _,v := range resp.List{
		fmt.Println(v)
		var item types.QueryByNameResponse
		item.Id = item.Id
		item.Price = v.Price
		item.Name = v.Name
		list  = append(list,&item)
	}

	// 手动代码结束
	return &types.QueryAllResp{
		BookList: list,
	},nil
}
