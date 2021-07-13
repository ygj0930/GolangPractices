package main

import (
	"fmt"
	"net/http"
)

//通过实现Handler接口，来自定义请求处理器
type MyHandler struct {}
func (myHandler *MyHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Println("自定义处理器处理请求：",r.URL.Path)

	fmt.Fprintln(w,"hello web by myHandler!")
}

func main() {
	//创建自定义请求处理器实例
	myHandler := MyHandler{}

	//注册http所用请求处理器
	http.Handle("/hello-myhandler",&myHandler)

	//启动web服务器
	http.ListenAndServe("localhost:8090",nil)

	//启动go程序，浏览器输入 localhost:8090/hello-myhandler，查看结果
}
