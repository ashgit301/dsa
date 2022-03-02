package tree

func MaxPathRootToSum(root *Node) int{
	if root == nil {
		return 0
	}
	lsum := MaxPathSum(root.Left)
	rsum := MaxPathSum(root.Right)
	return Max(lsum,rsum)
}