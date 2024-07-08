package main

func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	left, right := 0, sum
	for i := 0; i < len(nums); i++ {
		right -= nums[i]
		if left == right {
			return i
		}
		left += nums[i]
	}
	return -1
}

func main() {

}
