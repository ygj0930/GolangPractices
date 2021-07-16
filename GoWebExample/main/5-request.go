package main

import (
	"fmt"
	"net/http"
)

//请求处理
func reqHandler(w http.ResponseWriter,r *http.Request){
	//提取请求路径内容
	//fmt.Fprintln(w,"请求地址：",r.URL.Path)
	//fmt.Fprintln(w,"查询参数：",r.URL.RawQuery)

	//提取请求头信息
	//fmt.Fprintln(w,"请求头整体：",r.Header)
	//Get：拿到值字符串
	//fmt.Fprintln(w,"请求头ByGet-Accept：",r.Header.Get("Accept-Encoding"))
	//[]：拿到值切片，可以按下标提取具体值
	//fmt.Fprintln(w,"请求头By[]-Accept-Encoding：",r.Header["Accept-Encoding"])
	//fmt.Fprintln(w,"请求头By[]-Accept-Encoding[0]：",r.Header["Accept-Encoding"][0])


	//提取请求体信息：要用post来传才有请求体内容
	//bodyLenth := r.ContentLength
	//fmt.Fprintln(w,"请求体大小：",bodyLenth)
	//fmt.Fprintln(w,"请求体：",r.Body)

	//bodyContent := make([]byte,bodyLenth)
	//r.Body.Read(bodyContent)
	//fmt.Fprintln(w,"请求体内容：",string(bodyContent))

	//提取请求参数
	//方法一：进行转换，将URL中的请求参数对和body中的请求参数对转换到Form字段
	//r.ParseForm()
	//fmt.Fprintln(w,"Form：",r.Form)
	//通过Form字段提取具体参数
	//fmt.Fprintln(w,"Form field user：",r.Form.Get("user"))
	//fmt.Fprintln(w,"Form field password：",r.Form.Get("password"))

	//方法二：快速获取参数
	//获取参数值
	fmt.Fprintln(w,"FormValue user：",r.FormValue("user"))
	fmt.Fprintln(w,"FormValue password：",r.FormValue("password"))
}


func main() {
	http.HandleFunc("/test/request",reqHandler)

	http.ListenAndServe(":8090",nil)
	//浏览器测试：http://localhost:8090/test/request?user=1&login=enable
}
