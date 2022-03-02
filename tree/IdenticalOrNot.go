package tree

func IdenticalOrNot(root1,root2 *Node) bool{
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return root1 == root2 && IdenticalOrNot(root1.Left, root2.Left) && IdenticalOrNot(root1.Right, root2.Right)
}

//always the identity/mirror/symmetry goes leve by level