package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	/*
		文件操作：
		文件路径、创建目录、创建文件、打开文件(只读、读写)、关闭文件、删除文件或目录
	*/
	// 1.文件路径（绝对路径 和 相对路径）
	filePath1 := "D:\\JadenData\\jadenGoProject\\src\\basic\\test.txt"
	filePath2 := "test.txt"
	fmt.Println(filepath.IsAbs(filePath1)) //true
	fmt.Println(filepath.IsAbs(filePath2)) //false
	fmt.Println(filepath.Abs(filePath1))
	fmt.Println(filepath.Abs(filePath2)) // D:\JadenData\jadenGoProject\src\basic\test.txt
	fmt.Println("获取父目录：", filepath.Dir(filePath1))

	// 2.创建目录（单层目录、多层目录）
	/*err := os.Mkdir("D:\\JadenData\\jadenGoProject\\src\\basic\\bb", os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("单层目录创建成功...")*/

	/*err := os.MkdirAll("D:\\JadenData\\jadenGoProject\\src\\basic\\cc\\dd\\ee", os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("多层目录创建成功...")*/

	// 3.创建文件: os.Create(name) 采用模式0666(任何人可读 可写 不可执行)
	// 创建一个名为 name 的文件，如果文件已存在会截断它。（变为空文件）
	file1, err := os.Create("D:\\JadenData\\jadenGoProject\\src\\basic\\bb\\jaden.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件创建成功：", file1)

	// 4.打开文件
	// 只读模式-打开文件
	/*file3, err := os.Open(filePath2)
	if err != nil{
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file3)*/

	// 读写模式-打开文件
	/*file4, err := os.OpenFile(filePath2, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file4)*/

	// 5.关闭文件
	/*closeErr := file4.Close()
	if closeErr != nil {
		fmt.Println("关闭文件，错误信息：", closeErr)
	}
	fmt.Println("关闭文件成功")*/

	// 6.删除文件或目录
	// os.Remove 删除单个文件 或 空目录；os.RemoveAll 级联删除【谨慎使用！】
	removeErr := os.RemoveAll("cc")
	if removeErr != nil {
		fmt.Println("删除文件或目录，错误信息：", removeErr)
	}
	fmt.Println("删除文件或目录成功")

}
