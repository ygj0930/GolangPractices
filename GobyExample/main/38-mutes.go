package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	state := make(map[int]int)
	mutex := &sync.Mutex{}

	for i:=0;i<10;i++ {//并发读
		go func(){
			for{
				key := rand.Intn(100)

				mutex.Lock()//加锁
				val,ok := state[key]
				if ok {
					fmt.Printf("read key:%d,val:%d",key,val)
				}
				mutex.Unlock()//释放锁
				runtime.Gosched()
			}
		}()
	}

	for j:=0;j<5;j++ {//并发写
		go func(){
			for {
				key := rand.Intn(100)
				val := rand.Intn(500)

				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				runtime.Gosched()
			}
		}()
	}

	mutex.Lock()
	fmt.Println("before:",state)
	mutex.Unlock()

	time.Sleep(time.Second)

	mutex.Lock()
	fmt.Println("after:",state)
	mutex.Unlock()
}
