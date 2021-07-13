package main

import (
	"fmt"
	"math/rand"
	"time"
)

type readRequest struct {
	key  int
	resp chan int //用通道来回传响应：值
}

type writeRequest struct {
	key  int
	val  int
	resp chan bool //用通道来回传响应：是否操作成功
}

func main() {
	readChan := make(chan *readRequest)
	writeChan := make(chan *writeRequest)

	//写协程：启动100个协程发起写请求
	for i:=1;i<=10;i++ {
		go func(){
			for{
				writeReqBody := &writeRequest{
					key:rand.Intn(100),
					val: rand.Intn(100),
					resp: make(chan bool)}

				//发送写请求
				writeChan<-writeReqBody

				//通过响应通道获取结果
				res := <-writeReqBody.resp
				fmt.Printf("write result:%b\n",res)

				time.Sleep(time.Second)
			}
		}()
	}

	//读协程：启动100个协程发起读请求
	for i:=1;i<=10;i++ {
		go func(){
			for{
				readReqBody := &readRequest{
					key:rand.Intn(100),
					resp: make(chan int)}
				//发送读请求
				readChan<-readReqBody
				//通过响应通道获取结果
				res := <-readReqBody.resp
				fmt.Printf("read result:%d\n",res)

				time.Sleep(time.Second)
			}
		}()
	}


	//主线程：维护共享变量，监听共享变量的读、写请求并处理
	var stateMap = make(map[int]int)
	for {
		select {
		case readReq := <-readChan:
			readReq.resp <- stateMap[readReq.key]
			fmt.Printf("read key:%d\n", readReq.key)
		case writeReq := <-writeChan:
			stateMap[writeReq.key] = writeReq.val
			fmt.Printf("write key:%d,val:%d\n", writeReq.key, writeReq.val)
			writeReq.resp <- true
		}
	}

}
