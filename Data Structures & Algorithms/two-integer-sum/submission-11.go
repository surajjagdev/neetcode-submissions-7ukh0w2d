func twoSum(nums []int, target int) []int {
    // not sorted 
    hashMap := make(map[int]int)

    for i, v := range(nums) {
        remaining := target - v

        oth, ok := hashMap[remaining]

        if ok {
            return []int {oth, i}
        }

        hashMap[v] = i
    }

    return []int {}
}
