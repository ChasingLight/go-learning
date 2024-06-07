package main

import "fmt"

func main() {
	/*
		Part1：Go语言中，字符串本质上是一个字节切片。
	*/
	// 1.根据字节切片，构建字符串
	var slice = []byte{65, 66, 67, 68, 69}
	str := string(slice)
	fmt.Printf("由 字节切片 构建的 字符串为：%s\n", str) //ABCDE

	// 2.根据字符串，获取对应的字节切片。
	var name = "jadenoliver"
	slice2 := []byte(name)
	fmt.Println("由 字节切片 构建的 字符串为：", slice2) //[106 97 100 101 110 111 108 105 118 101 114]

	// 3.获取字符串，指定下标位置的字节。
	// 字符串不能修改
	fmt.Printf("字符串 name 首个下标对应字节为：%d，对应字节为：%c\n", name[0], name[0]) //106 'j'
	// name[0] = 'J'
	fmt.Println(name) //cannot assign to name[0]

	/*
		Part2：strings包-常见函数使用
	*/
	fmt.Println("-------------------")
}
