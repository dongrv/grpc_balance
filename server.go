// Package grpc_balance
// @Description
// @author      董荣旺
// @datetime    2022/10/6 13:17
package grpc_balance

import (
	pb "cat/grpc_balance/protocol"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
)

var port = flag.Int("port", 8927, "服务端口")

type server struct {
	pb.UnimplementedGreeterServer
	Addr string
}

func (s *server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := fmt.Sprintf("hello [%s]. from [%s]", request.GetName(), s.Addr)
	return &pb.HelloResponse{Reply: reply}, nil
}

func Run(wg sync.WaitGroup) {
	flag.Parse()
	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("failed to listen %s, err:%s\n", addr, err.Error())
		return
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{Addr: addr})
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("run server err:%s\n", err.Error())
		return
	}
	wg.Done()
}
