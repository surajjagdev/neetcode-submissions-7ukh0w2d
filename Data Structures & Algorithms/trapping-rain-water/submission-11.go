func trap(height []int) int {
    res := 0
    n := len(height)

    suffix := make([]int, n)
    prefix := make([]int, n)
    prefix[0] = height[0]
    suffix[n-1] = height[n-1]

    for i := 1; i < n; i++ {
        prefix[i] = max(height[i], prefix[i-1])
    }
    for i := n - 2; i >= 0; i-- {
        suffix[i] = max(height[i], suffix[i+1])
    }

    fmt.Println(prefix)
    fmt.Println(suffix)

    for i := 0; i < n; i++ {
        minHeight := min(prefix[i], suffix[i])
        res += (minHeight - height[i])
    }

    return res
}
