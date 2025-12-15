package main

import (
	"flag"
)

func main() {
	// println("Command:", os.Args[0])
	// println("Argument1:", os.Args[1])

	// Define flags
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "John", "Name of the user")
	flag.IntVar(&age, "age", 18, "Age of the user")
	flag.BoolVar(&male, "male", true, "Gender of the user")

	flag.Parse()

	println(name, age, male)
}
