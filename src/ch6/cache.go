package main

import "sync"

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	defer cache.Unlock()
	cache.Lock()
	v := cache.mapping[key]
	return v
}
