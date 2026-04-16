func searchMatrix(matrix [][]int, target int) bool {
	// find the correct row
	rows, cols := len(matrix), len(matrix[0])

	// 1. binary search to find correct row
	top, bottom := 0, rows - 1

	for top <= bottom {
		row := top + ((bottom - top) / 2)

		beginning, end := matrix[row][0], matrix[row][cols-1]

		if target > end {
			top = row + 1
		} else if target < beginning {
			bottom = row - 1
		} else {	
			// in range
			break
		}
	}

	// if we didnt find a range suitable for row
	if !(top <= bottom) {
		return false
	}

	left, right := 0, cols - 1
	row := top + ((bottom - top) / 2)

	for left <= right {
		mid := left + ((right - left) / 2)

		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}	

	return false
}
