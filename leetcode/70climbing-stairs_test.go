package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
通过次数159,309提交次数330,966

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestClimbStairs(t *testing.T) {
	var res int

	type Sample struct {
		Args   int
		Result int
	}

	sampleList := []Sample{
		{Args: 2, Result: 2},
		{Args: 3, Result: 3},
		{Args: 4, Result: 5},
	}
	for i := 0; i < len(sampleList); i++ {
		res = leetcode.ClimbStairs(sampleList[i].Args)
		require.Equal(t, sampleList[i].Result, res)
	}
}
