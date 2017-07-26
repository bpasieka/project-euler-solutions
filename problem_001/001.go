package main

import "fmt"

// Problem 1
// Multiples of 3 and 5
func main() {
	sum := 0

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
