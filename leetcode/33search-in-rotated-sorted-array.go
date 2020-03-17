package leetcode

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:

输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
通过次数84,268提交次数230,897

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// [4,5,6,7,0,1,2]   3
func DivisionSearch(nums []int, target int, firstNum int, numRange string, startOffset int, endOffset int) int {
	if startOffset == endOffset {
		if nums[startOffset] == target {
			return startOffset
		}

		return -1
	}

	if startOffset+1 == endOffset {
		if nums[startOffset] == target {
			return startOffset
		}

		if nums[endOffset] == target {
			return endOffset
		}

		return -1
	}

	midOffset := (endOffset + startOffset) / 2
	midNum := nums[midOffset]

	if midNum == target {
		return midOffset
	}

	if numRange == "front" {
		if midNum > target {
			// 往前搜索
			return DivisionSearch(nums, target, firstNum, numRange, startOffset, midOffset)
		}

		if midNum < target && midNum > firstNum {
			// 往后搜索
			return DivisionSearch(nums, target, firstNum, numRange, midOffset, endOffset)
		}

		if midNum < target && midNum < firstNum {
			// 往前搜索
			return DivisionSearch(nums, target, firstNum, numRange, startOffset, midOffset)
		}
	} else {
		if midNum < target {
			// 往后搜索
			return DivisionSearch(nums, target, firstNum, numRange, midOffset, endOffset)
		}

		if midNum > target && midNum < firstNum {
			// 往前搜索
			return DivisionSearch(nums, target, firstNum, numRange, startOffset, midOffset)
		}

		if midNum > target && midNum > firstNum {
			// 往后搜索
			return DivisionSearch(nums, target, firstNum, numRange, midOffset, endOffset)
		}
	}

	return -1
}

func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 && target == nums[0] {
		return 0
	}

	if len(nums) == 1 && target != nums[0] {
		return -1
	}

	if nums[0] == target {
		return 0
	}

	if nums[0] > target {
		// 搜后面的数字
		return DivisionSearch(nums, target, nums[0], "behind", 0, len(nums)-1)
	} else {
		// 搜前面的数字
		return DivisionSearch(nums, target, nums[0], "front", 0, len(nums)-1)
	}
}
