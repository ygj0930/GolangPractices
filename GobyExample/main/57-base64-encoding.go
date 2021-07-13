package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	p := fmt.Println

	//标准base64编解码
	dataStr := "abc123!?$*&()'-=@~"
	stdBase64 := base64.StdEncoding.EncodeToString([]byte(dataStr))
	p(stdBase64)

	decoStr, _ := base64.StdEncoding.DecodeString(stdBase64)
	p(string(decoStr))

	//URL兼容的base64编解码
	urlData := "abc123!?$*&()'-=@~"
	urlBase64 := base64.URLEncoding.EncodeToString([]byte(urlData))
	p(urlBase64)

	decoUrl, _ := base64.URLEncoding.DecodeString(urlBase64)
	p(string(decoUrl))
}
