func trap(height []int) int {
    res := 0
    n := len(height)

    for i := 0; i < n; i++ {
        leftMax, rightMax := height[i], height[i]

        for j := 0; j < i; j++ {
            leftMax = max(leftMax, height[j])
        }
        for j := i + 1; j < n; j++ {
            rightMax = max(rightMax, height[j])
        }

        fmt.Println(leftMax, " ", rightMax)
        res += min(leftMax, rightMax) - height[i]
    }

    return res
}
