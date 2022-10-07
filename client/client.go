// Package client
// @Description
// @author      董荣旺
// @datetime    2022/10/7 12:30
// 学习地址：https://www.liwenzhou.com/posts/Go/name-resolving-and-load-balancing-in-grpc/
package client

import (
	pb "cat/grpc_balance/protocol"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"sync"
	"time"
)

const (
	myScheme   = "dongrv"
	myEndpoint = "resolver.dongrv.com"
)

var addrs = []string{"127.0.0.1:8972", "127.0.0.1:8973"}

type dongrvResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *dongrvResolver) ResolveNow(o resolver.ResolveNowOptions) {
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrList := make([]resolver.Address, len(addrs))
	for i, s := range addrStrs {
		addrList[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrList})
}

func (*dongrvResolver) Close() {}

type dongrvBuilder struct{}

func (b *dongrvBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &dongrvResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			myEndpoint: addrs,
		},
	}
	r.ResolveNow(resolver.ResolveNowOptions{})
	return r, nil
}

func (*dongrvBuilder) Scheme() string { return myScheme }
func init() {
	// 注册
	resolver.Register(&dongrvBuilder{})
}

func Connect(wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := grpc.Dial(
		fmt.Sprintf("%s://%s/%s", myScheme, ``, myEndpoint),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.WithResolvers(&dongrvBuilder{}),
	)
	if err != nil {
		fmt.Printf("connect failed,err:%s\n", err.Error())
		return
	}

	c := pb.NewGreeterClient(conn)
	Send(c)
}

func Send(c pb.GreeterClient) {
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Pony"})
		if err != nil {
			panic(err)
		}
		fmt.Printf("get resp:%s\n", resp.GetReply())
	}
}
