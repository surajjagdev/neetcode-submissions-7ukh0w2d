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
	i := len(bucket) - 1

	for len(res) < k && i >= 0 {
		currBucket := bucket[i]
		j := len(currBucket) - 1

		for len(res) < k && j >= 0 {
			res = append(res, currBucket[j])
			j -= 1
		}

		i -= 1
	}

	return res
}
