package zeroMatrix

type void struct{}

func setZeroMatrix(matrix [][]int) [][]int {
	rowIdx := map[int]void{}
	colIdx := map[int]void{}
	lrow := len(matrix)
	lcol := len(matrix[0])

	for i := 0; i < lrow; i++ {
		for j := 0; j < lcol; j++ {
			if matrix[i][j] == 0 {
				rowIdx[i] = void{}
				colIdx[j] = void{}
			}
		}
	}

	for key, _ := range rowIdx {
		for j := 0; j < lcol; j++ {
			matrix[key][j] = 0
		}
	}

	for key, _ := range colIdx {
		for i := 0; i < lrow; i++ {
			matrix[i][key] = 0
		}
	}
	return matrix

}
