package main

import "fmt"

//money := 1000

func main() {
	//------------单变量声明-----------------
	// 第一种：定义变量，然后进行赋值---var 变量名 数据类型
	var num1 int
	num1 = 29
	fmt.Printf("num1的数值是：%d\n", num1)
	// 书写在一行---var 变量名 数据类型 = 值
	var num2 int = 30
	fmt.Printf("num2的数值是：%d\n", num2)

	// 第二种：类型推断---var 变量名 = 值
	var name = "JadenOliver"
	fmt.Printf("类型是：%T，数值是：%s\n", name, name)

	// 第三种：简短声明---变量名 := 赋值
	price := 12.09
	fmt.Println(price)

	//------------多变量声明-----------------
	var a, b, c int
	a = 10
	b = 20
	c = 30
	fmt.Println(a, b, c)

	var m, n int = 100, 200
	fmt.Println(m, n)

	// 多变量-类型推断
	var studentName, age, sex = "jaden", 30, "男"
	fmt.Printf("姓名：%s，年龄：%d，性别：%s\n", studentName, age, sex)

	// 集合类型
	var dormitoryNumber = "CO3-N525"
	var (
		name1 string = "邢华阳"
		name2 string = "陈龙基"
		name3 string = "靳浩东"
	)
	fmt.Printf("河工大 %s 宿舍成员：%s，%s，%s", dormitoryNumber, name1, name2, name3)

	//------------变量使用注意事项-----------------
	//var money float32 = 4000.13
}
