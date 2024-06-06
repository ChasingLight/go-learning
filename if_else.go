package main

import (
	"fmt"
)

func main() {
	/*
		需求：从控制台输入分数 score，然后执行程序给出对应评级 grade
		优---[90,100]
		良---[60,90)
		差---[0,60)
	*/
	var score int
	fmt.Println("请输入您的分数")
	_, err := fmt.Scanf("%d", &score)
	if err != nil {
		fmt.Println("您输入的分数有误！")
		return
	}

	var grade string
	if score < 60 {
		grade = "差"
	} else if score < 90 {
		grade = "良"
	} else {
		grade = "优"
	}
	fmt.Printf("您分数对应的评级为：%s\n", grade)

	/*
		--------if-else的特殊用法-------------
		if statement; condition {

		}
		说明：在statement中定义的变量，只能在if-else语句块中使用，否则编译器会报错。
	*/
	if age := 16; age < 18{
		fmt.Println("未成年")
	}else{
		fmt.Println("已成年")
	}
	//fmt.Println(age)

}
