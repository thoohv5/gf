package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	libgrpc "github.com/thoohv5/gf/grpc"
	pberror "github.com/thoohv5/gf/grpc/demo/sdk/error"
)

func Error() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		m, err := handler(ctx, req)

		if nil == err {
			return m, err
		}

		grpErr, ok := err.(*libgrpc.Error)
		if !ok {
			return m, err
		}

		st := status.New(codes.Unknown, "Unknown")
		ds, err := st.WithDetails(
			&pberror.Error{
				Code:    grpErr.Code(),
				Message: grpErr.Error(),
				Detail:  grpErr.Detail(),
			},
		)
		if nil != err {
			return m, err
		}

		return m, ds.Err()
	}
}
