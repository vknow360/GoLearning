package main

import "fmt"

func main() {
	var ptr *int //nil

	var a int = 10
	ptr = &a

	fmt.Println(a)
	modifyValue(ptr)
	fmt.Println(*ptr) //dereferencing
}

func modifyValue(ptr *int) {
	*ptr++
}
