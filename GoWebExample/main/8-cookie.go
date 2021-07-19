package main

import (
	"fmt"
	"net/http"
)

//设置cookie
func setCookieHandler(w http.ResponseWriter,r *http.Request){
	//创建cookie
	cookie1 := http.Cookie{
		Name:       "cookie1",
		Value:      "cookie1 by Header.Set()",
		HttpOnly:   true,
	}

	cookie2 := http.Cookie{
		Name:       "cookie2",
		Value:      "cookie2 by Header.Add()",
		MaxAge:     60,//如果不设置，就是会话级别有效期，浏览器关闭就丢失；设置了，则会持续有效时间，无论浏览器关闭与否。
		HttpOnly:   true,
	}

	cookie3 := http.Cookie{
		Name:       "cookie3",
		Value:      "cookie3 by http.SetCookie",
		HttpOnly:   true,
	}

	//方式一：通过响应头来设置
	//第一个cookie要用Set，创建一个切片出来
	w.Header().Set("Set-Cookie",cookie1.String())
	//后续的cookie用Add
	w.Header().Add("Set-Cookie",cookie2.String())

	//方式二：直接用封装好的http.SetCookie
	http.SetCookie(w,&cookie3)
}


//提取cookie
func getCookieHandler(w http.ResponseWriter,r *http.Request){
	//提取所有cookie
	cookies_beforeadd := r.Cookies()
	fmt.Printf("cookies_beforeadd Cookies:%+v\n",cookies_beforeadd)

	//提取特定cookie
	cookie1,_ := r.Cookie("cookie1")
	fmt.Printf("cookie1:%+v\n",cookie1)

	//请求中也可以新增cookie，往服务器的后续处理链条中传递
	cookie4 := http.Cookie{
		Name:       "cookie4",
		Value:      "cookie4 by request.AddCookie",
		HttpOnly:   true,
	}
	r.AddCookie(&cookie4)

	//提取所有cookie
	cookies_afteradd := r.Cookies()
	fmt.Printf("cookies_afteradd Cookies:%+v\n",cookies_afteradd)
}
func main() {
	http.HandleFunc("/setCookie",setCookieHandler)
	http.HandleFunc("/getCookie",getCookieHandler)

	http.ListenAndServe(":8080",nil)
}
