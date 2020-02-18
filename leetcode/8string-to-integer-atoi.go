package leetcode

import "math"

/*
请你来实现一个 atoi 函数，使其能将字符串转换成整数。

首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。

当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。

该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。

注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。

在任何情况下，若函数不能进行有效的转换时，请返回 0。

说明：

假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。

示例 1:

输入: "42"
输出: 42
示例 2:

输入: "   -42"
输出: -42
解释: 第一个非空白字符为 '-', 它是一个负号。
     我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。
示例 3:

输入: "4193 with words"
输出: 4193
解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字。
示例 4:

输入: "words and 987"
输出: 0
解释: 第一个非空字符是 'w', 但它不是数字或正、负号。
     因此无法执行有效的转换。
示例 5:

输入: "-91283472332"
输出: -2147483648
解释: 数字 "-91283472332" 超过 32 位有符号整数范围。
     因此返回 INT_MIN (−231) 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/string-to-integer-atoi
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var mapStrToInt = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

var mapIntToString = map[int]string{
	0: "0",
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
	9: "9",
}

func IsStr(str string) bool {
	if _, ok := mapStrToInt[str]; ok {
		return false
	}

	if str == "-" {
		return false
	}
	return true
}

func Itoa(args int) string {
	res := ""
	for {
		if args < 10 {
			res = mapIntToString[args] + res
			break
		}
		res = mapIntToString[args%10] + res
		args = (args - (args % 10)) / 10
	}
	return res
}

func MyAtoi(str string) int {
	/*
		{Args: "42", Result: 42},
		{Args: "   -42", Result: -42},
		{Args: "4193 with words", Result: 4193},
		{Args: "words and 987", Result: 0},
		{Args: "-91283472332", Result: 2147483648},
	*/

	isGetData := false
	start := 0
	end := 0
	isNegative := false
	// 去掉左边的空格
	for i := 0; i < len(str); i++ {
		if str[i:i+1] == " " {
			if isGetData {
				break
			}
			if i+1 < len(str) {
				start = i + 1
				end = i + 1
			}
			continue
		}

		if str[i:i+1] == "+" {
			if isGetData {
				break
			}
			isGetData = true
			if i-1 >= 0 && str[i-1:i] == "-" {
				return 0
			}
			if i+1 < len(str) {
				start = i + 1
				end = i + 1
			}
			continue
		}

		if IsStr(str[i : i+1]) {
			break
		}

		if str[i:i+1] == "-" {
			if isGetData {
				break
			}
			isGetData = true
			if i-1 >= 0 && str[i-1:i] == "+" {
				return 0
			}
			isNegative = true
			if i+1 < len(str) {
				start = i + 1
				end = i + 1
			}
			continue
		}
		isGetData = true
		if i+1 <= len(str) {
			end = i + 1
		}
	}

	tmpStr := str[start:end]

	isGetNum := false
	start = 0
	// 去除坐标的0
	for i := 0; i < len(tmpStr); i++ {
		if tmpStr[i:i+1] == "0" && isGetNum == false {
			if i+1 < len(tmpStr) {
				start = i + 1
			}
			continue
		}
		isGetNum = true
	}
	tmpStr = tmpStr[start:]

	/*
		判断是否大于max or min
	*/
	if isNegative == false {
		if len(tmpStr) > len(Itoa(math.MaxInt32)) {
			return math.MaxInt32
		}

		if len(tmpStr) == len(Itoa(math.MaxInt32)) {
			if tmpStr > Itoa(math.MaxInt32) {
				return math.MaxInt32
			}
		}
	} else {
		if len(tmpStr) > len(Itoa(-math.MinInt32)) {
			return math.MinInt32
		}

		if len(tmpStr) == len(Itoa(-math.MinInt32)) && tmpStr > Itoa(-math.MinInt32) {
			return math.MinInt32
		}
	}

	d := 1
	res := 0
	for i := len(tmpStr); i >= 1; i-- {
		res += mapStrToInt[tmpStr[i-1:i]] * d
		d = d * 10
	}

	if isNegative {
		res = -res
	}

	return res
}
