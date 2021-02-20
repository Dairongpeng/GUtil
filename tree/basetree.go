package tree

import (
	"fmt"
)

// 儿叉树节点类型
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 1、二叉树递归先序遍历
func Pre(head *Node) {
	if head == nil {
		return
	}
	// 这里直接用地址点值也是可以的，go内部做了处理，正常的应该是先对地址解引用成结构体类型，再点操作结构体具体的值
	// fmt.Printf("%d", head.Val)
	fmt.Printf("%d", (*head).Val)

	Pre(head.Left)
	Pre(head.Right)
}

// 2、二叉树中序遍历
func Mid(head *Node) {

	if head == nil {
		return
	}

	Mid(head.Left)
	fmt.Printf("%d", head.Val)
	Mid(head.Right)

}

// 3、二叉树后序遍历
func Pos(head *Node) {
	if head == nil {
		return
	}
	Pos(head.Left)
	Pos(head.Right)
	fmt.Printf("%d", head.Val)
}

// 4、非递归先序
func NotRPre(head *Node) {

	fmt.Printf("pre-order: ")

	if head != nil {

		// 用切片制作一个栈
		var stack []*Node
		stack = append(stack, head)

		for len(stack) != 0 {
			// 栈取出最近添加的数据。例如[1,5,7,2] ，len = 4
			x := stack[len(stack)-1] // 2
			// 切掉最近添加的数据，上一步和这一步模仿栈的pop。
			stack = stack[:len(stack)-1] // [1,5,7]

			fmt.Printf("%d", x.Val)

			// 右孩子不为空，先压入右孩子。右孩子就会后弹出
			if x.Right != nil {
				stack = append(stack, x.Right)
			}

			// 左孩子不为空，压入左孩子，左孩子在右孩子之后压栈，先弹出
			if x.Left != nil {
				stack = append(stack, x.Left)
			}

		}

	}

}

// 5、非递归中序
func NotRMid(head *Node) {

	fmt.Printf("mid-order: ")

	if head != nil {
		var stack []*Node
		for len(stack) != 0 || head != nil {
			// 整条左边界以此入栈
			if head != nil {
				stack = append(stack, head)
				// head滑到自己的左孩子，左孩子有可能为空，但空的节点不会加入栈，下一个分支会判空处理
				head = head.Left
				// 左边界到头弹出一个打印，来到该节点右节点，再把该节点的左树以此进栈
			} else { // head为空的情况，栈顶是上次头结点的现场，head等于栈顶，回到上一个现场。打印后，head往右树上滑动
				head = stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				fmt.Printf("%d", (*head).Val)
				head = head.Right
			}
		}
	}

}

// 6、按层遍历
func Level(head *Node) {

	if head == nil {
		return
	}

	// 用切片模仿队列
	var queue []*Node
	queue = append(queue, head)

	for len(queue) != 0 {
		// 队头弹出，再把队头切掉，模仿队列的poll操作
		cur := queue[0]
		queue = queue[1:]

		fmt.Printf("%d", (*cur).Val)

		// 当前节点有左孩子，加入左孩子进队列
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}

		// 当前节点有右孩子，加入右孩子进队列
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

}

// 7、二叉树先序序列化
func PreSerial(head *Node) []*Node {
	var queue []*Node
	pres(head, queue)
	return queue
}

func pres(head *Node, queue []*Node) {
	if head == nil {
		queue = append(queue, nil)
	} else {
		queue = append(queue, head)
		pres(head.Left, queue)
		pres(head.Right, queue)
	}
}
