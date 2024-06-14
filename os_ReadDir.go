package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		使用 os.ReadDir() 递归遍历文件夹
	*/
	dirName := "C:\\Users\\Administrator.DESKTOP-BHMLSU6\\Desktop\\云设备服务平台-日志排查"
	listFiles(dirName, 0)
}

func listFiles(dirName string, level int) {
	// level 记录递归的层级，生成有层次感的空格
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|  " + s
	}
	dirEntrys, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Println("遍历文件夹出错，", err)
		return
	}
	for _, dirEntry := range dirEntrys {
		fileName := dirName + "/" + dirEntry.Name()
		fmt.Printf("%s%s\n", s, fileName)
		if dirEntry.IsDir() {
			listFiles(fileName, level+1)
		}
	}
}
