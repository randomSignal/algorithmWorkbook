package leetcode

import "fmt"

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

func GetNextOffset(row, col, minRow, maxRow, minCol, maxCol int, direction string) (int, int, int, int, int, int, string) {
	if direction == "right" {
		if col == maxCol {
			minRow++
			return row + 1, col, minRow, maxRow, minCol, maxCol, "down"
		} else {
			return row, col + 1, minRow, maxRow, minCol, maxCol, direction
		}

	} else if direction == "down" {
		if row == maxRow {
			maxCol--
			return row, col - 1, minRow, maxRow, minCol, maxCol, "left"
		} else {
			return row + 1, col, minRow, maxRow, minCol, maxCol, direction
		}
	} else if direction == "left" {
		if col == minCol {
			maxRow--
			return row - 1, col, minRow, maxRow, minCol, maxCol, "up"
		} else {
			return row, col - 1, minRow, maxRow, minCol, maxCol, direction
		}
	} else {
		if row == minRow {
			minCol++
			return row, col + 1, minRow, maxRow, minCol, maxCol, "right"
		} else {
			return row - 1, col, minRow, maxRow, minCol, maxCol, direction
		}
	}
}

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	if len(matrix) == 1 {
		return matrix[0]
	}

	m := len(matrix)
	n := len(matrix[0])

	var res []int
	row := 0
	col := 0

	minCol := 0
	maxCol := len(matrix[0]) - 1

	minRow := 0
	maxRow := len(matrix) - 1

	direction := "right"
	for i := 0; i < m*n; i++ {
		fmt.Println(matrix[row][col])
		res = append(res, matrix[row][col])
		row, col, minRow, maxRow, minCol, maxCol, direction = GetNextOffset(row, col, minRow, maxRow, minCol, maxCol, direction)
	}
	return res
}
