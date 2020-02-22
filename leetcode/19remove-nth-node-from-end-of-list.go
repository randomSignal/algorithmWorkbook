package leetcode

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

通过次数123,752提交次数329,129

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	first := head
	prex := head
	num := 1
	totalNum := 1
	for {
		if head == nil {
			break
		}
		if num > n+1 {
			prex = prex.Next
		}

		head = head.Next
		num = num + 1

		if head != nil {
			totalNum++
		}
	}

	if totalNum == n {
		return first.Next
	}

	if prex.Next != nil {
		prex.Next = prex.Next.Next
	} else {
		prex = nil
	}

	return first
}
