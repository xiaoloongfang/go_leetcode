package main

import "sort"

func numberGame(nums []int) []int {
	sort.Ints(nums)
	res := make([]int, len(nums))
	for i := 0; i < len(nums)/2; i++ {
		res[2*i] = nums[2*i+1]
		res[2*i+1] = nums[2*i]
	}
	return res
}

func main() {

}
