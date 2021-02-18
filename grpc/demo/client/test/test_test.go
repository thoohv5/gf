package test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"

	libgrpc "github.com/thoohv5/gf/grpc"
	"github.com/thoohv5/gf/grpc/client"
	"github.com/thoohv5/gf/grpc/client/interceptor"
	pbuser "github.com/thoohv5/gf/grpc/demo/sdk/user"
	"github.com/thoohv5/gf/grpc/tracer"
)

func TestInfo(t *testing.T) {

	conn, err := client.NewClient().Client(&client.Config{
		Target: "127.0.0.1:8028",
		Opts: []grpc.DialOption{
			grpc.WithInsecure(),
			// grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"pick_first"}`),
			// grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
			grpc.WithChainUnaryInterceptor(interceptor.Logger(), interceptor.Error(),
				interceptor.Opentracing(),
				// grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(tracer)),
			),
		},
	})

	if nil != err {
		t.Fatal(err)
	}

	ctx := context.Background()
	newCtx, finish := tracer.Start("DoSomeThing", ctx)

	md := metadata.Pairs("authorization", fmt.Sprintf("%s %v", "scheme", "token"))
	nCtx := metautils.NiceMD(md).ToOutgoing(newCtx)

	resp, err := pbuser.NewUserClient(conn).Info(nCtx, &pbuser.InfoReq{
		Code: "12345",
	})
	finish(tracer.SpanWithError(errors.New("11111")))
	if grpcErr, ok := err.(*libgrpc.Error); ok {
		t.Log(grpcErr.Code())
	}
	t.Log(resp, err)

	time.Sleep(10 * time.Second)

}

func TestJaeger(t *testing.T) {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}
	closer, err := cfg.InitGlobalTracer(
		"serviceName",
	)
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	var ctx = context.TODO()
	span1, ctx := opentracing.StartSpanFromContext(ctx, "span_1")
	time.Sleep(time.Second / 2)
	span11, _ := opentracing.StartSpanFromContext(ctx, "span_1-1")
	time.Sleep(time.Second / 2)
	span11.Finish()
	span1.Finish()
	defer closer.Close()
}
