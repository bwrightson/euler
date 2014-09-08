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
	if n%2 == 0 {
		return false
	}
	for i := 3; i < int(math.Sqrt(float64(n)))+1; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// LCM returns the lowest common multiple of slice s
func LCM(s []int) int {
	lcm := 1
	primes := make([][]int, 0)
	for _, i := range s {
		if i == 1 {
			continue
		}
		temp := PrimeFactorization(i)
		for _, t := range temp {
			found := false
			for _, v := range primes {
				if t[0] == v[0] {
					found = true
					if t[1] > v[1] {
						v[1] = t[1]
					}
				}
			}
			if !found {
				primes = append(primes, t)
			}
		}
	}
	for _, v := range primes {
		lcm *= int(math.Pow(float64(v[0]), float64(v[1])))
	}
	return lcm
}

// NthDigit returns the nth digit of x, counting from right to left
func NthDigit(n, x int) int {
	i := int(math.Pow(10, float64(n-1)))
	return x / i % 10
}

// NthPrime returns the nth prime (starting at 2) using a trial division method
func NthPrime(n int) int {
	prime := 2
	for i, count := 3, 1; count < n; i = i + 2 {
		if IsPrime(i) {
			prime = i
			count = count + 1
		}
	}
	return prime
}

// NthPrimeWithSieve returns the nth prime (starting at 2) using a simple
// version of the Sieve of Eratosthenes
func NthPrimeWithSieve(n int) int {
	if n == 1 {
		return 2
	}
	multiples := make([][]int, 0)
	init := []int{2, 4}
	multiples = append(multiples, init)
	prime := 2
	for i, count, check := 3, 1, true; count < n; i = i + 1 {
		check = true
		for _, j := range multiples {
			if j[1] == i {
				j[1] = j[1] + j[0]
				check = false
			}
		}
		if check {
			prime = i
			count = count + 1
			inc := []int{i, i + i}
			multiples = append(multiples, inc)
		}
	}
	return prime
}

// PrimeFactorization returns a slice of int slices that represents
// the prime factorization of n
func PrimeFactorization(n int) [][]int {
	temp := n
	found := false
	primes := make([][]int, 0)
OuterLoop:
	for i := 2; ; {
		found = false
		if IsPrime(temp) {
			for _, v := range primes {
				if v[0] == temp {
					v[1] += 1
					break OuterLoop
				}
			}
			if !found {
				val := []int{temp, 1}
				primes = append(primes, val)
				break OuterLoop
			}
		}
		if IsPrime(i) {
			if temp%i == 0 {
				temp = temp / i
				for _, v := range primes {
					if v[0] == i {
						v[1] += 1
						found = true
						break
					}
				}
				if !found {
					val := []int{i, 1}
					primes = append(primes, val)
				}
				continue
			}
		}
		i += 1
	}
	return primes
}

// SetNthDigit sets the nth digit of x to newVal, counting from right to left
func SetNthDigit(n, newVal int, x *int) {
	cur := NthDigit(n, *x)
	diff := (newVal - cur) * int(math.Pow(10, float64(n-1)))
	*x = *x + diff
	return
}
