package interceptor

import (
	"sync"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"github.com/thoohv5/gf/grpc/tracer"
)

var (
	_tracerOnce sync.Once
)

func Opentracing() grpc.UnaryClientInterceptor {
	_tracerOnce.Do(func() {
		// tracer, _ := initJaeger("gf-client")
		// defer closer.Close()
		_tracer, _ := tracer.InitJaeger("gf-client")
		opentracing.InitGlobalTracer(_tracer)
	})
	return grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(opentracing.GlobalTracer()))
}
