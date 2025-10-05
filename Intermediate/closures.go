package main

import "fmt"

func main() {
	sequence := adder()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	//op: 5

	substractor := func() func(int) int {

		countdown := 99

		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(substractor(1))
	fmt.Println(substractor(2))
	fmt.Println(substractor(3))
	fmt.Println(substractor(4))
	//op: 89
}

func adder() func() int {
	i := 0
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("Added 1 to i")
		return i
	}
}
