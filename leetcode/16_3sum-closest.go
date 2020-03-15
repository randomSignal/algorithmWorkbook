package leetcode

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
通过次数82,046提交次数189,305

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func ThreeSumClosest(nums []int, target int) int {
	if len(nums) <= 2 {
		return 0
	}
	closestNum := nums[0] + nums[1] + nums[2]
	closestOffset := closestNum - target
	if closestOffset < 0 {
		closestOffset = -closestOffset
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				currentNum := nums[i] + nums[j] + nums[k]
				offset := target - currentNum
				if offset < 0 {
					offset = -offset
				}
				if offset < closestOffset {
					closestOffset = offset
					closestNum = currentNum
				}
				if offset == 0 {
					return closestNum
				}
			}
		}
	}
	return closestNum
}
