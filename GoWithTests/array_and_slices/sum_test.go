package array_and_slices

import "testing"

func TestSum(t *testing.T) {
	t.Run("Collection of 4 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		expected := 6

		if sum != expected {
			t.Errorf("expected %d but got %d given, %v", expected, sum, numbers)
		}
	})

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		sum := Sum(numbers)
		expected := 10

		if sum != expected {
			t.Errorf("expected %d but got %d given, %v", expected, sum, numbers)
		}
	})
}
