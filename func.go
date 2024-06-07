package main

import "fmt"

func main() {
	// Go语言 函数支持-多返回值
	perimeter, area := rectangle(6, 8)
	fmt.Printf("长方形的 周长：%.2f，面积：%.2f\n", perimeter, area)

	// 可使用 空白标识符（"_"），来舍弃函数多余的返回值
	// 只要周长，不要面积
	perimeter2, _ := rectangle(3, 5)
	fmt.Printf("长方形的 周长：%.2f\n", perimeter2)
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
