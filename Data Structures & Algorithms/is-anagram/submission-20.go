func isAnagram(s string, t string) bool {
	n, m := len(s), len(t)

	if n != m {
		return false
	}

	nArr := make([]rune, 26)
	mArr := make([]rune, 26)

	for i := 0; i < n; i++ {
		sIndex := rune(s[i]) - rune('a')
		tIndex := rune(t[i]) - rune('a')
		nArr[sIndex] += 1
		mArr[tIndex] += 1
	}

	for i := 0; i < 26; i++ {
		if nArr[i] != mArr[i] {
			return false
		}
	}

	return true
}
