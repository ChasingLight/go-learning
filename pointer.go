package main

import "fmt"

func main() {
	/*
		指针（pointer）是存储了 另一个变量的内存地址 的变量。
	*/
	var a = 6
	fmt.Printf("a 的数据类型为：%T，a 的数值为：%d，a 的内存地址是：%p\n", a, a, &a)

	// 1. 声明一个指针
	// 语法格式：var varName *varType
	var p *int
	fmt.Println("指针p 的默认值为：", p) //<nil>
	p = &a

	// 2. 指针的数据类型、数值、内存地址
	fmt.Printf("p 的数据类型为：%T，p 的数值为：%p，p 的内存地址是：%p\n", p, p, &p)

	// 3. 指针 指向的 变量的值
	fmt.Printf("p 指向的变量 a 的数值为：%d\n", *p)

	// 4. 通过指针，改变指向变量的值
	*p = 8
	fmt.Printf("通过指针改变指向变量的值，变量a的最新值为：%d\n", a)

	// 5. 指针的指针
	fmt.Println("--------------指针的指针-------------------")
	var p2 **int
	p2 = &p
	fmt.Printf("p2 的数据类型为：%T，p2 的数值为：%p，p2 的内存地址是：%p\n", p2, p2, &p2)
	fmt.Printf("p2 指向的变量 p 的数值为：%p\n", *p2)
	fmt.Printf("p2 指向的变量 p 的数值，再获取 p 指向的变量 a的数值 为：%d\n", **p2)

}
