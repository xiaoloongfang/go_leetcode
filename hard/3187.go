package main

func countOfPeaks(nums []int, queries [][]int) []int {
	n := len(nums)
	treeL := 1
	for treeL < 2*n+2 {
		treeL *= 2
	}
	tree := make([]int, treeL)
	for i := 0; i < n; i++ {
		if i == 0 || i == n-1 {
			continue
		}
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			build(tree, 0, n, i, 0, 1)
		}
	}
	res := make([]int, 0)
	for _, q := range queries {
		printTree(tree)
		if q[0] == 1 {
			l, r := q[1], q[2]
			res = append(res, find(tree, 0, n, l+1, r-1, 0))
		} else {
			index, val := q[1], q[2]
			if index >= 2 {
				if nums[index-1] > nums[index-2] && nums[index-1] > nums[index] {
					if nums[index-1] < val {
						build(tree, 0, n, index-1, 0, -1)
					}
				} else {
					if nums[index-1] > nums[index-2] && nums[index-1] > val {
						build(tree, 0, n, index-1, 0, 1)
					}
				}
			}
			if index < n-2 {
				if nums[index+1] > nums[index+2] && nums[index+1] > nums[index] {
					if nums[index+1] < val {
						build(tree, 0, n, index+1, 0, -1)
					}
				} else {
					if nums[index+1] > nums[index+2] && nums[index+1] > val {
						build(tree, 0, n, index+1, 0, 1)
					}
				}
			}
			if index-1 >= 0 && index+1 <= n-1 {
				if nums[index] > nums[index-1] && nums[index] > nums[index+1] {
					if val <= nums[index-1] || val <= nums[index+1] {
						build(tree, 0, n, index, 0, -1)
					}
				} else {
					if val > nums[index-1] && val > nums[index+1] {
						build(tree, 0, n, index, 0, 1)
					}
				}
			}
			nums[index] = val
		}
	}
	return res
}

func build(tree []int, s int, e int, index int, p int, add int) {
	tree[p] += add
	if s == e {
		return
	}
	mid := (s + e) / 2
	if mid >= index {
		build(tree, s, mid, index, 2*p+1, add)
	} else {
		build(tree, mid+1, e, index, 2*p+2, add)
	}

}

func find(tree []int, s int, e int, l int, r int, p int) int {
	if l > r {
		return 0
	}
	if s == e {
		return tree[p]
	}
	if s == l && e == r {
		return tree[p]
	}
	mid := (s + e) / 2
	if mid >= r {
		return find(tree, s, mid, l, r, 2*p+1)
	} else if mid < l {
		return find(tree, mid+1, e, l, r, 2*p+2)
	} else {
		return find(tree, s, mid, l, mid, 2*p+1) + find(tree, mid+1, e, mid+1, r, 2*p+2)
	}
}

func printTree(tree []int) {
	for _, item := range tree {
		print(item)
		print(" ")
	}
	println()
}

func main() {
	res := countOfPeaks([]int{3, 1, 4, 2, 5, 1, 3, 5, 4, 3, 4, 5, 6, 73, 4, 5, 6, 7, 3, 4, 5, 2, 3, 4, 6, 3, 4}, [][]int{{1, 0, 4}, {1, 3, 5}, {1, 11, 18}, {2, 6, 43}, {2, 10, 28}, {1, 0, 11}, {1, 0, 7}})
	printTree(res)
}
