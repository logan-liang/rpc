// Code generated by goctl. DO NOT EDIT!
// Source: demo.proto

package server

import (
	"context"

	"demo/test/rpc/demo"
	"demo/test/rpc/internal/logic"
	"demo/test/rpc/internal/svc"
)

type DemoServer struct {
	svcCtx *svc.ServiceContext
	demo.UnimplementedDemoServer
}

func NewDemoServer(svcCtx *svc.ServiceContext) *DemoServer {
	return &DemoServer{
		svcCtx: svcCtx,
	}
}

func (s *DemoServer) Ping(ctx context.Context, in *demo.Request) (*demo.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}