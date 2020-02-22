package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

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

func TestRemoveNthFromEnd(t *testing.T) {
	var res *leetcode.ListNode

	type Sample struct {
		Args   *leetcode.ListNode
		N      int
		Result *leetcode.ListNode
	}

	sampleList := []Sample{
		{
			Args: &leetcode.ListNode{
				Val: 1,
				Next: &leetcode.ListNode{
					Val: 2,
					Next: &leetcode.ListNode{
						Val: 3,
						Next: &leetcode.ListNode{
							Val: 4,
							Next: &leetcode.ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			N: 1,
			Result: &leetcode.ListNode{
				Val: 1,
				Next: &leetcode.ListNode{
					Val: 2,
					Next: &leetcode.ListNode{
						Val: 3,
						Next: &leetcode.ListNode{
							Val: 4,
						},
					},
				},
			},
		},
		{
			Args:   &leetcode.ListNode{Val: 1},
			N:      1,
			Result: nil,
		},

		{
			Args:   &leetcode.ListNode{Val: 1, Next: &leetcode.ListNode{Val: 2}},
			N:      1,
			Result: &leetcode.ListNode{Val: 1},
		},
		{
			Args:   &leetcode.ListNode{Val: 1, Next: &leetcode.ListNode{Val: 2}},
			N:      2,
			Result: &leetcode.ListNode{Val: 2},
		},
	}

	for i := 0; i < len(sampleList); i++ {
		res = leetcode.RemoveNthFromEnd(sampleList[i].Args, sampleList[i].N)
		for {
			if res == nil {
				break
			}
			if sampleList[i].Result == nil {
				require.Equal(t, sampleList[i].Result, res)
			} else {
				require.Equal(t, sampleList[i].Result.Val, res.Val)
			}

			res = res.Next
			sampleList[i].Result = sampleList[i].Result.Next
		}
	}

}
