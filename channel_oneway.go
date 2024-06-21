package main

import "fmt"

func main() {
	/*
		双向通道：
		ch := make(chan int)
		ch 是一个双向通道，既可以 发送数据(ch <- data)，也可以接受数据(data := <- ch)。

		单向通道，又称为定向通道
		chan <- T , 只支持写
		<- chan T , 只支持读
	*/
	ch := make(chan int)
	// 等待所有 子goroutine 都执行完毕后，才结束 main方法
	done := make(chan bool)

	go produce(ch, done)
	go consume(ch, done)

	for i := 1; i <= 2; i++ {
		<-done
	}
	fmt.Println("---main over---")
}

func produce(ch chan<- int, done chan<- bool) {
	// 生产者-只能往通道中写数据
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Println("---写入通道的的数据为：", i)
	}
	// 发送完毕后，关闭通道
	close(ch)
	done <- true
}

func consume(ch <-chan int, done chan<- bool) {
	// 消费者-只能往通道中读数据
	for v := range ch {
		fmt.Println("---从通道中读取到的数据为：", v)
	}
	done <- true
}
