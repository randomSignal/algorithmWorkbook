package leetcode

/*
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
通过次数77,562提交次数105,594

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func InsertChar(args string, offset int, char string) string {
	return args[:offset] + char + args[offset:]
}

func RemoveDuplicates(args []string) []string {
	mapData := map[string]bool{}
	for _, v := range args {
		mapData[v] = true
	}

	res := []string{}
	for k := range mapData {
		res = append(res, k)
	}
	return res
}

func GenerateParenthesisByString(list []string, n int) []string {
	if n == 0 {
		return list
	}

	var tmp []string
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i])+1; j++ {
			for k := j + 1; k < len(list[i])+2; k++ {
				res := InsertChar(list[i], j, "(")
				res = InsertChar(res, k, ")")
				tmp = append(tmp, res)
			}
		}
	}
	tmp = RemoveDuplicates(tmp)
	return GenerateParenthesisByString(tmp, n-1)
}

func GenerateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	return GenerateParenthesisByString([]string{"()"}, n-1)
}
