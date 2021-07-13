package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	p := fmt.Println

	//伪随机数：以下代码每次执行得到的随机数都是一样的，因为默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。
	//范围随机整数
	p(rand.Intn(99))
	p(rand.Intn(99))
	p(rand.Intn(99))
	p()

	//0~1随机小数
	p(rand.Float64())
	p()

	//范围随机小数：自行加工
	p(rand.Float64() * 5) //0~5范围小数
	p()

	//真随机数：每次生成新的种子
	//非加密随机数：使用系统时间作为种子
	timeSource := rand.NewSource(time.Now().UnixNano())
	timeRand := rand.New(timeSource)

	//使用时间随机数生成器进行随机数生成
	p(timeRand.Intn(99))
	p(timeRand.Intn(99))
	p(timeRand.Intn(99))
	p()
}
