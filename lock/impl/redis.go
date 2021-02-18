package impl

import (
	"time"

	"github.com/go-redsync/redsync"
	redigo "github.com/gomodule/redigo/redis"

	"github.com/thoohv5/gf/lock/standard"
)

type redis struct {
	*redsync.Mutex
}

// 可选参数
type parameter struct {
	// 连接池
	pool *redigo.Pool
	// key
	name string
	// 过期时间，默认 8s
	expiry time.Duration
	// 重试次数，默认 32
	tries int
	// 重试的间隔时间, 默认 500ms
	delayFunc redsync.DelayFunc
	// 随机值函数
	genValueFunc func() (string, error)
}

type Option interface {
	apply(*parameter)
}

const (
	__defaultKey = "default_mutex_key"
)

type optionFunc func(*parameter)

func (f optionFunc) apply(o *parameter) {
	f(o)
}

func WithRedisPool(pool *redigo.Pool) Option {
	return optionFunc(func(o *parameter) {
		o.pool = pool
	})
}

func WithName(name string) Option {
	return optionFunc(func(o *parameter) {
		o.name = name
	})
}

func WithExpiry(expiry time.Duration) Option {
	return optionFunc(func(o *parameter) {
		o.expiry = expiry
	})
}

func WithTries(tries int) Option {
	return optionFunc(func(o *parameter) {
		o.tries = tries
	})
}

func NewRedis(opts ...Option) standard.ILock {
	p := &parameter{
		// key
		name: __defaultKey,
		// 过期时间，默认 8s
		expiry: 8 * time.Second,
		// 重试次数，默认 32
		tries: 32,
	}
	for _, opt := range opts {
		opt.apply(p)
	}

	return &redis{
		Mutex: redsync.New([]redsync.Pool{p.pool}).NewMutex(p.name, redsync.SetExpiry(p.expiry), redsync.SetTries(p.tries)),
	}
}

func (r *redis) Lock() error {
	return r.Mutex.Lock()
}

func (r *redis) Unlock() {
	r.Mutex.Unlock()
}
