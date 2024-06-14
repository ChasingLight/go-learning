package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/*
		write 写入文件
		1> 使用 io 包下的 Writer 接口
		2> 其中 os 包下的 file.go 实现了 Writer 接口。
	*/

	// 1.打开文件
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	HandlerErr(err)

	// 3.关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件错误：", err)
			return
		}
		fmt.Println("关闭文件成功")
	}(file)

	// 2.写入文件
	bs := []byte{65, 66, 67, 68, 69}
	n, err := file.Write(bs)
	HandlerErr(err)
	fmt.Println(n)

	file.WriteString("\n")
	n, err = file.WriteString("靳浩东\n")
	HandlerErr(err)
	fmt.Println(n)

}

func HandlerErr(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
