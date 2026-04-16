func searchMatrix(matrix [][]int, target int) bool {
	// 1 - d search plane
	rows, cols := len(matrix), len(matrix[0])
	low, high := 0, (rows * cols) - 1

	for low <= high {
		mid := low + ((high - low) / 2)

		row, col := mid / cols, mid % cols
		val := matrix[row][col]

		if val == target {
			return true
		} else if val < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}
