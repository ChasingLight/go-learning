package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func init() {
	goVersion := runtime.Version()
	fmt.Printf("--init--Go 版本:[%s]\n", goVersion)
	fmt.Printf("--init--逻辑 CPU 个数 :[%d]\n", runtime.NumCPU())

	// 如果go1.8之前，设置最大可使用逻辑CPU数，发挥高性能
	isVerLow := true
	split := strings.Split(goVersion, ".")
	// 截取大版本号
	split[0] = (split[0])[2:3]
	for index, value := range split {
		// 大版本
		if index == 0 {
			bigVer, _ := strconv.Atoi(value)
			if bigVer > 1 {
				fmt.Println("go 大版本号大于1")
				isVerLow = false
				break
			}
		}
		// 中版本号
		if index == 1 {
			mVer, _ := strconv.Atoi(value)
			if mVer > 8 {
				fmt.Println("go 中版本号大于8")
				isVerLow = false
				break
			}
		}
	} //end for

	// go1.8之前手动设置逻辑CPU数，提高性能
	if isVerLow {
		gomaxprocs := runtime.GOMAXPROCS(runtime.NumCPU())
		fmt.Printf("--init--因为go版本低于1.8，所以手动设置逻辑CPU数:%d\n", gomaxprocs)
	} else {
		fmt.Printf("--init--go版本高于1.8，自动设置逻辑CPU数")
	}

}

func main() {
	/*
		runtime包初步使用
	*/
	fmt.Printf("----GOROOT 对应位置:[%s]\n", runtime.GOROOT())
	fmt.Printf("----操作系统信息:[%s]\n", runtime.GOOS)

}
