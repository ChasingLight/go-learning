package main

import "fmt"

func main() {

	/*
		switch 中的 break 和 fallthrough 语句
		break：
			强制结束 case 语句，从而结束 switch 分支；
		fallthrough：向下穿透switch
			当 switch 中某个 case 语句匹配成功后，执行 case 后的语句
			如果遇到 fallthrough，那么当前 case 后紧邻的case，无需匹配，穿透执行。
	*/
	var quarter = 2
	switch quarter {
	case 1:
		fmt.Println("第一季度")
		break
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
		fmt.Println("第二季度")
		fallthrough
	case 3:
		fmt.Println("第三季度")
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
		fmt.Println("第四季度")
	default:
		fmt.Println("季度值-输入有误！")
	}
}
