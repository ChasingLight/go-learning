package main

import "fmt"

func main() {
	/*
		Go语言支持函数式编程，可以将一个函数作为另一个函数的参数。
		假如将 func1 作为 func2 的参数，那么
		func1 被称为 "回调函数"
			作为另一个函数的参数。
		func2 被称为 "高阶函数"
			接收一个函数作为参数。
	*/

	// 1.process 是 高阶函数，sum 是 回调函数
	res := process(6, 8, sum)
	fmt.Println(res)

	// 2.匿名函数 作为 回调函数
	res2 := process(12, 6, func(a, b int) int {
		return a - b
	})
	fmt.Println(res2)
}

// 高阶函数，接收一个函数作为参数
func process(a, b int, callback func(int, int) int) int {
	return callback(a, b)
}

// 回调函数，作为另一个函数的参数
func sum(a, b int) int {
	return a + b
}
