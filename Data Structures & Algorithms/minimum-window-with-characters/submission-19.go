func getIdx(c byte) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a')
	}

	return int(c - 'A') + 26
}

func minFreqMatch(arr1 *[52]int, arr2 *[52]int) bool {
	for i := 0; i < len(*arr1); i++ {
		if (*arr2)[i] < (*arr1)[i] {
			return false
		}
	}

	return true
}

func minWindow(s string, t string) string {
    var charArrS [52]int
	var charArrT [52]int
	startIdx := 0
	minLen := math.MaxInt
	left := 0

	// create char arr for t
	for i := 0; i < len(t); i++ {
		idx := getIdx(t[i])
		charArrT[idx]++
	}

	for right := 0; right < len(s); right++ {
		idx := getIdx(s[right])
		charArrS[idx]++

		// when minFreq matches, try to reduce the window
		for minFreqMatch(&charArrT, &charArrS) {
			fmt.Println(startIdx, " ", minLen)
			currWindow := right - left + 1

			if currWindow < minLen {
				minLen = currWindow
				startIdx = left
			}

			// remove character from left
			leftIdx := getIdx(s[left])
			charArrS[leftIdx]--
			left++
		}
	}


	if minLen == math.MaxInt {
		return ""
	}

	return s[startIdx:startIdx+minLen]
}
