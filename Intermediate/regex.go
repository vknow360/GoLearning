package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("He said \"I am great\"")

	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	email := "user@email.com"
	email2 := "invalid_email"

	fmt.Println("Email:", re.MatchString(email))
	fmt.Println("Email:", re.MatchString(email2))
}
