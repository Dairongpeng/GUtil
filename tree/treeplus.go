package tree

import "math"

// 1、判断二叉树是否是平衡的
func IsBalanced(head *Node) bool {

	return IsBalancedProcess(head).IsBalanced

}

// 递归信息体
type BalancedInfo struct {
	IsBalanced bool
	Height     int
}

// 1)递归调用，head传入整体需要返回一个信息
// 2)解决当前节点的Info信息怎么得来
func IsBalancedProcess(head *Node) BalancedInfo {

	// base case
	if head == nil {
		return BalancedInfo{true, 0}
	}

	leftInfo := IsBalancedProcess(head.Left)
	rightInfo := IsBalancedProcess(head.Right)

	// 当前节点的高度，等于左右树最大的高度，加上当前节点高度1
	currentHeight := math.Max(float64(leftInfo.Height), float64(rightInfo.Height))
	isBalance := true

	if !leftInfo.IsBalanced || !rightInfo.IsBalanced || math.Abs(float64(leftInfo.Height)-float64(rightInfo.Height)) > 1 {
		isBalance = false
	}

	return BalancedInfo{isBalance, int(currentHeight)}

}

// 2、求二叉树中任意两个节点的最大距离
func MaxDistance(head *Node) int {

	return MaxDistanceProcess(head).MaxDistance

}

// 最大距离递归信息体
type MaxDistanceInfo struct {
	MaxDistance int
	height      int
}

func MaxDistanceProcess(head *Node) MaxDistanceInfo {

	if head == nil {
		return MaxDistanceInfo{0, 0}
	}

	leftInfo := MaxDistanceProcess(head.Left)
	rightInfo := MaxDistanceProcess(head.Right)

	// 当前节点为头的情况下，高度等于左右树较大的高度，加上1
	height := math.Max(float64(leftInfo.height), float64(rightInfo.height))

	// 当前节点为头的情况下，最大距离等于，左右树距离较大的那个距离(与当前节点无关的情况)
	// 和左右树高度相加再加上当前节点距离1的距离(与当前节点有关的情况)取这两种情况较大的那个

	maxDistance :=
		math.Max(math.Max(float64(leftInfo.MaxDistance), float64(rightInfo.MaxDistance)), float64(leftInfo.height+rightInfo.height+1))
	return MaxDistanceInfo{int(maxDistance), int(height)}
}

// 3、判断一颗二叉树是不是满二叉树
func IsFull(head *Node) bool {

	if head == nil {
		return true
	}

	all := IsFullProcess(head)

	// 递归拿到整棵树的头结点个数，根据满二叉树的公式，验证。（高度乘以2） - 1 = 节点个数
	return (1<<all.Height)-1 == all.Nodes

}

// 判断满二叉树的结构体
type IsFullInfo struct {
	Height int
	Nodes  int
}

func IsFullProcess(head *Node) IsFullInfo {

	if head == nil {
		return IsFullInfo{0, 0}
	}

	leftInfo := IsFullProcess(head.Left)
	rightInfo := IsFullProcess(head.Right)

	// 当前节点为头的树，高度
	height := math.Max(float64(leftInfo.Height), float64(rightInfo.Height))

	// 当前节点为头的树，节点数量
	nodes := leftInfo.Nodes + rightInfo.Nodes

	return IsFullInfo{int(height), nodes}

}
