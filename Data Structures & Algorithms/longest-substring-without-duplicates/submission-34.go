func lengthOfLongestSubstring(s string) int {
    mp := make(map[byte]int)
    n := len(s)
    left := 0
    maxLen := 0

    for right := 0; right < n; right++ {
        ch := s[right]

        // if char was already in map,  delete everything along the way
        if lastIndex, ok := mp[ch]; ok && lastIndex + 1 > left {
            left = lastIndex + 1
        }

        mp[ch] = right

        windowSize := right - left + 1
        maxLen = max(maxLen, windowSize)
    }

    fmt.Println(mp)

    return maxLen
}
