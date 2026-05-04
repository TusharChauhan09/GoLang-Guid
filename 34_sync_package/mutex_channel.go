package main

import (
	"fmt"
	"sync"
)

var count int
var mu sync.Mutex
var wg sync.WaitGroup

func worker() {
	defer wg.Done()

	mu.Lock()
	count++
	mu.Unlock()
}

func mutex() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker()
	}

	wg.Wait()
	fmt.Println("Final count:", count)
}

func channel() {
	ch := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup

	// one goroutine owns count
	go func() {
		count := 0

		for v := range ch {
			count += v
		}

		fmt.Println("Final count:", count)
		close(done)
	}()

	// 100 workers
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			ch <- 1
		}()
	}

	wg.Wait()
	close(ch)
	<-done
}