package problems

import "testing"

func TestProblem1(t *testing.T) {
	if Solve001() != 233168 {
		t.Error("Problem 1 is broken")
	}
}

func TestProblem2(t *testing.T) {
	if Solve002() != 4613732 {
		t.Error("Problem 2 is broken")
	}
}

func TestProblem3(t *testing.T) {
	if Solve003() != 6857 {
		t.Error("Problem 3 is broken")
	}
}

func TestProblem4(t *testing.T) {
	if Solve004() != 906609 {
		t.Error("Problem 4 is broken")
	}
}

func TestProblem4Alt(t *testing.T) {
	if Solve004Alt() != 906609 {
		t.Error("Problem 4Alt is broken")
	}
}

func TestProblem5(t *testing.T) {
	if Solve005() != 232792560 {
		t.Error("Problem 5 is broken")
	}
}

func TestProblem7(t *testing.T) {
	if Solve007() != 104743 {
		t.Error("Problem 7 is broken")
	}
}

func TestProblem7Alt(t *testing.T) {
	if Solve007Alt() != 104743 {
		t.Error("Problem 7Alt is broken")
	}
}

func TestProblem8(t *testing.T) {
	if Solve008() != 23514624000 {
		t.Error("Problem 8 is broken")
	}
}

func TestProblem9(t *testing.T) {
	if Solve009() != 31875000 {
		t.Error("Problem 9 is broken")
	}
}

func TestProblem9Alt(t *testing.T) {
	if Solve009Alt() != 31875000 {
		t.Error("Problem 9Alt is broken")
	}
}

func TestProblem10(t *testing.T) {
	if Solve010() != 142913828922 {
		t.Error("Problem 10 is broken")
	}
}

func TestProblem10Alt(t *testing.T) {
	if Solve010Alt() != 142913828922 {
		t.Error("Problem 10Alt is broken")
	}
}

func TestProblem11(t *testing.T) {
	if Solve011() != 70600674 {
		t.Error("Problem 11 is broken")
	}
}

func TestProblem12(t *testing.T) {
	if Solve012() != 76576500 {
		t.Error("Problem 12 is broken")
	}
}

func TestProblem13(t *testing.T) {
	answer := []int{5, 5, 3, 7, 3, 7, 6, 2, 3, 0}
	test := Solve013()
	if len(answer) != len(test) {
		t.Error("Problem 13 is broken")
	}
	for i := range answer {
		if answer[i] != test[i] {
			t.Error("Problem 13 is broken")
		}
	}
}

func TestProblem14(t *testing.T) {
	if Solve014() != 837799 {
		t.Error("Problem 14 is broken")
	}
}
