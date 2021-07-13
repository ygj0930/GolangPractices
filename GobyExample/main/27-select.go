package main

import (
	"fmt"
	"time"
)

func sendCh1(ch chan<- string){
	var i int
	for i=1;i<=5;i++{
		ch <- fmt.Sprintf("ch1 msg send...%d",i)
		time.Sleep(time.Second*time.Duration(i))
	}
}


func sendCh2(ch chan<- string){
	var i int
	for i=1;i<=5;i++{
		ch <- fmt.Sprintf("ch2 msg send...%d",i)
		time.Sleep(time.Second*time.Duration(i))
	}
}

func main() {
	ch1 := make(chan string,1)
	ch2 := make(chan string,1)

	go sendCh1(ch1)
	go sendCh2(ch2)

	for{//死循环监听通道是否有消息到来
		select {
		case msg1:=<-ch1:
			fmt.Println(msg1)
		case msg2:=<-ch2:
			fmt.Println(msg2)
		}
	}
}
