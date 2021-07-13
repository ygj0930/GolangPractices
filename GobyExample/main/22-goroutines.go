package main

import (
	"fmt"
	"time"
)

func f(from string){
	for i:=0;i<3;i++{
		fmt.Println(from,":",i)
	}
}

func main() {
	f("begin")

	go f("goroutine")

	go func(msg string){
		for i:=1;i<9;i++{
			fmt.Println("noname:",msg)
		}
	}("go noname")

	time.Sleep(5000000)
}
