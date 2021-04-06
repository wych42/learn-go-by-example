package safemap

import "sync"

type SafeMap struct {
	data map[interface{}]interface{}
	mu   *sync.RWMutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[interface{}]interface{}),
		mu:   &sync.RWMutex{},
	}
}

func (m *SafeMap) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.data)
}

// Put: return is key exists
func (m *SafeMap) Put(key, value interface{}) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.data[key]
	m.data[key] = value
	return ok
}

func (m *SafeMap) Get(key interface{}) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, ok := m.data[key]
	return value, ok
}

// Delete: return is key exists
func (m *SafeMap) Delete(key interface{}) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.data[key]
	if ok {
		delete(m.data, key)
	}
	return ok
}
