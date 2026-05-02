/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

    count := 1

	var search func(*TreeNode, int) int
 
	search = func(node *TreeNode, maxVal int) int {
		if node == nil {
			return 0
		}

		if node.Left != nil {
			count += search(node.Left, max(maxVal, node.Val))
		}
		if node.Right != nil {
			count += search(node.Right, max(maxVal, node.Val))
		}

		if node.Val >= maxVal {
			return 1
		}

		return 0
	}

	search(root, root.Val)

	return count
}
