package main

import (
	"fmt"
	"time"
)

func systemWork(ch chan string){
	fmt.Println("System work begin...")

	time.Sleep(time.Second * 2)

	//发送消息，通知主线程系统调用已完成
	ch<-"System work done"
}

func main() {
	ch := make(chan string,1)

	//发起系统调用
	go systemWork(ch)

	//监听系统调用结果
	select {
	case res:=<-ch:
		fmt.Println(res)
	case <-time.After(time.Second * 2)://用于阻塞select操作一定时间
		fmt.Println("timeout 2 sec")
	}
}
