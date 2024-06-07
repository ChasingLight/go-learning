package main

import "fmt"

func main() {
	/*
		闭包函数：
		当它们在函数体内引用了外部变量时，这些变量将被“绑定”到闭包中，
		即使这些变量超出了它们的原始作用域，闭包依然可以访问和修改它们。
	*/
	// 创建两个独立的计数器
	counter1 := createCounter()
	counter2 := createCounter()

	fmt.Println(counter1()) // 输出: 1
	fmt.Println(counter1()) // 输出: 2
	fmt.Println(counter1()) // 输出: 3
	fmt.Println(counter2()) // 输出: 1
	fmt.Println(counter2()) // 输出: 2

}

// 返回一个闭包函数
func createCounter() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}
