package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func bufferExample() {
	var buf bytes.Buffer
	buf.WriteString("Hello, Buffer!")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello")
	r2 := strings.NewReader(", World!")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

func pipeExample() {
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		_, err := pw.Write([]byte("Hello, Pipe!"))
		if err != nil {
			log.Fatal(err)
		}
	}()
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(pr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

func main() {
	fmt.Println("*** Read from Reader ***")
	readFromReader(strings.NewReader("Hello, Reader!"))
	fmt.Println("\n*** Write to Writer ***")
	writeToWriter(&bytes.Buffer{}, "Hello, Writer!")
	fmt.Println("\n*** Buffer Example ***")
	bufferExample()
	fmt.Println("\n*** MultiReader Example ***")
	multiReaderExample()
	fmt.Println("\n*** Close Resource Example ***")

}
