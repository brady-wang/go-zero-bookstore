package svc

import (
	"bookstore/api/internal/config"
	"bookstore/rpc/book/bookclient"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Book    bookclient.Book
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Book: bookclient.NewBook(zrpc.MustNewClient(c.Book)),
	}
}
