func isValid(s string) bool {
    stk := make([]rune, 0)
	mp := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		if val, ok := mp[ch]; ok {
			// closing brace
			if len(stk) == 0 {
				return false
			}

			lastEle := stk[len(stk)-1]
			stk = stk[:len(stk)-1]

			if lastEle != val {
				return false
			}
		} else {
			stk = append(stk, ch)
		}
	}


	return len(stk) == 0
}
