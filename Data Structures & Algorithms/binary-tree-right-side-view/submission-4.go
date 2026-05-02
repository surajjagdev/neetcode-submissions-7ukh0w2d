/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    res := make([]int, 0)

	if root == nil {
		return res
	}

	q := make([]*TreeNode, 0)
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

			if i == sz - 1 {
				res = append(res, node.Val)
			}
		}

		q = q[sz:]
	}

	return res
}
