func isValidSudoku(board [][]byte) bool {
	// keep a row-set
	// keep a col set
	// square must also be valid
	N := len(board)
	rowSet := make([]map[byte]struct{}, N)
	colSet := make([]map[byte]struct{}, N)
	squares := make(map[string]map[byte]struct{})

	for i := 0; i < N; i++ {
		rowSet[i] = make(map[byte]struct{})
		colSet[i] = make(map[byte]struct{})
	}

	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if board[r][c] == '.' {
				continue
			}

			val := board[r][c]
			squareIdx := fmt.Sprintf("%d,%d", r/3, c/3)

			if _, ok := squares[squareIdx]; !ok {
				squares[squareIdx] = make(map[byte]struct{})
			}

			if _, ok := rowSet[r][val]; ok {
				return false
			}
			if _, ok := colSet[c][val]; ok {
				return false
			}
			if _, ok := squares[squareIdx][val]; ok {
				return false
			}

			rowSet[r][val] = struct{}{}
			colSet[c][val] = struct{}{}
			squares[squareIdx][val] = struct{}{}
		}
	}


	return true
}
