package safemap

import (
	cmap "github.com/orcaman/concurrent-map"
)

type WrapConcurrentMap struct {
	cm cmap.ConcurrentMap
}

func NewWrapConcurrentMap() *WrapConcurrentMap {
	return &WrapConcurrentMap{
		cm: cmap.New(),
	}
}

func (m *WrapConcurrentMap) Len() int {
	return m.cm.Count()
}

// Put: return is key exists
func (m *WrapConcurrentMap) Store(key, value interface{}) {
	m.cm.Set(key.(string), value)
}

func (m *WrapConcurrentMap) Load(key interface{}) (interface{}, bool) {
	return m.cm.Get(key.(string))
}

func (m *WrapConcurrentMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	panic("not implemented")
}

func (m *WrapConcurrentMap) Delete(key interface{}) {
	m.cm.Remove(key.(string))
}

// Delete: return is key exists
func (m *WrapConcurrentMap) LoadAndDelete(key interface{}) (value interface{}, loaded bool) {
	panic("not implemented")
}

func (m *WrapConcurrentMap) Range(f func(key, value interface{}) (shouldContinue bool)) {
	for item := range m.cm.Iter() {
		ok := f(item.Key, item.Val)
		if !ok {
			return
		}
	}
}
