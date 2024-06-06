package main

import "fmt"

func main() {
	/*
		声明初始化map
	*/
	var map1 map[string]int	//仅声明，默认为nil
	var map2 = map[string]int{"jaden":30, "ribbon":29, "luna":18}
	var map3 = make(map[string]int)
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)

}
