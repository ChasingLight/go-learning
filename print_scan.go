package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		name string
		age  int
		sex  string
	)
	//---------使用fmt包读取------------
	fmt.Println("请输入你的姓名、年龄、性别")
	// 阻塞式：读取键盘输入，通过操作地址 赋值给 name age sex。
	fmt.Scanf("%s %d %s", &name, &age, &sex)
	fmt.Printf("采集到您的姓名：%s，年龄：%d，性别：%s\n", name, age, sex)

	//---------使用bufio包读取------------
	fmt.Println("请输入一个字符串")
	reader := bufio.NewReader(os.Stdin)
	// 读取输入，以 换行符( \n )结束
	readString, _ := reader.ReadString('\n')
	fmt.Printf("通过bufio包，读取到的数据：%s", readString)
}
