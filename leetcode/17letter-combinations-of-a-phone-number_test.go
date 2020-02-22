package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

通过次数79,685提交次数151,483

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestLetterCombinations(t *testing.T) {
	var res []string

	type Sample struct {
		Args   string
		Result []string
	}

	sampleList := []Sample{
		//{Args: "23", Result: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{Args: "2", Result: []string{"a", "b", "c"}},
	}

	for i := 0; i < len(sampleList); i++ {
		res = leetcode.LetterCombinations(sampleList[i].Args)
		require.Equal(t, sampleList[i].Result, res)
	}
}
