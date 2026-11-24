func majorityElement(nums []int) int {
    // we can keep a res and count
	// if there is a majoirty element it will eventually 
	// set itself as res 
	res, count := math.MinInt, 0

	for _, num := range nums {
		if res == num {
			count++
		} else {
			if count > 0 {
				count--
			}

			if count == 0 {
				res = num
				count++
			}
		}
	}

	return res
}
