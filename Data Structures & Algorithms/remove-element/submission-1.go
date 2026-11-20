func removeElement(nums []int, val int) int {
    left := 0
	
	for right := 0; right < len(nums); right++ {
		num := nums[right]

		if num != val {
			nums[left] = num
			left++
		}
	}

	return left
}
