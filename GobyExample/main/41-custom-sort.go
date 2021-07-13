package main

import (
	"fmt"
	"sort"
)

//自定义集合类型
type stringCollection []string

//为自定义的集合类型实现sort.Interface接口相关方法：Len、Less、Swap
func (s stringCollection) Len() int{
	return len(s)
}
//比较规则：长度小的在前
func (s stringCollection) Less(i,j int) bool{
	return len(s[i])<len(s[j])
}
//交换：交换集合中的2个元素
func (s stringCollection) Swap(i,j int) {
	s[i],s[j] = s[j],s[i]
}

func main() {
	stringArray :=[]string{"a","asafsf","asf","sfsf","adgfgadgffg","asvss"}

	sort.Sort(stringCollection(stringArray))//传入实现了排序接口的数据类型

	fmt.Println(stringArray)
}
