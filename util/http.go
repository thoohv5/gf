package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	nethttp "net/http"
	neturl "net/url"
	"time"
)

type (
	IHttp interface {
		// get
		Get(ctx context.Context, url string, result interface{}, opts ...Option) error
		// post
		Post(ctx context.Context, url string, param map[string]interface{}, result interface{}, opts ...Option) error
	}
	// 参数
	Parameter struct {
		// 请求方式
		method Method
		// 超时时间
		timeout time.Duration
		// header
		header map[string]string
		// param
		param map[string]interface{}
		// preDeal
		preDeal []func(r *Parameter) error
		// reader
		body io.Reader
	}
	Option interface {
		apply(*Parameter)
	}
	optionFunc func(*Parameter)
	http       struct {
		*Parameter
	}
)

func (f optionFunc) apply(o *Parameter) {
	f(o)
}

func (p *Parameter) SetBody(body io.Reader) {
	p.body = body
}

type Method string

const (
	HttpTimeOut        = 10 * time.Second
	HttpPost    Method = "POST"
	HttpGet     Method = "GET"
)

func NewHttp() IHttp {
	h := &http{
		Parameter: &Parameter{
			method:  HttpGet,
			timeout: HttpTimeOut,
			header: map[string]string{
				"Content-Type": "application/json",
			},
			param: map[string]interface{}{},
		},
	}
	return h
}

func (r *http) withOpt(opts ...Option) error {
	for _, o := range opts {
		o.apply(r.Parameter)
	}
	return nil
}

func WithMethod(method Method) Option {
	return optionFunc(func(r *Parameter) {
		r.method = method
	})
}

func WithTimeout(timeout time.Duration) Option {
	return optionFunc(func(r *Parameter) {
		r.timeout = timeout * time.Second
	})
}

func WithParam(params map[string]interface{}) Option {
	return optionFunc(func(r *Parameter) {
		for key, val := range params {
			r.param[key] = val
		}
	})
}

func WithHeader(headers map[string]string) Option {
	return optionFunc(func(r *Parameter) {
		for key, val := range headers {
			r.header[key] = val
		}
	})
}

func WithPreDeal(preDeal func(r *Parameter) error) Option {
	return optionFunc(func(r *Parameter) {
		r.preDeal = append(r.preDeal, preDeal)
	})
}

// HttpPost
func (r *http) request(ctx context.Context, url string, result interface{}, opts ...Option) error {

	// 可选参数
	if err := r.withOpt(opts...); nil != err {
		return fmt.Errorf("request withOpt err, opts: %v, %w", opts, err)
	}

	// 预处理
	for _, preDeal := range r.preDeal {
		if err := preDeal(r.Parameter); nil != err {
			return fmt.Errorf("request callback err, r: %v, %w", r, err)
		}
	}

	// 组装request
	req, err := nethttp.NewRequestWithContext(ctx, string(r.method), url, r.body)
	if nil != err {
		return fmt.Errorf("request NewRequestWithContext err, url: %s, body: %s, %w", url, r.body, err)
	}

	// 组装header
	for key, value := range r.header {
		req.Header.Set(key, value)
	}

	// 发送请求
	client := &nethttp.Client{Timeout: r.timeout}
	resp, err := client.Do(req)
	if nil != err {
		return fmt.Errorf("request do err, param: %v, %w", req, err)
	}
	defer func() {
		if err = resp.Body.Close(); nil != err {
			fmt.Println(err)
		}
	}()

	// 读取请求
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return fmt.Errorf("request read err, bodyByte: %v, %w", bodyByte, err)
	}

	// 解析请求值
	err = json.Unmarshal(bodyByte, &result)
	if nil != err {
		return fmt.Errorf("request json un err, result: %v, %w", result, err)
	}

	return nil
}

// Get
func (r *http) Get(ctx context.Context, url string, result interface{}, opts ...Option) error {

	// withMethod, WithPreDeal
	opts = append(opts, WithMethod(HttpGet), WithPreDeal(func(r *Parameter) error {
		// 组装url
		params := neturl.Values{}
		netUrl, err := neturl.Parse(url)
		if err != nil {
			return fmt.Errorf("get json ma err, param: %s, %w", url, err)
		}
		for key, value := range r.param {
			// todo 这儿可以优化
			params.Add(key, fmt.Sprintf("%s", value))
		}
		netUrl.RawQuery = params.Encode()
		url = netUrl.String()
		return nil
	}))

	return r.request(ctx, url, result, opts...)
}

// Post
func (r *http) Post(ctx context.Context, url string, param map[string]interface{}, result interface{}, opts ...Option) error {

	// withMethod, withParam, WithPreDeal
	opts = append(opts, WithMethod(HttpPost), WithParam(param), WithPreDeal(func(r *Parameter) error {
		if nil == param {
			return nil
		}
		// 组装param
		data, err := json.Marshal(param)
		if nil != err {
			return fmt.Errorf("post json ma err, param: %s, %w", param, err)
		}
		r.SetBody(bytes.NewBuffer(data))
		return nil
	}))

	return r.request(ctx, url, result, opts...)
}

// Get
func Get(ctx context.Context, url string, result interface{}, opts ...Option) error {
	return NewHttp().Get(ctx, url, result, opts...)
}

// Post
func Post(ctx context.Context, url string, param map[string]interface{}, result interface{}, opts ...Option) error {
	return NewHttp().Post(ctx, url, param, result, opts...)
}
