package main

import "fmt"

func main() {
	/*
		接口断言（interface assertion）：
			方式一：
				1. instance := 接口对象.(实际类型)	//不安全，会panic()
				2. instance, ok := 接口对象.(实际类型)	//安全
			方式二：switch
				switch instance := 接口对象.(type){
					case 实际类型1:
						...
					case 实际类型2:
						...
				}//end switch
	*/
	superman := Superman{"Clark Kent", 500}
	bee := Bee{"尼莫"}
	testFlyInterface(superman)
	testFlyInterface(bee)

	fmt.Println("------------方式一：接口断言-------------------")
	interfaceAssertion1(superman)
	interfaceAssertion1(bee)

	fmt.Println("------------方式二：接口断言-------------------")
	interfaceAssertion2(superman)
	interfaceAssertion2(bee)

}

// Fly 接口
type Fly interface {
	flying()
}

// Superman 结构体，并实现接口
type Superman struct {
	name string
	age  int
}

func (s Superman) flying() {
	fmt.Println(s.name, "超人使用超能力在飞...")
}

// Bee 结构体，并实现接口
type Bee struct {
	name string
}

func (b Bee) flying() {
	fmt.Println(b.name, "蜜蜂使用翅膀在飞...")
}

// 接口测试方法
func testFlyInterface(f Fly) {
	f.flying()
}

func interfaceAssertion1(f Fly) {
	superman, isSuperman := f.(Superman)
	if isSuperman {
		fmt.Printf("超人的属性，姓名：%s，年纪：%d\n", superman.name, superman.age)
	}

	bee, isBee := f.(Bee)
	if isBee {
		fmt.Printf("蜜蜂的属性，姓名：%s\n", bee.name)
	}

	f.flying()
}

func interfaceAssertion2(f Fly) {
	switch instance := f.(type) {
	case Superman:
		fmt.Printf("超人的属性，姓名：%s，年纪：%d\n", instance.name, instance.age)
	case Bee:
		fmt.Printf("蜜蜂的属性，姓名：%s\n", instance.name)
	default:
		fmt.Println("未知数据类型")
	}
	f.flying()
}
