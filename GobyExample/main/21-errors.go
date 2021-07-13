package main

import (
	"errors"
	"fmt"
)

func f1(a int)(int,error){
	if a < 5 {
		return -1,errors.New("buildin error")
	}

	return a+1,nil
}

type argError struct {
	errorCode int
	errorMsg string
}

func (e argError) Error() string{
	return fmt.Sprintf("%d - %s",e.errorCode,e.errorMsg)
}

func f2(a int)(int,error){
	if a<5 {
		return -1,argError{-1,"custom error"}
	}
	return 1,nil
}
func main() {
	for i:=3;i<9;i++{
		if code,e := f1(i);code<0 {
			fmt.Println("Code:",code,",error:",e)
		}else{
			fmt.Println("Code:",code)
		}

		if code,e := f2(i);code<0 {
			fmt.Println("Code:",code,",error:",e)
		}else{
			fmt.Println("Code:",code)
		}
	}
}
