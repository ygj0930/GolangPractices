package main

import (
	"fmt"
	"regexp"
)

func main() {
	//编译一个正则结构体
	//变量编译
	pattern, err := regexp.Compile("p([a-z]+)ch")
	if err != nil {
		panic(err)
	}
	fmt.Println(pattern)
	//常量编译
	constPattern := regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(constPattern)

	//===字符串匹配===
	//匹配校验
	fmt.Println(pattern.MatchString("peach"))
	//查找第一次匹配的字符串
	fmt.Println(pattern.FindString("peach puncha patch"))
	//查找第一个匹配的字符串的下标索引：开始、结束索引
	fmt.Println(pattern.FindStringIndex("peach puncha patch"))
	//查找完全匹配与局部匹配的字符串
	fmt.Println(pattern.FindStringSubmatch("peach puncha patch"))
	//指定查找多少匹配项：-1 查所有，大于1则查指定数量匹配项返回
	fmt.Println(pattern.FindAllString("peach puncha patch", -1))
	fmt.Println(pattern.FindAllString("peach puncha patch", 2))

	//===[]byte数组匹配===
	byteArray := []byte{'p', 'e', 'a', 'c', 'h', 'g', 'i'}
	fmt.Println(constPattern.Match(byteArray))

	//===字符串替换===
	fmt.Println(pattern.ReplaceAllString("a peach", "banana"))
}
