package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User2 struct {
	Id int
	UserName string
	PassWord string
	Email string
}

func templateHandler(w http.ResponseWriter,r *http.Request){
	//解析模板
	//fileNames参数的文件路径：如果是goland直接运行文件，需要先配置Configurations中Working directory为当前包
	//或者编译成exe再执行
	//否则会报错：no such file or directory
	t,errParse := template.ParseFiles("template.html")
	if errParse != nil{
		fmt.Println("ParseFiles err:",errParse)
	}

	//渲染模板并输出
	//模板引擎中的字段用 {{.xx}} 来访问
	userData := User2{
		Id: 12,
		UserName: "admin",
		PassWord: "123456",
		Email: "admin@qwe.com",
	}
	err := t.Execute(w,userData)
	if err != nil{
		fmt.Println("Execute err:",err)
	}
}
func main() {
	http.HandleFunc("/template",templateHandler)

	http.ListenAndServe(":8080",nil)
}
