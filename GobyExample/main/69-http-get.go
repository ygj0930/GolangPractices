package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	//发起请求
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//获取响应状态相关信息
	fmt.Println("Resp status:", resp.Status)

	//读取响应流内容
	httpScanner := bufio.NewScanner(resp.Body)
	for httpScanner.Scan() {
		fmt.Println(httpScanner.Text())
	}
}
