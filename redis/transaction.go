package redis

import redigo "github.com/gomodule/redigo/redis"

func ErrHandler(redisConn redigo.Conn, task Task) (i interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = xerror.Wrap(xerror.New(fmt.Sprint(e)), "redis handler recover err")
		}
	}()
	return task(redisConn)
}

func ExecTrans(param ParamOption, trans ...Task) error {
	// 链接redis
	redisConn, err := Get(param.RedisConnName)
	if nil != err {
		return xerror.Wrap(err, "redis conn err")
	}
	defer func() {
		if err := redisConn.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	if _, err := redisConn.Do("MULTI"); err != nil {
		return xerror.Wrapf(err, "redis multi err")
	}
	for _, task := range trans {
		if _, err := ErrHandler(redisConn, task); err != nil {
			if _, err := redisConn.Do("DISCARD"); err != nil {
				return xerror.Wrapf(err, "redis discard err : %s", err.Error())
			}
			return xerror.Wrapf(err, "redis handler err : %s", err.Error())
		}
	}

	if _, err := redisConn.Do("EXEC"); err != nil {
		if _, err := redisConn.Do("DISCARD"); err != nil {
			return xerror.Wrapf(err, "redis discard err : %s", err.Error())
		}
		return xerror.Wrapf(err, "redis exec err : %s", err.Error())
	}
	return nil
}
