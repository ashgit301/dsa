package linkedList

type Node struct {
	Val int
	Next *Node
 }
 func NewNode(value int, next *Node) *Node{
	var n Node
	n.Val = value
	n.Next = next
	return &n
 }