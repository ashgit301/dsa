package tree
var prev *Node
func Flatten(root *Node) { //dont have to return anything cus we are modifying the tree itself
	if root == nil {
		return
	}
	Flatten(root.Right)
	Flatten(root.Left)
	root.Left = nil
	root.Right = prev
	prev = root.Left

}


// func Flatten(root *Node) { //dont have to return anything cus we are modifying the tree itself
// 	stack := {root}
// 	for stack != empty {
//        curr := stack.top()
//        if curr.Right != nil {
// 		   stack.Push{curr.Right}
// 		}
// 	   if curr.Left != nil {
// 		   stack.Push{curr.Left}
// 	    }
// 		if stack != empty {
// 			curr.Left = nil 
// 			curr.Right = stack.top()
// 		}
// 	}

// }


