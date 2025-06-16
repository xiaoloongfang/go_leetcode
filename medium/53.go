package medium

func maxSubArray(nums []int) int {
	// sum, mn := make([]int, len(nums)),  make([]int, len(nums))
	sum, mn := 0, 0

	res := -1000000009
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		res = max(res, sum-mn)
		mn = min(mn, sum)
	}
	return res
}
