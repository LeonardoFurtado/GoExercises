package array_and_slices

const arrayLength = 3

func Sum(arr [arrayLength]int) int {
	var result int
	for _, number := range arr {
		result += number
	}
	return result
}
