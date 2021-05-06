package main

import (
	"fmt"
	"sync"
)

var x int64 // 临界区

var wg sync.WaitGroup
var lock sync.Mutex // 互斥锁

func add() {
	for i := 0; i < 5000; i++ {
		// lock.Lock()   // 上锁
		x = x + 1 // 非原子操作
		// lock.Unlock() // 解锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	// 多个goroutine同时操作一个资源（临界区），发生竞态问题
	go add()
	go add()

	wg.Wait()
	// 理想情况下x=10000，但发生竞态问题，极有可能是x<10000
	// 可以通过互斥锁的上锁操作和解锁操作解决
	// 非原子操作之前上锁，之后解锁
	fmt.Println(x)
}
