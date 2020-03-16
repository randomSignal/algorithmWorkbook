package leetcode

/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
通过次数84,642提交次数171,310

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func MergeTwoKLists(listA *ListNode, listB *ListNode) *ListNode {
	headerA := listA
	if listA == nil {
		return listB
	}

	if listB == nil {
		return listA
	}

	if listA.Val > listB.Val {
		return MergeTwoKLists(listB, listA)
	}

	for {
		if listA.Next == nil {
			break
		}

		if listA.Next.Val > listB.Val {
			itemB := listB
			listB = listB.Next
			itemB.Next = listA.Next
			listA.Next = itemB
		}

		if listA.Next == nil {
			break
		}
		if listB == nil {
			break
		}
		listA = listA.Next
	}
	if listA.Next == nil {
		listA.Next = listB
	}
	return headerA
}

func MergeOverTwoKLists(list *ListNode, lists []*ListNode) *ListNode {
	if len(lists) >= 2 {
		return MergeTwoKLists(list, MergeOverTwoKLists(lists[0], lists[1:]))
	} else {
		return MergeTwoKLists(list, lists[0])
	}
}

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	return MergeOverTwoKLists(lists[0], lists[1:])
}
