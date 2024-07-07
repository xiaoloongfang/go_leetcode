package main

func maximumPrimeDifference(nums []int) int {
	primeMap := make(map[int]int)
	primeMap[1] = 2
	primeMap[2] = 1
	start, end := 0, len(nums)-1
	for start < len(nums) {
		if isPrime(nums[start], primeMap) {
			break
		}
		start++
	}

	for end >= 0 {
		if isPrime(nums[end], primeMap) {
			break
		}
		end--
	}

	return end - start
}

func isPrime(n int, primeMap map[int]int) bool {
	flag := primeMap[n]
	if flag == 0 {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				primeMap[n] = 2
				return false
			}
		}
		return true
	} else {
		return flag == 1
	}
}
