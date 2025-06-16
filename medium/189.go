package medium

func rotate(nums []int, k int) {
	tmp := make([]int, len(nums))
	for i := 0; i < len(tmp); i++ {
		tmp[(i+k)%len(tmp)] = nums[i]
	}
	for i := 0; i < len(tmp); i++ {
		nums[i] = tmp[i]
	}

}
