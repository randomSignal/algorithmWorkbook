package leetcode

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
通过次数156,269提交次数421,572

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func FindMedianSortedArray(nums1 []int) float64 {
	if len(nums1) == 0 {
		return 0
	}

	if len(nums1)%2 == 0 {
		return float64(nums1[len(nums1)/2-1]+nums1[len(nums1)/2]) / 2
	} else {
		return float64(nums1[(len(nums1)+1)/2-1])
	}
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return FindMedianSortedArray(nums2)
	}
	if len(nums2) == 0 {
		return FindMedianSortedArray(nums1)
	}

	if len(nums1) == 1 && len(nums2) == 1 {
		return float64(nums1[0]+nums2[0]) / 2
	}

	if nums2[0] <= nums1[0] && nums2[len(nums2)-1] >= nums1[len(nums1)-1] {

		if len(nums2) == 1 {
			return FindMedianSortedArrays(nums1, nil)
		} else {
			return FindMedianSortedArrays(nums1, nums2[1:len(nums2)-1])
		}

	} else if nums2[0] >= nums1[0] && nums2[len(nums2)-1] <= nums1[len(nums1)-1] {
		if len(nums1) == 1 {
			return FindMedianSortedArrays(nums2, nil)
		} else {
			return FindMedianSortedArrays(nums2, nums1[1:len(nums1)-1])
		}
	} else if nums2[0] <= nums1[0] && nums2[len(nums2)-1] <= nums1[len(nums1)-1] {
		if len(nums2) == 1 {
			return FindMedianSortedArrays(nil, nums1[:len(nums1)-1])
		}

		if len(nums1) == 1 {
			return FindMedianSortedArrays(nil, nums2[1:])
		}

		return FindMedianSortedArrays(nums1[:len(nums1)-1], nums2[1:])
	} else if nums2[0] >= nums1[0] && nums2[len(nums2)-1] >= nums1[len(nums1)-1] {
		if len(nums1) == 1 {
			return FindMedianSortedArrays(nil, nums2[:len(nums2)-1])
		}

		if len(nums2) == 1 {
			return FindMedianSortedArrays(nil, nums1[1:])
		}

		return FindMedianSortedArrays(nums1[1:], nums2[:len(nums2)-1])
	}
	return 0
}
