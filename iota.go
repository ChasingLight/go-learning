package main

import "fmt"

func main() {
	/*
		iota：特殊的常量，可被编译器自动修改的常量。
			每当定义一个const，iota值初始化为0；
			每当第一个常量值，iota自动累加1；
			直到下一个const出现，iota重置为0。
	*/
	const (
		A = iota
		B = iota
		C = iota
	)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	// 一组常量中，如果某个常量没有初始值，默认和上行保持一致
	const (
		D = iota
		E
	)
	fmt.Println(D)
	fmt.Println(E)
}
