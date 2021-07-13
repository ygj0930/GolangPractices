package main

import (
	"fmt"
	"sync"
	"time"
)

func watiworker(workerId int,waitGroup *sync.WaitGroup){
	fmt.Printf("Worker- %d is working...\n",workerId)
	time.Sleep(time.Second * time.Duration(workerId))

	fmt.Printf("Worker- %d is done...\n",workerId)
	waitGroup.Done()//完成任务，计数-1
}

func main() {
	waitGroup := sync.WaitGroup{}//相当于java的countDownLatch

	for i:=1;i<=5;i++ {
		waitGroup.Add(1)//设置计数+1

		go watiworker(i,&waitGroup)
	}

	waitGroup.Wait()//相当于carrier栅栏，等待waitGroup上的协程全部完成

	fmt.Println("All done")
}
