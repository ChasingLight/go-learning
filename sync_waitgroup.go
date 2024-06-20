package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		sync.WaitGroup 是一个用于等待一组 goroutine 完成的同步原语。
		它允许你发起多个 goroutine，并且在所有 goroutine 完成之前阻塞主程序的执行。
	*/

	// 创建 WaitGroup 实例
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		// 添加计数
		wg.Add(1)
		go worker(i, &wg)
	}

	// 阻塞直到所有 wg.Done() 调用完成
	wg.Wait()
	fmt.Println("--main--All workers done...")
}

func worker(id int, wg *sync.WaitGroup) {
	// 执行完成后，调用 Done 减少计数
	// wg.Done() 等价于 wg.Add(-1)
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Printf("--goroutine--worker %d working...\n", id)
}
