package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestLongestPalindrome(t *testing.T) {
	type Sample struct {
		Args   string
		Result []string
	}

	sampleList := []Sample{
		{Args: "babad", Result: []string{"bab", "aba"}},
		//{Args: "cbbd", Result: []string{"bb"}},

		//{Args:"ac", Result:"c"},
		//{Args:"a", Result:"a"},
		//{Args:"babad", Result:"aba"},
		//{Args:"cbbd", Result:"bb"},
		//{Args:"asbadc", Result:"asba"},
		//{Args:"abccasda", Result:"asda"},
		//{Args: "abb", Result: "bb"},
		//{Args: "ccc", Result: "ccc"},
		//{Args: "abcda", Result: "abcda"},
	}

	for i := 0; i < len(sampleList); i++ {
		res := leetcode.LongestPalindrome(sampleList[i].Args)
		equal := false
		for j := 0; j < len(sampleList[i].Result); j++ {
			if sampleList[i].Result[j] == res {
				equal = true
			}
		}
		require.Equal(t, true, equal)
	}
}

func TestCheckPalindrome(t *testing.T) {
	type Sample struct {
		Args   string
		Result bool
	}

	sampleList := []Sample{
		{Args: "a", Result: true},
		{Args: "aba", Result: true},
		{Args: "abba", Result: true},
		{Args: "abcba", Result: true},
		{Args: "abccba", Result: true},

		{Args: "ab", Result: false},
		{Args: "aaba", Result: false},
		{Args: "abbac", Result: false},
		{Args: "cabcba", Result: false},
		{Args: "cabccba", Result: false},
	}

	for _, v := range sampleList {
		res := leetcode.CheckPalindrome(v.Args)
		require.Equal(t, v.Result, res)
	}
}
