package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id int
	UserName string
	PassWord string
	Email string
}

func jsonHandler(w http.ResponseWriter,r *http.Request){
	//设置响应类型
	w.Header().Set("Content-Type","application/json")

	//创建响应结果
	user := User{
		Id: 12,
		UserName: "admin",
		PassWord: "123456",
		Email: "admin@qwe.com",
	}

	//转换为json格式返回
	jsonRes,_ := json.Marshal(user)

	w.Write(jsonRes)
}

func redirectHandler(w http.ResponseWriter,r *http.Request){
	//设置重定向目标地址
	w.Header().Set("Location","https://www.baidu.com")
	//设置状态码
	w.WriteHeader(302)
}
func main() {
	http.HandleFunc("/json",jsonHandler)
	http.HandleFunc("/redirect",redirectHandler)

	http.ListenAndServe(":8090",nil)
}
