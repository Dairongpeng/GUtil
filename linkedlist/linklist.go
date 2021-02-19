package linkedlist

// 链表节点结构体
type ListNode struct {
	Val int
	// 指向ListNode类型的指针Next。
	Next *ListNode
}

// 1、检查链表是否有环。全部操作指针。
func hasCycle(head *ListNode) bool {

	// 如果链表小于等于两个节点。一定不存在环
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	// for循环在go里面由三个表达式e1, e2, e3组成，e1在循环开始之前调用，e3在每轮循环结束之时调用。
	// 如果省略e1和e3就是while循环的功能
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

// 2、传入链表头结点的地址，翻转链表
func ReverseLinkedList(head *ListNode) *ListNode {

	// 注意：new开辟了内存空间，会给各个变量0值。这里需要的是不初始化的nil值
	//pre := new(ListNode)
	var pre *ListNode
	var next *ListNode

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre

}

// 3、移除链表中等于值的节点
func RemoveValue(head *ListNode, num int) *ListNode {

	// 从链表的头开始，舍弃掉开头的且连续的等于num的节点
	for head != nil {
		if head.Val != num {
			break
		}

		head = head.Next
	}

	// head来到 第一个不需要删的位置
	pre := head
	cur := head

	// 快慢指针
	for cur != nil {
		if cur.Val == num { // 快指针cur向下滑动，如果值等于num，则暂时把下一个节点给慢指针的下一个指向。从而跳过等于num的节点
			pre.Next = cur.Next
		} else { // cur此时到了不等于num的节点，则慢指针追赶上去。达到的效果就是等于num的节点都被删掉了
			pre = cur
		}

		// 快指针向下滑动
		cur = cur.Next
	}

	return head

}
