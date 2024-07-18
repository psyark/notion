package objects

import (
	"sync"
)

type SyncMap[K comparable, V any] struct{ m sync.Map }

func (sm *SyncMap[K, V]) Load(key K) (V, bool) {
	v, ok := sm.m.Load(key)
	if ok {
		return v.(V), ok
	}

	var zero V
	return zero, false
}

func (sm *SyncMap[K, V]) Store(key K, value V) { sm.m.Store(key, value) }

func (sm *SyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	a, loaded := sm.m.LoadOrStore(key, value)
	return a.(V), loaded
}

func (sm *SyncMap[K, V]) Delete(key K) { sm.m.Delete(key) }

func (sm *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	sm.m.Range(func(k, v any) bool { return f(k.(K), v.(V)) })
}
