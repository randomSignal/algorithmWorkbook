package leetcode

import (
	"math"
	"strconv"
)

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func ReverseString(s string) string {
	if s == "" {
		return ""
	}
	isNegative := false
	if s[0:1] == "-" {
		isNegative = true
		s = s[1:]
	}

	if s == "" {
		return ""
	}

	res := ""

	for i := 0; i < len(s); i++ {
		res = s[i:i+1] + res
	}

	result := ""
	for i := 0; i < len(res); i++ {
		if res[i:i+1] != "0" {
			result = res[i:]
			break
		}
	}
	if isNegative {
		return "-" + result
	} else {
		return result
	}
}

func Reverse(x int) int {
	s := strconv.Itoa(x)
	reverseStr := ReverseString(s)

	res, err := strconv.Atoi(reverseStr)
	if err != nil {
		return 0
	}
	if float64(res) > math.Pow(float64(2), float64(31))-1 {
		return 0
	}
	if float64(res) < -math.Pow(float64(2), float64(31)) {
		return 0
	}
	return res
}
