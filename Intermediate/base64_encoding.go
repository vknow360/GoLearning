package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Hello, base64 encoding"

	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(decoded))

	url := "q=what is the value of 1 + 2"
	urlSafeEncoding := base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println(urlSafeEncoding)
	urlDecoded, err := base64.URLEncoding.DecodeString(urlSafeEncoding)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(urlDecoded))
}
