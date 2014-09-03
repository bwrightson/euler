package problems

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
