package tree

func Ceil(key int, root *Node) int {
	ceil:= -1
	for root != nil {
		if key == root.Val {
			return root.Val
		}
		if key > root.Val {
			root = root.Right
		}
		if key < root.Val { 
			ceil= root.Val
			root = root.Left
		}
	}
	return ceil

}