package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		defer 延迟执行函数
		说明：有点类似 Java语言中的 finally
	*/
	// 1. 多个defer语句，按照后进先出（LIFO）的顺序执行
	deferExample1()

	// 2. defer 虽然会延迟函数执行，但是参数会立即传递
	deferExample2()

	// 3. defer 典型用法：关闭文件、释放互斥锁、记录方法执行时间
	deferExample3()

}

func deferExample1() {
	fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	fmt.Println("5") // 1、5、4、3、2
}

func deferExample2() {
	fmt.Println("-----------------")
	x := 10
	defer fmt.Println(x)
	x = 20
	fmt.Println(x) //20、10
}

func deferExample3() {
	fmt.Println("-----------------")
	start := time.Now()
	fmt.Printf("方法开始执行的时间: %s\n", start)
	defer fmt.Printf("Execution time: %s\n", time.Since(start))
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
