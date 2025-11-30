package channel

import "context"

type Task func()

type TaskPool struct {
	tasks chan Task
	close chan struct{}
}

// numG是goroutine的数量
// size是任务队列的大小
// 要有退出close的方法，不然会goroutine泄漏
func NewTaskPool(numG int, size int) *TaskPool {
	res := &TaskPool{
		tasks: make(chan Task, size),
		close: make(chan struct{}),
	}
	for i := 0; i < numG; i++ {
		go func() {
			for {
				select {
				// 如果是零值直接退出
				case <-res.close:
					return
				case t := <-res.tasks:
					t()
				}
			}
		}()
	}
	return res
}

func (p *TaskPool) Submit(ctx context.Context, t Task) error {
	select {
	case p.tasks <- t:
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func (p *TaskPool) Close() error {
	close(p.close)
	return nil
}
