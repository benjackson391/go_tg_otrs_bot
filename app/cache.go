package app

import (
	"sync"
	"time"
)

type CacheItem struct {
	Value      string
	Expiration int64
}

type InterfaceItem struct {
	Value      interface{}
	Expiration int64
}

type CacheMap struct {
	items           map[string]CacheItem
	interface_itmes map[string]InterfaceItem
	mutex           sync.RWMutex
}

func NewCacheMap() *CacheMap {
	return &CacheMap{
		items:           make(map[string]CacheItem),
		interface_itmes: make(map[string]InterfaceItem),
	}
}

func (m *CacheMap) Set(key string, value string, duration time.Duration) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.items[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(duration).UnixNano(),
	}
}

func (m *CacheMap) Get(key string) (string, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	item, found := m.items[key]
	if !found {
		return "", false
	}
	if time.Now().UnixNano() > item.Expiration {
		delete(m.items, key)
		return "", false
	}
	return item.Value, true
}
