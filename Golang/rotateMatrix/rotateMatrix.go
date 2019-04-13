package rotateMatrix

func rotateMatrix(matrix [][]int) [][]int {
	rowCount := len(matrix)

	colCount := len(matrix[0])

	if rowCount != colCount {
		return [][]int{}
	}

	for i := 0; i < rowCount/2; i++ {
		for j := i; j < rowCount-i-1; j++ {
			buf := matrix[i][j]
			matrix[i][j] = matrix[rowCount-1-j][i]
			matrix[rowCount-1-j][i] = matrix[rowCount-1-i][rowCount-1-j]
			matrix[rowCount-1-i][rowCount-1-j] = matrix[j][rowCount-1-i]
			matrix[j][rowCount-1-i] = buf

		}
	}

	return matrix
}
