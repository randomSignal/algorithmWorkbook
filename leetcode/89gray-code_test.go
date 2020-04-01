package leetcode_test

import (
	"fmt"
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
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

func TestGrayCode(t *testing.T) {
	var res []int

	type Sample struct {
		Args   int
		Result []int
	}

	sampleList := []Sample{
		//{Args: 2, Result: []int{0, 2, 3, 1}},
		{Args: 3, Result: []int{0, 1, 3, 2, 6, 7, 5, 4}},
		{Args: 4, Result: []int{0, 1, 3, 2, 6, 7, 5, 4}},
		//{Args: 0, Result: []int{0}},
	}

	for i := 0; i < len(sampleList); i++ {
		res = leetcode.GrayCode(sampleList[i].Args)
		fmt.Println(res)
		require.Equal(t, sampleList[i].Result, res)
	}
}
