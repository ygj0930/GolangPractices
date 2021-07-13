package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//设置环境变量
	os.Setenv("Foo", "custom env val")

	//获取某环境变量值
	fmt.Println("Foo:", os.Getenv("Foo"))

	//获取所有环境变量
	all := os.Environ()
	fmt.Println("=========")
	fmt.Println(all)
	fmt.Println("=========")

	//遍历环境变量键值对
	for _, kv := range os.Environ() {
		fmt.Println("key-value:", kv)
		pair := strings.Split(kv, "=")
		fmt.Println("key:", pair[0])
		fmt.Println("value:", pair[1])
	}
}
