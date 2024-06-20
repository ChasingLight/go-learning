package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		channel初步使用
		说明：
		main goroutine执行结束后，程序就终止了，其他子goroutine将不再执行。
		为了保证 所有的子gouroutine执行完毕后，main gouroutine才终止。
		有2种方法：
		1> sync.WaitGroup;
		2> channel;
	*/
	done := make(chan bool)
	fmt.Printf("%T, %v\n", done, done)
	for i := 0; i < 5; i++ {
		go worker(i, done)
	}

	// 从channel中获取响应次数的结果
	for i := 0; i < 5; i++ {
		<-done
	}
	fmt.Println("--main--All workers done...")
}

func worker(id int, done chan bool) {
	time.Sleep(time.Second)
	fmt.Printf("--goroutine--worker %d working...\n", id)
	done <- true
}
