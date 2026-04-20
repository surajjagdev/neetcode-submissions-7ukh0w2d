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

	var dfs func(int, int)

	dfs = func(row, col int) {
		// mark cell as water
		grid[row][col] = '0'

		// travel in all 4 directions
		for k := 0; k < len(directions); k++ {
			nr := row + directions[k][0]
			nc := col + directions[k][1]

			if nr >= 0 && nr < ROWS && nc >= 0 && nc < COLS && grid[nr][nc] == '1' {
				dfs(nr, nc)
			}
		}
	}

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if grid[i][j] == '1' {
				res += 1

				dfs(i, j)
			}
		}
	}


	return res
}
