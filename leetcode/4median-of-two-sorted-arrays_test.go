package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
通过次数156,269提交次数421,572

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestFindMedianSortedArrays(t *testing.T) {
	type Sample struct {
		ArgA   []int
		ArgB   []int
		Result float64
	}

	sampleList := []Sample{
		//{ArgA: []int{1, 3}, ArgB: []int{2}, Result: 2},
		{ArgA: []int{1, 2}, ArgB: []int{3, 4}, Result: 2.5},
	}

	for _, v := range sampleList {
		res := leetcode.FindMedianSortedArrays(v.ArgA, v.ArgB)
		require.Equal(t, v.Result, res)
	}
}
