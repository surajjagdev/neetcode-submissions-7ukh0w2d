func maxArea(heights []int) int {
    // area of how much we can hold
    n := len(heights)
    left, right := 0, n - 1
    maxArea := 0

    for left < right {
        width := right - left
        height := min(heights[left], heights[right])
        currArea := height * width
        maxArea = max(maxArea, currArea)

        // do we go left or right
        if heights[left] < heights[right] {
            left += 1
        } else {
            right -= 1
        }
    }

    return maxArea
}
