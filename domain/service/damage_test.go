package service

import "testing"

func TestCalculateDamge(t *testing.T) {
	min, max := CalculateDamage(50, 100, 182, 189)
	if min != 37 {
		t.Errorf("min expected: 37 result: %d", min)
	}
	if max != 44 {
		t.Errorf("min expected: 44 result: %d", max)
	}
}
