package main

import "fmt"

func main() {
	/*
		1. for 循环-标准写法
		for 表达式1; 表达式2; 表达式3 {
			循环体
		}
	*/
	for i := 1; i <= 5; i++ {
		fmt.Println("功德+1")
	}

	/*
		2. for循环 同时省略表达式1 和 表达式3。
		相当于 while(条件) 循环
		for 表达式2 {
			循环体
		}
	*/
	var j = 1
	for j <= 5 {
		fmt.Println("---功德+1")
		j++
	}
	fmt.Println("--->", j)

	/*
		3. 同时省略3个表达式
		相当于 while(true) 死循环
		for {
			循环体
		}
	*/
	for {
		fmt.Println("******功德+1")
	}
}
