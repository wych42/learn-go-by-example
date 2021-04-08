package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"sync/atomic"
	"time"
)

var mlock sync.RWMutex
var wg sync.WaitGroup

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go gets()
	}

	wg.Wait()
	if a > 0 {
		fmt.Println("here", a)
	}
}

func gets() {
	for i := 0; i < 100000; i++ {
		get(i)
	}
	wg.Done()
}

var a int64

func get(i int) {
	beginTime := time.Now()
	mlock.RLock()
	tmp1 := time.Since(beginTime).Nanoseconds() / 1000000
	if tmp1 > 100 { // 超过100ms
		atomic.AddInt64(&a, 1)
	}
	mlock.RUnlock()
}
