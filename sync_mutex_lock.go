package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// 全局变量，表示100张票
var ticket = 10

func main() {
	/*
		4个售票窗口，共同售卖100张票。
		使用 Go 语言的 waitGroup、Mutex 保证票不会被超卖。
	*/
	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go sellTicket(i, &wg, &lock)
	}

	wg.Wait()
	fmt.Println("---main over...")
}

func sellTicket(id int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	// 死循环-进行售票
	for {
		lock.Lock()
		if ticket > 0 {
			// 模拟售卖代码执行时间
			time.Sleep(time.Duration(rand.IntN(1000)) * time.Millisecond)
			fmt.Printf("售票口%d，售出票：%d\n", id, ticket)
			ticket--
			lock.Unlock()
		} else {
			lock.Unlock()
			fmt.Printf("售票口%d，售罄了\n", id)
			break
		}
	}
}
