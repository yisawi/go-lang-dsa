func replaceElements(arr []int) []int {
	maxSoFar := -1
	result := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		result[i] = maxSoFar
		maxSoFar = max(maxSoFar, arr[i])
	}
	return result
}
