package main

import "fmt"

func main() {
	/*
		接口的嵌套
	*/
	bird := Bird{"金丝雀"}
	bird.test1()
	bird.test2()
	bird.test3()

	fmt.Println("------------")
	var b B = bird
	b.test1()

	fmt.Println("------------")
	var c C = bird
	c.test2()

	fmt.Println("------------")
	var d D = bird
	d.test1()
	d.test2()
	d.test3()

}

type B interface {
	test1()
}

type C interface {
	test2()
}

type D interface {
	B
	C
	test3()
}

// Bird 结构体
type Bird struct {
	name string
}

func (bird Bird) test1() {
	fmt.Println(bird.name, "test1...")
}

func (bird Bird) test2() {
	fmt.Println(bird.name, "test2...")
}

func (bird Bird) test3() {
	fmt.Println(bird.name, "test3...")
}
