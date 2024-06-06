package main

import (
	"fmt"
	"math"
)

func main() {
	// int8 和 uint8 区别
	// 前者有符号位（最高位是符号位，0代表正数，1代表复数，其余7为是数据位），后者无符号位。
	var minInt8 = math.MinInt8
	var maxInt8 = math.MaxInt8
	fmt.Printf("int8 的取值范围：%d 到 %d\n", minInt8, maxInt8)

	var minUint8 = 0
	var maxUint8 = math.MaxUint8
	fmt.Printf("uint8 的取值范围：%d 到 %d\n", minUint8, maxUint8)

	//------数据类型转换---------
	// 语法格式：Type(value)
	// 注意点：1> 常数在有需要的时候，自动转型；2> 变量需要手动转型
	var f1 = 4.83
	fmt.Printf("%T , %f\n", f1, f1)

	// 常数100，自动转换为float64
	sum := f1 + 100
	fmt.Printf("%T , %f\n", sum, sum)

	// 变量c 手动转型
	var c int
	c = int(f1)
	fmt.Printf("%T , %d\n", c, c)
}
