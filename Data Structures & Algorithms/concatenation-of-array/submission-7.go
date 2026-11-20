func getConcatenation(nums []int) []int {
    N := len(nums)

	if N == 0 {
		return []int{}
	}

	res := make([]int, 2 * N)

	for i := 0; i < N; i++ {
		num := nums[i]

		res[i] = num
		res[i + N] = num
	}

	return res
}
