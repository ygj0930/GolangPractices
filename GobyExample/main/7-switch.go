package main

import (
	"fmt"
)

func main() {
	whatAmI(1)
	whatAmI(float32(1.00))
	whatAmI(1+2i)
	whatAmI(false)
	whatAmI("str")

	//i:=1
	//
	//switch i {
	//case 1:
	//	fmt.Println("1")
	//case 2:
	//	fmt.Println("2")
	//case 3:
	//	fmt.Println("3")
	//default:
	//	fmt.Println("others")
	//}
	//
	//switch time.Now().Weekday() {
	//case time.Saturday,time.Sunday:
	//	fmt.Println("It's weekend")
	//default:
	//	fmt.Println("It's weekday")
	//}
	//
	//t:=time.Now()
	//switch {
	//case t.Hour()<12:
	//	fmt.Println("before noon")
	//	fallthrough
	//case t.Weekday()==time.Sunday:
	//	fmt.Println(" weekend")
	//default:
	//	fmt.Println("no weekend no before noon")
	//}

}

func whatAmI(i interface{}){
	switch t:=i.(type) {
	case bool:
		fmt.Println("is bool")
	case int:
		fmt.Println("is int")
	case string:
		fmt.Println("is string")
	default:
		fmt.Printf("%T\n",t)
	}
}