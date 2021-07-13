package main

import (
	"encoding/json"
	"fmt"
)

type Object1 struct {
	Code    int
	Content string
}

type Object2 struct {
	Code    int    `json:"code"`
	Content string `json:"content"`
}

func main() {

	//基本数据类型的json编码
	bolToJson, _ := json.Marshal(true)
	fmt.Println("bolToJson byte:", bolToJson)        //Marshal返回的是[]byte数组
	fmt.Println("bolToJson str:", string(bolToJson)) //需要自行转换为字符串

	intToJson, _ := json.Marshal(123)
	fmt.Println("intToJson:", string(intToJson))

	floatToJson, _ := json.Marshal(123.4)
	fmt.Println("floatToJson:", string(floatToJson))

	strToJson, _ := json.Marshal("string content")
	fmt.Println("strToJson:", string(strToJson))

	//切片、map等容器的json编码
	arr := []string{"apple", "peach", "pear"}
	sliceToJson, _ := json.Marshal(arr)
	fmt.Println("sliceToJson:", string(sliceToJson))

	stringMap := make(map[string]string)
	stringMap["key1"] = "value1"
	stringMap["key2"] = "value2"
	mapToJson, _ := json.Marshal(stringMap)
	fmt.Println("mapToJson:", string(mapToJson))

	//结构体对象的json编码
	obj1 := Object1{
		Code:    200,
		Content: "Success"}
	obj1ToJson, _ := json.Marshal(obj1)
	fmt.Println("obj1ToJson:", string(obj1ToJson))

	obj2 := Object2{
		Code:    404,
		Content: "Failed"}
	obj2ToJson, _ := json.Marshal(obj2)
	fmt.Println("obj2ToJson:", string(obj2ToJson))

	//json字符串解码：解码为map对象
	jsonStr1 := `{"Code":200,"data":["data1","data2"]}`
	//map容器，用于保存解码结果：值数据类型解码为interface{}，访问时再做类型转换
	var jsonMap map[string]interface{}
	//进行转换，把map地址传进去，接收转换结果
	if err := json.Unmarshal([]byte(jsonStr1), &jsonMap); err != nil {
		panic(err)
	}
	fmt.Println("jsonToMap:", jsonMap)
	//访问转换结果字段：需要做类型转换（弊端，不方便）
	var code float64 = jsonMap["Code"].(float64)
	fmt.Println("jsonMap.Code:", code)

	var datas []interface{} = jsonMap["data"].([]interface{})
	fmt.Println("jsonMap.data:", datas)
	var data1 string = datas[0].(string)
	fmt.Println("jsonMap.data[0]:", data1)

	//json字符串解码：解码为自定义结构体对象
	//好处：访问内容时，无需做类型转换
	jsonStr2 := `{"code":404,"content":"some content"}`
	var jsonToObj1 Object1
	if err := json.Unmarshal([]byte(jsonStr2), &jsonToObj1); err != nil {
		panic(err)
	}
	fmt.Println("jsonToObj1:", jsonToObj1)
	fmt.Println("jsonToObj1.Code:", jsonToObj1.Code)

	var jsonToObj2 Object2
	if err := json.Unmarshal([]byte(jsonStr2), &jsonToObj2); err != nil {
		panic(err)
	}
	fmt.Println("jsonToObj2:", jsonToObj2)
}
