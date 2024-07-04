package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "<h1 style='color:red'>Hello Golang!</h1>")
}

func main() {
	/*
		go语言 http包简单使用
	*/
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server start error:%v \n", err)
		return
	}
}
