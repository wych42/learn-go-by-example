package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Printf("working on job %d \n", id)
	time.Sleep(2 * time.Second)
}

// 限制并发数量
var limit = make(chan struct{}, 2)

// 等任务完成
var wait = make(chan int)

func main() {
	for i := 0; i < 5; i++ {
		go func(id int) {
			fmt.Printf("start job %d\n", id)
			// 尝试写 channel，能写进去就可以执行
			limit <- struct{}{}
			worker(id)
			// 释放
			<-limit
			wait <- id
		}(i)
	}
	for i := 0; i < 5; i++ {
		id := <-wait
		fmt.Printf("finish job %d\n", id)
	}
}
