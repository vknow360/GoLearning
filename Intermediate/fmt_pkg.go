package main

import "fmt"

func main() {

	fmt.Print("Hello ")
	fmt.Print("World!")
	fmt.Print(12.456)

	fmt.Println("Hello ")
	fmt.Println("World!")
	fmt.Println(12.456)

	name := "John"
	age := 25
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	s := fmt.Sprint("Hello ", "World!", 123, 456)
	fmt.Print(s)
	s = fmt.Sprintln("Hello", "World!", 123, 456)
	fmt.Println(s)

	sf := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(sf)

	var nm string
	var ag int

	fmt.Print("Enter your name and age:")
	fmt.Scan(&nm, &ag)
	fmt.Printf("Name: %s, Age: %d\n", nm, ag)

	var ln string
	fmt.Scanln(&ln)
	fmt.Println(ln)
}
