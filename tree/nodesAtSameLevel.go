package tree
var leafLevel int
func SameLevel(root *Node, level int) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		if leafLevel == 0 {
			level = leafLevel
			return true
		}
		return leafLevel == level

	}
	return SameLevel(root.Left, level+1) && SameLevel(root.Right, level+1)
}

