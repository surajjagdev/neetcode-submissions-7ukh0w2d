func getMax(arr []int) int {
	res := math.MinInt

	for _, val := range arr {
		res = max(res, val)
	}

	return res
}

func minEatingSpeed(piles []int, h int) int {
	// find minimum k banana eating speed
	low, high := 1, getMax(piles)

	totalHoursTaken := func (k int) int {
		res := 0
		for _, val := range piles {
			res += int(math.Ceil(float64(val) / float64(k)))
		}

		return res
	}

	for low < high {
		mid := low + (high - low) / 2

		hoursTaken := totalHoursTaken(mid)

		// can we finish all bananas given mid value ?
		if hoursTaken <= h {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return high
}
