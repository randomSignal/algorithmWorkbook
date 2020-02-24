package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

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

func ListEqual(left *leetcode.ListNode, right *leetcode.ListNode) bool {
	for {
		if left == nil || right == nil {
			if right != nil || left != nil {
				return false
			}
			return true
		}

		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
}

func TestMergeTwoLists(t *testing.T) {
	var res *leetcode.ListNode

	type Sample struct {
		Left   []int
		Right  []int
		Result []int
	}

	sampleList := []Sample{
		{Left: []int{1, 2, 4}, Right: []int{1, 3, 4}, Result: []int{1, 1, 2, 3, 4, 4}},
	}
	for i := 0; i < len(sampleList); i++ {
		res = leetcode.MergeTwoLists(leetcode.ArrayToList(sampleList[i].Left), leetcode.ArrayToList(sampleList[i].Right))
		require.Equal(t, leetcode.ListToArray(res), sampleList[i].Result)
	}

}
