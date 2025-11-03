package main

import (
	"errors"
	"fmt"
	"math"
)

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

func SquareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("cannot take the square root of a negative number")
	}
	return math.Sqrt(x), nil
}

func main() {
	// err := &CustomError{
	// 	Code:    404,
	// 	Message: "Not Found",
	// }
	// fmt.Println(err)

	result, err := SquareRoot(-4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root:", result)
	}
}
