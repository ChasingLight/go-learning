package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		错误（error）、异常（panic）区分
		1. 错误是意料之中的问题。比如打开一个文件时，因为文件不存在而失败。
		2. 异常是意料之外的问题。比如引用了空指针。
		说明：错误是业务过程中的一部分，而异常不是。

		Go语言中 错误 也是一种类型。使用内置的error表示。
		就像其他的类型 int、float64，错误可以存储在变量中、从函数中返回等。
	*/
	f, err := os.Open("test.txt")
	if err != nil {
		// log.Fatal(err)	//日志打印错误信息
		fmt.Println(err) //open test.txt: The system cannot find the file specified.
		return
	}
	fmt.Println(f.Name(), "打开文件成功...")

}
