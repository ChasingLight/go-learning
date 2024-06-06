package main

import "fmt"

func main() {
	/*
		Part1：声明初始化map
	*/
	// 1. 仅声明，默认为nil
	var map1 map[string]int
	fmt.Println(map1)
	fmt.Println(map1 == nil)

	// 2. 声明 + 初始化
	var map2 = map[string]int{"jaden": 30, "ribbon": 29, "luna": 18}
	fmt.Println(map2)
	fmt.Println(map2 == nil)

	// 3. make函数 创建map
	var map3 = make(map[string]int)
	fmt.Println(map3)
	fmt.Println(map3 == nil)

	/*
			Part2：map的使用
		    具体操作：添加/修改、遍历、获取、删除、长度
	*/

	fmt.Println("---------map使用---------------")
	// 添加/修改
	var nameScoreMap = make(map[string]float64)
	nameScoreMap["张三"] = 85.6
	nameScoreMap["李四"] = 95
	nameScoreMap["王五"] = 92.5
	fmt.Println(nameScoreMap)
	nameScoreMap["张三"] = 75.6
	fmt.Println(nameScoreMap)

	// 遍历---map中键值对是无序的，多次打印顺序可能不同。
	fmt.Println("--------for-range遍历map--------------")
	for name, score := range nameScoreMap {
		fmt.Printf("学生姓名：%s，成绩：%.2f\n", name, score)
	}

	// 获取
	score := nameScoreMap["李四"]
	fmt.Printf("map 中'李四'的成绩为: %.2f\n", score)

	score, isKeyExist := nameScoreMap["乌拉乌拉"]
	if isKeyExist {
		fmt.Printf("key 在 map 中存在，对应的value为：%f", score)
	} else {
		fmt.Printf("key 在 map 中不存在，对应的value为 值-数据类型的默认值：%.2f\n", score)
	}

	// 删除---如果key存在，正常删除；如果key不存在，map不变。
	delete(nameScoreMap, "李四")
	fmt.Println(nameScoreMap)

	// 长度---即：键值对个数
	len := len(nameScoreMap)
	fmt.Printf("map 中键值对的个数：%d\n", len)
}
