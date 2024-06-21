package main

import "fmt"

func main() {
	/*
		发送者-关闭通道
		close(channel)

		接收者-读取通道数据，2种方式
		1> 读取通道数据时，使用额外的变量 isOpen 判断通道是否关闭；
			v, isOpen := <- ch1
		2> for range, 从通道接收值，直到关闭为止。【推荐使用这种方式】
			for v := range ch1{...}
	*/
	ch1 := make(chan int)

	go sendData(ch1)

	// Way1：额外变量，判断通道是否关闭
	// 死循环-读取通道中数据，直到 发送者-关闭通道
	/*for {
		value, isOpen := <-ch1
		if !isOpen {
			fmt.Println("---通道关闭，已读取所有的数据了", value, isOpen)
			break
		}
		fmt.Println("通道开启，读取到的数据：", value, isOpen)
	}*/

	// Way2：for range 读取通道中数据
	for v := range ch1 { // v <- ch1
		fmt.Println("读取到的数据：", v)
	}
}

func sendData(ch1 chan int) {
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
	// 关闭通道
	close(ch1)
}
