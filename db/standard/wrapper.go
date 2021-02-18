package standard

import (
	"math/rand"
	"sync"
	"time"
)

// 数据库包装类
type Wrapper struct {
	sync.Mutex
	dsn   IBuilder
	slave []IBuilder
}

var (
	_random *rand.Rand
)

func init() {
	_random = rand.New(rand.NewSource(time.Now().Unix()))
}

func NewWrap(dsn IBuilder, slave ...IBuilder) *Wrapper {
	return &Wrapper{
		Mutex: sync.Mutex{},
		dsn:   dsn,
		slave: slave,
	}
}

// 写库
func (wrapper *Wrapper) Write() IBuilder {
	return wrapper.dsn
}

// 读库
func (wrapper *Wrapper) Read() IBuilder {
	defer wrapper.Unlock()
	wrapper.Lock()
	if len(wrapper.slave) == 0 {
		return wrapper.Write()
	}
	return wrapper.slave[_random.Intn(len(wrapper.slave))]
}
