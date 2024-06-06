package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	/*
		使用Go 随机数-生成4位短信验证码
	*/
	var code int
	for i := 1; i <= 3; i++ {
		code = rand.IntN(9000) + 1000
		fmt.Printf("随机数-生成的4位短信验证码：%d\n", code)
	}
}
