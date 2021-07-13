package main

import (
	"fmt"
	"net/url"
)

func main() {
	p := fmt.Println

	urlStr := "http://www.baidu.com:8080/index?k=v#f"

	//url字符串转换为URL对象
	u, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	p(u)

	//通过URL对象访问url相关组成内容
	p(u.Scheme)
	p(u.Host)
	p(u.Path)
	p(u.RawQuery)

	//提取url请求参数
	param, _ := url.ParseQuery(u.RawQuery)
	p(param)
	p(param["k"][0])

}
