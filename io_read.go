package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		read 读取文件
		1> 使用 io 包下的 Reader 接口
		2> 其中 os 包下的 file.go 实现了 Reader接口。
	*/
	// 1.打开文件
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("打开文件失败，err:", err)
	}
	// 3.关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件错误：", err)
			return
		}
		fmt.Println("关闭文件成功")
	}(file)

	// 2.读取文件
	bs := make([]byte, 1024, 1024)
	n := -1
	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("读取到文件的末尾，结束读取操作")
			break
		}
		fmt.Println(string(bs))
	}
}
