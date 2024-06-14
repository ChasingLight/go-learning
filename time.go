package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		time包：go语言中专门操作时间相关的包。
	*/

	// 1.获取当前时间
	t1 := time.Now()
	fmt.Printf("%T\n", t1)
	fmt.Println(t1)

	// 2.获取指定时间
	t2 := time.Date(1994, 5, 22, 18, 20, 50, 0, time.Local)
	fmt.Println(t2)

	// 3.time--->string
	// 格式模版，记忆小妙招：“2006-1-2-3-4-5”
	s1 := t2.Format(time.DateTime)
	s2 := t2.Format("2006年 01月 02日 15时 04分 05秒")
	fmt.Println(s1)
	fmt.Println(s2)

	// 4.string--->time
	s3 := "1994-06-08 12:00:00"
	t3, err := time.Parse("2006-01-02 15:04:05", s3)
	if err != nil {
		fmt.Println("字符串转换time错误", err)
	}
	fmt.Printf("字符串 转换 time，结果数据类型：%T，对应数值为：%v\n", t3, t3)

	// 5.根据当前时间，获取指定内容
	year, month, day := t1.Date()
	fmt.Printf("通过 time.Time 的 Date() 函数，获取当前时间的 年：%d， 月：%d，日：%d\n", year, int(month), day)
	hour, min, sec := t1.Clock()
	fmt.Printf("通过 time.Time 的 Clock() 函数，获取当前时间的 时：%d， 分：%d，秒：%d\n", hour, min, sec)
	// 获取-独立-指定内容
	fmt.Println("年：", t1.Year())
	fmt.Println("月：", int(t1.Month()))
	fmt.Println("日：", t1.Day())
	fmt.Println("时：", t1.Hour())
	fmt.Println("分：", t1.Minute())
	fmt.Println("秒：", t1.Second())
	fmt.Println("星期几：", t1.Weekday())

	// 6.时间戳：指定的日期，距离1970-01-01 00:00:00 的时间差值。秒、纳秒
	t4 := time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC)
	fmt.Println(t4.Unix()) //秒
	fmt.Println(t1.Unix())
	fmt.Println(t4.UnixNano()) //纳秒
	fmt.Println(t1.UnixNano())

	// 7.时间间隔
	t5 := t1.Add(-2 * time.Hour)
	fmt.Println(t1)
	fmt.Println(t5)
	duration := t5.Sub(t1)
	fmt.Printf("%T\n", duration)
	fmt.Println(duration)

	// 8.睡眠
	time.Sleep(3 * time.Second)
	fmt.Println("main over...")

}
