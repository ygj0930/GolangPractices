package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[4]=100
	fmt.Println(a)

	fmt.Println("len:",len(a))

	b:=[5]int{1,2,3,4,5}
	fmt.Println(b)

	var twoDArray [3][3]int
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			twoDArray[i][j]=i+j
		}
	}
	fmt.Println(twoDArray)
}
