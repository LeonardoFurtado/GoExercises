package array_and_slices

func Sum(arr []int) int {
	var result int
	for _, number := range arr {
		result += number
	}
	return result
}
