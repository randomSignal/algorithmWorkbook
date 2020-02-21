package leetcode

import (
	"sort"
	"strconv"
)

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
通过次数160,090提交次数625,341

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	argsMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := argsMap[v]; ok {
			argsMap[v]++
		} else {
			argsMap[v] = 1
		}
	}

	if v, ok := argsMap[0]; ok {
		if v >= 3 {
			res = append(res, []int{0, 0, 0})
		}
	}

	var list []int
	for key, value := range argsMap {
		list = append(list, key)
		if value >= 2 && key != 0 {
			if _, ok := argsMap[-2*key]; ok {
				item := []int{key, key, -2 * key}
				sort.Ints(item)
				res = append(res, item)
			}
		}
	}

	resMap := make(map[string]bool)
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			value := -(list[i] + list[j])
			if value == list[i] || value == list[j] {
				continue
			}

			if _, ok := argsMap[value]; ok {
				item := []int{list[i], list[j], value}
				sort.Ints(item)
				mapKey := strconv.Itoa(item[0]) + strconv.Itoa(item[1]) + strconv.Itoa(item[2])
				if _, ok := resMap[mapKey]; !ok {
					resMap[mapKey] = true
					res = append(res, item)
				}
			}
		}
	}

	//resMap := make(map[string]bool)
	//for i := 0; i < len(nums); i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		for k := j + 1; k < len(nums); k++ {
	//			if nums[i]+nums[j]+nums[k] != 0 {
	//				continue
	//			}
	//			item := []int{nums[i], nums[j], nums[k]}
	//			sort.Ints(item)
	//			mapKey := strconv.Itoa(item[0]) + strconv.Itoa(item[1]) + strconv.Itoa(item[2])
	//			if _, ok := resMap[mapKey]; !ok {
	//				resMap[mapKey] = true
	//				res = append(res, item)
	//			}
	//		}
	//	}
	//}
	return res
}
