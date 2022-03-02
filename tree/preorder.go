package tree

func PreorderTraversal(Node *Node, result*[]int){
	if Node == nil {
		return
	} 
	*result = append(*result, Node.Val)
	PreorderTraversal(Node.Left, result)
    PreorderTraversal(Node.Right,result)
}