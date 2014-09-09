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
