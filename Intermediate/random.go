package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// seed: the starting point for generating random numbers
	// fmt.Println(rand.Intn(100) + 1)

	// val := rand.New(rand.NewSource(time.Now().Unix()))
	// fmt.Println(val.Intn(101))

	for {
		fmt.Println("Welcome to the Dice Game!")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice")
			continue
		}

		if choice == 2 {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}

		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		fmt.Printf("You rolled a %d and a %d\n", die1, die2)
		fmt.Println("Total: ", die1+die2)

		fmt.Print("Do you want to roll again? (y/n): ")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("Invalid input, assuming no.")
			break
		}

		if rollAgain == "n" {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}
	}

}
