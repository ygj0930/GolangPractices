package main

import "fmt"

func ping(pings chan<- string,msg string){
	pings<-msg
}

func pong(pings <-chan string,pongs chan<- string){
	msg := <-pings
	pongs <- msg
}
func main() {
	//第二个参数必须要写，不写就会导致deadlock，因为该通道无缓冲作用
	pings := make(chan string,2)
	pongs := make(chan string,2)


	ping(pings,"msg send")
	pong(pings,pongs)


	fmt.Println(<-pongs)
}
