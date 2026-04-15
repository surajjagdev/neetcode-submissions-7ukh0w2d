func carFleet(target int, position []int, speed []int) int {
	// cars cannot pass one another, but if car beyond
	// will reach target before car infront they become one
	// fleet

	pairs := make([][]int, 0)
	stk := make([]float64, 0)
	n := len(position)

	for i := 0; i < n; i++ {
		pairs = append(pairs, []int{position[i], speed[i]})
	}

	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a][0] > pairs[b][0]
	})

	getTime := func(p, s int) float64{
		return (float64(target) - float64(p)) / float64(s)
	}

	for i := 0; i < n; i++ {
		currTime := getTime(pairs[i][0], pairs[i][1])

		if len(stk) > 0 {
			lastTime := stk[len(stk)-1]

			if currTime <= lastTime {
				continue
			}
		}

		stk = append(stk, currTime)
	}

	return len(stk)
}
