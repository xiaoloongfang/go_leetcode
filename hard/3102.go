package main

import (
	"math"
	"sort"
)

type Point struct {
	x, y, index int
}

func minimumDistance(points [][]int) int {
	a, b := 0, 0
	c, d := int(1e9), int(1e9)
	check := min(100, len(points))
	ppoints := make([]Point, len(points))
	for i, point := range points {
		if point[0] > a {
			a = point[0]
		}
		if point[1] > b {
			b = point[1]
		}
		if point[0] < c {
			c = point[0]
		}
		if point[1] < d {
			d = point[1]
		}
		ppoints[i] = Point{point[0], point[1], i}
	}

	sort.Slice(ppoints, func(i, j int) bool {
		return ppoints[i].x-c+ppoints[i].y-d < ppoints[i].x-c+ppoints[i].y-d
	})

	near := ppoints[:check]

	sort.Slice(ppoints, func(i, j int) bool {
		return ppoints[i].x-c+b-ppoints[i].y < ppoints[j].x-c+b-ppoints[j].y
	})

	near = append(near, ppoints[:check]...)

	sort.Slice(ppoints, func(i, j int) bool {
		return a-ppoints[i].x+ppoints[i].y-d < a-ppoints[j].x+ppoints[j].y-d
	})

	near = append(near, ppoints[:check]...)

	sort.Slice(ppoints, func(i, j int) bool {
		return a-ppoints[i].x+b-ppoints[i].y < a-ppoints[j].x+b-ppoints[j].y
	})

	near = append(near, ppoints[:check]...)

	//for _, point := range near {
	//	println(point.x, " ", point.y, " ", point.index)
	//}

	res := int(1e9 + 1)
	for i := range near {
		nnear := make([]Point, len(near))
		copy(nnear, near)
		s := append(nnear[:i], nnear[i+1:]...)
		cal := findMax(s, near[i].index)
		if cal >= 0 {
			res = min(res, cal)
		}
	}
	return res
}

func findMax(points []Point, index int) int {
	res := -1
	for i := range points {
		if points[i].index == index {
			continue
		}
		for j := i + 1; j < len(points); j++ {
			if points[j].index == index {
				continue
			}
			res = max(res, int(math.Abs(float64(points[i].x-points[j].x)))+int(math.Abs(float64(points[i].y-points[j].y))))
		}
	}
	return res
}

func main() {
	println(minimumDistance([][]int{{3, 10}, {5, 15}, {10, 2}, {4, 4}}))
	println(minimumDistance([][]int{{1, 1}, {1, 1}, {1, 1}}))
	println(minimumDistance([][]int{{3, 2}, {3, 9}, {7, 10}, {4, 4}, {8, 10}, {2, 7}}))
}
