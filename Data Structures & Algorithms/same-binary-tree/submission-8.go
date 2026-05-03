/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var dfs func(*TreeNode, *TreeNode) bool

	dfs  = func(node1 *TreeNode, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}

		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}

		return dfs(node1.Left, node2.Left) && dfs(node1.Right, node2.Right)
	}

	return dfs(p, q)
}
