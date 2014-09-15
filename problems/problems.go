package problems

import (
	"container/list"
	"strconv"
	"strings"
)

import localmath "github.com/bwrightson/euler/math"

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
		if val%i == 0 && localmath.IsPrime(i) {
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
		j := localmath.NthDigit(3, pal)
		if j > 0 {
			localmath.SetNthDigit(3, j-1, &pal)
			localmath.SetNthDigit(4, j-1, &pal)
		} else {
			k := localmath.NthDigit(2, pal)
			if k > 0 {
				localmath.SetNthDigit(2, k-1, &pal)
				localmath.SetNthDigit(5, k-1, &pal)
				localmath.SetNthDigit(3, 9, &pal)
				localmath.SetNthDigit(4, 9, &pal)
			} else {
				l := localmath.NthDigit(1, pal)
				if l > 0 {
					localmath.SetNthDigit(1, l-1, &pal)
					localmath.SetNthDigit(6, l-1, &pal)
					localmath.SetNthDigit(2, 9, &pal)
					localmath.SetNthDigit(5, 9, &pal)
					localmath.SetNthDigit(3, 9, &pal)
					localmath.SetNthDigit(4, 9, &pal)
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
			if localmath.IsPalindrome(p) && p > pal {
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
	return localmath.LCM(series)
}

func Solve007() int {
	val := 10001
	return localmath.NthPrimeWithSieve(val)
}

func Solve007Alt() int {
	val := 10001
	return localmath.NthPrime(val)
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

func Solve009() int {
	for a := 2; a < 333; a++ {
		aSquared := a * a
		for b := a + 1; b < 1000; b++ {
			bSquared := b * b
			for c := b + 1; c < 1000; c++ {
				cSquared := c * c
				if a+b+c == 1000 && aSquared+bSquared == cSquared {
					return a * b * c
				}
				if cSquared > aSquared+bSquared || a+b+c > 1000 {
					break
				}
			}
		}
	}
	return 1
}

func Solve009Alt() int {
	a, b, c := 0, 0, 0
	for m := 1; m < 250; m++ {
		if 500%m != 0 {
			continue
		}
		n := (500/m - m)
		if m > n {
			a = m*m - n*n
			b = 2 * m * n
			c = m*m + n*n
			break
		}
	}
	return a * b * c
}

func Solve010() int {
	sieve := localmath.GeneratePrimesToN(2000000)
	sum := 2
	for i := 3; i < len(sieve); i += 2 {
		if sieve[i] {
			sum += i
		}
	}
	return sum
}

func Solve010Alt() int {
	sum := 2
	for i := 3; i < 2000000; i += 2 {
		if localmath.IsPrime(i) {
			sum += i
		}
	}
	return sum
}

func Solve011() int {
	input := `08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08
49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00
81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65
52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91
22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80
24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50
32 98 81 28 64 23 67 10 26 38 40 67 59 54 70 66 18 38 64 70
67 26 20 68 02 62 12 20 95 63 94 39 63 08 40 91 66 49 94 21
24 55 58 05 66 73 99 26 97 17 78 78 96 83 14 88 34 89 63 72
21 36 23 09 75 00 76 44 20 45 35 14 00 61 33 97 34 31 33 95
78 17 53 28 22 75 31 67 15 94 03 80 04 62 16 14 09 53 56 92
16 39 05 42 96 35 31 47 55 58 88 24 00 17 54 24 36 29 85 57
86 56 00 48 35 71 89 07 05 44 44 37 44 60 21 58 51 54 17 58
19 80 81 68 05 94 47 69 28 73 92 13 86 52 17 77 04 89 55 40
04 52 08 83 97 35 99 16 07 97 57 32 16 26 26 79 33 27 98 66
88 36 68 87 57 62 20 72 03 46 33 67 46 55 12 32 63 93 53 69
04 42 16 73 38 25 39 11 24 94 72 18 08 46 29 32 40 62 76 36
20 69 36 41 72 30 23 88 34 62 99 69 82 67 59 85 74 04 36 16
20 73 35 29 78 31 90 01 74 31 49 71 48 86 81 16 23 57 05 54
01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48`
	product := 1
	grid := make([][]int, 0)
	lines := strings.Split(input, "\n")
	for _, i := range lines {
		line := strings.Split(i, " ")
		temp := make([]int, 0)
		for j := range line {
			num, err := strconv.Atoi(line[j])
			if err != nil {
				return 0
			}
			temp = append(temp, num)
		}
		grid = append(grid, temp)
	}
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			// Handle right
			if j < 17 {
				candidate := grid[i][j] * grid[i][j+1] * grid[i][j+2] * grid[i][j+3]
				if candidate > product {
					product = candidate
				}
			}
			// Handle down
			if i < 17 {
				candidate := grid[i][j] * grid[i+1][j] * grid[i+2][j] * grid[i+3][j]
				if candidate > product {
					product = candidate
				}
			}
			// Handle diagonal down-right
			if j < 17 && i < 17 {
				candidate := grid[i][j] * grid[i+1][j+1] * grid[i+2][j+2] * grid[i+3][j+3]
				if candidate > product {
					product = candidate
				}
			}
			// Handle diagonal down-left
			if j > 2 && i < 17 {
				candidate := grid[i][j] * grid[i+1][j-1] * grid[i+2][j-2] * grid[i+3][j-3]
				if candidate > product {
					product = candidate
				}
			}
		}
	}
	return product
}

func Solve012() int {
	n := 0
	primes := localmath.GeneratePrimesToN(100000000)
	for i := 1; ; i++ {
		n += i
		if n%2 != 0 {
			continue
		}
		if localmath.DivisorCount(n, primes) > 500 {
			return n
		}
	}
}
