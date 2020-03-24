package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
通过次数50,365提交次数126,027

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func TestRotateRight(t *testing.T) {
	var res *leetcode.ListNode

	type Sample struct {
		Left   []int
		K      int
		Result []int
	}

	sampleList := []Sample{
		{Left: []int{1, 2, 3, 4, 5}, K: 2, Result: []int{4, 5, 1, 2, 3}},
		{Left: []int{0, 1, 2}, K: 4, Result: []int{2, 0, 1}},
	}
	for i := 0; i < len(sampleList); i++ {
		res = leetcode.RotateRight(leetcode.ArrayToList(sampleList[i].Left), sampleList[i].K)
		require.Equal(t, leetcode.ListToArray(res), sampleList[i].Result)
	}
}
