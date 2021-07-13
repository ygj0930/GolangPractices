package main

import (
	"fmt"
	"net/http"
)

//请求处理函数
//handler 函数有两个参数，`http.ResponseWriter` 和 `http.Request`
//response writer 被用于写入 HTTP 响应数据
func helloWeb(respWriter http.ResponseWriter, req *http.Request) {
	//返回应答内容
	fmt.Fprintf(respWriter, "hello go web!\n")

	//读取请求头并返回
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(respWriter, "got header,%v:%v", name, h)
		}
	}
}

func main() {
	//注册请求handler函数到应用路由：将某路径的api请求交给对应handler函数处理
	http.HandleFunc("/hello", helloWeb)

	//启动监听:addr指定监听端口，nil告诉它使用我们刚刚设置的默认路由器。
	http.ListenAndServe(":8090", nil)
}

//进行访问：curl localhost:8090/hello
