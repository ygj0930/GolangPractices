package main

import "fmt"

func main() {
	nums:=[]int{2,3,4,5}
	sum:=0
	for _,num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for index,num := range nums {
		fmt.Println("index:",index,",num:",num)
	}

	kvs:=map[string]string{"a":"apple","b":"boy"}
	for k,v := range kvs {
		fmt.Println("key:",k,",val:",v)
	}

	for k := range kvs{
		fmt.Println(k)
	}

	str:="abcdefg"
	for i,c := range str{
		fmt.Println("index:",i,",char:",c)
	}
}
