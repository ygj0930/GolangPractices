package main

import "fmt"

func main() {
	s:=make([]string,3)
	fmt.Println("init:",s)

	s[0]="a"
	s[2]="2"
	fmt.Println(s)

	fmt.Println("len:",len(s))

	s = append(s,"d")
	s = append(s,"e","f","g")
	fmt.Println(s)

	c:=make([]string,len(s))
	copy(c,s)
	fmt.Println("cpy",c)

	l:=s[2:4]
	fmt.Println(l)

	h:=s[:3]
	fmt.Println(h)

	ini:=[]string{"a","b","c"}
	fmt.Println(ini)

	twoD:=make([][]int,3)
	for i:=0;i<=2;i++{
		twoD[i] = make([]int,i+1)
		for j:=0;j<=9;j++{
			twoD[i] = append(twoD[i],j)
		}
	}
	fmt.Println(twoD)
}
