package main

import "fmt"

const (
	maxNumber int = 100
)

// Problem 6
// Sum square difference
func main() {
	sum := 0
	sumOfSquares := 0

	for i := 1; i <= maxNumber; i++ {
		sum += i
		sumOfSquares += i * i
	}

	fmt.Println(sum*sum - sumOfSquares)
}
