package main

import "fmt"

func plus(a,b int) int {
	return a+b
}

func swap(a,b int) (int,int) {
	return b,a
}

func sum(nums ...int) int{
	numSum := 0
	for _,num := range nums{
		numSum += num
	}
	return numSum
}

func fact(n int) int{
	if n==0{
		return 1
	}
	return n * fact(n-1)
}

func main() {
	res := plus(2342,2355)
	fmt.Println(res)

	a,b := swap(123,456)
	fmt.Println(a,b)

	fmt.Println(sum(1,2,3))
	fmt.Println(sum(1,2,3,4,5,6))
	fmt.Println(sum(1,2,3,4,5,6,7,8,9))

	fmt.Println(fact(7))
}
