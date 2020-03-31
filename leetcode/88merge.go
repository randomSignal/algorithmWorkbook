package leetcode

import "sort"

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 num1 成为一个有序数组。

 

说明:

    初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
    你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

 

示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]

通过次数128,322
提交次数272,419
*/
func Merge(nums1 []int, m int, nums2 []int, n int)  {
    offset := 0
    for i := 0; i < len(nums1); i++ {
	if offset >= n {
            break
	}
        if nums1[i] == 0 {
            nums1[i] = nums2[offset]
	    offset ++
	}
    }
    sort.Ints(nums1)
}
