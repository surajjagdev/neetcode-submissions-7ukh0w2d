func search(nums []int, target int) int {
	low, high := 0, len(nums) - 1

	for low <= high {
		mid := low + ((high - low) / 2)

		fmt.Println(mid)

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low++
		} else {
			high--
		}
	}

	return -1
}
