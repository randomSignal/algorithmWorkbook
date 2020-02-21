package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

通过次数189,644提交次数523,526

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestLongestCommonPrefix(t *testing.T) {
	var res string

	type Sample struct {
		Args   []string
		Result string
	}

	sampleList := []Sample{
		{Args: []string{"flower", "flow", "flight"}, Result: "fl"},
		{Args: []string{"dog", "racecar", "car"}, Result: ""},
		{Args: []string{"dog", "dogracecar", ""}, Result: ""},
		{Args: []string{"dog1", "dogracecar", "dog"}, Result: "dog"},
	}

	for i := 0; i < len(sampleList); i++ {
		res = leetcode.LongestCommonPrefix(sampleList[i].Args)
		require.Equal(t, sampleList[i].Result, res)
	}
}
