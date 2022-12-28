package notedemo

import (
	"sync"

	"github.com/cloudwego/kitex/client"
)

var (
	// todo edit custom config
	defaultClient     NoteDemoClient
	defaultDstService = "demonote"
	defaultClientOpts = []client.Option{
		client.WithMuxConnection(1),
		client.WithHostPorts("127.0.0.1:8888"),
	}
	once       sync.Once
	mockClient = &MockClientStruct{}
)

func init() {
	DefaultClient()
}

func DefaultClient() NoteDemoClient {
	once.Do(func() {
		defaultClient = newClient(defaultDstService, defaultClientOpts...)
	})
	return defaultClient
}

func newClient(dstService string, opts ...client.Option) NoteDemoClient {
	c, err := NewNoteDemoClient(dstService, opts...)
	if err != nil {
		panic("failed to init client: " + err.Error())
	}
	c.SetMockClient(mockClient)
	return c
}

func InitClient(dstService string, opts ...client.Option) {
	defaultClient = newClient(dstService, opts...)
}
