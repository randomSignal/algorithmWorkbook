package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
通过次数82,046提交次数189,305

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestThreeSumClosest(t *testing.T) {
	var res int

	type Sample struct {
		Args   []int
		Target int
		Result int
	}

	sampleList := []Sample{
		//{Args: []int{-1, 2, 1, -4}, Target: 1, Result: 2},
		{Args: []int{-3, -2, -5, 3, -4}, Target: -1, Result: -2},
	}

	for i := 0; i < len(sampleList); i++ {
		res = leetcode.ThreeSumClosest(sampleList[i].Args, sampleList[i].Target)
		require.Equal(t, sampleList[i].Result, res)
	}
}
