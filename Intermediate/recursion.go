package main

import "fmt"

func main() {
	num := 5
	fmt.Println("Factorial of", num, "is:", factorial(num))

	num = 12345
	fmt.Println("Sum:", sumOfDigits(num))
}

func factorial(n int) int {
	if n < 2 {
		return n
	}
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}
