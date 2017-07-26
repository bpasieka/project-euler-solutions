package main

import "fmt"

// Problem 2
// Even Fibonacci numbers
func main() {
	sum := 0
	previous := 1
	current := 2

	for current < 4000000 {
		if current%2 == 0 {
			sum += current
		}

		previous, current = current, current+previous
	}

	fmt.Println(sum)
}
