package standard

// 锁标准
type ILock interface {
	// 加锁
	Lock() error
	// 释放锁
	Unlock()
}
