package util

import (
	"context"
	"testing"
)

// Post
func TestPost(t *testing.T) {
	type args struct {
		ctx    context.Context
		url    string
		param  map[string]interface{}
		result interface{}
		opts   []Option
	}

	ret := new(struct {
		Args struct {
			Age  string `json:"age"`
			Name string `json:"name"`
		} `json:"args"`
		Headers struct {
			Accept                    string `json:"Accept"`
			Accept_Encoding           string `json:"Accept-Encoding"`
			Accept_Language           string `json:"Accept-Language"`
			Dnt                       string `json:"Dnt"`
			Host                      string `json:"Host"`
			Upgrade_Insecure_Requests string `json:"Upgrade-Insecure-Requests"`
			User_Agent                string `json:"User-Agent"`
			X_Amzn_Trace_Id           string `json:"X-Amzn-Trace-Id"`
		} `json:"headers"`
		Origin string `json:"origin"`
		URL    string `json:"url"`
	})

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Post",
			args{
				ctx: context.Background(),
				url: "http://httpbin.org/post",
				param: map[string]interface{}{
					"name": "zhaofan",
					"age":  23,
				},
				result: ret,
				opts:   nil,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Post(tt.args.ctx, tt.args.url, tt.args.param, tt.args.result, tt.args.opts...); (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		t.Log(tt.args.result)
	}
}

// Get
func TestGet(t *testing.T) {
	type args struct {
		ctx    context.Context
		url    string
		result interface{}
		opts   []Option
	}

	ret := new(struct {
		Args struct {
			Age  string `json:"age"`
			Name string `json:"name"`
		} `json:"args"`
		Headers struct {
			Accept                    string `json:"Accept"`
			Accept_Encoding           string `json:"Accept-Encoding"`
			Accept_Language           string `json:"Accept-Language"`
			Dnt                       string `json:"Dnt"`
			Host                      string `json:"Host"`
			Upgrade_Insecure_Requests string `json:"Upgrade-Insecure-Requests"`
			User_Agent                string `json:"User-Agent"`
			X_Amzn_Trace_Id           string `json:"X-Amzn-Trace-Id"`
		} `json:"headers"`
		Origin string `json:"origin"`
		URL    string `json:"url"`
	})

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Get",
			args{
				ctx:    context.Background(),
				url:    "http://httpbin.org/get?name=zhaofan&age=23",
				result: ret,
				opts:   nil,
			},
			false,
		},
		{
			"Get1",
			args{
				ctx:    context.Background(),
				url:    "http://httpbin.org/get",
				result: ret,
				opts: []Option{
					WithParam(map[string]interface{}{
						"name": "zhaofan",
						"age":  23,
					}),
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Get(tt.args.ctx, tt.args.url, tt.args.result, tt.args.opts...); (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		t.Log(tt.args.result)
	}
}
