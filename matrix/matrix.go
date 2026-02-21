package matrix

import "math/rand"
import "fmt"
import "math"

func MatrixSize(matrix [][]float64) (int, int) {
	rows := len(matrix)
	cols := len(matrix[0])
	return rows, cols
}

func GenerateSequencialVector(size int) []float64 {
	seq_vec := make([]float64, size)
	for i := range size {
		seq_vec[i] = float64(i + 1)
	}
	return seq_vec
}

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
	rows, cols := MatrixSize(matrix)
	t_matrix := GenerateZeroMatrix(cols, rows)
	for j := range cols {
		for i := range rows {
			t_matrix[j][i] = matrix[i][j]
		}
	}
	return t_matrix
}

func MatrixMultiplication(matrixA [][]float64, matrixB [][]float64) ([][]float64, error) {
	rowA, colA := MatrixSize(matrixA)
	rowB, colB := MatrixSize(matrixB)
	if colA != rowB {
		return nil, fmt.Errorf("cannot multiply a matrix (%d, %d) by a matrix (%d, %d)", rowA, colA, rowB, colB)
	}

	multMatrix := GenerateZeroMatrix(rowA, colB)
	for i := 0; i < rowA; i++ {
		for j := 0; j < colB; j++ {
			for k := 0; k < colA; k++ {
				multMatrix[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}
	return multMatrix, nil
}

// InvertMatrix computes the inverse of a square matrix using
// Gaussian elimination with partial pivoting. Returns an error
// if the matrix is singular or not square.
func InvertMatrix(m [][]float64) ([][]float64, error) {
	n, cols := MatrixSize(m)
	if n != cols {
		return nil, fmt.Errorf("cannot invert a non-square matrix (%d x %d)", n, cols)
	}

	// Build an augmented matrix [m | I].
	aug := GenerateZeroMatrix(n, 2*n)
	for i := range n {
		for j := range n {
			aug[i][j] = m[i][j]
		}
		aug[i][n+i] = 1.0 // identity block
	}

	// Forward elimination with partial pivoting.
	for col := range n {
		// Find the row with the largest absolute value in this column.
		pivot := col
		for row := col + 1; row < n; row++ {
			if math.Abs(aug[row][col]) > math.Abs(aug[pivot][col]) {
				pivot = row
			}
		}
		aug[col], aug[pivot] = aug[pivot], aug[col]

		if math.Abs(aug[col][col]) < 1e-12 {
			return nil, fmt.Errorf("matrix is singular and cannot be inverted")
		}

		// Scale the pivot row so the diagonal becomes 1.
		scale := aug[col][col]
		for j := range 2 * n {
			aug[col][j] /= scale
		}

		// Eliminate all other rows.
		for row := range n {
			if row == col {
				continue
			}
			factor := aug[row][col]
			for j := range 2 * n {
				aug[row][j] -= factor * aug[col][j]
			}
		}
	}

	// Extract the right-hand block as the inverse.
	inv := GenerateZeroMatrix(n, n)
	for i := range n {
		for j := range n {
			inv[i][j] = aug[i][n+j]
		}
	}
	return inv, nil
}

// ExpMatrix raises each element of a matrix to the given exponent (element-wise).
// For a true matrix inverse use InvertMatrix instead.
func ExpMatrix(m [][]float64, exponent float64) [][]float64 {
	nRow, nCol := MatrixSize(m)
	expMatrix := GenerateZeroMatrix(nRow, nCol)
	for i := range nRow {
		for j := range nCol {
			expMatrix[i][j] = math.Pow(m[i][j], exponent)
		}
	}
	return expMatrix
}

func PrettyPrint(matrix [][]float64) {
	for _, row := range matrix {
		fmt.Printf("[")
		for _, val := range row {
			fmt.Printf("%10.4f ", val)
		}
		fmt.Printf("]\n")
	}
	fmt.Println()
}

func FlattenMatrix(m [][]float64) []float64 {

	nRow, nCol := MatrixSize(m)
	flatVector := make([]float64, (nRow + nCol) - 1)

	for i := range nRow {
		for j := range nCol{
			flatVector[i + j] = m[i][j]
		}
	}

	return flatVector

}
