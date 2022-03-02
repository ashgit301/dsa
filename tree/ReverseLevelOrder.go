package tree

import (
	
)

func ReverseLevelOrder(root *Node, Result *[]int) {
	var Queue , Stack []*Node
	Queue = append(Queue, root)
	for len(Queue) != 0 {
		curr := Queue[0]
		Queue = Queue[1:]
		Stack = append(Stack, curr) //keep puching the node to stack. Print stack from last to first
		if curr.Right != nil {
			Queue = append(Queue, curr.Right)
		}
		if curr.Left != nil {
			Queue = append(Queue, curr.Left)
		}
	}
	for i := len(Stack)-1 ; i >= 0 ; i-- {
		node := Stack[i]
		*Result = append(*Result, node.Val)
	}
}