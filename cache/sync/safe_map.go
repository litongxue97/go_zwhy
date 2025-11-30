package sync

import "sync"

type SafeMap[K comparable, V any] struct {
	data  map[K]V
	mutex sync.RWMutex
}

func (s *SafeMap[K, V]) Put(key K, value V) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data[key] = value
}

func (s *SafeMap[K, V]) Get(key K) (any, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	res, ok := s.data[key]
	return res, ok
}

// double-check写法
func (s *SafeMap[K, V]) LoadOrStore(key K, newValue V) (val V, loaded bool) {
	s.mutex.RLock()
	res, ok := s.data[key]
	s.mutex.RUnlock()
	if ok {
		return res, true
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	// double-check begin
	res, ok = s.data[key]
	if ok {
		return res, true
	}
	// double-check end

	s.data[key] = newValue
	return newValue, false
}
