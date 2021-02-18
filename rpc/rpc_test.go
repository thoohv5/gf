package rpc

import (
	"context"
	"testing"
)

type JsonRpc struct {
}

func (jr JsonRpc) GetConfig(name ModuleName) *Config {
	return &Config{
		Host:    "https://d2d-qa.medlinker.com",
		TimeOut: 100,
	}
}

type SvrRpc struct {
	JsonRpc
}

func (svr SvrRpc) ModuleName() ModuleName {
	return "d2d"
}

type InquiryService interface {
	GetCountInPool(ctx context.Context, param *GetCountInPoolReq, ret *GetCountInPoolResp) error
}

// 请求值
type GetCountInPoolReq struct {
	DoctorId uint32
}

// 返回值
type GetCountInPoolResp struct {
	Count uint32
}

// 服务
type inquiryService struct {
	SvrRpc
}

const (
	inquiryServiceName ServiceName = "rpc/inquiry_v2"
)

// 初始化服务
func NewInquiry() InquiryService {
	return &inquiryService{}
}

func (s *inquiryService) ServiceName() ServiceName {
	return inquiryServiceName
}

func (s *inquiryService) GetCountInPool(ctx context.Context, param *GetCountInPoolReq, ret *GetCountInPoolResp) error {

	// 参数拼接
	params := map[string]interface{}{
		"doctorId": param.DoctorId,
	}

	// 参数构造
	reqArgs := &ReqArgs{
		Method: "getCountInPool",
		Params: params,
	}

	// 请求
	err := NewRpc(s).Call(ctx, reqArgs, &ret.Count)
	if err != nil {
		return err
	}

	return nil
}

func TestCountInPool(t *testing.T) {
	ret := new(GetCountInPoolResp)
	err := NewInquiry().GetCountInPool(context.Background(), &GetCountInPoolReq{DoctorId: 877803565}, ret)
	t.Log(err, ret)
}
