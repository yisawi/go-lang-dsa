func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)

	for i, num := range nums {
		comp := target - num

		if prevIndex, found := hashMap[comp]; found {
		return []int{prevIndex, i}
	}
	hashMap[num] = i
	}
	return []int{}
}
