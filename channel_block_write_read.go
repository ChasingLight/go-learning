package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		channel 通道的 发送、读取是阻塞的
		发送：chan <- data
		读取：data := <- chan

		输出结果展示：
		---子 goroutine开始执行---
		---main goroutine开始执行---
		---从 ch1 通道中 读取到的数据： JadenOliver
		---main goroutine over---
	*/

	ch1 := make(chan string)

	// 匿名函数-协程
	go func() {
		fmt.Println("---子 goroutine开始执行---")
		ch1 <- "JadenOliver"
		fmt.Println("---子 goroutine over---")
	}()

	// 主 goroutine
	time.Sleep(3 * time.Second)
	fmt.Println("---main goroutine开始执行---")
	data := <-ch1
	fmt.Println("---从 ch1 通道中 读取到的数据：", data)
	fmt.Println("---main goroutine over---")
}
