package main

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	m, n := len(board), len(board[0])
	moves := [8][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for _, mov := range moves {
		dx, dy := mov[0], mov[1]
		count := 0
		nx, ny := rMove, cMove
		for true {
			nx, ny = nx+dx, ny+dy
			count++
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				break
			} else if board[nx][ny] == '.' {
				break
			} else if board[nx][ny] != color {
				continue
			} else if board[nx][ny] == color && count >= 2 {
				return true
			} else {
				break
			}
		}
	}
	return false
}

func main() {

}
