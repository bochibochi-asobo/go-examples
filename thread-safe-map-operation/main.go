package main

import (
	"fmt"
	"sync"
)

type KeyValue struct {
	cache map[string]string
	mu    sync.RWMutex // 排他制御のための mutex
}

func CreateNewCache() *KeyValue {
	return &KeyValue{cache: make(map[string]string)}
}

func (kv *KeyValue) Put(key, val string) {
	kv.mu.Lock()         // lock
	defer kv.mu.Unlock() // defer unlock
	kv.cache[key] = val
}

func (kv *KeyValue) Get(key string) (string, bool) {
	kv.mu.RLock()         // read lock
	defer kv.mu.RUnlock() // defer unlock
	val, ok := kv.cache[key]
	return val, ok
}

func main() {
	cache := CreateNewCache()

	cache.Put("key1", "value1")
	if value, ok := cache.Get("key1"); ok {
		fmt.Println(value)
	}
}
