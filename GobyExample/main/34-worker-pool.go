package main

import (
	"fmt"
	"time"
)

func workerdo(workerId int,jobs <-chan string,results chan<- string){
	for job := range jobs {
		fmt.Println("worker:",workerId,"doing job:",job)
		time.Sleep(time.Second)
		results<-fmt.Sprintf("Job-%s,done by worker-%d",job,workerId)
	}
}

func main() {
	jobs := make(chan string,10)
	results := make(chan string,10)

	for w:=1;w<=3;w++ {
		go workerdo(w,jobs,results)
	}

	for j:=1;j<=9;j++{
		jobs <- fmt.Sprintf("jobcontent-%d",j)
	}
	//close(jobs)

	for j:=1;j<=9;j++{
		select {
		case res:=<-results:
			fmt.Println(res)
		}
	}
}
