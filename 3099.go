package main

func sumOfTheDigitsOfHarshadNumber(x int) int {
	now := x
	sum := 0
	for now > 0 {
		sum += now % 10
		now = now / 10
	}
	if x%sum == 0 {
		return sum
	}
	return -1
}
