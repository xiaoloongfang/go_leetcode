package easy

func minMaxDifference(num int) int {
	nums := make([]int, 0)
	for num > 0 {
		mod := num % 10
		nums = append(nums, mod)
		num /= 10
	}
	least := make([]int, len(nums))
	high := -1
	x := nums[len(least)-1]
	for i := len(nums) - 1; i >= 0; i-- {

		if nums[i] == x {
			least[i] = 0
		} else {
			least[i] = nums[i]
		}
		if high == -1 && nums[i] != 9 {
			high = nums[i]
		}
		if high == nums[i] {
			nums[i] = 9
		}
	}

	max := 0
	for i := len(nums) - 1; i >= 0; i-- {
		max = max*10 + (nums[i] - least[i])
	}

	return max

}
