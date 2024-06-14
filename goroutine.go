package main

import "fmt"

func main() {
	/*
		goroutine 初步使用
	*/
	// 1.创建并启动 子goroutine，执行printNum
	go printNum()

	// 2.main中打印字母
	for i := 0; i < 100; i++ {
		fmt.Printf("|---主goroutine中打印的字母：A %d\n", i)
	}
}

func printNum() {
	for i := 0; i < 100; i++ {
		fmt.Printf("子goroutine中打印的数字：%d\n", i)
	}
}
