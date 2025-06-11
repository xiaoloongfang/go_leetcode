package medium

import "sort"

func threeSum(nums []int) [][]int {

	res := make([][]int, 0)

	sort.Ints(nums)

	for i := len(nums) - 1; i >= 2; i-- {
		left := 0
		right := i - 1
		n1 := 1000005
		n2 := 1000005
		n3 := nums[i]
		for left < right {
			if nums[left]+nums[right]+n3 == 0 && nums[left] != n1 && nums[right] != n2 {
				n1 = nums[left]
				n2 = nums[right]
				res = append(res, []int{n1, n2, n3})
				left++
				right--
			} else if nums[left]+nums[right]+n3 > 0 {
				n1 = nums[left]
				n2 = nums[right]
				right--
			} else if nums[left]+nums[right]+n3 < 0 {
				n1 = nums[left]
				n2 = nums[right]
				left++
			} else {
				n1 = nums[left]
				n2 = nums[right]
				left++
				right--
			}
		}
	}

	return res

}
