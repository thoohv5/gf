package lock

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

func TestLock(t *testing.T) {

	lk := InitLock(Redis)

	wg := new(sync.WaitGroup)
	count := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer lk.Unlock()
			if err := lk.Lock(); nil != err {
				log.Println("lk fail")
				return
			}
			count++
			log.Println(fmt.Sprintf("count: %v", count))
		}(i)

	}
	wg.Wait()

}
