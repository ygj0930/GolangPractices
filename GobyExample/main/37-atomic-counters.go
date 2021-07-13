package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	var counter uint64 = 0

	//并发更新：计数器+1
	for i:=0;i<50;i++ {
		go func(){
			//进行原子计数
			atomic.AddUint64(&counter,1)
			fmt.Println(counter)
			//协程让步，让其他协程执行
			runtime.Gosched()
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("final:",counter)
}
