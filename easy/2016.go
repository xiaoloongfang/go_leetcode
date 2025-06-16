package easy

func maximumDifference(nums []int) int {
	res, mn := 0, 1000000007
	for _, num := range nums {
		mn = min(num, mn)
		res = max(res, num-mn)
	}
	if res == 0 {
		res = -1
	}
	return res
}
