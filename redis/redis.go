package redis

import (
	redigo "github.com/gomodule/redigo/redis"

	"github.com/thoohv5/gf/util"
)

type redis struct {
}

type Task func(redisConn redigo.Conn) (interface{}, error)

func ErrHandler(redisConn redigo.Conn, task Task) (i interface{}, err error) {
	util.WrapRecover()
	return task(redisConn)
}
