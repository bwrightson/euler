package problems

import (
	"github.com/bwrightson/euler/math"
)

func Solve001() int {
	sum := 0
	for i := 3; i < 1000; i += 3 {
		sum += i
	}
	for i := 5; i < 1000; i += 5 {
		if i%3 != 0 {
			sum += i
		}
	}
	return sum
}

func Solve002() int {
	sum := 0
	for val, prev := 2, 1; val < 4000000; {
		sum += val
		val, prev = 3*val+2*prev, 2*val+prev
	}
	return sum
}

func Solve003() int {
	val := 600851475143
	factor := 0
	for i := 2; ; {
		if val%i == 0 && math.Prime(i) {
			val = val / i
			factor = i
			continue
		}
		if val == 1 {
			break
		}
		if i > 2 {
			i += 2
		} else {
			i += 1
		}
	}
	return factor
}
