func trap(height []int) int {
    water := 0
    left, right := 0, len(height) - 1
    leftMax, rightMax := height[left], height[right]


    for left < right {
        if leftMax < rightMax {
            left += 1
            leftMax = max(leftMax, height[left])
            water += leftMax - height[left]
        } else {
            right -= 1
            rightMax = max(rightMax, height[right])
            water += rightMax - height[right]
        }
    }

    return water
}
