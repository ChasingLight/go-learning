package main

import "fmt"

func main() {
	// 在Go语言中，++、--针对整数操作。
	// 和java语言不同，Go语言中不支持在表达式中进行运算。Eg：--c、++c这种操作
	var c = 6
	c++
	fmt.Println(c)

	c--
	fmt.Println(c)
}
