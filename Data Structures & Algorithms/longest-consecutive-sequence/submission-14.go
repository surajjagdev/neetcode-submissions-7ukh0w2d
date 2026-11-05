func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func longestConsecutive(nums []int) int {
	// we can make a set
	s := make(map[int]struct{})

	for _, num := range(nums) {
		s[num] = struct{}{}
	}

	res := 0

	for _, num := range(nums) {
		if _, ok := s[num-1]; ok {
			continue
		}

		// first
		j := 1
		for {
			if _, ok := s[num+j]; !ok {
				break
			}
			j++
		}

		res = max(res, j)
	}

	return res
}
