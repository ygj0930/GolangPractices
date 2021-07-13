package main

import (
	"fmt"
	"time"
)

func main() {
	//窍门：简化fmt.Println调用
	p := fmt.Println

	//当前时间输出
	now := time.Now()
	p(now)

	//根据年月日自行构建一个time
	customTime := time.Date(2021, 07, 07, 11, 11, 58, 23414214, time.UTC)
	p(customTime)

	//提取某个time的年月日、时分秒、毫秒、时区
	p(customTime.Year())
	p(customTime.Month())
	p(customTime.Day())

	p(customTime.Hour())
	p(customTime.Minute())
	p(customTime.Second())
	p(customTime.Nanosecond())

	p(now.Location())

	//提取某个time的周几信息
	p(now.Weekday())

	//时间比较：先、后、同时
	p(customTime.Before(now))
	p(customTime.After(now))
	p(customTime.Equal(now))

	//时间间隔计算
	delta := now.Sub(customTime)
	p(delta)
	p(delta.Hours())
	p(delta.Minutes())
	p(delta.Seconds())

	//时间操作：增加时间间隔
	p(customTime.Add(time.Duration(time.Hour * 24 * 2)))
	p(customTime.AddDate(1, 1, 2))
}
