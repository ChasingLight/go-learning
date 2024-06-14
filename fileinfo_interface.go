package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	/*
		FileInfo 文件信息，对应源代码如下：
		type FileInfo interface {
			Name() string       // base name of the file
			Size() int64        // length in bytes for regular files; system-dependent for others
			Mode() FileMode     // file mode bits
			ModTime() time.Time // modification time
			IsDir() bool        // abbreviation for Mode().IsDir()
			Sys() any           // underlying data source (can return nil)
		}
	*/

	fileInfo, err := os.Stat("./test.txt")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	fmt.Printf("%T\n", fileInfo)

	fmt.Println("------------通过FileInfo接口，获取文件相关信息------------------")
	fmt.Printf("文件名：%s\n", fileInfo.Name())
	fmt.Printf("文件大小（单位：字节）：%d\n", fileInfo.Size())
	fmt.Printf("权限：%v\n", fileInfo.Mode())
	fmt.Printf("修改时间：%s\n", fileInfo.ModTime().Format(time.DateTime))
	fmt.Printf("是否是目录：%v\n", fileInfo.IsDir())

}
