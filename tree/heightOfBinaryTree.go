package tree

func Height(root *Node)int{
	if root == nil {
		return 0
	}
	lh := Height(root.Left)
	rh := Height(root.Right)
	return(Max(lh,rh))+1
}

func Max(a int, b int)int {
	if a > b {
		return a 
	}else if b > a {
		return b 
	}else {
		return a
	}
}