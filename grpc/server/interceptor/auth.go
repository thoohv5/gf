package interceptor

import (
	"context"
	"fmt"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

func Auth() grpc.UnaryServerInterceptor {
	return grpc_auth.UnaryServerInterceptor(exampleAuthFunc)
}

// exampleAuthFunc is used by a middleware to authenticate requests
func exampleAuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "scheme")
	if err != nil {
		return nil, err
	}

	fmt.Println(token)

	return ctx, nil
}
