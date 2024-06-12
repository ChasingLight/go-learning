package main

import "fmt"

// CallbackFunc 使用type关键字，来定义函数类型（使代码结构更加清晰）
// type CallbackFunc func(int, int) int	//创建了一个新类型（不同的类型）
type CallbackFunc = func(int, int) int //创建了一个类型别名（同一个类型，不同的名字）

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

	var f CallbackFunc = sum
	fmt.Printf("%T, %T\n", f, sum)

}

// 高阶函数，接收一个函数作为参数
func process(a, b int, callback CallbackFunc) int {
	return callback(a, b)
}

// 回调函数，作为另一个函数的参数
func sum(a, b int) int {
	return a + b
}
