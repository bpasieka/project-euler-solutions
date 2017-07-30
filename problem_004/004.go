package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Problem 4
// Largest palindrome product
func main() {
	max := 0
	start, end := 100, 999

	for i := start; i <= end; i++ {
		for j := i; j <= end; j++ {
			if n := i * j; isPalindromic(n) && n > max {
				max = n
			}
		}
	}

	fmt.Println(max)
}

func isPalindromic(n int) bool {
	data := strings.Split(strconv.Itoa(n), "")
	dataLen := len(data)

	if dataLen == 1 {
		return true
	}

	partLen := dataLen / 2

	for i := 1; i <= partLen; i++ {
		if data[i-1] != data[dataLen-i] {
			return false
		}
	}

	return true
}
