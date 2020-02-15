package leetcode

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

func CheckPalindrome(s string) bool {
	res := true
	for i := 0; i < (len(s) / 2); i++ {
		if s[i] != s[len(s)-1-i] {
			res = false
			break
		}
	}
	return res
}

func LongestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	maxLen := 0
	maxStr := ""
	tmp := ""
	tmpLen := 0
	for start := 0; start < len(s); start++ {
		if len(s)-start < maxLen {
			break
		}
		for end := start + 1; end <= len(s); end++ {
			tmp = s[start:end]
			tmpLen = end - start
			if tmpLen <= maxLen {
				continue
			}
			if CheckPalindrome(tmp) {
				maxStr = tmp
				maxLen = tmpLen
			}
		}
	}

	return maxStr
}
