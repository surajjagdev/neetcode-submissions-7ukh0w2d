func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

    // we can loop through, but we need to start with the 
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		j := 0

		for j < len(prefix) && j < len(strs[i]) {
			if prefix[j] != strs[i][j] {
				break
			}

			j++
		}

		prefix = prefix[:j]
	}

	return prefix
}
