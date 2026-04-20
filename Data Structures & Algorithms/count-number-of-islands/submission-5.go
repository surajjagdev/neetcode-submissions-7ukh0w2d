func numIslands(grid [][]byte) int {
    // need a function to count how many connected components we have together

	res := 0
	ROWS := len(grid)
	COLS := len(grid[0])
	directions := [][]int {
		{-1,0},
		{1,0},
		{0,-1},
		{0,1},
	}

	bfs := func(row, col int) {
		q := make([][]int, 0)
		q = append(q, []int{row, col})
		grid[row][col] = '0'

		for len(q) != 0 {
			sz := len(q)

			for i := 0; i < sz; i++ {
				r, c := q[0][0], q[0][1]
				// pop from queue
				q = q[1:]

				for _, direction := range directions {
					nr, nc := r + direction[0], c + direction[1]

					if nr >= 0 && nr < ROWS && nc >= 0 && nc < COLS && grid[nr][nc] == '1' {
						q = append(q, []int{nr, nc})
						grid[nr][nc] = '0'
					}
				} 
			}
		}
	}

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if grid[i][j] == '1' {
				res += 1

				bfs(i, j)
			}
		}
	}


	return res
}
