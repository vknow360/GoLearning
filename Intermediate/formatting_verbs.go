package main

import "fmt"

func main() {
	name := "Sunny"
	age := 20
	isDev := true
	score := 95
	pi := 3.1415926535
	scores := []int{90, 85, 88}

	// --- General formatting verbs
	// %v - Prints value in default format
	fmt.Printf("Name: %v, Age: %v, Developer: %v\n", name, age, isDev)

	// %#v - Prints value in Go-syntax format (useful for debugging)
	fmt.Printf("Go-syntax representation of scores: %#v\n", scores)

	// %T - Prints the type of value
	fmt.Printf("Type of 'scores': %T\n", scores)

	// %% - Prints a literal percent sign
	fmt.Printf("Success rate: %v%%\n", 95)

	// You can also mix these
	fmt.Printf("User: %#v (%T)\n\n", name, name)

	// --- Integer formatting verbs
	fmt.Printf("Decimal: %d\n", score) // Base 10
	fmt.Printf("Binary: %b\n", score)  // Base 2
	fmt.Printf("Octal: %o\n", score)   // Base 8
	fmt.Printf("Hexadecimal (lowercase): %x\n", score)
	fmt.Printf("Hexadecimal (uppercase): %X\n", score)
	fmt.Printf("With sign: %+d\n\n", age)

	// --- String formatting verbs
	fmt.Printf("Normal string: %s\n", name)
	fmt.Printf("Quoted string: %q\n", name)
	fmt.Printf("Width 10 (right align): %10s\n", name)
	fmt.Printf("Width 10 (left align): %-10s<--\n\n", name)

	// --- Float formatting verbs
	fmt.Printf("Default: %v\n", pi)
	fmt.Printf("Fixed-point (2 decimals): %.2f\n", pi)
	fmt.Printf("Scientific (lowercase): %e\n", pi)
	fmt.Printf("Scientific (uppercase): %E\n", pi)
	fmt.Printf("Compact (auto choose): %g\n\n", pi)

	// --- Boolean formatting verbs
	fmt.Printf("isDev: %t\n", isDev)
	fmt.Printf("Negated: %t\n", !isDev)
}
