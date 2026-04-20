func rob(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

    one := nums[0]
	two := max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		temp := max(two, one + nums[i])
		one = two
		two = temp
	}

	return two
}
