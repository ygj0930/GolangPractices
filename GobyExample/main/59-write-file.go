package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"time"
)

//小技巧：精简错误检查
func writeFileErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	p := fmt.Println

	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05") + "\r"

	//===操作File进行写入===
	//1、打开文件，进行写入：需要使用OpenFile而不是Open函数
	//具体区别：https://blog.csdn.net/benben_2015/article/details/80607425
	f, err := os.OpenFile("defer.txt", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	writeFileErrorCheck(err)
	defer f.Close()

	fStr := "Step 1:wrote by file @" + timeStr
	fwriteSize, ferr := f.Write([]byte(fStr))
	//fwriteSize, err := f.WriteString(fStr) //两种方式都可以
	writeFileErrorCheck(ferr)
	p("File Write size:", fwriteSize)
	f.Sync()
	time.Sleep(time.Second * 2)

	//===使用bufio包写文件===
	bStr := "Step 2:wrote by bufio @" + timeStr
	bwriter := bufio.NewWriter(f)
	bwriteSize, berr := bwriter.WriteString(bStr)
	writeFileErrorCheck(berr)
	p("Bufio Write size:", bwriteSize)
	bwriter.Flush()
	time.Sleep(time.Second * 2)

	//===使用ioutil包写文件===
	uStr := "Step 3:wrote by ioutil @" + timeStr
	content := []byte(uStr)
	//实践发现：ioutil.WriteFile会清空源文件内容再写入，即使采用ModeAppend模式也不行
	//要进行追加写入，还是要采取上面OpenFile再写入的方式才行
	err = ioutil.WriteFile("defer.txt", content, fs.ModeAppend)
	writeFileErrorCheck(err)
}
