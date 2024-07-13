package main

import "math/bits"

func canSortArray(nums []int) bool {
	start, end := 0, 1
	seenMax := nums[0]
	rmin := make([]int, len(nums))
	rmin[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		rmin[i] = min(rmin[i+1], nums[i])
	}
	for end < len(nums) {
		if calOntBitNum(nums[end]) != calOntBitNum(nums[start]) {
			if seenMax > rmin[end] {
				return false
			}
			start = end
			seenMax = nums[start]
		} else {
			seenMax = max(seenMax, nums[end])
		}
		end++
	}
	return true
}

func calOntBitNum(num int) int {
	return bits.OnesCount(uint(num))
}

func main() {

}
