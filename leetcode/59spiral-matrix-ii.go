package leetcode

/*
给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
通过次数29,035提交次数37,783

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func GenerateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	if n == 1 {
		return [][]int{{1}}
	}

	var res [][]int
	for i := 0; i < n; i++ {
		var item []int
		for j := 0; j < n; j++ {
			item = append(item, 0)
		}
		res = append(res, item)
	}

	row := 0
	col := 0

	minCol := 0
	maxCol := len(res[0]) - 1

	minRow := 0
	maxRow := len(res) - 1

	direction := "right"
	for i := 0; i < n*n; i++ {
		res[row][col] = i + 1
		row, col, minRow, maxRow, minCol, maxCol, direction = GetNextOffset(row, col, minRow, maxRow, minCol, maxCol, direction)
	}

	return res
}
