package sync

import (
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}
	var result int
	for i := 0; i < 10; i++ {
		// 每个goroutine都要增加waitgroup的计数
		wg.Add(1)
		go func(i int) {
			// 每个goroutine都要在结束时减少waitgroup的计数
			defer wg.Done()
			result += i
		}(i)
	}
	wg.Wait()
	t.Log(result)
}
