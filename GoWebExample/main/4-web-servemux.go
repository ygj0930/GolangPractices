package main

import (
	"fmt"
	"net/http"
)

//定义请求处理函数
func customMuxHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"hello web by custom serveMux!")
}

func defaultMuxHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"hello web by default serveMux!")
}

//通过实现Handler接口，自定义请求处理器
type MyHandler3 struct {}
func (myHandler *MyHandler3) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Println("自定义多路转接器处理请求：",r.URL.Path)

	fmt.Fprintln(w,"hello web by 自定义多路转接器处理请求!")
}

func main() {
	//自行创建一个多路转接器
	myServeMux := http.NewServeMux()

	//为多路转接器绑定路由路径、映射的请求处理函数或请求处理器
	myServeMux.HandleFunc("/helloFunc",customMuxHandler)//函数用HandleFunc绑定
	myServeMux.Handle("/helloHandler",&MyHandler3{})//请求处理器用Handle函数绑定

	//测试默认多路转接器是否还会工作
	http.HandleFunc("/helloDefault",defaultMuxHandler)//404 page not found：此时默认多路转接器不起效

	//启动服务器，传入自行创建的多路转接器
	http.ListenAndServe(":8090",myServeMux)

}
