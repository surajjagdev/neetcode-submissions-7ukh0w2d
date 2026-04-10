func mostFreq(arr *[26]int) int {
    k := 0

    for _, val := range(*arr) {
        k = max(k, val)
    }

    return k
}

func characterReplacement(s string, k int) int {
    maxLen := 0
    left := 0
    n := len(s)
    var charArray [26]int

    for right := 0; right < n; right++ {
        ch := s[right]
        idx := ch - byte('A')
        charArray[int(idx)]++

        // get freq of most frequent char
        freq := mostFreq(&charArray)

        // how many chars can we replace
        currSize := right - left + 1
        remaining := currSize - freq

        if remaining > k {
            // remove a character from window
            leftIdx := s[left] - byte('A')
            charArray[int(leftIdx)]--
            left++
        }

        maxLen = max(maxLen, (right - left + 1))
    }

    return maxLen
}
