package tree

func SumTree(root *Node) int {
	if root == nil {
		return 0
	}
	return root.Val + SumTree(root.Left) + SumTree(root.Right)
}