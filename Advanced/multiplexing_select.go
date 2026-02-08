package main

import "time"

func main() {
	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	ch1 <- 1
	// }()
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	ch2 <- 2
	// }()

	// select {
	// case msg := <-ch1:
	// 	println("Received from channel 1:", msg)
	// case msg := <-ch2:
	// 	println("Received from channel 2:", msg)
	// 	// default:
	// 	// 	println("No messages received, default case executed")
	// }

	ch := make(chan int)
	go func() {
		//time.Sleep(2 * time.Second)
		ch <- 42
		close(ch)
	}()

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				println("Channel closed")
				return
			}
			println("Received:", msg)
		case <-time.After(3 * time.Second):
			println("Timeout: No message received within 3 seconds")
		}
	}
}
