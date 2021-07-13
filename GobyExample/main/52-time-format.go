package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	t := time.Now()

	//按内置的常量模式格式化：time.RFCXXX
	p(t.Format(time.RFC822))
	p(t.Format(time.RFC3339))

	//自定义模式进行格式化
	//自定义模式的符号：按Mon Jan 2 15:04:05 MST 2006格式
	//自定义年月日输出
	p(t.Format("2006-01-02"))
	//时分秒输出
	p(t.Format("15:04:05"))

	//按来源模式将时间字符串转回时间类型值
	timeStrPattern := "2006/01/02 15:04:05"
	t2, err := time.Parse(timeStrPattern, "2018/07/01 09:59:59")
	if err != nil {
		panic(err)
	}
	p(t2)
}
