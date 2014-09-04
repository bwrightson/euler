package math

import (
	"math"
)

// Prime takes int n and returns true if it's prime and false if it's not
func Prime(n int) bool {
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
