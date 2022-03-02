package tree

func InorderTraversal(Node *Node, result*[]int){
	if Node == nil {
		return
	} 
	InorderTraversal(Node.Left, result)
	*result = append(*result, Node.Val)
    InorderTraversal(Node.Right,result)
}