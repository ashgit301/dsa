package tree

func Floor(key int, root *Node) int {
	floor := -1
	for root != nil {
		if key == root.Val {
			return root.Val
		}
		if key > root.Val {
			floor = root.Val
			root = root.Right
		}
		if key < root.Val { //dont upate floor value
			root = root.Left
		}
	}
	return floor

}