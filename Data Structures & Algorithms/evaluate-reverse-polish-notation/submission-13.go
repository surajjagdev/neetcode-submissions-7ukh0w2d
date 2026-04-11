func toNum(val string) int {
	v, _ := strconv.Atoi(val)

	return v
}

func getValue(operand string, stk *[]string) int{
	res := toNum((*stk)[len(*stk)-2])
	currVal := toNum((*stk)[len(*stk)-1])

	if operand == "+" {
		res += (currVal)
	} else if operand == "-" {
		res -= (currVal)
	} else if operand == "*" {
		res *= (currVal)
	} else {
		res /= (currVal)
	}

	return res
}

func evalRPN(tokens []string) int {
	stk := make([]string, 0)

	for _, val := range tokens {
		if val == "+" || val == "*" || val == "-" || val == "/" {
			res := getValue(val, &stk)
			stk = stk[:len(stk)-2]
			cv := strconv.Itoa(res)
			stk = append(stk, cv)
		} else {
			stk = append(stk, val)
		}
	}

	v, _ := strconv.Atoi(stk[len(stk)-1])

	return v
}
