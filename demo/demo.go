// Code generated by goctl. DO NOT EDIT!
// Source: demo.proto

package demo

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Demo interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultDemo struct {
		cli zrpc.Client
	}
)

func NewDemo(cli zrpc.Client) Demo {
	return &defaultDemo{
		cli: cli,
	}
}

func (m *defaultDemo) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := NewDemoClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
