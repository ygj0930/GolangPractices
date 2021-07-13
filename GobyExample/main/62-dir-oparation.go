package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func dirErrCheck(e error) {
	if e != nil {
		panic(e)
	}
}

//创建空文件
func createEmptyFile(path string) {
	initData := []byte("")
	dirErrCheck(ioutil.WriteFile(path, initData, 0644))
}

//判断文件是否存在
func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	//判断Stat返回的err是否为os抛出的NotExitError
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func main() {
	//创建目录
	dirStr1 := "dir1"
	if pathExists(dirStr1) {
		os.RemoveAll(dirStr1)
	}
	mErr := os.Mkdir(dirStr1, 0755)
	dirErrCheck(mErr)

	//目录下创建文件
	fileStr1 := "dir1/file1.txt"
	if pathExists(fileStr1) {
		os.RemoveAll(fileStr1)
	}
	createEmptyFile(fileStr1)

	//创建嵌套目录
	dirPath := filepath.Join("dir1", "dir2", "dir3")
	if pathExists(dirPath) {
		os.RemoveAll(dirPath)
	}
	aErr := os.MkdirAll(dirPath, 0755)
	dirErrCheck(aErr)
}
