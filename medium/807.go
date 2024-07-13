package main

func maxIncreaseKeepingSkyline2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	west, east, north, south := make([][]int, m), make([][]int, m), make([][]int, m), make([][]int, m)
	for i := 0; i < m; i++ {
		west[i] = make([]int, n)
		east[i] = make([]int, n)
		south[i] = make([]int, n)
		north[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		west[i][0] = grid[i][0]
		for j := 1; j < n; j++ {
			west[i][j] = max(grid[i][j], west[i][j-1])
		}
	}

	for i := 0; i < m; i++ {
		east[i][n-1] = grid[i][n-1]
		for j := n - 2; j >= 0; j-- {
			east[i][j] = max(grid[i][j], east[i][j+1])
		}
	}

	for j := 0; j < n; j++ {
		north[0][j] = grid[0][j]
		for i := 1; i < m; i++ {
			north[i][j] = max(grid[i][j], north[i-1][j])
		}
	}

	for j := 0; j < n; j++ {
		south[m-1][j] = grid[m-1][j]
		for i := m - 2; i >= 0; i-- {
			south[i][j] = max(grid[i][j], south[i+1][j])
		}
	}

	res := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res += min(south[i][j], east[i][j], north[i][j], west[i][j]) - grid[i][j]
		}
	}
	return res

}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	west, north := make([]int, m), make([]int, n)

	for i := 0; i < m; i++ {
		west[i] = grid[i][0]
		for j := 1; j < n; j++ {
			west[i] = max(grid[i][j], west[i])
		}
	}

	for j := 0; j < n; j++ {
		north[j] = grid[0][j]
		for i := 1; i < m; i++ {
			north[j] = max(grid[i][j], north[j])
		}
	}

	res := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res += min(north[i], west[j]) - grid[i][j]
		}
	}
	return res

}

func main() {

}
