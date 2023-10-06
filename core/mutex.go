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
func (ds *MutexDS[T, K]) Set(key T, val K) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.m[key] = val
}
func (ds *MutexDS[T, K]) GetInner() map[T]K {
	return ds.m
}
func (ds *MutexDS[T, K]) Clear() {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.m = map[T]K{}
}
