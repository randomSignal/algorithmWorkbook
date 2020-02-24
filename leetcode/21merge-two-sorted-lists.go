package leetcode

import "sort"

/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
通过次数187,577提交次数313,418

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func ArrayToList(list []int) *ListNode {
	var head *ListNode
	var last *ListNode
	for i := 0; i < len(list); i++ {
		node := ListNode{Val: list[i]}
		if head == nil {
			head = &node
			last = head
			continue
		}
		last.Next = &node
		last = last.Next
	}
	return head
}

func ListToArray(list *ListNode) []int {
	var res []int
	for {
		if list == nil {
			break
		}

		res = append(res, list.Val)
		list = list.Next
	}
	return res
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	tmpArray := append(ListToArray(l1), ListToArray(l2)...)
	sort.Ints(tmpArray)
	return ArrayToList(tmpArray)
}
