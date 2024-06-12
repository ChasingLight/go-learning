package main

import "fmt"

func main() {
	/*
		OOP(Object Oriented Programming)-面向对象编程，三大特性 之 继承性：
		如果两个类存在继承关系，其中一个是子类，另一个是父类，那么：
		1. 子类可以直接访问父类的 属性 和 方法；
		2. 子类可以新增自己的 属性 和 方法：
		3. 子类可以重写（Override）父类的方法；

		Go语言的结构体嵌套（struct nest）：
		1. 模拟继承性：is a
		【B 继承 A】
		type A struct{
			field
		}
		type B struct{
			A	//匿名结构体字段
		}
		2. 模拟聚合关系：has a
		【D 拥有 C】
		type C struct{
			field
		}
		type D struct{
			c C	//聚合关系
		}
	*/

	// 0.父类对象，访问自身的 属性 和 方法
	p1 := Person11{"小头爸爸", 35}
	fmt.Printf("父类对象，访问父类的字段。姓名：%s，年龄：%d\n", p1.name, p1.age)
	p1.eat() //父类对象，访问父类方法

	// 1.子类可以 访问父类的 属性 和 方法
	s1 := Student11{Person11{"大头儿子", 8}, "希望小学"}
	// s1.Person11.name 可简写为 s1.name。
	// 另外 Person11 中的name、age字段，对Student11来说属于 "提升字段"。
	fmt.Printf("子类对象，访问父类的字段。姓名：%s，年龄：%d，学校：%s\n", s1.name, s1.age, s1.school)
	s1.eat() //子类对象，访问父类方法

	// 2.子类新增自己的方法
	s1.study()
	s1.eat()
}

type Person11 struct {
	name string
	age  int
}

type Student11 struct {
	Person11 //结构体嵌套（匿名结构体字段），模拟继承性
	school   string
}

func (p Person11) eat() {
	fmt.Printf("%s 吃窝窝头...\n", p.name)
}

func (s Student11) eat() {
	fmt.Printf("子类重写父类的方法：%s 吃炸鸡喝啤酒...\n", s.name)
}

func (s Student11) study() {
	fmt.Printf("子类新增的方法，%s学习啦...\n", s.name)
}
