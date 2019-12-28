package memo

import (
	"fmt"
	"sync"
	"time"
)

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	f     Func
	cache map[string]*entry
	mu    sync.Mutex
}

func New(f Func) *Memo {
	return &Memo{f, map[string]*entry{}, sync.Mutex{}}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	var start time.Time
	memo.mu.Lock()
	var test int
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		test = 100
		start = time.Now()
		memo.mu.Unlock()
		<-e.ready
	}
	fmt.Printf("%s, %s, %d\n", key, time.Since(start), test)
	return e.res.value, e.res.err
}
