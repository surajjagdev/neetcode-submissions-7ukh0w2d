func isAnagram(s string, t string) bool {
	arrS := [26]int{}
	arrT := [26]int{}

	for _, v := range s {
		i := int(v - 'a')
		arrS[i] += 1
	}
	for _, v := range t {
		i := int(v - 'a')
		arrT[i] += 1
	}

	for i := 0; i < 26; i++ {
		if arrS[i] != arrT[i] {
			return false
		}
	}

	return true
}
