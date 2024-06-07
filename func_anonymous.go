package main

import "fmt"

func main() {
	/*
		匿名函数（没有名字的函数），通常只使用一次。
	*/
	// 1. 定义匿名函数 + 调用
	res1 := func(a, b int) int {
		return a + b
	}(3, 3)
	fmt.Println(res1)

	// 2. 仅定义匿名函数 + 引用传递
	res2 := func(a, b int) int {
		return a + b
	}
	fmt.Println(res2(4, 4))
	fmt.Println(res2(5, 5))
	fmt.Println(res2(8, 8))

}
