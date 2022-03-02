package tree
//easy if you trace with strivers example
func Delete(root *Node, key int) *Node {
	if root.Val == key {
		return Helper(root)
	}
	node := root
	for root != nil {
		if key > root.Val {
			if root.Right != nil && root.Right.Val == key {
				root.Right = Helper(root)
				break
			}else {
				root = root.Right
			}
		}else {
			if root.Left != nil && root.Left.Val == key {
				root.Left = Helper(root)
				break
			}else {
				root = root.Left
			}
		}
	}
	return node
}



func Helper(root *Node) *Node {
	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}
	rightSubTree := root.Right
	LastRight := FindLast(root.Left)
	LastRight.Right = rightSubTree
	return root.Left

}

func FindLast(root *Node) *Node {
	if root.Right == nil {
		return root
	}
	return FindLast(root.Right)
}