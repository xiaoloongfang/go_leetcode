package main

func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 && target == 0 {
		return 1
	} else if len(nums) == 0 {
		return 0
	}
	return findTargetSumWays(nums[1:], target-nums[0]) + findTargetSumWays(nums[1:], target+nums[0])
}

func main() {

}
