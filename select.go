package main

import "fmt"

func main() {
	/*
		select 语句 可以监测 通道上的数据流动
	*/

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 10
	}()

	go func() {
		ch2 <- 20
	}()

	select {
	case data := <-ch1:
		fmt.Println("case1 可以执行。", data)
	case data := <-ch2:
		fmt.Println("case2 可以执行。", data)
	default:
		fmt.Println("select 中 所有 case 语句阻塞，执行 default 语句")
	}
}
