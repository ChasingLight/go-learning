package main

import "fmt"

func main() {
	/*
		数组的数据类型：
			[size]type
		值类型：存储的数值本身。将数值传递给其他变量时，传递的是数据的副本。
				int,float,string,bool,array
		引用类型：存储的是 数据的内存值
				slice,map 等
	*/
	arr1 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("数组的数据类型：%T\n", arr1) //[6]int

	arr2 := arr1
	arr2[0] = 100
	fmt.Println(arr1) //[1 2 3 4 5 6]
	fmt.Println(arr2) //[100 2 3 4 5 6]
}
