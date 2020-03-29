package leetcode

/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
通过次数73,497提交次数95,544

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func CheckExist(args [][]int, k []int) bool {
	for i := 0; i < len(args); i++ {
		equal := true
		if len(args[i]) != len(k) {
			equal = false
			continue
		}
		for j := 0; j < len(args[i]); j++ {
			if args[i][j] != k[j] {
				equal = false
			}
		}

		if equal {
			return true
		}
	}
	return false
}

func Deduplication(args [][]int) [][]int {
	res := [][]int{}
	for i := 0; i < len(args); i++ {
		if !CheckExist(res, args[i]) {
			res = append(res, args[i])
		}
	}
	return res
}

func Sub(nums []int) [][]int {
	var res [][]int
	res = append(res, nums)
	for i := 0; i < len(nums); i++ {
		item := make([]int, len(nums))
		copy(item, nums)
		deleteRes := append(item[:i], item[i+1:]...)
		var deleteResSubSet [][]int
		if len(deleteRes) == 1 {
			deleteResSubSet = [][]int{{deleteRes[0]}}
		} else {
			deleteResSubSet = Sub(deleteRes)
		}
		res = append(res, deleteResSubSet...)
	}
	return res
}

func Subsets(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}, {}}
	}

	res := Sub(nums)

	res = Deduplication(res)
	res = append(res, []int{})
	return res
}
