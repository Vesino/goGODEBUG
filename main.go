package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1) // increases WaitGroup
		go work() // calls a function as goroutine
	}

	wg.Wait() // waits until WaitGroup is <= 0
}

func work() {
	time.Sleep(time.Second)

	var counter int

	for i := 0; i < 1e10; i++ {
		counter++
	}

	wg.Done()
}

// $ export GODEBUG=schedtrace=1000

// gomaxprocs: Processors configured
// idleprocs: Processors are not in use. Goroutine running.
// threads: Threads in use.
// idlethreads: Threads are not in use.
// runqueue: Goroutines in the global queue.
// [1 0 0 0]: Goroutines in each processor's local run queue.

// https://betterprogramming.pub/deep-dive-into-concurrency-of-go-93002344d37b
// https://morsmachine.dk/go-scheduler
// https://www.ardanlabs.com/blog/2015/02/scheduler-tracing-in-go.html
