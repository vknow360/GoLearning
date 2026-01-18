// communicate between goroutines using channels.

package main

import (
	"fmt"
)

func main() {
	// var := make(chan type)

	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
		greeting <- "World"

		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}
	}()

	receiver := <-greeting
	fmt.Println(receiver)
	receiver = <-greeting
	fmt.Println(receiver)

	for _ = range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}

	//time.Sleep(1 * time.Second)
}
