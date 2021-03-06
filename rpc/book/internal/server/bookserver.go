// Code generated by goctl. DO NOT EDIT!
// Source: book.proto

package server

import (
	"context"

	"bookstore/rpc/book/book"
	"bookstore/rpc/book/internal/logic"
	"bookstore/rpc/book/internal/svc"
)

type BookServer struct {
	svcCtx *svc.ServiceContext
}

func NewBookServer(svcCtx *svc.ServiceContext) *BookServer {
	return &BookServer{
		svcCtx: svcCtx,
	}
}

func (s *BookServer) AddBook(ctx context.Context, in *book.AddRequest) (*book.AddResponse, error) {
	l := logic.NewAddBookLogic(ctx, s.svcCtx)
	return l.AddBook(in)
}

func (s *BookServer) QueryByName(ctx context.Context, in *book.QueryByNameRequest) (*book.QueryByNameResponse, error) {
	l := logic.NewQueryByNameLogic(ctx, s.svcCtx)
	return l.QueryByName(in)
}

func (s *BookServer) QueryAll(ctx context.Context, in *book.QueryAllRequest) (*book.QueryAllResponse, error) {
	l := logic.NewQueryAllLogic(ctx, s.svcCtx)
	return l.QueryAll(in)
}
