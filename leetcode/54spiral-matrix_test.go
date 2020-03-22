package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
通过次数47,260提交次数120,746

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestSpiralOrder(t *testing.T) {
	var res []int

	type Sample struct {
		Args   [][]int
		Result []int
	}

	sampleList := []Sample{
		//{Args: [][]int{{1, 2, 3 },{4, 5, 6 },{7, 8, 9}}, Result: []int{1,2,3,6,9,8,7,4,5}},
		//{Args: [][]int{{1, 2, 3, 4 },{5, 6, 7, 8},{9,10,11,12}}, Result: []int{1,2,3,4,8,12,11,10,9,5,6,7}},
		{Args: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, Result: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}},
	}
	for i := 0; i < len(sampleList); i++ {
		res = leetcode.SpiralOrder(sampleList[i].Args)
		require.Equal(t, sampleList[i].Result, res)
	}
}
