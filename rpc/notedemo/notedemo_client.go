package notedemo

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/notedemo"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/notedemo/noteservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type NoteDemoClient interface {
	KitexClient() noteservice.Client
	Service() string
	MockClient() *MockClientStruct
	SetMockClient(mockClient *MockClientStruct)
	CreateNote(ctx context.Context, req *notedemo.CreateNoteRequest, callOptions ...callopt.Option) (r *notedemo.CreateNoteResponse, err error)
	MGetNote(ctx context.Context, req *notedemo.MGetNoteRequest, callOptions ...callopt.Option) (r *notedemo.MGetNoteResponse, err error)
	DeleteNote(ctx context.Context, req *notedemo.DeleteNoteRequest, callOptions ...callopt.Option) (r *notedemo.DeleteNoteResponse, err error)
	QueryNote(ctx context.Context, req *notedemo.QueryNoteRequest, callOptions ...callopt.Option) (r *notedemo.QueryNoteResponse, err error)
	UpdateNote(ctx context.Context, req *notedemo.UpdateNoteRequest, callOptions ...callopt.Option) (r *notedemo.UpdateNoteResponse, err error)
}

func NewNoteDemoClient(dstService string, opts ...client.Option) (NoteDemoClient, error) {
	kitexClient, err := noteservice.NewClient(dstService, opts...)
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
	kitexClient noteservice.Client
	mockClient  *MockClientStruct
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() noteservice.Client {
	return c.kitexClient
}

func (c *clientImpl) MockClient() *MockClientStruct {
	return c.mockClient
}

func (c *clientImpl) SetMockClient(mockClient *MockClientStruct) {
	if mockClient == nil {
		panic("cannot set a nil mock client")
	}
	c.mockClient = mockClient
}

func (c *clientImpl) CreateNote(ctx context.Context, req *notedemo.CreateNoteRequest, callOptions ...callopt.Option) (r *notedemo.CreateNoteResponse, err error) {
	if c.mockClient != nil && c.mockClient.CreateNoteFunc != nil {
		return c.mockClient.CreateNoteFunc(ctx, req, callOptions...)
	}
	return c.kitexClient.CreateNote(ctx, req, callOptions...)
}

func (c *clientImpl) MGetNote(ctx context.Context, req *notedemo.MGetNoteRequest, callOptions ...callopt.Option) (r *notedemo.MGetNoteResponse, err error) {
	if c.mockClient != nil && c.mockClient.MGetNoteFunc != nil {
		return c.mockClient.MGetNoteFunc(ctx, req, callOptions...)
	}
	return c.kitexClient.MGetNote(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteNote(ctx context.Context, req *notedemo.DeleteNoteRequest, callOptions ...callopt.Option) (r *notedemo.DeleteNoteResponse, err error) {
	if c.mockClient != nil && c.mockClient.DeleteNoteFunc != nil {
		return c.mockClient.DeleteNoteFunc(ctx, req, callOptions...)
	}
	return c.kitexClient.DeleteNote(ctx, req, callOptions...)
}

func (c *clientImpl) QueryNote(ctx context.Context, req *notedemo.QueryNoteRequest, callOptions ...callopt.Option) (r *notedemo.QueryNoteResponse, err error) {
	if c.mockClient.QueryNoteFunc != nil {
		return c.mockClient.QueryNoteFunc(ctx, req, callOptions...)
	}
	return c.kitexClient.QueryNote(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateNote(ctx context.Context, req *notedemo.UpdateNoteRequest, callOptions ...callopt.Option) (r *notedemo.UpdateNoteResponse, err error) {
	if c.mockClient.UpdateNoteFunc != nil {
		return c.mockClient.UpdateNoteFunc(ctx, req, callOptions...)
	}
	return c.kitexClient.UpdateNote(ctx, req, callOptions...)
}

type MockClientStruct struct {
	CreateNoteFunc func(ctx context.Context, req *notedemo.CreateNoteRequest, callOptions ...callopt.Option) (r *notedemo.CreateNoteResponse, err error)
	MGetNoteFunc   func(ctx context.Context, req *notedemo.MGetNoteRequest, callOptions ...callopt.Option) (r *notedemo.MGetNoteResponse, err error)
	DeleteNoteFunc func(ctx context.Context, req *notedemo.DeleteNoteRequest, callOptions ...callopt.Option) (r *notedemo.DeleteNoteResponse, err error)
	QueryNoteFunc  func(ctx context.Context, req *notedemo.QueryNoteRequest, callOptions ...callopt.Option) (r *notedemo.QueryNoteResponse, err error)
	UpdateNoteFunc func(ctx context.Context, req *notedemo.UpdateNoteRequest, callOptions ...callopt.Option) (r *notedemo.UpdateNoteResponse, err error)
}
