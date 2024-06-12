package main

import "fmt"

func main() {
	/*
		panic：词义恐慌
		recover：词义恢复
		Go语言中利用 panic()，recover() 实现程序中的 “异常处理”。
		1. panic()：让当前程序进入恐慌，中断程序的执行；
		2. recover()：让程序恢复，必须在 defer 延迟函数 中执行。

		Go语言中 panic-recover-defer 类似于 Java中的 try-catch-finally
	*/
	funA()
	defer myprint("defer main：3...")
	funB()
	defer myprint("defer main：4...")
	fmt.Println("main over...")

}

func myprint(s string) {
	fmt.Println(s)
}

func funA() {
	fmt.Println("我是函数funA()...")
}

func funB() {
	// defer + 匿名函数 + recover()
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println("程序恢复了")
		}
	}()
	fmt.Println("我是函数funB()...")
	defer myprint("defer funB()：1...")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		if i == 5 {
			// 让程序中断
			panic("funB函数，恐慌了")
		}
	}
	defer myprint("defer funB()：2...")
}
