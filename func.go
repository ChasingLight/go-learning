package main

import "fmt"

func main() {
	// 函数支持-多返回值
	perimeter, area := rectangle(6, 8)
	fmt.Printf("长方形的 周长：%.2f，面积：%.2f\n", perimeter, area)

	// 可使用 空白标识符（"_"），来舍弃函数多余的返回值
	// 只要周长，不要面积
	perimeter2, _ := rectangle(3, 5)
	fmt.Printf("长方形的 周长：%.2f\n", perimeter2)

	// 函数 是一种复合数据类型，本质上是一个变量。
	// 函数调用：func()
	// 函数作为变量：func
	/*
		 函数是一种复合数据类型，可以看做是一种特殊的变量。
			函数名(): 对函数进行调用，执行函数中的代码，然后将return结果返回调用处。
			函数名: 是指向 函数体 的内存地址 的特殊变量（本质是个指针）。
	*/
	fmt.Println("------------函数的数据类型和本质-----------------")
	fmt.Printf("函数的数据类型为：%T，函数名 指向的 函数体内存地址为：%p\n", add, add)
	res1 := add(1, 2)
	fmt.Printf("对函数进行调用，函数返回结果: %d\n", res1)

	res2 := add
	fmt.Println(add, res2)
	fmt.Println(res2(3, 4))

}

/*
*
求长方形的 周长、面积
*/
func rectangle(len, wid float64) (float64, float64) {
	perimeter := (len + wid) * 2
	area := len * wid
	return perimeter, area
}

func add(a, b int) int {
	return a + b
}
