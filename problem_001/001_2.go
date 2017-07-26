package main

import "fmt"

var maxNumber = 999

// Problem 1
// Multiples of 3 and 5
//
// Sn = n*(a+an)/2, where: n - number of elements, a - first term, an - n-th element
// an = a+(n-1)*d, where d - common difference, n - number of elements
func main() {
	// calculating sum of numbers that are multiples of 3
	sum3 := sum(3)
	// calculating sum of numbers that are multiples of 5
	sum5 := sum(5)
	// calculating sum of numbers that are multiples of 3 AND 5
	sum15 := sum(15)

	// numbers from sum15 have been already included in sum3 and sum5
	sum := sum3 + sum5 - sum15
	fmt.Println(sum)
}

// Get a sum of all the multiples of 'm'
// First element and common difference == 'm'
func sum(m int) int {
	n := int(maxNumber / m)
	sum := (2*m + m*(n-1)) * n / 2

	return sum
}
