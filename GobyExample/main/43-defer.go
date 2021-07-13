package main

import (
	"fmt"
	"os"
)

func createFile(path string) *os.File{
	fmt.Println("creat file...")

	f,err := os.Create(path)
	if err!=nil{
		panic(err)
	}
	return f
}

func writeFile(f *os.File){
	fmt.Println("writing file...")
	fmt.Fprintln(f,"write some data")
}

func closeFile(f *os.File){
	fmt.Println("close file...")
	f.Close()
}

func main() {
	f := createFile("defer.txt")
	defer closeFile(f)

	writeFile(f)
}
