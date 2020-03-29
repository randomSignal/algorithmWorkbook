package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
通过次数73,497提交次数95,544

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestSubsets(t *testing.T) {
	var res [][]int

	type Sample struct {
		Args   []int
		Result [][]int
	}

	sampleList := []Sample{
		{Args: []int{1, 2, 3}, Result: [][]int{{3}, {1}, {2}, {1, 2, 3}, {1, 3}, {1, 2}, {2, 3}}},
		//{Args: []int{2,3,4,5,6,7,8,10,0}, Result: [][]int{{3}, {1}, {2}, {1, 2, 3}, {1, 3}, {1, 2}, {2, 3}}},
	}
	for i := 0; i < len(sampleList); i++ {
		res = leetcode.Subsets(sampleList[i].Args)
		require.Equal(t, sampleList[i].Result, res)
	}
}
