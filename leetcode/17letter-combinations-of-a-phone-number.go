package leetcode

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

var mapLetter = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func GetTwoLetterCombinations(args [][]string) []string {
	if len(args) == 0 {
		return []string{}
	} else if len(args) == 1 {
		return args[0]
	} else if len(args) == 2 {
		var res []string
		for i := 0; i < len(args[0]); i++ {
			for j := 0; j < len(args[1]); j++ {
				res = append(res, args[0][i]+args[1][j])
			}
		}
		return res
	} else {
		var res []string
		for i := 0; i < len(args[0]); i++ {
			for j := 0; j < len(args[1]); j++ {
				res = append(res, args[0][i]+args[1][j])
			}
		}
		var tmpArgs [][]string
		tmpArgs = append(tmpArgs, res)
		tmpArgs = append(tmpArgs, args[2:]...)
		return GetTwoLetterCombinations(tmpArgs)
	}
}

func LetterCombinations(digits string) []string {
	var tmpString [][]string
	for i := 0; i < len(digits); i++ {
		tmpString = append(tmpString, mapLetter[digits[i:i+1]])
	}

	return GetTwoLetterCombinations(tmpString)
}
