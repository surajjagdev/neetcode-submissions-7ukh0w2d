func findMin(nums []int) int {
	low, high := 0, len(nums) - 1

	for low < high {
		mid := low + (high - low) / 2

		// compare this value to the number on the left
		// compare value to number on right
		// are we in left sorted or right sorted?
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return nums[high]
}