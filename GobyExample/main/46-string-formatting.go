package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	//=======格式化输出==========
	//结构体
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	//fmt.Printf("%+v\n", p)
	//fmt.Printf("%#v\n", p)
	//fmt.Printf("%T\n", p)

	//////基本数据类型
	////布尔类型
	//fmt.Printf("%t\n", false)
	//
	////整形
	//fmt.Printf("%d\n", 123)
	//fmt.Printf("%b\n", 123)
	//fmt.Printf("%x\n", 123)
	//fmt.Printf("%c\n", 123)
	//fmt.Printf("|%3d|%6d|\n", 12, 345)

	//
	////浮点型
	//fmt.Printf("%f\n", 78.9)
	//fmt.Printf("%e\n", 12340000000000.0)
	//fmt.Printf("%E\n", 12340000000000.0)
	//fmt.Printf("%.2f\n", 78.9)
	//fmt.Printf("|%6.2f|%6.2f\n", 1.214, 3.45142)  //默认右对齐
	//fmt.Printf("|%-6.2f|%-6.2f\n", 1.243, 3.4514) //左对齐

	//
	////字符串
	//fmt.Printf("%s\n", "\"str\"")
	//fmt.Printf("%q\n", "\"str\"")
	//fmt.Printf("%x\n", "\"str\"")
	//fmt.Printf("|%6s|%6s|\n", "foo", "bar")
	//fmt.Printf("|%-6s|%-6s|\n", "foo", "bar")

	//指针类型值
	//fmt.Printf("%p\n", &p)

	//=======格式化并返回字符串==========
	s := fmt.Sprintf("a string:%s\n", "字符串内容")
	fmt.Println(s)

	//=======格式化并输出到io流==========
	fmt.Fprintf(os.Stderr, "an error str:%s\n", "有异常")
}
