package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestAddTwoNumbers(t *testing.T) {
	res := leetcode.AddTwoNumbers(&leetcode.ListNode{
		Val: 2,
		Next: &leetcode.ListNode{
			Val: 4,
			Next: &leetcode.ListNode{
				Val: 3,
			},
		},
	}, &leetcode.ListNode{
		Val: 5,
		Next: &leetcode.ListNode{
			Val: 6,
			Next: &leetcode.ListNode{
				Val: 4,
			},
		},
	})

	eqList := &leetcode.ListNode{
		Val: 7,
		Next: &leetcode.ListNode{
			Val: 0,
			Next: &leetcode.ListNode{
				Val: 8,
			},
		},
	}

	for {
		if eqList == nil || res == nil {
			break
		}

		require.Equal(t, res.Val, eqList.Val)
		eqList = eqList.Next
		res = res.Next
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	res := leetcode.AddTwoNumbers(&leetcode.ListNode{
		Val: 2,
		Next: &leetcode.ListNode{
			Val: 4,
			Next: &leetcode.ListNode{
				Val: 5,
			},
		},
	}, &leetcode.ListNode{
		Val: 5,
		Next: &leetcode.ListNode{
			Val: 6,
			Next: &leetcode.ListNode{
				Val: 4,
				Next: &leetcode.ListNode{
					Val: 1,
				},
			},
		},
	})

	eqList := &leetcode.ListNode{
		Val: 7,
		Next: &leetcode.ListNode{
			Val: 0,
			Next: &leetcode.ListNode{
				Val: 0,
				Next: &leetcode.ListNode{
					Val: 2,
				},
			},
		},
	}

	for {
		if eqList == nil || res == nil {
			break
		}

		require.Equal(t, res.Val, eqList.Val)
		eqList = eqList.Next
		res = res.Next
	}
}

func TestAddTwoNumbers3(t *testing.T) {
	res := leetcode.AddTwoNumbers(&leetcode.ListNode{
		Val: 5,
	}, &leetcode.ListNode{
		Val: 5,
	})

	eqList := &leetcode.ListNode{
		Val: 0,
		Next: &leetcode.ListNode{
			Val: 1,
		},
	}

	for {
		if eqList == nil || res == nil {
			break
		}

		require.Equal(t, res.Val, eqList.Val)
		eqList = eqList.Next
		res = res.Next
	}
}
