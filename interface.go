package main

import "fmt"

func main() {
	/*
		接口（interface）：
		在Go语言中，接口是一组方法的签名。
		当某个类型实现了接口中 所有的方法，它被称为实现接口。

		Go语言中，接口和类型的实现关系，是非侵入式。
		Java语言中，接口和实现类之间，是侵入式的。
		Eg：class Mouse implements USB{...}
	*/
	// 1.创建Mouse
	m1 := Mouse{"罗技M330"}
	fmt.Println(m1.name)
	// 2.创建FlashDisk
	f1 := FlashDisk{"闪迪16G"}
	fmt.Println(f1.name)

	testInterface(m1)
	// 如果FlashDisk没有全部实现 USB接口中方法，会提示如下错误：
	// cannot use f1 (variable of type FlashDisk) as USB value in argument to testInterface: FlashDisk does not implement USB (missing method end)
	testInterface(f1)

	var usb USB = m1
	usb.start()
	usb.end()
	//usb.name	//usb.name undefined
}

// 1.定义接口
type USB interface {
	start() //USB设备开始工作
	end()   //USB设备结束工作
}

// 2.实现类
type Mouse struct {
	name string
}

type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "开始工作...点点点")
}

func (m Mouse) end() {
	fmt.Println(m.name, "结束工作...")
}

func (f FlashDisk) start() {
	fmt.Println(f.name, "开始数据读写工作...")
}

func (f FlashDisk) end() {
	fmt.Println(f.name, "可以弹出...")
}

// 3.测试方法
func testInterface(usb USB) {
	usb.start()
	usb.end()
}
