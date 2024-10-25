package core

import (
	"sync"
)

type MutexDS[T comparable, K any] struct {
	mu sync.RWMutex
	m  map[T]K
}

func NewMutexDS[T comparable, K any]() *MutexDS[T, K] {
	return &MutexDS[T, K]{
		mu: sync.RWMutex{},
		m:  map[T]K{},
	}
}

func (ds *MutexDS[T, K]) Get(key T) K {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return ds.m[key]
}
func (ds *MutexDS[T, K]) Delete(key T) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	delete(ds.m, key)
}
func (ds *MutexDS[T, K]) Exists(key T) bool {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	_, ok := ds.m[key]
	return ok
}
func (ds *MutexDS[T, K]) Set(key T, val K) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.m[key] = val
}
func (ds *MutexDS[T, K]) Keys() (ans []T) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	for k := range ds.m {
		ans = append(ans, k)
	}
	return ans
}
func (ds *MutexDS[T, K]) GetInner() map[T]K {
	return ds.m
}
func (ds *MutexDS[T, K]) SetInner(in map[T]K) {
	ds.m = in
}
func (ds *MutexDS[T, K]) Clear() {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.m = map[T]K{}
}
