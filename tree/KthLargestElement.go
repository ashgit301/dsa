package tree

func Largest(k int, root *Node, result []int) int{
	InorderTraversalForKth(root, &result)
	return result[k]

}

func InorderTraversalForKth(Node *Node, result*[]int){
	if Node == nil {
		return
	} 
	InorderTraversal(Node.Left, result)
	*result = append(*result, Node.Val)
    InorderTraversal(Node.Right,result)
}
