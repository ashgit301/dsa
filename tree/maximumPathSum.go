package tree
var maxi int
// func MaxPathSum(root *Node, sum int) int {
// 	if root == nil {
// 		return 0
// 	}
// 	sum = sum + root.Val
// 	maxi = Max(maxi, sum)
// 	MaxPathSum(root.Left,sum)
// 	MaxPathSum(root.Right,sum)
// 	return maxi
// }


func MaxPathSum(root *Node) int{
	if root == nil {
		return 0
	}
	lsum := MaxPathSum(root.Left)
	rsum := MaxPathSum(root.Right)
	maxi = max(maxi,lsum+rsum+root.Val)
	return lsum+rsum+root.Val
}