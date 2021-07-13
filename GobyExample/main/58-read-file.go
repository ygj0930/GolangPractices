package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

//小技巧：精简错误检查
func readFileErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	p := fmt.Println

	//1、直接读取文件所有内容
	fileData, err := ioutil.ReadFile("defer.txt")
	readFileErrorCheck(err)
	p("Whole file:", string(fileData))

	//2、进行精细的读取
	//文件打开
	f, err := os.Open("defer.txt")
	defer f.Close()
	readFileErrorCheck(err)

	//进行个性化读取
	//按字节读取：每次一定数量的字节
	slot := make([]byte, 18)
	readSize, err := f.Read(slot)
	readFileErrorCheck(err)
	p("Slot read,size:", strconv.Itoa(readSize), "content:", string(slot))

	//移到某位置再读取
	beginPos, err := f.Seek(18, 0)
	readFileErrorCheck(err)
	slot2 := make([]byte, 5)
	readSize2, err := f.Read(slot2)
	readFileErrorCheck(err)
	p("Seek read,Size:", strconv.Itoa(readSize2), ",begin position:", beginPos, ",content:", string(slot2))

	//文件读取指针重置：回到第一位
	_, err = f.Seek(0, 0)
	readFileErrorCheck(err)

	//3、使用io包、bufio包进行文件读取
	slot3 := make([]byte, 5)
	readSize3, err := io.ReadAtLeast(f, slot3, 5)
	readFileErrorCheck(err)
	p("Io read,size:", strconv.Itoa(readSize3), "content:", string(slot3))

	bufRead := bufio.NewReader(f)
	bufContent, err := bufRead.Peek(5)
	readFileErrorCheck(err)
	p("BufIO read:", "content:", string(bufContent))
}
