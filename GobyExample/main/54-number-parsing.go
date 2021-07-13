package main

import (
	"fmt"
	"strconv"
)

func main() {
	p := fmt.Println

	//数字转字符串
	//高频使用：int转字符串
	p(strconv.Itoa(123))
	//其他整型转字符串
	p(strconv.FormatInt(123, 10))
	p(strconv.FormatUint(123, 10))

	//浮点数转字符串
	// f：要转换的浮点数
	//fmt：格式标记（b、e、E、f、g、G), 一般用‘f’ (-ddd.dddd，没有指数)
	//prec：精度，小数位数
	// bitSize：指定浮点类型（32:float32、64:float64）
	p(strconv.FormatFloat(123.43578762124, 'f', 5, 64))

	//字符串转数字
	//高频使用：int整数字符串转整数
	intParse, _ := strconv.Atoi("123")
	p(intParse)
	//其他整型转换
	int64Parse1, _ := strconv.ParseInt("1234", 10, 64)
	p(int64Parse1)
	//参数0：表示自动根据字符串表示的数字来推断进制
	int64Parse2, _ := strconv.ParseInt("1234", 0, 64)
	p(int64Parse2)

	uintParse, _ := strconv.ParseUint("789", 0, 64)
	p(uintParse)

	//浮点数字符串解析得浮点数
	flt, fe := strconv.ParseFloat("1.223445", 64)
	if fe != nil {
		panic(fe)
	}
	p(flt)
}
