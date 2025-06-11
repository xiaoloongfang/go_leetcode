package main

func trap(height []int) int {

	n := len(height)

	now := 0
	left := make([]int, n)
	right := make([]int, n)

	for i := 0; i < n; i++ {
		if height[i] >= now {
			now = height[i]
			left[i] = 0
		} else if height[i] < now {
			left[i] = now - height[i]
		}
	}

	now = 0
	for i := n - 1; i >= 0; i-- {
		if height[i] >= now {
			now = height[i]
			right[i] = 0
		} else if height[i] < now {
			right[i] = now - height[i]
		}
	}

	res := 0
	for i := 0; i < len(left); i++ {
		res += min(left[i], right[i])
	}

	return res

}

func main() {
	println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
