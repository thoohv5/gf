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

func Logger() grpc.UnaryServerInterceptor {
	// opts := []grpc_logrus.Option{
	// 	grpc_logrus.WithDurationField(grpc_logrus.DurationToDurationField),
	// }
	// logger := logrus.New()
	// logger.Out = os.Stdout
	// logger.Formatter = &logrus.JSONFormatter{DisableTimestamp: true}
	// logger.Level = logrus.DebugLevel // a lot of our stuff is on debug level by default
	// return grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger), opts...)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()
		fields := newServerLoggerFields(ctx, info.FullMethod, startTime)
		resp, err := handler(ctx, req)
		fields = append(fields,
			standard.NewField("grpc.time_ms", time.Since(startTime)),
			standard.NewField("grpc.code", status.Code(err)),
		)
		log.Info("grpc server:", fields...)
		return resp, err
	}
}

func newServerLoggerFields(ctx context.Context, fullMethodString string, start time.Time) []standard.Field {
	service := path.Dir(fullMethodString)[1:]
	method := path.Base(fullMethodString)

	fields := make([]standard.Field, 0)
	fields = append(fields, standard.NewField("grpc.start_time", start.Format(time.RFC3339)))
	if d, ok := ctx.Deadline(); ok {
		fields = append(fields, standard.NewField("grpc.request.deadline", d.Format(time.RFC3339)))
	}
	return []standard.Field{
		standard.NewField("system", "grpc"),
		standard.NewField("span.kind", "server"),
		standard.NewField("grpc.service", service),
		standard.NewField("grpc.method", method),
	}
}
