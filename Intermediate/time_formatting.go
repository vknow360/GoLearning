package main

import (
	"fmt"
	"time"
)

func main() {
	// Mon Jan 2 15:04:05 MST 2006
	layout := "2006-01-02T15:04:05Z07:00"

	str := "2024-06-15T14:30:00+02:00"

	// Parse the string into a Time object
	t, err := time.Parse(layout, str)
	if err != nil {
		panic(err)
	}

	fmt.Println(t)

	str1 := "Jul 20, 2024 03:45:00 PM"
	layout1 := "Jan 2, 2006 03:04:05 PM"

	t1, err := time.Parse(layout1, str1)
	if err != nil {
		panic(err)
	}
	fmt.Println(t1)
}
