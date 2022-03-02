package tree
var predecessor,successor int
func Pred_Succ(key int, root *Node) (int, int) {
	if key == root.Val {
		predecessor := root.Val
		successor := root.Val
		if root.Right != nil {
			tmp := root.Right
			for tmp.Left != nil {
				tmp = tmp.Left
			}
			successor = tmp.Val 
		}
		if root.Left != nil {
			tmp := root.Left
			for tmp.Right != nil {
				tmp = tmp.Right
			}
			predecessor = tmp.Val 
		}
		return predecessor,successor
	}
	if key > root.Val {
		Pred_Succ(key, root.Right)
	}
	if key < root.Val {
		Pred_Succ(key, root.Left)
	}
	return predecessor,successor
}