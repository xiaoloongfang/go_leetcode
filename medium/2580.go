package main

import "sort"

func countWays(ranges [][]int) int {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[j][1]
		} else {
			return ranges[i][0] < ranges[j][0]
		}
	})

	res := 2
	mod := 1000000007

	lastRight := ranges[0][1]
	for i := 0; i < len(ranges); i++ {
		if ranges[i][0] <= lastRight {
			lastRight = max(ranges[i][1], lastRight)
		} else {
			lastRight = ranges[i][1]
			res *= 2
			res %= mod
		}
	}

	return res

}
