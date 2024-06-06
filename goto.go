package main

import (
	"fmt"
	"strconv"
)

func main() {

	/**
	使用 goto 语句 进行 统一错误处理
	注意：
		请谨慎使用 goto 会导致代码执行逻辑混乱，代码可读性变差！
	*/
	// 1. 字符串 转换为 数字
	var value = "abc"
	atoi, err := strconv.Atoi(value)
	if err != nil {
		goto onExit
	}
	fmt.Println(atoi)

onExit:
	{
		fmt.Println(err)
		return
	}

}
