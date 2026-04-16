func getMax(arr []int) int {
	res := math.MinInt

	for _, val := range arr {
		res = max(res, val)
	}

	return res
}

func totalHoursTaken(arr []int, k int) int {
	res := 0
	for _, val := range arr {
		res += int(math.Ceil(float64(val) / float64(k)))
	}

	return res
}

func minEatingSpeed(piles []int, h int) int {
	// find minimum k banana eating speed
	low, high := 1, getMax(piles)

	fmt.Println(low, high)

	for low < high {
		mid := low + (high - low) / 2

		hoursTaken := totalHoursTaken(piles, mid)

		// can we finish all bananas given mid value ?
		if hoursTaken <= h {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return high
}
