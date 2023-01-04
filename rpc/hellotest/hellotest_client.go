package hellotest

import (
	"context"

	"github.com/cloudwego/biz-demo/easy_note/kitex_gen/hello/example"
	"github.com/cloudwego/biz-demo/easy_note/kitex_gen/hello/example/helloservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() helloservice.Client
	Service() string
	HelloMethod(ctx context.Context, request *example.HelloReq, callOptions ...callopt.Option) (resp *example.HelloResp, err error)

	HelloMethod1(ctx context.Context, request *example.HelloReq, callOptions ...callopt.Option) (resp *example.HelloResp, err error)

	HelloMethod2(ctx context.Context, request *example.HelloReq, callOptions ...callopt.Option) (resp *example.HelloResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := helloservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient helloservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() helloservice.Client {
	return c.kitexClient
}

func (c *clientImpl) HelloMethod(ctx context.Context, request *example.HelloReq, callOptions ...callopt.Option) (resp *example.HelloResp, err error) {
	return c.kitexClient.HelloMethod(ctx, request, callOptions...)
}

func (c *clientImpl) HelloMethod1(ctx context.Context, request *example.HelloReq, callOptions ...callopt.Option) (resp *example.HelloResp, err error) {
	return c.kitexClient.HelloMethod1(ctx, request, callOptions...)
}

func (c *clientImpl) HelloMethod2(ctx context.Context, request *example.HelloReq, callOptions ...callopt.Option) (resp *example.HelloResp, err error) {
	return c.kitexClient.HelloMethod2(ctx, request, callOptions...)
}
