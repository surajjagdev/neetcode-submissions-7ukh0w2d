/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    depth := 0
	q := make([]*TreeNode, 0)

	if root == nil {
		return depth
	}

	q = append(q, root)

	for len(q) != 0 {
		sz := len(q)

		for i := 0; i < sz; i++ {
			node := q[i]

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[sz:]
		depth += 1
	}

	return depth
}
