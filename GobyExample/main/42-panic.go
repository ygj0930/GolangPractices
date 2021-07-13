package main

import "os"

func main() {
	//panic("a problem")

	_,err := os.Open("/temp/file")
	if err!=nil{
		panic(err)
	}
}
