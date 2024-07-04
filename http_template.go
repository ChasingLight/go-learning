package main

import (
	"fmt"
	"html/template"
	"math/rand/v2"
	"net/http"
	"strconv"
)

type User struct {
	Name   string
	Age    int
	Gender string
	Code   string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 1.定义模版 同目录下的 hello.html------{{ . }} 这里的点 代表后端传递过来的 当前对象

	// 2.解析模版
	t, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Printf("parse template fail, err:%v\n", err)
	}
	// 传递多个参数，在模版中进行渲染---结构体、map
	code := rand.IntN(9000) + 1000
	user := User{Name: "靳浩东", Age: 30, Gender: "男", Code: strconv.Itoa(code)}

	map1 := make(map[string]interface{})
	map1["name"] = "张瑞斌"
	map1["age"] = 30
	map1["gender"] = "女"
	map1["code"] = strconv.Itoa(rand.IntN(9000) + 1000)

	dataMap := make(map[string]interface{})
	dataMap["user"] = user
	dataMap["map1"] = map1

	// 3.渲染模版
	err = t.Execute(w, dataMap)
	if err != nil {
		fmt.Printf("render template fail, err:%v\n", err)
	}
}

func main() {
	/*
		http包 + template 简单使用
	*/
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server start error:%v \n", err)
		return
	}
}
