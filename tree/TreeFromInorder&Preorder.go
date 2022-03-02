package tree

func PreOrder_Inorder(Preorder, Inroder []int)*Node{
	curr := Preorder[0]
	tmp := 0 
	for i := 0 ; i < len(Inroder) ; i++ {
		if curr == Inroder[i]{
			tmp = i
		}
	}
	return  &Node{
		Val:   curr,
		Left:  PreOrder_Inorder(Preorder[1:tmp+1], Inroder[:tmp]),
		Right: PreOrder_Inorder(Preorder[tmp+1:], Inroder[tmp+1:]),
	}
}