package matrix

import (
	"errors"
	"fmt"
)

type Matrix struct {
	nrow int
	ncol int
	data [][]float64
}

// Constructor

func New(nrow int, ncol int) Matrix {
	m := Matrix{nrow: nrow, ncol: ncol}
	m.Generate()
	return m
}

// Basic Operations

func (m *Matrix) Generate() {
	m.data = make([][]float64, m.nrow)
	for i := range m.data {
		m.data[i] = make([]float64, m.ncol)
	}
}

// Scalar Operations

func (m *Matrix) ScalarMult(x float64) {
	for i := range m.nrow {
		for j := range m.ncol {
			m.data[i][j] = m.data[i][j] * x
		}
	}
}

func (m *Matrix) ScalarSum(x float64) {
	for i := range m.nrow {
		for j := range m.ncol {
			m.data[i][j] = m.data[i][j] + x
		}
	}
}

// Matrix Operations

func (m *Matrix) Transpose() Matrix {

	// Define copy with transposed size
	tm := Matrix{nrow: m.ncol, ncol: m.nrow}
	tm.Generate()

	for i := range m.nrow {
		for j := range m.ncol {
			tm.data[j][i] = m.data[i][j]
		}
	}

	return tm
}

func (m1 *Matrix) MatrixMult(m2 Matrix) (Matrix, error) {

	if m1.ncol != m2.nrow {
		return Matrix{}, errors.New("matrix: can only multiply a matrix (m, n) by another (n, p)")
	}

	m3 := New(m1.nrow, m2.ncol)

	// could be m2.nrow as well
	n := m1.ncol

	for i := range m1.nrow {
		for j := range m2.ncol {
			curr := 0.0
			for k := range n {
				curr += (m1.data[i][k] * m2.data[k][j])
			}
			m3.data[i][j] = curr
		}
	}

	return m3, nil

}

// Utils

func (m Matrix) Print() {
	for _, row := range m.data {
		fmt.Printf("[")
		for _, val := range row {
			fmt.Printf("%10.4f ", val)
		}
		fmt.Printf("]\n")
	}
	fmt.Println()
}
