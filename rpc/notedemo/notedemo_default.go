package notedemo

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/notedemo"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func CreateNote(ctx context.Context, req *notedemo.CreateNoteRequest, callOptions ...callopt.Option) (resp *notedemo.CreateNoteResponse, err error) {
	resp, err = defaultClient.CreateNote(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateNote call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func QueryNote(ctx context.Context, req *notedemo.QueryNoteRequest, callOptions ...callopt.Option) (resp *notedemo.QueryNoteResponse, err error) {
	resp, err = defaultClient.QueryNote(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MGetNote call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MGetNote(ctx context.Context, req *notedemo.MGetNoteRequest, callOptions ...callopt.Option) (resp *notedemo.MGetNoteResponse, err error) {
	resp, err = defaultClient.MGetNote(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MGetNote call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteNote(ctx context.Context, req *notedemo.DeleteNoteRequest, callOptions ...callopt.Option) (resp *notedemo.DeleteNoteResponse, err error) {
	resp, err = defaultClient.DeleteNote(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteNote call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateNote(ctx context.Context, req *notedemo.UpdateNoteRequest, callOptions ...callopt.Option) (resp *notedemo.UpdateNoteResponse, err error) {
	resp, err = defaultClient.UpdateNote(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MGetNote call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
