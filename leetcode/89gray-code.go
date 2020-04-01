package leetcode

import (
	"math"
)

/*
格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。

给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。格雷编码序列必须以 0 开头。

示例 1:

输入: 2
输出: [0,1,3,2]
解释:
00 - 0
01 - 1
11 - 3
10 - 2

对于给定的 n，其格雷编码序列并不唯一。
例如，[0,2,3,1] 也是一个有效的格雷编码序列。

00 - 0
10 - 2
11 - 3
01 - 1
示例 2:

输入: 0
输出: [0]
解释: 我们定义格雷编码序列必须以 0 开头。
     给定编码总位数为 n 的格雷编码序列，其长度为 2n。当 n = 0 时，长度为 20 = 1。
     因此，当 n = 0 时，其格雷编码序列为 [0]。
通过次数21,652提交次数31,866

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/gray-code
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func GrayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	mapData := make(map[int][]bool)
	for i := 0; i < int(math.Pow(2, float64(n))); i++ {
		mapData[i] = Int2BoolArray(i)
	}
	var res []int
	lastKey := 0
	lastValue := mapData[0]
	res = append(res, lastKey)
	delete(mapData, lastKey)

	for {
		if len(mapData) == 0 {
			break
		}
		for key, value := range mapData {
			mi := math.Log2(math.Abs(float64(lastKey - key)))
			if float64(int(mi)) != mi {
				continue
			}
			if IsRelate(value, lastValue) {
				lastKey = key
				lastValue = value

				delete(mapData, key)
				res = append(res, key)
			}
		}
	}
	return res
}

func IsRelate(A []bool, B []bool) bool {
	if len(A) > len(B) {
		return IsRelate(B, A)
	}
	if len(A) == len(B) {
		diffNum := 0
		for i := 0; i < len(A); i++ {
			if A[i] != B[i] {
				diffNum++
			}
		}
		if diffNum >= 2 {
			return false
		} else if diffNum == 1 {
			return true
		} else {
			return false
		}
	} else if len(A)+1 == len(B) {
		diffNum := 1
		B = B[1:]
		for i := 0; i < len(A); i++ {
			if A[i] != B[i] {
				diffNum++
			}
		}
		if diffNum >= 2 {
			return false
		} else if diffNum == 1 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func Int2BoolArray(arg int) []bool {
	var res []bool
	i := 2
	for {
		if arg == 0 {
			break
		}
		if arg%i != 0 {
			res = append([]bool{true}, res...)
		} else {
			res = append([]bool{false}, res...)
		}
		arg = arg - (arg % i)
		i = i * 2
	}
	return res
}
