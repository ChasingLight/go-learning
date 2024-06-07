package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		strconv包下，处理字符串和基本类型相互转换。
	*/
	// 1. string 和 int 相互转换
	i, err := strconv.Atoi("-42")
	if err != nil {
		fmt.Printf("字符串 转换为 int 有异常，异常信息：%s\n", err)
	}
	fmt.Println(i)
	a := strconv.Itoa(68)
	fmt.Println(a)

	// 2.string 转 float---parse
	f, err := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f)

	// 3.float 转 string---format
	s := strconv.FormatFloat(3.1415926, 'G', -1, 64)
	fmt.Println(s)

}
