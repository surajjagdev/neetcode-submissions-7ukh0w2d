func getRow(matrix [][]int, target int, rows int, cols int) int {
	top, bottom := 0, rows - 1

	for top <= bottom {
		row := top + ((bottom - top) / 2)

		// check if range is valid
		if target > matrix[row][cols-1] {
			top = row + 1
		} else if target < matrix[row][0] {
			bottom = row - 1
		} else {
			// this is our row
			return row
		}
	}

	return -1
}


func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])

	// get the correct row
	row := getRow(matrix, target, rows, cols)

	if row == -1 {
		return false
	}

	low, high := 0, cols - 1

	for low <= high {
		mid := low + ((high - low) / 2)

		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}

