package main

import(
	"fmt"
	"math"
)

const strConst string = "constant string"

func main() {
	fmt.Println(strConst)

	const n = 5000000000
	const d = 3e20/n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
