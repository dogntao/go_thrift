package main

import (
	"context"
	"fmt"
	"hello"

	"github.com/apache/thrift/lib/go/thrift"
)

type HelloServerImp struct {
}

func (h *HelloServerImp) Echo(ctx context.Context, req *hello.HelloReq) (*hello.HelloRes, error) {
	fmt.Printf("message from client: %v\n", req.GetMsg())

	res := &hello.HelloRes{
		Msg: req.GetMsg(),
	}

	return res, nil
}

func main() {
	transport, err := thrift.NewTServerSocket("127.0.0.1:3000")
	if err != nil {
		panic(err)
	}

	processor := hello.NewHelloProcessor(&HelloServerImp{})
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		thrift.NewTBufferedTransportFactory(8192),
		thrift.NewTCompactProtocolFactory(),
	)
	if err := server.Serve(); err != nil {
		panic(err)
	}
}
