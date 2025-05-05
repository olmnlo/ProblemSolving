package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	// Base cases
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	// Recursively compare left and right subtrees
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
