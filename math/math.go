package math

import (
	"math"
	"strconv"
)

// IsPalindrome returns true if n is a palindrome and false if it's not
func IsPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// IsPrime returns true if n is prime and false if it's not
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	for i := 3; i < int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// NthDigit returns the nth digit of x, counting from right to left
func NthDigit(n, x int) int {
	i := int(math.Pow(10, float64(n-1)))
	return x / i % 10
}

// SetNthDigit sets the nth digit of x to newVal, counting from right to left
func SetNthDigit(n, newVal int, x *int) {
	cur := NthDigit(n, *x)
	diff := (newVal - cur) * int(math.Pow(10, float64(n-1)))
	*x = *x + diff
	return
}
