func replaceElements(arr []int) []int {
    n := len(arr)
    maxSoFar := -1

	for i := n - 1; i >= 0; i-- {
		temp := arr[i]
		arr[i] = maxSoFar
        if temp > maxSoFar {
        maxSoFar = temp
    }
	}
	return arr
}
