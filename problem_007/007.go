package main

import (
	"fmt"
	"math"
)

const (
	numberPosition = 10001
)

// Problem 7
// 10001st prime
func main() {
	result := 0
	count := 1

	for i := 3; count != numberPosition; i += 2 {
		if isPrime(i) {
			result = i
			count++
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
