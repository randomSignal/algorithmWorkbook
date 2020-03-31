package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 num1 成为一个有序数组。



说明:

    初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
    你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。



示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]

通过次数128,322
提交次数272,419

*/
func TestMerge(t *testing.T) {
	type Sample struct {
		ArgsA []int
		KA    int
		ArgsB []int
		KB    int
	}

	sampleList := []Sample{
		{ArgsA: []int{1, 2, 3, 0, 0, 0}, KA: 3, ArgsB: []int{2, 5, 6}, KB: 3},
	}

	for i := 0; i < len(sampleList); i++ {
		leetcode.Merge(sampleList[i].ArgsA, sampleList[i].KA, sampleList[i].ArgsB, sampleList[i].KB)
		require.Equal(t, sampleList[i].ArgsA, []int{1,2,2,3,5,6})
	}
}

