func twoSum(numbers []int, target int) []int {
	mp := make(map[int]int)

	for i, v := range numbers {
		diff := target - v

		if k, ok := mp[diff]; ok {
			return []int{k+1, i + 1}
		}

		mp[v] = i
	}

	return []int{}
}
