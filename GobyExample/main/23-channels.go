package main

import "fmt"

func main() {
	ch := make(chan string)

	go func(){
		ch<-"ping"
	}()

	msg := <- ch

	fmt.Println(msg)

	messages := make(chan string,2)
	messages <- "buffer"
	messages <- "chan"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	messages <- "chan2"
	fmt.Println(<-messages)
}
