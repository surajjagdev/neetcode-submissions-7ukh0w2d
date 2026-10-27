func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	mp := make(map[[26]int][]string)

	for _, v := range strs {
		arr := [26]int{}
		
		for _, c := range(v) {
			idx := int(c - 'a')
			arr[idx] += 1
		}

		mp[arr] = append(mp[arr], v)
	}

	for _, arr := range(mp) {
		res = append(res, arr)
	}

	return res
}
