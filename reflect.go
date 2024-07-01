package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		reflect 反射 使用
		转化为 reflect 对象，根据不同情况调用
		t := reflect.TypeOf(i)	//得到类型元数据，通过 t 可获取类型定义的所有元素
		v := reflect.ValueOf(i)	//得到实际值，通过 v 还可以改变值
	*/

	// 反射操作：通过反射，获取一个 interface 类型变量的 类型和值
	x := 66.88
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Printf("%T,%v\n", t, t)
	fmt.Printf("%T,%v\n", v, v)
	fmt.Println("kind of float64:", v.Kind())
	xx := v.Interface()
	fmt.Printf("%T,%v\n", xx, xx)
	tt := v.Type()
	fmt.Printf("%T,%v\n", tt, tt)

	var all interface{}
	all = x
	fmt.Println(all)
}
