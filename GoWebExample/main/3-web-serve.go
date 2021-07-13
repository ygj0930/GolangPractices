package main

import (
	"fmt"
	"net/http"
	"time"
)

//通过实现Handler接口，来自定义请求处理器
type MyHandler2 struct {}
func (myHandler *MyHandler2) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Println("自定义配置服务器处理请求：",r.URL.Path)

	fmt.Fprintln(w,"hello web by 自定义配置服务器!")
}

func main() {

	//创建自定义请求处理器实例
	myHandler := MyHandler2{}

	//自定义配置serve
	myserve := http.Server{
		Addr: "localhost:8088",//自定义服务器监听端口
		Handler: &myHandler,
		ReadHeaderTimeout: 2*time.Second,
	}

	//启用自定义配置的服务器进行监听服务
	myserve.ListenAndServe()
}
