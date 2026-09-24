/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    var res []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		elem := queue[0]
		if elem == nil {
			break
		}
		queue = queue[1:]
		res = append(res, elem.Val)
		if elem.Left != nil{
			queue = append(queue, elem.Left)	
		}
		if elem.Right != nil{
			queue = append(queue, elem.Right)	
		}
	}
	sort.Ints(res)
	return res[k-1]
}
