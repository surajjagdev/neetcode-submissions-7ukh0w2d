func getIdx(c byte) int {
    return int(c - 'a')
}

func checkInclusion(s1 string, s2 string) bool {
    var charArrS1 [26]int
    var charArrS2 [26]int
    left := 0

    // how many unique chars do we need and have matched
    
    for i := 0; i < len(s1); i++ {
        idx := getIdx(s1[i])
        charArrS1[idx]++
    }

    for right := 0; right < len(s2); right++{
        ch := s2[right]
        idx := getIdx(ch)
        charArrS2[idx]++

        // what if our window is beyond str len of s1?
        if (right - left + 1) > len(s1) {
            // shrink window
            leftIdx := getIdx(s2[left])
            charArrS2[leftIdx]--

            left++
        }


        if  charArrS1 == charArrS2 {
            return true
        }
    }

    return false
}
