package easy

import "sort"

type Number struct {
	index int
}

func twoSum(nums []int, target int) []int {
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1
	res := make([]int, 2)
	for left < right {
		if nums[left]+nums[right] < target {
			left++
		} else if nums[left]+nums[right] > target {
			right--
		} else {
			res[0] = nums[left]
			res[1] = nums[right]
			break
		}
	}
	return res
}
