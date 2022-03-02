package tree

func lowestCommonAncestorOfBST(root, p, q *Node) *Node {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorOfBST(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorOfBST(root.Right, p, q)
	}
	return root
}