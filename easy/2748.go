package main

func countBeautifulPairs(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			a := nums[i]
			for a >= 10 {
				a /= 10
			}
			b := nums[j] % 10
			if gcd(a, b) == 1 {
				result++
			}
		}
	}
	return result
}

func gcd(a, b int) int {
	if a < b {
		return gcd(b, a)
	}
	mod := a % b
	if mod == 0 {
		return b
	}
	return gcd(b, mod)
}

func main() {

}
