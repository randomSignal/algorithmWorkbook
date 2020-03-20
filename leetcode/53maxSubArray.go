package leetcode

/*

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

通过次数183,842提交次数369,128
*/

func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	maxNum := nums[0]

	totalNum := 0
	for start := 0; start < len(nums); start++ {
		if nums[start] > maxNum {
			maxNum = nums[start]
		}

		if nums[start] < 0 {
			continue
		}
		totalNum = nums[start]
		for end := start + 1; end < len(nums); end++ {
			totalNum += nums[end]
			if nums[end] < 0 {
				continue
			}
			if totalNum > maxNum {
				maxNum = totalNum
			}
		}
	}

	return maxNum
}
