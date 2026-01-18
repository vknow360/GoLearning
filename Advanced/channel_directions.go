package main

import "fmt"

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}

func producer(ch chan<- int) {
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println(value)
	}
}
