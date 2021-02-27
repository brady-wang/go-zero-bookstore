// Code generated by goctl. DO NOT EDIT!
// Source: book.proto

package main

import (
	"flag"
	"fmt"

	"bookstore/rpc/book/book"
	"bookstore/rpc/book/internal/config"
	"bookstore/rpc/book/internal/server"
	"bookstore/rpc/book/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/book.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewBookServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		book.RegisterBookServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}