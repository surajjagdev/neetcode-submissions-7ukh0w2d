func getIdx(c byte) int {
    return int(c - 'a')
}

func checkInclusion(s1 string, s2 string) bool {
    var charArrS1 [26]int
    var charArrS2 [26]int
    left := 0

    // how many unique chars do we need and have matched
    required, matches := 0, 0
    
    for i := 0; i < len(s1); i++ {
        idx := getIdx(s1[i])
        
        // unique
        if charArrS1[idx] == 0 {
            required++
        }
        charArrS1[idx]++
    }

    for right := 0; right < len(s2); right++{
        ch := s2[right]
        idx := getIdx(ch)

        // add this char to arr
        charArrS2[idx]++

        if charArrS1[idx] == charArrS2[idx] {
            matches++
        }
        if charArrS1[idx] == charArrS2[idx] - 1 {
            // what if adding it doesnt match anymore?
            matches--
        }

        // what if our window is beyond str len of s1?
        if (right - left + 1) > len(s1) {
            // shrink window
            leftIdx := getIdx(s2[left])
            charArrS2[leftIdx]--

            // if removing char made things equal ?
            if charArrS1[leftIdx] == charArrS2[leftIdx] {
                matches++
            }
            if charArrS1[leftIdx] == charArrS2[leftIdx] + 1 {
                matches--
            }

            left++
        }


        if required == matches {
            return true
        }
    }

    return false
}
