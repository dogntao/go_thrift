package main

import (
	"context"
	"fmt"
	"hello"
	"os"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"
)

func main() {
	fmt.Println("hello")
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket("127.0.0.1:3000")
	if err != nil {
		fmt.Println(err)
		fmt.Errorf("NewTSocket failed. err: [%v]\n", err)
		return
	}

	transport, err = thrift.NewTBufferedTransportFactory(8192).GetTransport(transport)
	if err != nil {
		fmt.Errorf("NewTransport failed. err: [%v]\n", err)
		return
	}
	defer transport.Close()

	if err := transport.Open(); err != nil {
		fmt.Errorf("Transport.Open failed. err: [%v]\n", err)
		return
	}

	protocolFactory := thrift.NewTCompactProtocolFactory()
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	client := hello.NewHelloClient(thrift.NewTStandardClient(iprot, oprot))

	var res *hello.HelloRes
	res, err = client.Echo(context.Background(), &hello.HelloReq{
		Msg: strings.Join(os.Args[1:], " "),
	})
	if err != nil {
		fmt.Errorf("client hello failed. err: [%v]", err)
		return
	}

	fmt.Printf("message from server: %v", res.GetMsg())
}
