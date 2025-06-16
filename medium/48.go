package medium

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			tmp := matrix[i][j]

			for x := 0; x < 4; x++ {
				ni, nj := i, j
				nexti := nj
				nextj := n - ni - 1
				matrix[nexti][nextj], tmp = tmp, matrix[nexti][nextj]
			}

		}
	}

	if n%2 == 1 {

	}

}
