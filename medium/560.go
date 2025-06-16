package medium

func subarraySum(nums []int, k int) int {

	if len(nums) == 1 {
		if nums[0] == k {
			return 1
		} else {
			return 0
		}
	}

	left, right := 0, 0
	sum := 0

	res := 0

	for right <= len(nums) {
		if sum == k {
			res++
			if right == len(nums) {
				break
			}
			sum -= nums[left]
			left++
			sum += nums[right]
			right++
		} else if sum < k {
			if right == len(nums) {
				break
			}
			sum += nums[right]
			right++
		} else {
			sum -= nums[left]
			left++
		}
	}
	return res

}
