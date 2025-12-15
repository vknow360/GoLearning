package main

import (
	"embed"
)

//go:embed example.txt

var content string

//go:embed basics
var basicsFolder embed.FS

func main() {

	println("Embedded content:", content)
	content, err := basicsFolder.ReadFile("basics/hello.txt")
	if err != nil {
		panic(err)
	}
	println(string(content))
}
