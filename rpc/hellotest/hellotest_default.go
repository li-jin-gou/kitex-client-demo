package hellotest

import (
	"context"

	"github.com/cloudwego/biz-demo/easy_note/kitex_gen/hello/example"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func HelloMethod(ctx context.Context, request *example.HelloReq, callOptions ...callopt.Option) (resp *example.HelloResp, err error) {
	resp, err = defaultClient.HelloMethod(ctx, request, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "HelloMethod call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func HelloMethod1(ctx context.Context, request *example.HelloReq, callOptions ...callopt.Option) (resp *example.HelloResp, err error) {
	resp, err = defaultClient.HelloMethod1(ctx, request, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "HelloMethod1 call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func HelloMethod2(ctx context.Context, request *example.HelloReq, callOptions ...callopt.Option) (resp *example.HelloResp, err error) {
	resp, err = defaultClient.HelloMethod2(ctx, request, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "HelloMethod2 call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
