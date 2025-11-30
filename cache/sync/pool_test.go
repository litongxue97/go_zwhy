package sync

import (
	"sync"
	"testing"
)

// 资源位被复用
func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			t.Log("创建资源了。。。。")
			return "hello"
		},
	}
	str := pool.Get().(string)
	t.Log(str)
	pool.Put(str)
	str = pool.Get().(string)
	t.Log(str)
	pool.Put(str)
}
