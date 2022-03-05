package array_and_slices

import "testing"

func TestSum(t *testing.T) {
	numbers := [3]int{1, 2, 3}
	sum := Sum(numbers)
	expected := 6

	if sum != expected {
		t.Errorf("expected %d but got %d given, %v", expected, sum, numbers)
	}
}
