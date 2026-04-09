func kSum(nums []int, k int, start int, target int) [][]int{
    res := make([][]int, 0)
    n := len(nums)

    if k == 2 {
        left, right := start, n - 1

        for left < right {
            curr := nums[left] + nums[right]

            if curr == target {
                res = append(res, []int{nums[left], nums[right]})

                left += 1
                right -=1

                for left < right && nums[left] == nums[left - 1] {
                    left += 1
                }
            } else if curr < target {
                left += 1
            } else {
                right -= 1
            }
        }
    } else {
        for i := start; i < n - k + 1; i++ {

            if i > start && nums[i] == nums[i-1] {
                continue
            }

            inner := kSum(nums, k - 1, i + 1, target - nums[i])

            for _, subset := range inner {
                comb := append([]int{nums[i]}, subset...)
                res = append(res, comb)
            }
        }
    }

    return res
}

func threeSum(nums []int) [][]int {
    sort.Slice(nums, func(a, b int) bool {
        return nums[b] > nums[a]
    })

    k := 3
    return kSum(nums, k, 0, 0)
}
