/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    q := make([]*TreeNode, 0)
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	q = append(q, root)

	for len(q) != 0 {
		sz := len(q)
		temp := make([]int, sz)

		for i := 0; i < sz; i++ {
			curr := q[i]
			temp[i] = curr.Val

			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}

		q = q[sz:]

		res = append(res, temp)
	}

	return res
}
