// 34_sync_package.go
// Topic: sync package — Mutex, RWMutex, Once, WaitGroup, Cond, Pool, Map
//        sync/atomic — lock-free primitives
//
// PRIMITIVES
// ----------
//   sync.Mutex       Lock(), Unlock()      — exclusive
//   sync.RWMutex     RLock/RUnlock + Lock/Unlock  — many readers OR one writer
//   sync.WaitGroup   Add(n), Done(), Wait()
//   sync.Once        Do(func())            — runs once across goroutines
//   sync.Cond        Wait/Signal/Broadcast — condition vars (rare)
//   sync.Pool        Get/Put               — reusable temp object pool
//   sync.Map         Load/Store/Delete     — concurrent map (specialized)
//   atomic.AddInt64, LoadInt64, StoreInt64, CompareAndSwap, atomic.Value
//
// RULES
//   - Don't COPY a Mutex. Use pointer or embed in struct.
//   - Always Unlock — `defer m.Unlock()` right after Lock.
//   - Prefer channels for ownership transfer; mutex for shared state.
//
// Run: go run -race 34_sync_package.go

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Mutex example: safe counter
type SafeCounter struct {
	mu sync.Mutex
	n  int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

// RWMutex: many readers, exclusive writer
type Cache struct {
	mu sync.RWMutex
	m  map[string]string
}

func (c *Cache) Get(k string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.m[k]
}

func (c *Cache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[k] = v
}

// Once: lazy init
var (
	once   sync.Once
	config map[string]string
)

func loadConfig() {
	once.Do(func() {
		fmt.Println("loading config...")
		config = map[string]string{"env": "prod"}
	})
}

func main() {
	// Counter
	c := &SafeCounter{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); c.Inc() }()
	}
	wg.Wait()
	fmt.Println("counter:", c.Value())

	// Cache
	cache := &Cache{m: map[string]string{}}
	cache.Set("a", "1")
	fmt.Println(cache.Get("a"))

	// Once
	loadConfig()
	loadConfig() // no-op
	fmt.Println(config)

	// atomic
	var x int64
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); atomic.AddInt64(&x, 1) }()
	}
	wg.Wait()
	fmt.Println("atomic x:", atomic.LoadInt64(&x))

	// CAS
	var ready int32
	swapped := atomic.CompareAndSwapInt32(&ready, 0, 1)
	fmt.Println("swapped:", swapped, "ready:", ready)

	// sync.Map (good for read-heavy / disjoint key sets)
	var sm sync.Map
	sm.Store("k", 1)
	v, ok := sm.Load("k")
	fmt.Println(v, ok)
	sm.Range(func(k, v any) bool {
		fmt.Println("sm:", k, v)
		return true
	})

	// sync.Pool (reuse short-lived objects)
	pool := &sync.Pool{
		New: func() any { return make([]byte, 1024) },
	}
	buf := pool.Get().([]byte)
	_ = buf
	pool.Put(buf)
}
