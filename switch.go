package main

import "fmt"

func main() {
	//--------switch标准用法------------
	// 与 Java 语言的 switch 语句相比，Go语言中 case 语句默认有 break，不必显式额外设置。
	var quarter int = 3
	switch quarter {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
	case 3:
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
	default:
		fmt.Println("季度值-输入有误！")
	}

	//--------switch特殊用法 之 省略 switch 后的变量。
	// 相当于直接作用在 switch true {} 上------------
	var age = 18
	switch {
	case age > 0 && age < 18:
		fmt.Println("未成年")
	case age >= 18:
		fmt.Println("已成年")
	default:
		fmt.Println("年纪输入有误")
	}

	//--------switch特殊用法 之 case 后同时跟随多个数值
	var year = 2024
	var month = 2
	var dayOfMonth int
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		dayOfMonth = 31
	case 4, 6, 9, 11:
		dayOfMonth = 30
	case 2:
		if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
			dayOfMonth = 29
		} else {
			dayOfMonth = 28
		}
	default:
		fmt.Println("月份month有误")
	}
	fmt.Printf("%d 年 %d 月 的天数是：%d 天\n", year, month, dayOfMonth)
}
