package main

import "fmt"

func main() {
	i := 1
	for i<3 {
		fmt.Println(i)
		i++
	}

	for j:=1;j<=9;j++ {
		fmt.Println(j)
	}

	l:=1
	for{
		fmt.Println("loop")
		if l>9 {
			break
		}else{
			l++
		}
	}

	for j:=1;j<=9;j++{
		if j%2==0{
			continue
		}
		fmt.Println(j)
	}
}
