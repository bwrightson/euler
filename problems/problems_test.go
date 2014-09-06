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
