package leetcode

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
通过次数91,151提交次数122,400

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func SplitArray(nums []int, offset int) (int, []int) {
	var resArray []int
	for i := 0; i < len(nums); i++ {
		if i == offset {
			continue
		}
		resArray = append(resArray, nums[i])
	}
	return nums[offset], resArray
}

func Permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	var res [][]int

	for i := 0; i < len(nums); i++ {
		intVar, arrayVar := SplitArray(nums, i)
		listList := Permute(arrayVar)

		for _, v := range listList {
			var item []int
			item = append(item, intVar)
			item = append(item, v...)
			res = append(res, item)
		}
	}
	return res
}
