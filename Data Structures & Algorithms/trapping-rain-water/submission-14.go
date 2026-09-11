func trap(height []int) int {

    // get left max and right max
    leftMax := make([]int, len(height))
    rightMax := make([]int, len(height))

    for i := 0; i < len(height); i++ {
        prevElement := height[0]
        if i > 0 {
            prevElement = leftMax[i-1]
        }

        leftMax[i] = max(height[i], prevElement)
    }
    for i := len(height)-1; i >= 0; i-- {
        nextElement := height[len(height) - 1]
        if i < len(height) - 1 {
            nextElement = rightMax[i+1]
        }

        rightMax[i] = max(height[i], nextElement)
    }

    water := 0

    for i := 0; i < len(height); i++ {
        surronding := min(leftMax[i], rightMax[i])
        trapped := surronding - height[i]

        water += trapped
    }

    return water
}
