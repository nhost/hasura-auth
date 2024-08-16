package ratelimit

import (
	"sync"
	"time"
)

type InMemoryStoreValue struct {
	v    int
	time time.Time
}

type InMemoryStore struct {
	data map[string]InMemoryStoreValue
	mx   sync.RWMutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		data: make(map[string]InMemoryStoreValue),
		mx:   sync.RWMutex{},
	}
}

func (i *InMemoryStore) DeleteExpired() {
	i.mx.Lock()
	defer i.mx.Unlock()

	for k, v := range i.data {
		if time.Now().After(v.time) {
			delete(i.data, k)
		}
	}
}

func (i *InMemoryStore) Get(key string) int {
	i.DeleteExpired()

	i.mx.RLock()
	defer i.mx.RUnlock()

	if v, ok := i.data[key]; ok {
		return v.v
	}
	return 0
}

func (i *InMemoryStore) Increment(key string, expire time.Duration) int {
	i.DeleteExpired()

	i.mx.Lock()
	defer i.mx.Unlock()

	current := i.Get(key)
	i.data[key] = InMemoryStoreValue{
		v:    current + 1,
		time: time.Now().Add(expire),
	}

	return i.data[key].v
}
