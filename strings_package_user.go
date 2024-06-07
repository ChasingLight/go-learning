package main

import (
	"fmt"
	"strings"
)

func main() {
	/**
	strings包下 操作字符串的 常用函数
	*/
	str := "hello world"
	// 1.包含
	fmt.Println(strings.Contains(str, "hell"))
	fmt.Println(strings.Contains(str, "abc"))

	// 2.前缀、后缀
	fileName := "20240607课堂笔记.txt"
	if strings.HasPrefix(fileName, "202406") {
		fmt.Println("是 2024年6月份的文件")
	}
	if strings.HasSuffix(fileName, ".txt") {
		fmt.Println("是 文本文件")
	}

	// 3.索引
	index1 := strings.Index(str, "world")
	fmt.Println(index1) //6
	index2 := strings.Index(str, "dddd")
	fmt.Println(index2) //-1

	// 4.分割
	split := strings.Split(str, " ")
	fmt.Printf("字符串分割后，返回类型为：%T\n", split)
	for index, value := range split {
		fmt.Printf("下标：%d, 值：%s\n", index, value)
	}

	// 5.替换
	newStr := strings.ReplaceAll(str, "l", "*")
	fmt.Println(newStr)
	// 延展：使用替换 去除 空格
	noSpaceStr := strings.ReplaceAll(str, " ", "")
	fmt.Println(noSpaceStr)

	// 6.大小写转换
	upperStr := strings.ToUpper(str)
	fmt.Printf("转换为大写：%s\n", upperStr)
	lowerStr := strings.ToLower(str)
	fmt.Printf("转换为小写：%s\n", lowerStr)

	/*7.截取子串
		Go语言中没有substring函数
	 	subStr := str[startIndex:endIndex]
			包含 startIndex 所在的字符，不包含 endIndex 的字符。
			即：左闭右开。
	*/
	subStr := str[0:5]
	fmt.Println(subStr)
}
