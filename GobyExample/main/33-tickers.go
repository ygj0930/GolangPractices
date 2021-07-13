package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	time.Sleep(time.Second * 3)
	go func(){
		for t:= range ticker.C {
			fmt.Println("Tick:",t)
		}
	}()
	ticker.Stop()
	fmt.Println("Ticker stopped")
	time.Sleep(time.Second * 3)
}
