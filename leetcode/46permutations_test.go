package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
通过次数91,151提交次数122,400

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestPermute(t *testing.T) {
	var res [][]int

	type Sample struct {
		ArgsA  []int
		Result [][]int
	}

	sampleList := []Sample{
		{ArgsA: []int{1, 2, 3}, Result: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{ArgsA: []int{1, 2}, Result: [][]int{{1, 2}, {2, 1}}},
	}
	for i := 0; i < len(sampleList); i++ {
		res = leetcode.Permute(sampleList[i].ArgsA)
		require.Equal(t, res, sampleList[i].Result)
	}
}
