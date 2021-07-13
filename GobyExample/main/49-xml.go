package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`   //xml根标签
	Id      int      `xml:"id,attr"` //xml标签属性

	Name   string   `xml:"name"`   //子标签
	Origin []string `xml:"origin"` //子标签
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v,name=%v,origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	//结构体转xml字符串
	coffee := Plant{Id: 27, Name: "Coffee", Origin: []string{"Brazil", "China"}}
	out, _ := xml.MarshalIndent(coffee, " ", " ")
	fmt.Println(string(out))
	fmt.Println(xml.Header + string(out)) //增加xml文档头

	//xml字符串转结构体
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)
}
