package main

import (
	"crypto/md5"
	"fmt"
)

//技巧：提炼方便打印
func p(data interface{}) {
	fmt.Println(data)
}

func main() {
	//散列的作用：用于签名，作为版本标识

	//待散列源内容
	s := "origin string"

	//使用go相关散列实现创建散列器
	//hashing := sha1.New()
	hashing := md5.New()

	//进行签名
	hashing.Write([]byte(s))
	res := hashing.Sum(nil)
	//转成16进制字符串表示
	shaStr := fmt.Sprintf("%x", res)

	p(s)
	p(res)
	p(shaStr)

}
