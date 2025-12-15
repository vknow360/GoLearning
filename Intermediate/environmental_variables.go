package main

import (
	"os"
	"strings"
)

func main() {
	user := os.Getenv("user")
	home := os.Getenv("HOME")

	println("User:", user)
	println("Home Directory:", home)

	for _, e := range os.Environ() {
		kvpair := strings.SplitN(e, "=", 2)
		println(kvpair[0], ":", kvpair[1])
	}

}
