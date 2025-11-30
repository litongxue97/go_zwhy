package channel

import (
	"errors"
	"sync"
)

type Broker struct {
	mutex sync.RWMutex
	chans []chan Msg
}

func (b *Broker) Send(m Msg) error {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	for _, ch := range b.chans {
		select {
		case ch <- m:
		default:
			return errors.New("消息队列已经满了")
		}
	}
	return nil
}

func (b *Broker) Subscribe(cap int) (<-chan Msg, error) {
	res := make(chan Msg, cap)
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.chans = append(b.chans, res)
	return res, nil
}

func (b *Broker) Close() error {
	b.mutex.Lock()
	chans := b.chans
	b.chans = nil
	b.mutex.Unlock()
	// 避免重复 close channel
	for _, ch := range chans {
		close(ch)
	}
	return nil
}

type Msg struct {
	Content string
}
