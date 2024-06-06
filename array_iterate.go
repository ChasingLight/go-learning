package main

import "fmt"

func main() {
	/*
		Array数组遍历
	*/
	var scoreArray = []float64{98.1, 65, 78.3, 86.7}
	// 1. for + 数组下标 遍历
	for i := 0; i < len(scoreArray); i++ {
		fmt.Printf("数组下标 %d 对应的值 %.1f\n", i, scoreArray[i])
	}

	fmt.Println("-------------")
	// 2. for + range 遍历
	for index, value := range scoreArray {
		fmt.Printf("数组下标 %d 对应的值 %.1f\n", index, value)
	}

	// 3. for + range 遍历（只需要值，忽略下标）
	sum := 0.0
	for _, score := range scoreArray {
		sum += score
	}
	average := sum / float64(len(scoreArray))
	fmt.Printf("总成绩：%.2f，平均分：%.2f\n", sum, average)
}
