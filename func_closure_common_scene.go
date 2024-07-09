package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		闭包函数定义：是指在一个函数内部定义的函数。这个内部函数（即闭包函数）可引用外部函数中的变量。
					即使外部函数已经执行完毕并返回，内部函数依然可以访问这些变量。
		闭包函数-常见使用场景：数据隐藏-迭代器、工厂函数、回调、缓存。
	*/
	//-----------数据隐藏-迭代器-------------
	counter1 := counter()
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())
	counter2 := counter()
	fmt.Println(counter2())
	fmt.Println(counter2())

	//-----------工厂函数-------------
	sumCal := calculator("sum")
	fmt.Println(sumCal(1, 2))
	fmt.Println(sumCal(6, 8))
	mulCal := calculator("mul")
	fmt.Println(mulCal(22, 8))
	fmt.Println(mulCal(9, 9))

	//-----------回调-------------
	task := func() {
		fmt.Println("Hello World")
	}
	scheduleTask(task, 2*time.Second)

	//-----------简单缓存机制-------------
	memorizedFunc := memorize(slowFunction)
	fmt.Println(memorizedFunc(5)) //计算并缓存结果
	fmt.Println(memorizedFunc(5)) //从缓存map中直接获取结果返回
}

func counter() func() int {
	// count 变量被封装在 counter() 函数作用域内，只能通过闭包函数访问和修改
	count := 0
	return func() int {
		// 迭代增加
		count++
		return count
	}
}

func calculator(calType string) func(x, y int) int {
	// 根据计算类型-标识符，返回对应的闭包函数实例
	return func(x, y int) int {
		result := 0
		switch calType {
		case "sum":
			result = x + y
		case "sub":
			result = x - y
		case "mul":
			result = x * y
		case "div":
			result = x / y
		}
		return result
	}
}

func scheduleTask(task func(), delay time.Duration) {
	// task 闭包在 scheduleTask 延迟执行
	time.Sleep(delay)
	task()
}

func slowFunction(x int) int {
	time.Sleep(3 * time.Second)
	return x * x
}
func memorize(f func(int) int) func(int) int {
	cacheMap := make(map[int]int)
	return func(x int) int {
		if val, found := cacheMap[x]; found {
			return val
		}
		result := f(x)
		cacheMap[x] = result
		return result
	}
}
