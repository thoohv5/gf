package lock

import (
	"time"

	redigo "github.com/gomodule/redigo/redis"

	"github.com/thoohv5/gf/lock/impl"
	"github.com/thoohv5/gf/lock/standard"
)

type Type string

const (
	Redis Type = "redis"
)

func InitLock(lockType Type) standard.ILock {
	var lock standard.ILock
	switch lockType {
	case Redis:
		lock = impl.NewRedis(impl.WithRedisPool(&redigo.Pool{
			MaxIdle:     3,
			IdleTimeout: time.Duration(24) * time.Second,
			Dial: func() (redigo.Conn, error) {
				c, err := redigo.Dial("tcp", "127.0.0.1:6379")
				if err != nil {
					panic(err.Error())
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redigo.Conn, t time.Time) error {
				_, err := c.Do("PING")
				if err != nil {
					return err
				}
				return err
			},
		},
		),
		// impl.WithTries(1),
		)

	default:
		lock = impl.NewRedis()
	}
	return lock
}
