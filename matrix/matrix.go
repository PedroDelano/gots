package matrix

import "math/rand"

func GenerateZeroMatrix(nRow int, nCol int) [][]float64 {
    matrix := make([][]float64, nRow)
    for i := range matrix {
        matrix[i] = make([]float64, nCol)
    }
    return matrix
}

func GenerateRandomMatrix(nRow int, nCol int) [][]float64 {
	m := GenerateZeroMatrix(nRow, nCol)
	for i := range nRow {
		for j := range nCol {
			m[i][j] = rand.Float64()
		}
	}
	return m
}

func TransposeMatrix(matrix [][]float64) [][]float64 {
	rows := len(matrix)
	cols := len(matrix[0])
	t_matrix := GenerateZeroMatrix(rows, cols)
	for j := range cols {
		for i := range rows {
			t_matrix[j][i] = matrix[i][j]
		}
	}
	return t_matrix
}
