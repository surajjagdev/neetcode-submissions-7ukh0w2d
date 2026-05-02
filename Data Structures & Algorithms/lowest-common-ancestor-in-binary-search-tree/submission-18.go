/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    // do we go left, or do we go right
	if root == nil {
		return nil
	}

	rootVal := root.Val
	minVal := min(p.Val, q.Val)
	maxVal := max(p.Val, q.Val)
	

	if maxVal < rootVal {
		return lowestCommonAncestor(root.Left, p, q)
	} else if minVal > rootVal {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}
