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
		if val%i == 0 && math.IsPrime(i) {
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

func Solve004() int {
	for pal := 997799; ; {
		for i := 999; ; i-- {
			if pal%i == 0 {
				return pal
			}
			if pal/i > i {
				break
			}
		}
		j := math.NthDigit(3, pal)
		if j > 0 {
			math.SetNthDigit(3, j-1, &pal)
			math.SetNthDigit(4, j-1, &pal)
		} else {
			k := math.NthDigit(2, pal)
			if k > 0 {
				math.SetNthDigit(2, k-1, &pal)
				math.SetNthDigit(5, k-1, &pal)
				math.SetNthDigit(3, 9, &pal)
				math.SetNthDigit(4, 9, &pal)
			} else {
				l := math.NthDigit(1, pal)
				if l > 0 {
					math.SetNthDigit(1, l-1, &pal)
					math.SetNthDigit(6, l-1, &pal)
					math.SetNthDigit(2, 9, &pal)
					math.SetNthDigit(5, 9, &pal)
					math.SetNthDigit(3, 9, &pal)
					math.SetNthDigit(4, 9, &pal)
				}
			}
		}
	}
}

func Solve004Alt() int {
	pal, p := 0, 0
	for i := 999; i > 100; i = i - 1 {
		for j := i; j > 100; j = j - 1 {
			p = i * j
			if math.IsPalindrome(p) && p > pal {
				pal = p
				break
			}
			if p < 100001 || p < pal {
				break
			}
		}
	}
	return pal
}
