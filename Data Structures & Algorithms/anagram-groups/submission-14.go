func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	mp := make(map[string][]string)

	for _, v := range strs {
		arr := [26]int{}
		
		for _, c := range(v) {
			idx := int(c - 'a')
			arr[idx] += 1
		}

		key := ""

		for j := 0; j < 26; j++ {
			key += (string(arr[j]) + ",")
		}

		if _, ok := mp[key]; !ok {
			mp[key] = make([]string, 0)
		}

		mp[key] = append(mp[key], v)
	}

	for _, arr := range(mp) {
		res = append(res, arr)
	}

	return res
}
