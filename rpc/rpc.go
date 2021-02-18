package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/thoohv5/gf/util"
)

type (
	// 模块名称
	ModuleName string
	// 服务名称
	ServiceName string
	// 模块名称
	Method string
	// 服务标准
	Rpc interface {
		GetConfig(name ModuleName) *Config
		ModuleName() ModuleName
		ServiceName() ServiceName
	}
	// 标准
	IRpc interface {
		// 调用
		Call(ctx context.Context, args *ReqArgs, result interface{}) error
	}
	// 服务
	rpc struct {
		Rpc
		*Parameter
	}
	// 参数
	Parameter struct {
		// 配置
		Config
		// 服务名称
		serviceName ServiceName
		// 方法名称
		methodName Method
	}
	Option interface {
		apply(*Parameter)
	}
	optionFunc func(*Parameter)
	// 配置
	Config struct {
		Host    string `toml:"host"`
		TimeOut uint32 `toml:"timeout"`
	}
	// 请求参数
	ReqArgs struct {
		JsonRpc string      `json:"jsonrpc"`
		Method  Method      `json:"method"`
		Params  interface{} `json:"params"`
		ID      uint8       `json:"id"`
	}
	// 返回值
	Resp struct {
		JsonRpc string      `json:"jsonrpc"`
		Result  interface{} `json:"result"`
		ID      uint8       `json:"id"`
		Error   interface{} `json:"error"`
	}
)

const (
	Version   = "2.0"
	DefaultID = 1
	TimeOut   = 3
)

func New(opts ...Option) IRpc {
	return NewRpc(nil, opts...)
}

func NewRpc(sRpc Rpc, opts ...Option) IRpc {
	svr := &rpc{
		Rpc: sRpc,
		Parameter: &Parameter{
			Config: Config{
				TimeOut: TimeOut,
			},
			methodName: Method(defaultFunc()),
		},
	}

	if svr.Rpc != nil {
		moduleName := svr.ModuleName()
		serviceName := svr.ServiceName()
		rpcConfig := svr.GetConfig(moduleName)
		host := rpcConfig.Host
		timeOut := rpcConfig.TimeOut
		opts = append(opts, WithHost(host), WithTimeOut(timeOut), WithServiceName(serviceName))
	}

	for _, opt := range opts {
		opt.apply(svr.Parameter)
	}

	return svr
}

func (f optionFunc) apply(o *Parameter) {
	f(o)
}

// 支持回调
func (r *rpc) Call(ctx context.Context, args *ReqArgs, result interface{}) error {

	url := fmt.Sprintf("%s/%s", r.Host, r.serviceName)
	args.JsonRpc = Version
	args.ID = DefaultID
	if len(args.Method) == 0 {
		args.Method = Method(defaultFunc())
	}

	// 返回值
	rpcResp := Resp{}
	if nil != result {
		rpcResp.Result = &result
	}

	// Post
	if err := util.Post(ctx, url, nil, &rpcResp, util.WithTimeout(time.Duration(r.TimeOut)), util.WithPreDeal(func(r *util.Parameter) error {
		// 组装param
		data, err := json.Marshal(args)
		if nil != err {
			return fmt.Errorf("call json ma err, param: %v, %w", args, err)
		}
		r.SetBody(bytes.NewBuffer(data))
		return nil
	})); nil != err {
		return fmt.Errorf("call Post error: url: %s, args:%v, %w", url, args, err)
	}

	// 检查返回值
	if nil != rpcResp.Error {
		rpcErrorByte, err := json.Marshal(rpcResp.Error)
		return fmt.Errorf("call json ma error: url:%s, rpcErr:%v, %w", url, string(rpcErrorByte), err)
	}

	return nil
}

func WithHost(host string) Option {
	return optionFunc(func(r *Parameter) {
		r.Host = host
	})
}

func WithTimeOut(timeOut uint32) Option {
	return optionFunc(func(r *Parameter) {
		r.TimeOut = timeOut
	})
}

func WithServiceName(serviceName ServiceName) Option {
	return optionFunc(func(r *Parameter) {
		r.serviceName = serviceName
	})
}

// 默认方法
func defaultFunc() string {
	method := ""
	pc, _, _, _ := runtime.Caller(2)
	if name := runtime.FuncForPC(pc).Name(); len(name) > 0 {
		nameArr := strings.Split(name, ".")
		method = nameArr[len(nameArr)-1]
	}
	return method
}
