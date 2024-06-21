package main

import "fmt"

func main() {
	/*
		缓冲通道
		ch1 := make(chan int, 5)
		1> 对于发送来讲，只有缓冲区满时才被阻塞；
		2> 对于接收来讲，只有缓冲区为空时才被阻塞。
	*/

	// 1. 通道的 len 长度 和 cap 容量
	ch1 := make(chan int)
	ch2 := make(chan int, 5)
	fmt.Printf("非缓冲通道，初始 len ：%d, cap ：%d\n", len(ch1), cap(ch1))
	fmt.Printf("缓冲通道，初始 len ：%d, cap ：%d\n", len(ch2), cap(ch2))

	// 2. 缓冲通道-死锁
	//<- ch2	//对于接收来讲，只有缓冲区为空时才被阻塞

	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	ch2 <- 4
	ch2 <- 5
	//ch2 <- 6	//对于发送来讲，只有缓冲区满时才被阻塞

}
