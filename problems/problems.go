package problems

import (
	"container/list"
	"github.com/bwrightson/euler/math"
	"strconv"
	"strings"
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

func Solve005() int {
	series := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	return math.LCM(series)
}

func Solve007() int {
	val := 10001
	return math.NthPrimeWithSieve(val)
}

func Solve007Alt() int {
	val := 10001
	return math.NthPrime(val)
}

func Solve008() int {
	input := `73167176531330624919225119674426574742355349194934
96983520312774506326239578318016984801869478851843
85861560789112949495459501737958331952853208805511
12540698747158523863050715693290963295227443043557
66896648950445244523161731856403098711121722383113
62229893423380308135336276614282806444486645238749
30358907296290491560440772390713810515859307960866
70172427121883998797908792274921901699720888093776
65727333001053367881220235421809751254540594752243
52584907711670556013604839586446706324415722155397
53697817977846174064955149290862569321978468622482
83972241375657056057490261407972968652414535100474
82166370484403199890008895243450658541227588666881
16427171479924442928230863465674813919123162824586
17866458359124566529476545682848912883142607690042
24219022671055626321111109370544217506941658960408
07198403850962455444362981230987879927244284909188
84580156166097919133875499200524063689912560717606
05886116467109405077541002256983155200055935729725
71636269561882670428252483600823257530420752963450`
	lines := strings.Split(input, "\n")
	product := 0
	l := list.New()
	for _, v := range lines {
		for j := 0; j < len(v); j++ {
			candidate := 1
			k, err := strconv.Atoi(string(v[j]))
			if err != nil {
				return 0
			}
			if k == 0 {
				l = list.New()
				continue
			}
			l.PushBack(k)
			if l.Len() > 13 {
				l.Remove(l.Front())
			}
			for e := l.Front(); e != nil; e = e.Next() {
				candidate = candidate * e.Value.(int)
			}
			if candidate > product {
				product = candidate
			}
		}
	}
	return product
}
