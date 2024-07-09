package main

import "fmt"

// 定义一个返回闭包函数的外部函数
func outerFunction(x int) func(int) int {
	// x 是外部函数中的变量
	return func(y int) int {
		return x + y
	}
}

func main() {
	/*
		闭包函数定义：是指在一个函数内部定义的函数。这个内部函数(即 闭包函数)可引用外部函数中的变量。
				即使外部函数已经执行完毕并返回，内部函数依然可以访问这些变量。
	*/

	// 创建一个闭包函数实例
	closureInstance := outerFunction(10)
	// 调用闭包
	result := closureInstance(5)
	fmt.Println(result)
}
