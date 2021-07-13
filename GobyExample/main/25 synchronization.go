package main

import (
	"fmt"
	"time"
)

func worker(done chan bool)  {
	fmt.Println("working......")

	time.Sleep(time.Second * 10)

	fmt.Println("done!")

	done <- true
}

func main() {
	done := make(chan bool)

	go worker(done)

	fmt.Println("waiting...")
	//result := <-done
	//if result{
	//	os.Exit(1)
	//}
}
