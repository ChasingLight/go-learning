package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		自定义错误（custom error）
		方式1：使用errors包
			errors.New("年龄不合法")
		方式2：fmt.Errorf()函数
			fmt.Errorf("年龄是：%d，不合法", age)
	*/
	age := 130
	err := checkAge(age)
	// 错误处理，属于业务过程的一部分
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("年龄合法，继续代码执行...")

}

func checkAge(age int) error {
	if age < 0 {
		// 方式1：error包
		err := errors.New("年龄不合法")
		return err
	}
	if age > 120 {
		// 方式2：fmt.Errorf
		err := fmt.Errorf("年龄：%d，不合法", age)
		return err
	}
	fmt.Println("年龄是：", age)
	return nil
}
