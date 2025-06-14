package easy

func maxAdjacentDistance(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res = max(res, nums[i]-nums[(i-1+len(nums))%len(nums)], nums[(i-1+len(nums))%len(nums)]-nums[i])
	}
	return res
}
