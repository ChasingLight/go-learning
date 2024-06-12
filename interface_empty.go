package main

import "fmt"

func main() {
	/*
		空接口（interface{}）
			不包含任何方法的接口。因此所有的类型都实现了空接口，所以空接口可以代表任意类型。
			【类似Java语言中的Object---万物皆对象，万物皆空接口】
		用途：
			1. 函数参数列表，使用空接口，代表可以接受任意类型的参数；
			2. 容器类元素数据类型，使用空接口，代表容器内可以存储任意类型的数据。
	*/

	// 1.函数参数列表，使用空接口，可接受任意类型的参数
	fmt.Println("--------空接口在函数中的使用-------------")
	printInfo1(123)
	printInfo1("jaden")
	printInfo2(Dog{"小虎", "土黄色", 4})
	printInfo2(3.1415926)
	// fmt包下print相关函数，参数列表也都是 空接口
	// func Println(a ...interface{}) (n int, err error)
	fmt.Println(2, 3, 4)

	// 2.容器类元素数据类型，使用空接口，可存储任意类型数据
	fmt.Println("--------空接口在容器中的使用-------------")
	// 切片
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, 666, "云中君", 8.88, Dog{"lucky", "金黄色", 8})
	fmt.Println(len(slice1))
	fmt.Println(slice1)
}

// A 空接口（有别名）
type A interface{}

type Dog struct {
	name  string
	color string
	age   int
}

func printInfo1(a A) {
	fmt.Println(a)
}

// 匿名空接口
func printInfo2(a interface{}) {
	fmt.Println(a)
}
