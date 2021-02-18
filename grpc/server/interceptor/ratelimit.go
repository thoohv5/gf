package interceptor

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	"google.golang.org/grpc"
)

func RateLimit() grpc.UnaryServerInterceptor {
	return ratelimit.UnaryServerInterceptor(NewLimit())
}

type limit struct{}

func NewLimit() ratelimit.Limiter {
	return &limit{}
}

func (*limit) Limit() bool {
	return false
}
