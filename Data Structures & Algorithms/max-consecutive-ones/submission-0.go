func findMaxConsecutiveOnes(nums []int) int {
	counter := 0
	maxCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			counter += 1
		} else {
			counter = 0
		}
		if counter > maxCount {
			maxCount = counter
		}
	}
	return maxCount
}
