package context

import (
	"context"
	"testing"
	"time"
)

type mykey struct {
}

func TestContext(t *testing.T) {
	// 调用起点
	ctx := context.Background()
	ctx = context.WithValue(ctx, mykey{}, "value")

	val := ctx.Value(mykey{}).(string)
	t.Log(val)

	// 不存在的key
	newVal := ctx.Value("不存在的key")
	// 防止panic
	val, ok := newVal.(string)
	if !ok {
		t.Log("类型不对")
		return
	}
	t.Log(val)
}

func TestContext_WithCancel(t *testing.T) {
	// 调用起点
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	// defer cancel()

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	<-ctx.Done()
	t.Log("hello, cancel: ", ctx.Err())
}

func TestContext_WithDeadline(t *testing.T) {
	// 调用起点
	ctx := context.Background()

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	deadline, _ := ctx.Deadline()
	t.Log("deadline: ", deadline)
	defer cancel()

	<-ctx.Done()
	t.Log("hello, deadline: ", ctx.Err())
}

func TestContext_WithTimeOut(t *testing.T) {
	// 调用起点
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	deadline, _ := ctx.Deadline()
	t.Log("deadline: ", deadline)
	defer cancel()

	<-ctx.Done()
	t.Log("hello, timeout: ", ctx.Err())
}
