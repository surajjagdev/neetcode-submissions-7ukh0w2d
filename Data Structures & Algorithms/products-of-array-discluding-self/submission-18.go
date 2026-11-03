func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	suffix := make([]int, n)
	prefix[0] = 1
	suffix[n-1] = 1
	
	for i := 1; i < n; i++ {
		prefix[i] = nums[i-1] * prefix[i-1]
	}
	for i := n - 2; i > -1 ; i-- {
		suffix[i] = nums[i+1] * suffix[i+1]
	}
	for i := 0; i < n; i++ {
		prefix[i] = prefix[i] * suffix[i]
	}

	return prefix
}
