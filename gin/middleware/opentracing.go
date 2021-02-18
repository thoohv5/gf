package middleware

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"github.com/thoohv5/gf/grpc/tracer"
)

var (
	_tracerOnce sync.Once
)

func Opentracing() gin.HandlerFunc {
	_tracerOnce.Do(func() {
		// tracer, _ := initJaeger("gf-client")
		// defer closer.Close()
		_tracer, _ := tracer.InitJaeger("gf-http")
		opentracing.InitGlobalTracer(_tracer)
	})
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		span := opentracing.GlobalTracer().StartSpan(path,
			ext.SpanKindRPCServer)
		ext.HTTPUrl.Set(span, path)
		ext.HTTPMethod.Set(span, ctx.Request.Method)
		c := opentracing.ContextWithSpan(ctx, span)
		ctx.Set("ctx", c)
		ctx.Next()
		ext.HTTPStatusCode.Set(span, uint16(ctx.Writer.Status()))
		span.Finish()
	}
}
