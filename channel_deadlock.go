package main

import "fmt"

func main() {
	/*
		channel 通道 死锁---2种方式
		1> 在 main goroutine，发送通道，但是没有其他 goroutine 读取；
			ch1 <- true
		2> 在 main goroutine，接收通道，但是没有其他 goroutine 发送；
			<- ch1
		[额外说明]
			在 子 goroutine，仅发送 或 仅接收，是不会产生死锁的。
			因为 main goroutine 代码终止，其他 子 goroutine 也会终止运行。不会产生死锁。
	*/

	ch1 := make(chan bool)
	fmt.Println("---main goroutine 开始执行---") //fatal error: all goroutines are asleep - deadlock!
	//ch1 <- true
	<-ch1
	fmt.Printf("---main goroutine over---")

}
