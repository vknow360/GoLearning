package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Parsed integer:", num)

	numistr, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Parsed integer:", numistr)

	floatStr := "3.14"
	floatVal, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Parsed float: %.3f\n", floatVal)

	binaryStr := "101010101"
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Parsed decimal:", decimal)
}
