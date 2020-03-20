package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
通过次数49,199提交次数117,296

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/multiply-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestMultiply(t *testing.T) {
	var res string

	type Sample struct {
		ArgsA  string
		ArgsB  string
		Result string
	}

	sampleList := []Sample{
		{ArgsA: "2", ArgsB: "3", Result: "6"},
		{ArgsA: "123", ArgsB: "1", Result: "123"},
		{ArgsA: "123", ArgsB: "456", Result: "56088"},
		{ArgsA: "9133", ArgsB: "0", Result: "0"},
		{ArgsA: "8432708", ArgsB: "87372547", Result: "0"},
	}
	for i := 0; i < len(sampleList); i++ {
		res = leetcode.Multiply(sampleList[i].ArgsA, sampleList[i].ArgsB)
		_ = res
		require.Equal(t, sampleList[i].Result, res)
	}
}

func TestAdding(t *testing.T) {
	var res string

	type Sample struct {
		ArgsA  string
		ArgsB  string
		Result string
	}

	sampleList := []Sample{
		{ArgsA: "20", ArgsB: "3", Result: "23"},
		{ArgsA: "2", ArgsB: "", Result: "2"},
		{ArgsA: "2", ArgsB: "3", Result: "5"},
		{ArgsA: "123", ArgsB: "456", Result: "579"},
		{ArgsA: "423", ArgsB: "656", Result: "1079"},
	}
	for i := 0; i < len(sampleList); i++ {
		res = leetcode.Adding(sampleList[i].ArgsA, sampleList[i].ArgsB)
		require.Equal(t, res, sampleList[i].Result)
	}
}
