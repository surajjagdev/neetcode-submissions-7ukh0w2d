func isAnagram(s string, t string) bool {
	n, m := len(s), len(t)

	var indexer func (rune) int

	indexer = func(a rune) int {
		return int(a - 'a')
	}

	if n != m {
		return false
	}

	nArr := make([]rune, 26)
	mArr := make([]rune, 26)

	for i := 0; i < n; i++ {
		sIndex := indexer(rune(s[i]))
		tIndex := indexer(rune(t[i]))
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
