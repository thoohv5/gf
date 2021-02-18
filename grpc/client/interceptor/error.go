package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	libgrpc "github.com/thoohv5/gf/grpc"
	pberror "github.com/thoohv5/gf/grpc/demo/sdk/error"
)

func Error() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			s := status.Convert(err)
			if s.Code() == codes.Unknown {
				for _, d := range s.Details() {
					switch info := d.(type) {
					case *pberror.Error:
						err = libgrpc.NewError(libgrpc.Code(info.Code), info.Message)
					default:
						err = libgrpc.NewError(libgrpc.UnknownErrCode, "")
					}
				}
			} else if s.Code() == codes.InvalidArgument {
				err = libgrpc.NewError(libgrpc.InvalidArgument, "")
			}
		}
		return err
	}
}
