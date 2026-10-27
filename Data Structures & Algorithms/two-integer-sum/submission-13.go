func twoSum(nums []int, target int) []int {
    mp := make(map[int]int)

	for i, v := range nums {
		diff := target - v

		if v, ok := mp[diff]; ok {
			return []int {v, i}
		}

		mp[v] = i
	}

	return []int{}
}
