package main

import "fmt"

func main() {
	var n = 6
	res := getFibonacci(n)
	fmt.Printf("菲波那切数列，第 %d 位 是数字 %d。\n", n, res)
}

// 获取菲波那切数列，第 n 位的值
func getFibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return getFibonacci(n-2) + getFibonacci(n-1)
}
