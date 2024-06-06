package main

import "fmt"

func main() {
	// 常数：固定的数值
	fmt.Println(100)
	fmt.Println("hello")

	//------------单常量声明-----------------
	const PATH string = "https://www.baidu.com"
	const PI = 3.14
	//fmt.Println(PATH)
	fmt.Println(PI)

	// 若修改常量的值，会报错（cannot assign to PATH (neither addressable nor a map index expression)）
	//PATH = "https://www.sina.com"

	//------------多常量声明-----------------
	const C1, C2, C3 = 100, "小花", 16.88
	const (
		MALE    = 0
		FEMALE  = 1
		UNKNOWN = 2
	)

	// 一组常量中，如果某个常量没有初始值，默认和上行保持一致
	const (
		A int = 100
		B
		C string = "good"
		D
		E
	)
	fmt.Printf("%T,%d\n", A, A)
	fmt.Printf("%T,%d\n", B, B)
	fmt.Printf("%T,%s\n", C, C)
	fmt.Printf("%T,%s\n", D, D)
	fmt.Printf("%T,%s\n", E, E)

	// 枚举类型：使用常量组作为枚举类型
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)
	fmt.Println(SPRING, SUMMER, AUTUMN, WINTER)
}
