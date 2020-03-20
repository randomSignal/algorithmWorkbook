package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

/*

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

通过次数183,842提交次数369,128


*/

func TestMaxSubArray(t *testing.T) {
	var res int

	type Sample struct {
		Args   []int
		Result int
	}

	sampleList := []Sample{
		{Args: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, Result: 6},
		{Args: []int{-2, -3, 4, -1}, Result: 4},
		{Args: []int{-2, 1}, Result: 1},
	}
	for i := 0; i < len(sampleList); i++ {
		res = leetcode.MaxSubArray(sampleList[i].Args)
		require.Equal(t, sampleList[i].Result, res)
	}
}
