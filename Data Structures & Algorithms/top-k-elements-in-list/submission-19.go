func topKFrequent(nums []int, k int) []int {
	// create freq map. Number's freq cannot exceed 
	// number of elements in list
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	bucket := make([][]int, len(nums) + 1)

	for key, value := range freq {
		bucket[value] = append(bucket[value], key)
	}

	res := make([]int, 0, k)

	// go through from end of bucket till we fill up res
	for i := len(bucket) - 1; i >= 1; i-- {
		for _, num := range bucket[i] {
			res = append(res, num)

			if len(res) == k {
				return res
			}
		}
	}

	return res
}
