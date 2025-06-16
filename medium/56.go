package medium

import "sort"

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return intervals[i][0] < intervals[j][0]
		}
	})

	res := make([][]int, 0)

	left, right := intervals[0][0], intervals[0][1]
	for _, interval := range intervals {
		if interval[0] <= right {
			right = max(right, interval[1])
		} else {
			res = append(res, []int{left, right})
			left, right = interval[0], interval[1]
		}
	}

	res = append(res, []int{left, right})

	return res

}
