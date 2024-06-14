package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		拷贝文件
		三种方式，实现拷贝文件
		1. 使用 io 包下 Reader.read 和 Writer.write 函数；
		2. 使用 io.copy 函数；
		3. 使用 os 包下 ReadFile 和 WriteFile 函数；
	*/
	srcFilePath := "test.txt"
	destFilePath := "test_copy.txt"
	// 方式1
	//total, err := CopyFile1(destFilePath, srcFilePath)
	// 方式2
	//total, err := CopyFile2(destFilePath, srcFilePath)
	// 方式3
	total, err := CopyFile3(destFilePath, srcFilePath)
	fmt.Println(total, err)
}

// CopyFile1 函数：使用io操作实现文件拷贝，返回值是拷贝 总字节数,错误
func CopyFile1(destFilePath, srcFilePath string) (int, error) {
	// 1.打开源文件 和 目标文件
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return 0, err
	}
	destFile, err := os.OpenFile(destFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return 0, err
	}

	// 3.关闭文件
	defer srcFile.Close()
	defer destFile.Close()

	// 2.边读边写
	bs := make([]byte, 1024, 1024)
	n := -1 //单次读取的数据量
	total := 0
	for {
		n, err = srcFile.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完成")
			break
		} else if err != nil {
			fmt.Println("读取文件报错了")
			return total, err
		}
		total += n
		_, err = destFile.Write(bs[:n])
		if err != nil {
			fmt.Println("写入文件报错了")
			return total, err
		}
	} //end for
	return total, nil
}

// CopyFile2 函数：使用 io.copy 函数实现文件拷贝
func CopyFile2(destFilePath, srcFilePath string) (int64, error) {
	// 1.打开源文件 和 目标文件
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return 0, err
	}
	destFile, err := os.OpenFile(destFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return 0, err
	}

	// 3.关闭文件
	defer srcFile.Close()
	defer destFile.Close()

	// 2.拷贝文件
	return io.Copy(destFile, srcFile)
}

// CopyFile3 函数：使用 ioutil 包下 ReadFile 和 WriteFile 函数实现文件拷贝
func CopyFile3(destFilePath, srcFilePath string) (int, error) {
	// 一次性读取，【如果文件是大文件，有内存溢出的风险】
	input, err := os.ReadFile(srcFilePath)
	if err != nil {
		return 0, err
	}

	// 一次性写入
	err = os.WriteFile(destFilePath, input, os.ModePerm)
	if err != nil {
		return 0, err
	}
	return len(input), nil
}
