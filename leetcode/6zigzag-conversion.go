package leetcode

/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zigzag-conversion
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type Direction uint8

// 分析选项
const (
	Down Direction = iota
	Up
)

func getNextLocation(row int, column int, numRows int, direction Direction) (Direction, int, int) {
	if direction == Down {
		if row == numRows-1 {
			direction = Up
			row = row - 1
			column = column + 1
		} else {
			row = row + 1
		}
		return direction, row, column
	} else {
		if row == 0 {
			direction = Down
			row = row + 1
		} else {
			row = row - 1
			column = column + 1
		}
		return direction, row, column
	}
}

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var list [][]string

	for i := 0; i < numRows; i++ {
		var row []string
		for j := 0; j < len(s); j++ {
			row = append(row, "")
		}
		list = append(list, row)
	}

	row := 0
	column := 0
	direction := Down
	for i := 0; i < len(s); i++ {
		list[row][column] = s[i : i+1]
		direction, row, column = getNextLocation(row, column, numRows, direction)
	}

	res := ""
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i]); j++ {
			if list[i][j] != "" {
				res = res + list[i][j]
			}
		}
	}
	return res
}
