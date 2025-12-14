package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("Hello, bufio package!How are you doing?\n"))

	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(line)

	writer := bufio.NewWriter(os.Stdout)
	data = []byte("Hello, bufio package\n")
	n, err = writer.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)
	writer.Flush()

	str := "This is a string\n"
	writer.WriteString(str)
	writer.Flush()

}
