package main

import "fmt"

func main() {
	// 1. 数组 和 切片 数据类型不同。	[3]string、[]string
	var array = [3]string{"小红", "小芳", "小强"}
	var slice = []string{"jaden", "jack"}
	fmt.Printf("数组的数据类型为：%T，切片的数据类型为：%T\n", array, slice)

	// 2. slice 是 引用数据类型，存储的是 数据的内存值。
	slice2 := slice
	slice2[0] = "JadenOliver"
	fmt.Println(slice)
	fmt.Println(slice2)

	// 3. 切片 append 内置函数使用---多个元素
	// slice[2] = "bob" //panic: runtime error: index out of range [2] with length 2
	slice = append(slice, "bob", "linda")
	fmt.Println(slice)

	// 4. 切片 append --- 将一个新的切片-打开(...) 拼接到 当前切片的尾部
	slice3 := []string{"lucas", "selina"}
	slice = append(slice, slice3...)
	fmt.Println(slice)

}
