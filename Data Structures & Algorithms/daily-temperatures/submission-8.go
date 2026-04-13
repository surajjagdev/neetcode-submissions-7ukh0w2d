func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := make([]int, 0)

	for i := 0; i < n; i++ {
		temp := temperatures[i]

		for len(stack) != 0 && temp > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx] = i - idx
		}

		stack = append(stack, i)
	}
	

	return res
}
