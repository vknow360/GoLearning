// buffered channels allow asynchrounous communication
// only block when buffer is full

package main

import (
	"fmt"
	"time"
)

// Blocking on send only if buffer is full
// func main() {
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("Received:", <-ch)
// 	}()
// 	ch <- 3
// 	fmt.Println("Received:", <-ch)
// 	fmt.Println("Received:", <-ch)
// 	fmt.Println("Buffered Channels")
// }

// Blocking on send only if receiver is empty
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Receiving from buffer")
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch)
	}()
	fmt.Println("Blocking starts")
	ch <- 3
	// fmt.Println("Received:", <-ch)
	// fmt.Println("Received:", <-ch)
}
