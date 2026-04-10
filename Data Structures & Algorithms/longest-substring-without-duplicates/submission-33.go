func lengthOfLongestSubstring(s string) int {
    mp := make(map[byte]int)
    n := len(s)
    left := 0
    maxLen := 0

    for right := 0; right < n; right++ {
        ch := s[right]

        // if char was already in map,  delete everything along the way
        if _, ok := mp[ch]; ok {
            for {
                delete(mp, s[left])
                left++

                if _, ok2 := mp[ch]; !ok2 {
                    break
                }
            }
        }

        mp[ch] = right

        windowSize := right - left + 1
        maxLen = max(maxLen, windowSize)
    }

    fmt.Println(mp)

    return maxLen
}
