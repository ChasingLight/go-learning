package main

import "fmt"

func main() {
	/*
		方法：method
			一个方法 就是一个包含了接受者 的函数，且接受者可以在方法的内部访问。
			接受者可以是 命名类型、结构体类型、指针。
			所有给定类型的方法 属于 该类型的方法集。

		语法对比：
		func (t Type) methodName(parameter list)(return list){
			// 方法体
		}

		func funcName()(parameter list)(return list){
			// 函数体
		}

		总结：方法（method）和函数类似，区别在于需要接收者（也就是调用者）。
	*/
	fmt.Println("----------方法的调用------------------")
	w1 := Worker{"jack", 20}
	w1.work()
	w2 := &Worker{"bob", 18}
	w2.work()
	w1.rest()
	w2.rest()

	fmt.Println("----------方法的重写------------------")
	w1.printInfo()
	cat := Cat{"xixi", "蓝色"}
	cat.printInfo()

}

type Worker struct {
	name string
	age  int
}

type Cat struct {
	name  string
	color string
}

func (w Worker) work() {
	// 在 方法内部 可以访问 接受者
	fmt.Println(w.name, "在努力工作...")
}

func (w *Worker) rest() {
	// 在 方法内部 可以访问 接受者
	fmt.Println(w.name, "正在休息...")
}

// 类似 java语言的 方法重载 Overload（在一个类里面，方法名字相同，而参数列表不同。返回类型可以相同也可以不同）
func (w Worker) printInfo() {
	fmt.Printf("工人的姓名：%s，工人的年纪：%d\n", w.name, w.age)
}

func (c Cat) printInfo() {
	fmt.Printf("猫咪的名字：%s，猫咪的毛色：%s\n", c.name, c.color)
}
