package main

import "fmt"

func main() {
	/*
		函数指针：一个指针，是一个指向函数的指针
			go语言中，fuc 默认看作一个指针，没有*

		指针函数：一个函数，该函数的返回值是一个指针。
	*/
	var a func()
	a = printInfo
	a()

	fmt.Println("----------指针函数---------------")
	// 数组指针
	p := pointerFunc()
	fmt.Printf("p 的数据类型为：%T，p 的数值为：%p，p 的内存地址是：%p\n", p, p, &p)
	// 根据指针，修改数组中的值
	(*p)[0] = 100
	p[1] = 200 //简写
	fmt.Println(*p)
}

func printInfo() {
	fmt.Println("------")
}

func pointerFunc() *[4]int {
	arr := [4]int{1, 2, 3, 4}
	fmt.Printf("函数内-数组的内存地址：%p\n", &arr)
	return &arr
}
