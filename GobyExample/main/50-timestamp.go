package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p("now:", now)

	//获取时间的Unix纪元表示：时间戳
	secs := now.Unix() //秒级时间戳
	p("unix:", secs)
	nanos := now.UnixNano() //纳秒时间戳
	p("nanos:", nanos)

	//时间戳转时间
	p(time.Unix(secs, 0))
	p(time.Unix(0, nanos))
	p(time.Unix(secs, nanos))
}
