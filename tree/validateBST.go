package tree

func ValidateBST(root, Max , Min *Node) bool {

	if root == nil {
		return true
	}
	if (Max != nil && root.Val > Max.Val) ||(Min != nil && Min.Val > root.Val){
		return false
	}
	return ValidateBST(root.Left, root, Min) && ValidateBST(root.Right, Max, root)
}