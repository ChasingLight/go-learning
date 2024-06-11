package main

import "fmt"

func main() {
	/*
		struct 是 值类型（值传递）。

		struct 结构体指针：
		Go语言内建函数 new(T) 返回一个指向T类型变量的指针。
	*/
	fmt.Println("---------结构体（值传递）-------------")
	person1 := Person{"jaden", 30, "男"}
	person2 := person1
	fmt.Printf("person1的内存地址：%p, person2的内存地址：%p\n", &person1, &person2)
	person2.name = "JadenOliver"
	fmt.Println("修改person2的字段，不会影响person1，证明：strut是值传递。", person1, person2)

	fmt.Println("---------结构体指针（引用传递 使用new(T)内建函数）-------------")
	person3 := new(Person)
	fmt.Printf("person3的数据类型是：%T\n", person3)
	person3.name = "云中君"
	person3.age = 1000
	person3.sex = "男"
	person4 := person3
	fmt.Printf("person3的内存地址：%p, person4的内存地址：%p\n", person3, person4)
	person4.name = "大司命"
	fmt.Println("修改person4的字段，会影响person3，证明：new(Person) 得到的person3 是引用传递。", person3, person4)
}

type Person struct {
	name string
	age  int
	sex  string
}
