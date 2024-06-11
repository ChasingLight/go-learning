package main

import "fmt"

func main() {
	/*
		结构体嵌套：
			如果一个结构体，至少有一个字段的数据类型也是结构体，就称为"结构体嵌套"。
	*/
	addr := Address{"告成镇", "452477"}
	student := Student2{name: "靳浩东", age: 18, addr: &addr}
	fmt.Printf("学生的姓名：%s，年纪：%d，地址结构体-城市名：%s，地址结构体-邮政编码：%s\n",
		student.name, student.age, student.addr.city, student.addr.zipCode)
}

type Address struct {
	city    string
	zipCode string
}

type Student2 struct {
	name string
	age  int
	// 指向一个 地址结构体变量 的指针
	addr *Address
}
