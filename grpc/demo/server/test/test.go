package test

import (
	"context"

	pbuser "github.com/thoohv5/gf/grpc/demo/sdk/user"
)

// 测试服务
type testServer struct {
	pbuser.UnimplementedUserServer
}

func NewServer() *testServer {
	return &testServer{}
}

// 用户信息
func (ts *testServer) Info(ctx context.Context, req *pbuser.InfoReq) (*pbuser.InfoResp, error) {
	resp := new(pbuser.InfoResp)
	resp.Name = "name"
	resp.Code = "code"
	return resp, nil
}
