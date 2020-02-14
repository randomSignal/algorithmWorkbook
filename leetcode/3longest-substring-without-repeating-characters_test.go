package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestLengthOfLongestSubstring(t *testing.T) {
	var res int

	type Sample struct {
		Args   string
		Result int
	}

	sampleList := []Sample{
		{Args: "abcabcbb", Result: 3},
		{Args: "bbbbb", Result: 1},
		{Args: "pwwkew", Result: 3},
		{Args: "ppwwkeww", Result: 3},
		{Args: "ppwwkews", Result: 4},
	}
	for i := 0; i < len(sampleList); i++ {
		res = leetcode.LengthOfLongestSubstring(sampleList[i].Args)
		require.Equal(t, res, sampleList[i].Result)
	}
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	leetcode.LengthOfLongestSubstring("abc")
}
