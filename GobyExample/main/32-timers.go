package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	<- timer1.C
	fmt.Println("Timer1 is expired")

	time2 := time.NewTimer(time.Second)
	go func(){
		<-time2.C
		fmt.Println("Timer2 is expired")
	}()
	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("Timer2 is stopped")
	}

}
