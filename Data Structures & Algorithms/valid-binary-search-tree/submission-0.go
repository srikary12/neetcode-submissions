/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    if root == nil {
		return true
	}
	var dfs func(node *TreeNode, left int64, right int64) bool
	dfs = func(node *TreeNode, left int64, right int64) bool {
		if node == nil {
			return true
		}
		val := int64(node.Val)
		if val <= left || val>= right {
			return false
		}
		return dfs(node.Left, left, val) && dfs(node.Right, val, right)
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}
