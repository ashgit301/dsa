package tree
import ()
func PostorderTraversal(Node *Node, result *[]int){
	if Node == nil {
		return
	}
	PostorderTraversal(Node.Left,result)
	PostorderTraversal(Node.Right,result)
	*result=append(*result,Node.Val)
}