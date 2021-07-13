package main

import (
	"fmt"
	"net/http"
)

//定义请求处理函数
func helloHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"hello web!")
}

func main() {
	//注册请求处理路由
	http.HandleFunc("/hello",helloHandler)

	//注册web服务器监听端口以及所使用的请求处理器
	//传入nil：使用默认的请求处理器 DefaultServeMux
	http.ListenAndServe("localhost:8090",nil)

	//启动go程序，浏览器输入 localhost:8090/hello，查看结果
}
