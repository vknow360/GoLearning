package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "Hello, Go!"

	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(rawMessage)

	fmt.Println("Length of message var is", len(message))
	fmt.Println("The first char in message var is", message[0]) //ASCII

	greeting := "Hello "
	name := "Alice!"

	fmt.Println(greeting + name)

	str1 := "Apple"
	str2 := "banana"
	fmt.Println("Are equal:", str1 == str2)

	for i, char := range message {
		fmt.Printf("%d %c\n", i, char)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(message))

	var ch rune = 'a'

	cstr := string(ch)
	fmt.Printf("%T\n", ch)
}
