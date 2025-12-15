package main

import "os"

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.CreateTemp("", "temporaryFile")
	checkError(err)

	println("Temp file created:", file.Name())
	defer os.Remove(file.Name())
	defer file.Close()
}
