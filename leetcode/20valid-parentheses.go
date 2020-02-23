package leetcode

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
通过次数208,110提交次数508,766

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var charMap = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
}

func MiddleEqual(left string, right string) bool {
	if len(left) != len(right) {
		return false
	}

	for i := 0; i < len(left); i++ {
		key := left[i : i+1]
		value := right[len(left)-i-1 : len(left)-i]

		if v, ok := charMap[key]; ok {
			if v != value {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func DeleteEqualString(args string) (string, bool) {
	if args == "" {
		return "", true
	}
	if len(args)%2 != 0 {
		return "", false
	}

	offset := 0
	hasEqual := false
	for i := 0; i < len(args); i++ {
		key := args[i : i+1]
		if i+1 >= len(args) {
			return "", false
		}
		value := args[i+1 : i+2]

		if v, ok := charMap[key]; ok {
			if v == value {
				hasEqual = true
				offset = i
				break
			}
		} else {
			return "", false
		}
	}

	if hasEqual {
		leftString := args[:offset]
		rightString := args[offset+2:]
		str := leftString + rightString
		return DeleteEqualString(str)
	}
	return "", false
}

func IsValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}

	_, ok := DeleteEqualString(s)
	if ok {
		return true
	} else {
		return false
	}
}
