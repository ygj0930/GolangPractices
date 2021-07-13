package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	p := fmt.Println

	//构建可移植的文件路径(可跨操作系统)
	//使用Join来代替手动拼接/或\，因为其自动适配不同操作系统的分隔符
	pathStr := filepath.Join("dir1", "dir2", "dir3", "textFile.txt")
	p(pathStr)

	//绝对路径与相对路径判断
	p(filepath.IsAbs("dir/file"))
	p(filepath.IsAbs("/dir/file"))

	//提取路径内容
	p("Dir(path):", filepath.Dir(pathStr))   //从路径中提取目录信息
	p("Base(path):", filepath.Base(pathStr)) //从路径中提取文件信息
	p("Ext(path):", filepath.Ext(pathStr))   //从路径中提取文件扩展名
}
