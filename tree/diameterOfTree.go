package tree
func Diameter(root *Node)int {
	if root == nil {
		return 0
	}
	lh := Height(root.Left)
	rh := Height(root.Right)
	maxi := Max(lh,rh)
	Diameter(root.Left)
	Diameter(root.Right)
	return maxi
}