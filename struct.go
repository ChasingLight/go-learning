package main

import "fmt"

// Student 结构体定义
type Student struct {
	name string
	age  int
	sex  string
}

func main() {
	/*
		结构体：是由一系列具有相同类型 或 不同类型的数据 构成的 数据集合。
				结构体成员 被称为 '成员变量' 或 '字段'。
				（这些成员变量的类型 也可以为 结构体，即"结构体嵌套"）
		语法格式：
				type structVariableType struct{
					member type
					member type
					...
					member type
				}
	*/
	// 2. 初始化
	stu1 := Student{"邢华阳", 31, "男"}                 //按顺序提供初始化值
	stu2 := Student{age: 30, sex: "男", name: "陈隆基"} //通过field:value的方式进行初始化，可以任意顺序
	stu3 := Student{name: "靳浩东", age: 30, sex: "男"}

	// 3. 访问成员变量，使用点操作符（.）
	var dormitory []Student
	students := append(dormitory, stu1, stu2, stu3)
	for _, student := range students {
		fmt.Printf("学生姓名：%s，年龄：%d，性别：%s\n", student.name, student.age, student.sex)
	}
}
