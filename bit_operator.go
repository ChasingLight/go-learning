package main

import "fmt"

func main() {
	var a = 60
	var b = 13
	fmt.Printf("a 的十进制值: %d , 二进制值：%b\n", a, a)
	fmt.Printf("b 的十进制值: %d , 二进制值：%b\n", b, b)

	/*
		a: 60---0011 1100
		b: 13---0000 1101
		&		0000 1100
		|		0011 1101
		二元^	0011 0001
		&^		0011 0000
	*/
	res1 := a & b
	fmt.Println(res1) //12

	res2 := a | b
	fmt.Println(res2) //61

	res3 := a ^ b
	fmt.Println(res3) //49

	res4 := a &^ b
	fmt.Println(res4) //48

	res5 := ^a
	fmt.Println(res5) //-61

	//--------位移运算符--------
	var c = 8

	res6 := c << 2
	fmt.Println(res6)

	res7 := c >> 2
	fmt.Println(res7)
}
