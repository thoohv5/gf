package client

import (
	"github.com/pkg/errors"

	"google.golang.org/grpc"
)

type Grpc struct {
}

type Client interface {
	Client(clientConfig *Config) (*grpc.ClientConn, error)
}

type Config struct {
	Target string
	Opts   []grpc.DialOption
}

func NewClient() Client {
	return &Grpc{}
}

func (s *Grpc) Client(clientConfig *Config) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(clientConfig.Target, clientConfig.Opts...)
	if err != nil {
		return nil, errors.Wrap(err, "grpc Dial err")
	}
	return conn, nil
}
