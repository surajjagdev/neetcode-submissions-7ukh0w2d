func isAlphaNum(s byte) bool {
	return (s >= 'a' && s <= 'z') ||
		(s >= 'A' && s <= 'Z') ||
		(s >= '0' && s <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}

	return c
}

func isPalindrome(s string) bool {
	n := len(s)
	left, right := 0, n - 1

	for left < right {
		for left < right && !isAlphaNum(s[left]) {
			left++
		}
		for left < right && !isAlphaNum(s[right]) {
			right--
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}
