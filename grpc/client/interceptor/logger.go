package interceptor

import (
	"context"
	"path"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/thoohv5/gf/log"
	"github.com/thoohv5/gf/log/standard"
)

func Logger() grpc.UnaryClientInterceptor {
	// opts := []grpc_logrus.Option{
	// 	grpc_logrus.WithDurationField(grpc_logrus.DurationToDurationField),
	// }
	// logger := logrus.New()
	// logger.Out = os.Stdout
	// logger.Formatter = &logrus.JSONFormatter{DisableTimestamp: true}
	// logger.Level = logrus.DebugLevel // a lot of our stuff is on debug level by default
	// return grpc_logrus.UnaryClientInterceptor(logrus.NewEntry(logger), opts...)

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		startTime := time.Now()
		fields := newClientLoggerFields(ctx, method, startTime)
		err := invoker(ctx, method, req, reply, cc, opts...)

		fields = append(fields,
			standard.NewField("grpc.time_ms", time.Since(startTime)),
			standard.NewField("grpc.code", status.Code(err)),
		)
		log.Info("grpc client:", fields...)
		return err
	}
}

func newClientLoggerFields(ctx context.Context, fullMethodString string, start time.Time) []standard.Field {
	service := path.Dir(fullMethodString)[1:]
	method := path.Base(fullMethodString)

	fields := make([]standard.Field, 0)
	fields = append(fields, standard.NewField("grpc.start_time", start.Format(time.RFC3339)))
	if d, ok := ctx.Deadline(); ok {
		fields = append(fields, standard.NewField("grpc.request.deadline", d.Format(time.RFC3339)))
	}
	return []standard.Field{
		standard.NewField("system", "grpc"),
		standard.NewField("span.kind", "client"),
		standard.NewField("grpc.service", service),
		standard.NewField("grpc.method", method),
	}
}
