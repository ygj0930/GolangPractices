package main

import (
	"fmt"
	"time"
)

func main() {
	requestsCh := make(chan int,5)
	for i:=1;i<=5;i++ {
		requestsCh <- i
	}
	close(requestsCh)

	//基本限流：通过打点器来控制请求处理的间隔
	limiter := time.Tick(time.Second * 2)
	for req := range requestsCh {
		<-limiter//通过打点器来阻塞请求处理，每2秒处理1个

		fmt.Println("doing request:",req,time.Now())
	}


	//脉冲限流：通过带缓冲的脉冲通道来控制请求处理的进行；通过打点器以一定速率生成脉冲信号
	pulseLimiter := make(chan time.Time,3)
	for i := 1; i <= 3; i++ {
		pulseLimiter <- time.Now()
	}

	go func(){
		for t:=range time.Tick(time.Second) {
			pulseLimiter <- t //每秒生成1个脉冲信号
		}
	}()

	//生成请求
	pulseRequests := make(chan int,5)
	for i := 1; i <= 5; i++ {
		pulseRequests <- i
	}
	close(pulseRequests)

	//处理请求
	for req := range pulseRequests {
		<-pulseLimiter
		fmt.Println("doing pulseLimiter request:",req,time.Now())
		//time.Sleep(time.Second * 2)
	}
}
