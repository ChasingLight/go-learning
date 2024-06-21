package main

import "fmt"

func main() {
	// 1.仅声明
	var balance1 [4]float64 //等价于 var balance1 = [4]float64{}
	fmt.Println(balance1)

	// 2.声明 + 全量初始化（显式指定数组长度）
	var balance2 = [5]float64{1000.0, 2.0, 3.4, 7.0, 5.0}
	fmt.Println(balance2)

	// 3.声明 + 全量初始化（不指定数组长度）---本质是slice数据类型
	// 依据元素个数，设置数组大小
	fmt.Println("------")
	a := []int{1, 2}
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// 4.声明 + 全量初始化（不指定数组长度）
	// 依据元素个数，设置数组大小
	b := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// 5.声明 + 部分初始化。
	// 输出数组 长度 和 容量
	var nameArray = [4]string{"靳浩东", "陈龙基", "邢华阳"}
	fmt.Println(nameArray)
	fmt.Println(len(nameArray))
	fmt.Println(cap(nameArray))

	// 6.声明 + 指定下标初始化
	c := []string{0: "Java", 5: "php"} //数组c 长度为6
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))
}
