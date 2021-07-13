package main

import "fmt"

func main() {
	ch:=make(chan string)

	select {
	case ch<-"send":
		fmt.Println("send msg")
	default:
		fmt.Println("No message can send")
	}

	select {
	case msg:=<-ch:
		fmt.Println(msg)
	default:
		fmt.Println("No message can read")
	}


}
