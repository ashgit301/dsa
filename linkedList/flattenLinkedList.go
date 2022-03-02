package linkedList

func FlattenLinkedList(root *Node) *Node {
	if root == nil || root.Next == nil {
		return root
	}
	root.Next = FlattenLinkedList(root.Next)
	root = MergeLinkedList(root, root.Next)
	return root
}