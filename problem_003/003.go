package main

import (
	"fmt"
	"math"
)

var number = 600851475143

// Problem 3
// Largest prime factor
func main() {
	result := 1

	root := int(math.Sqrt(float64(number)))
	for i := 2; i <= root; i++ {
		// not a factor
		if number%i != 0 {
			continue
		}

		potentialResult := number / i
		if isPrime(potentialResult) {
			result = potentialResult
			// can not get larger factor
			break
		}

		if isPrime(i) {
			result = i
			// still can get larger factor
		}
	}

	fmt.Println(result)
}

// Check if number is prime
func isPrime(number int) bool {
	root := int(math.Sqrt(float64(number)))
	for i := 2; i <= root; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
