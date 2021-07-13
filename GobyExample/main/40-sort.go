package main

import (
	"fmt"
	"sort"
)

func main() {
	//内置数据类型排序：排序是原地更新的
	//不同类型的排序
	strArray := []string{"e","f","g","h","k"}
	sort.Strings(strArray)
	fmt.Println(strArray)

	intArray := []int{4,8,5,1,9,3}
	sort.Ints(intArray)
	fmt.Println(intArray)

	//检查序列是否有序
	isSorted := sort.IntsAreSorted(intArray)
	fmt.Println("IntArray sorted:",isSorted)
}
