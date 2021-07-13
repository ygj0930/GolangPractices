package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int,5)
	//done := make(chan bool)


	for i:=1;i<=5;i++{//发送工作任务
		jobs <- i
		fmt.Println("send job:",i)
	}
	close(jobs)//发送完毕，关闭通道
	fmt.Println("send done，close chan")

	//go func(){//处理工作
	//
	//}()

	for {
		job,ok := <- jobs

		if ok {
			fmt.Println("Read job:",job)
			time.Sleep(time.Second * 2)
		}else{
			fmt.Println("No more job")
			//done <- true
			break
		}
	}

	for elem := range jobs {
		fmt.Println(elem)
	}



	//res := <-done
	//if res {
	//	fmt.Println("done signal,exit")
	//	return
	//}
}
