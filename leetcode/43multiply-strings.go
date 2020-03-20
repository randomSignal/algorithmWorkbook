package leetcode

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
链接：https://leetcode-cn.com/problems/Multiply-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var strMap = map[string]int{
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
	//"10": 10,
	//"11": 11,
	//"12": 12,
	//"13": 13,
	//"14": 14,
	//"15": 15,
	//"16": 16,
	//"17": 17,
	//"18": 18,
}

var numMap = map[int]string{
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

func Adding(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		return Adding(num2, num1)
	}

	extraFlag := 0
	res := 0
	for i := len(num1) - 1; i >= 0; i-- {
		res = 0
		if i-(len(num1)-len(num2)) >= 0 {
			res = extraFlag + strMap[num1[i:i+1]] + strMap[num2[i-(len(num1)-len(num2)):i+1-(len(num1)-len(num2))]]
		} else {
			res = extraFlag + strMap[num1[i:i+1]]
		}
		extraFlag = 0
		if res >= 10 {
			extraFlag = 1
			res = res - 10
		}
		num1 = num1[:i] + numMap[res] + num1[i+1:]
	}
	if extraFlag == 1 {
		num1 = "1" + num1
	}
	return num1
}

func Multiply(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		return Multiply(num1, num2)
	}
	if num2 == "0" || num1 == "0" {
		return "0"
	}

	ret := ""
	suffix := ""
	j := 0
	cur := ""
	for i := 0; i < len(num2); i++ {
		suffix = ""
		for j = i; j < len(num2)-1; j++ {
			suffix = suffix + "0"
		}

		num1Item := num1
		extraFlag := 0
		res := 0
		for j = len(num1Item) - 1; j >= 0; j-- {
			res = extraFlag + strMap[num1Item[j:j+1]]*strMap[num2[i:i+1]]
			extraFlag = 0
			if res >= 10 {
				extraFlag = res / 10
				res = res % 10
			}
			num1Item = num1Item[:j] + numMap[res] + num1Item[j+1:]
		}
		if extraFlag != 0 {
			num1Item = numMap[extraFlag] + num1Item
		}

		cur = num1Item + suffix
		ret = Adding(ret, cur)
	}

	return ret
}
