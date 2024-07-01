package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		结构体、方法，reflect反射 使用
	*/
	p1 := People{"靳浩东", 17, "男"}

	// 结构体变量 -> reflect.Type
	reflectType := reflect.TypeOf(p1)
	fmt.Println("反射类型的名称：", reflectType.Name()) //People
	fmt.Println("反射类型的种类：", reflectType.Kind()) //struct

	// 结构体变量 -> reflect.Value
	reflectValue := reflect.ValueOf(p1)
	fmt.Println("反射值的所有的字段：", reflectValue) //{靳浩东 17 男}

	fmt.Println("------通过反射，获取结构体-所有字段的名称、字段的类型、字段的值------")
	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		filedValue := reflectValue.Field(i).Interface()
		j := i + 1
		fmt.Printf("第%d个字段：字段名称：%s，字段类型：%s，字段数值：%v\n", j, field.Name, field.Type, filedValue)
	}

	fmt.Println("------通过反射，获取结构体-所有方法------")
	for i := 0; i < reflectType.NumMethod(); i++ {
		method := reflectType.Method(i)
		j := i + 1
		fmt.Printf("第%d个方法：方法的名称：%s，方法的类型：%s\n", j, method.Name, method.Type)
	}

	fmt.Println("------通过反射，修改结构体中的值------")
	pointer := reflect.ValueOf(&p1) //注意点！！！ 要传递指针地址
	elem := pointer.Elem()
	fmt.Println("类型：", elem.Type())         //main.People
	fmt.Println("是否可以修改数据：", elem.CanSet()) //true
	// 通过字段下标进行修改
	elem.Field(0).SetString("张瑞斌")
	// 通过字段名称进行修改
	elem.FieldByName("Age").SetInt(16)
	elem.FieldByName("Sex").SetString("女")
	fmt.Println(p1) //{张瑞斌 16 女}

	fmt.Println("------------通过反射，调用方法------------")
	methodValue1 := reflectValue.MethodByName("PrintInfo")
	fmt.Printf("PrintInfo 方法---kind:%s, type:%s\n", methodValue1.Kind(), methodValue1.Type()) //kind:func, type:func()
	// 无参调用：可直接写 nil 或 创建一个空的切片
	methodValue1.Call(nil)
	args1 := make([]reflect.Value, 0)
	methodValue1.Call(args1)

	methodValue2 := reflectValue.MethodByName("Say")
	fmt.Printf("Say 方法：kind:%s, type:%s\n", methodValue2.Kind(), methodValue2.Type()) //kind:func, type:func(string)
	args2 := []reflect.Value{reflect.ValueOf("今天天气好晴朗，处处好风光！")}
	methodValue2.Call(args2)

	methodValue3 := reflectValue.MethodByName("DoAddNum")
	fmt.Printf("DoAddNum 方法：kind:%s, type:%s\n", methodValue3.Kind(), methodValue3.Type())
	args3 := []reflect.Value{reflect.ValueOf(120.5), reflect.ValueOf(126.3)}
	resultValue := methodValue3.Call(args3)
	fmt.Printf("%T\n", resultValue)
	fmt.Println(len(resultValue))
	sum := resultValue[0].Interface().(float64)
	fmt.Println("DoAddNum 方法的返回值：", sum)
}

type People struct {
	Name string
	Age  int
	Sex  string
}

func (p People) PrintInfo() {
	fmt.Printf("姓名：%s，年龄：%d，性别：%s\n", p.Name, p.Age, p.Sex)
}

func (p People) Say(msg string) {
	fmt.Println("说话：", msg)
}

// DoAddNum 方法名称首字母必须大写，才能使用 reflect 反射进行调用
func (p People) DoAddNum(a, b float64) (result float64) {
	return a + b
}
