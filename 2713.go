package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Val   int
	Index int
}

func maxIncreasingCells(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	size := m * n
	nodeList := make([][]int, size)
	// 构图
	nodeMat1 := make([][]Node, m)
	nodeMat2 := make([][]Node, n)

	for i := 0; i < m; i++ {
		nodeMat1[i] = make([]Node, n)
		for j := 0; j < n; j++ {
			nodeMat1[i][j] = Node{
				Index: i*n + j,
				Val:   mat[i][j],
			}
		}
	}

	for i := 0; i < n; i++ {
		nodeMat2[i] = make([]Node, m)
		for j := 0; j < m; j++ {
			nodeMat2[i][j] = Node{
				Index: j*n + i,
				Val:   mat[j][i],
			}
		}
	}

	for i := 0; i < m; i++ {
		sort.Slice(nodeMat1[i], func(a, b int) bool {
			return nodeMat1[i][a].Val < nodeMat1[i][b].Val
		})
	}

	for i := 0; i < n; i++ {
		sort.Slice(nodeMat2[i], func(a, b int) bool {
			return nodeMat2[i][a].Val < nodeMat2[i][b].Val
		})
	}

	for i := 0; i < m; i++ {
		lastIndex := []int{}
		nextIndex := []int{nodeMat1[i][0].Index}
		for j := 1; j < n; j++ {
			if nodeMat1[i][j].Val > nodeMat1[i][j-1].Val {
				lastIndex = nextIndex
				nextIndex = []int{}
				for _, index := range lastIndex {
					nodeList[index] = append(nodeList[index], nodeMat1[i][j].Index)
				}
				nextIndex = append(nextIndex, nodeMat1[i][j].Index)
			} else {
				for _, index := range lastIndex {
					nodeList[index] = append(nodeList[index], nodeMat1[i][j].Index)
				}
				nextIndex = append(nextIndex, nodeMat1[i][j].Index)
			}
		}
	}

	for i := 0; i < n; i++ {
		lastIndex := []int{}
		nextIndex := []int{nodeMat2[i][0].Index}
		for j := 1; j < m; j++ {
			if nodeMat2[i][j].Val > nodeMat2[i][j-1].Val {
				lastIndex = nextIndex
				nextIndex = []int{}
				for _, index := range lastIndex {
					nodeList[index] = append(nodeList[index], nodeMat2[i][j].Index)
				}
				nextIndex = append(nextIndex, nodeMat2[i][j].Index)
			} else {
				for _, index := range lastIndex {
					nodeList[index] = append(nodeList[index], nodeMat2[i][j].Index)
				}
				nextIndex = append(nextIndex, nodeMat2[i][j].Index)
			}
		}
	}

	printMatrix(nodeList)

	visited := make([]bool, size)
	memory := make([]int, size)
	res := 0
	for i := 0; i < size; i++ {
		res = max(dfs(visited, nodeList, i, memory), res)
	}
	return res
}

func dfs(vis []bool, nodeList [][]int, index int, memory []int) int {
	if vis[index] {
		return memory[index]
	}
	vis[index] = true
	nodes := nodeList[index]
	res := 1
	for _, node := range nodes {
		if !vis[node] {
			res = max(res, dfs(vis, nodeList, node, memory)+1)
		} else {
			res = max(res, memory[node]+1)
		}
	}
	memory[index] = res
	return res
}

func printMatrix(mat [][]int) {
	for i, row := range mat {
		print(i, ": ")
		for _, col := range row {
			fmt.Print(col, " ")
		}
		print("\n")
	}
}

func main() {
	println(maxIncreasingCells([][]int{{3, 1, 6}, {-9, 5, 7}}))
	println(maxIncreasingCells([][]int{{9, -3, -2, 0, -3}, {-1, -4, 7, 5, 9}, {1, 8, 0, -7, -6}}))
}
