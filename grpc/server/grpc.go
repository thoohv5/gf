package server

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"

	clientinterceptor "github.com/thoohv5/gf/grpc/client/interceptor"
	"github.com/thoohv5/gf/grpc/server/interceptor"
)

type Grpc struct {
}

type RegisterGrpc func(s *grpc.Server) error
type RegisterHttp func(mux *http.ServeMux) error
type RegisterGrpcHttp func(ctx context.Context, mux *runtime.ServeMux, addr string, dialOption []grpc.DialOption) error
type RegisterValidate func(validator.Validate) error

type Server interface {
	Server(config *Config, registerServer RegisterGrpc, registerGrpcHttp RegisterGrpcHttp, registerHttp RegisterHttp) error
}

type Config struct {
	Network string
	Address string
}

func NewServer() Server {
	return &Grpc{}
}

func (s *Grpc) Server(config *Config, registerGrpc RegisterGrpc, registerGrpcHttp RegisterGrpcHttp, registerHttp RegisterHttp) error {
	// lis, err := net.Listen(config.Network, config.Address)
	// if err != nil {
	// 	return errors.Wrap(err, "net Listen err")
	// }
	// grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor.RateLimit(), interceptor.Logger(), interceptor.Error(), interceptor.Auth(), interceptor.Recovery(), interceptor.Validator()))
	//
	// if err := registerGrpc(grpcServer); err != nil {
	// 	return errors.Wrap(err, "grpc Grpc RegisterGrpc err")
	// }
	//
	// if err := grpcServer.Serve(lis); err != nil {
	// 	return errors.Wrap(err, "grpc Grpc Serve err")
	// }

	// net listen
	conn, err := net.Listen(config.Network, config.Address)
	if nil != err {
		return errors.Wrap(err, "net Listen err")
	}

	// 中间件
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.RateLimit(),
			interceptor.Logger(),
			interceptor.Error(),
			// interceptor.Auth(),
			interceptor.Recovery(),
			interceptor.Validator(),
			interceptor.Opentracing(),
			// grpc_opentracing.UnaryServerInterceptor(),
		),
	)

	// grpc
	if err := registerGrpc(grpcServer); err != nil {
		return errors.Wrap(err, "grpc Grpc RegisterGrpc err")
	}

	/**
	http start
	*/
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	runtimeServerMux := runtime.NewServeMux()

	// grpc http
	if err := registerGrpcHttp(ctx, runtimeServerMux, config.Address, []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(
			clientinterceptor.Logger(),
			clientinterceptor.Error(),
			clientinterceptor.Opentracing(),
		),
	}); nil != err {
		return errors.Wrap(err, "grpc Grpc RegisterGrpc err")
	}

	httpServeMux := http.NewServeMux()
	httpServeMux.Handle("/", runtimeServerMux)

	// http
	if err := registerHttp(httpServeMux); nil != err {
		return errors.Wrap(err, "http RegisterHttp err")
	}

	httpServer := &http.Server{
		Addr:    config.Address,
		Handler: serverHandlerFunc(grpcServer, httpServeMux),
	}
	/**
	http end
	*/

	/**
	http server
	grpcServer => httpServer
	*/
	if err := httpServer.Serve(conn); nil != err {
		return errors.Wrap(err, "http Serve err")
	}

	return err
}

// server hand func
func serverHandlerFunc(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
