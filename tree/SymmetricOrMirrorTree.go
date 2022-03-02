package tree

func Symmetric(Left, Right *Node)bool{
	if Left == nil || Right == nil {
		return false
	}else if Left.Val != Right.Val {
		return false
	}
	return Left == Right && Symmetric(Left.Left, Right.Right) && Symmetric(Left.Right, Right.Left)


}