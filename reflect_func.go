package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		函数(func) 反射使用
		复习：
			函数 是一种复合数据类型，本质上是一个变量。
			函数调用：func()
			函数作为变量：func
	*/
	funcValue := reflect.ValueOf(CalAverage)                                             //核心位置：函数本质是一个变量，反射直接传入函数名！！！
	fmt.Printf("CalAverage 函数---kind:%s, type:%s\n", funcValue.Kind(), funcValue.Type()) //kind:func, type:func(...float64) float64
	// 函数直接调用 Call 即可
	args1 := []reflect.Value{reflect.ValueOf(12.3), reflect.ValueOf(12.4), reflect.ValueOf(12.5)}
	resultValue := funcValue.Call(args1)
	average := resultValue[0].Interface().(float64)
	fmt.Println("CalAverage 函数的返回值：", average)

}

// CalAverage 计算平均值
func CalAverage(a ...float64) (average float64) {
	sum := 0.0
	for _, value := range a {
		sum += value
	}
	average = sum / float64(len(a))
	return average
}
