package medium

func productExceptSelf(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	forward, backward := make([]int, len(nums)), make([]int, len(nums))
	forward[0] = nums[0]
	backward[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		forward[i] = forward[i-1] * nums[i]
		backward[len(nums)-1-i] = backward[len(nums)-i] * nums[len(nums)-i-1]
	}

	res := make([]int, len(nums))
	res[0] = backward[1]
	res[len(nums)-1] = forward[len(nums)-2]

	for i := 1; i < len(nums)-1; i++ {
		res[i] = forward[i-1] * backward[i+1]
	}

	return res

}
