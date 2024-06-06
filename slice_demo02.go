package main

import "fmt"

func main() {
	/*
		切片 slice：
		1. 每个切片引用一个底层数组；
		2. 切片本身不存储任何数据，是这个底层数组存储；
		3. 当向切片中添加数据时，如果没有超过容量cap，直接添加；如果超过容量，自动扩容（？？？扩容底层细节待学习）；
		4. 切片一旦扩容，就重新指向一个底层数组。
	*/
	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	fmt.Printf("len: %d, cap: %d\n", len(s1), cap(s1)) //len: 3, cap: 3
	fmt.Printf("引用地址值：%p\n", s1)

	s1 = append(s1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)
	fmt.Println(s1)
	fmt.Printf("len: %d, cap: %d\n", len(s1), cap(s1)) //len: 13, cap: 14
	fmt.Printf("引用地址值：%p\n", s1)

	// 基于数组创建切片
	arr := [5]int{1, 2, 3, 4, 5}
	s4 := arr[1:4]  //左闭，右开
	fmt.Println(s4) // 输出: [2 3 4]
}
