func largestRectangleArea(heights []int) int {
	// for a rectangle we can keep expanding, so long as our
	// next index's height is >=. If it is less than we have to stop
	maxArea := 0
	n := len(heights)
	stk := make([][]int, 0) // [0=height, 1 = index]

	for i := 0; i < n; i++ {
		currHeight := heights[i]
		minWidth := i

		// how far can we expand each height's width?
		for len(stk) != 0 && stk[len(stk)-1][0] > currHeight {
			height, idx := stk[len(stk)-1][0], stk[len(stk)-1][1]
			stk = stk[:len(stk)-1]
			width := i - idx

			currArea := height * width
			maxArea = max(maxArea, currArea)

			// we can actually expand our curr line back
			minWidth = min(minWidth, idx)
		}

		stk = append(stk, []int{currHeight, minWidth})
	}

	// There is left over heights, that go till end
	for len(stk) != 0 {
		height, idx := stk[len(stk)-1][0], stk[len(stk)-1][1]
		stk = stk[:len(stk)-1]
		width := n - idx

		currArea := height * width
		maxArea = max(maxArea, currArea)
	}

	return maxArea
}
