package tree
func MinimumVal(root *Node, min *int)int {
	if root == nil {
		return 100000
	}
	if root.Val < *min {
		*min = root.Val
	}
	MinimumVal(root.Left, min)
	MinimumVal(root.Right, min)
	return *min

}

//same pattern for max_value
//can also do inorder & find the pick first & last element