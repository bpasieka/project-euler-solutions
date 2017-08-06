package main

import "fmt"

const (
	minNumber int = 1
	maxNumber int = 20
)

// Problem 5
// Smallest multiple
func main() {
	result := 0

	for i := maxNumber; ; i += maxNumber {
		if isDivisible(i) {
			result = i
			break
		}
	}

	fmt.Println(result)
}

func isDivisible(n int) bool {
	for i := maxNumber - 1; i > minNumber; i-- {
		if n%i != 0 {
			return false
		}
	}

	return true
}
