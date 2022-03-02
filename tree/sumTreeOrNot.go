package tree

func SumTreeOrNot(root *Node) bool {
	ls := SumTree(root.Left)
	rs := SumTree(root.Right)
	return (root.Val == ls+rs) && SumTreeOrNot(root.Left) && SumTreeOrNot(root.Right)
}