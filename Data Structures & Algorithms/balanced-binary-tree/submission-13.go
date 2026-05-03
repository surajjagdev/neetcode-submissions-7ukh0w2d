/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func abs(a int) int {
	if a < 0 {
		return -1 * a
	}

	return a
 }

func isBalanced(root *TreeNode) bool {
	// get the height at each level of the tree, and 
	// see if they are balanced
	var dfs func(*TreeNode) (bool, int)

	dfs = func(node *TreeNode) (bool, int) {
		if node == nil {
			return true, 0
		}

		// get heights of left and right
		leftBalanced, leftHeight := dfs(node.Left)
		rightBalanced, rightHeight := dfs(node.Right)

		isBalanced := false

		if leftBalanced && rightBalanced && abs(leftHeight - rightHeight) <= 1 {
			isBalanced = true
		}

		maxHeight := 1 + max(rightHeight, leftHeight)

		return isBalanced, maxHeight
	}


	isBalanced, _ := dfs(root)

	return isBalanced
}
