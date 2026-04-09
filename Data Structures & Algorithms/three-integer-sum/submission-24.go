func kSum(nums []int, k int, start int, target int, path []int, res *[][]int){
    n := len(nums)

    if k == 2 {
        left, right := start, n - 1

        for left < right {
            curr := nums[left] + nums[right]

            if curr == target {
                tmp := append(append([]int{}, path...), nums[left], nums[right])
                *res = append(*res, tmp)

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

            path = append(path, nums[i])
            kSum(nums, k - 1, i + 1, target - nums[i], path, res)
            path = path[:len(path)-1] // backtrack
        }
    }
}

func threeSum(nums []int) [][]int {
    sort.Slice(nums, func(a, b int) bool {
        return nums[b] > nums[a]
    })

    k := 3
    res := make([][]int, 0)
    kSum(nums, k, 0, 0, []int{}, &res)

    return res
}
