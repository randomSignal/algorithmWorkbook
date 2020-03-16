package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

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

func TestMergeKLists(t *testing.T) {
	var res *leetcode.ListNode

	type Sample struct {
		Args   [][]int
		Result []int
	}
	//1->1->2->3->4->4->5->6

	sampleList := []Sample{
		{Args: [][]int{{-1, 5, 11}, {6, 10}}, Result: []int{-1, 5, 6, 10, 11}},
		//{Args: [][]int{{1, 4, 5},{1, 3, 4},{2, 6}}, Result: []int{1, 1, 2, 3, 4, 4,5,6}},
		//{Args: [][]int{{},{-1,5,11},{},{6,10}}, Result: []int{-1,5,6,10,11}},
	}
	for i := 0; i < len(sampleList); i++ {
		var list []*leetcode.ListNode
		for _, v := range sampleList[i].Args {
			list = append(list, leetcode.ArrayToList(v))
		}

		res = leetcode.MergeKLists(list)
		require.Equal(t, leetcode.ListToArray(res), sampleList[i].Result)
	}

}
