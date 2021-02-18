package interceptor

import (
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func Recovery() grpc.UnaryServerInterceptor {
	return grpc_recovery.UnaryServerInterceptor(recoveryInterceptor())
}

// recoveryInterceptor panic时返回Unknown错误吗
func recoveryInterceptor() grpc_recovery.Option {
	return grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		return grpc.Errorf(codes.Unknown, "panic triggered: %v", p)
	})
}
