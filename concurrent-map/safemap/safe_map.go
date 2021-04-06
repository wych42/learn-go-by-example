package safemap

import "sync"

type SafeMap struct {
	data map[interface{}]interface{}
	mu   sync.RWMutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[interface{}]interface{}),
		mu:   sync.RWMutex{},
	}
}

func (m *SafeMap) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.data)
}

// Put: return is key exists
func (m *SafeMap) Store(key, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = make(map[interface{}]interface{})
	}
	m.data[key] = value
}

func (m *SafeMap) Load(key interface{}) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, ok := m.data[key]
	return value, ok
}

func (m *SafeMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	actual, loaded = m.data[key]
	if !loaded {
		if m.data == nil {
			m.data = make(map[interface{}]interface{})
		}
		m.data[key] = value
		actual = value
	}
	return
}

func (m *SafeMap) Delete(key interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.data[key]
	if ok {
		delete(m.data, key)
	}
	return
}

// Delete: return is key exists
func (m *SafeMap) LoadAndDelete(key interface{}) (value interface{}, loaded bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	value, loaded = m.data[key]
	if loaded {
		delete(m.data, key)
	}
	return
}

func (m *SafeMap) Range(f func(key, value interface{}) (shouldContinue bool)) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		ok := f(k, v)
		if !ok {
			return
		}
	}
}
