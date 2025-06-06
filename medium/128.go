package medium

import "sort"

func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	res := 1
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			count++
		} else if nums[i] == nums[i-1] {
			continue
		} else {
			res = max(res, count)
			count = 1
		}
	}

	return res

}
