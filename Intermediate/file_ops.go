package main

import (
	"fmt"
	"os"
)

func main() {
	//write()
	read()
}

func write() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// write data to file
	data := []byte("Hello world!\n")
	file.Write(data)

	file1, err := os.Create("writeString.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file1.Close()
	file1.WriteString("Hello Go!\n")
}

func read() {
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		file.Close()
	}()

	data := make([]byte, 1024)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
