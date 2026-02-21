package polyfit

import (
	"errors"
	"math"
	"example.com/matrix"
)


func PolynomialFit(y []float64, polynomial_degree int) ([]float64, error) {

	if (polynomial_degree < 1) {
		return nil, errors.New("polynomial degree must be > 1")
	}
	if (polynomial_degree > len(y)) {
		return nil, errors.New("polynomial degree must be less than the size of y")
	}

	// Generate the Vandermonde matrix
	V := matrix.GenerateZeroMatrix(len(y), polynomial_degree)
	for i := range len(y) {
		V[i] = matrix.GenerateSequencialVector(polynomial_degree)
	}

	Vi, Vj := matrix.MatrixSize(V)
	for i := range Vi {
		for j := range Vj {
			V[i][j] = math.Pow(float64(i + 1), float64(j))
		}
	}

	// factor = (VT * V)^(-1)
	Vt := matrix.TransposeMatrix(V)

	alpha, err := matrix.MatrixMultiplication(Vt, V)
	if (err != nil) {
		panic(err)
	}

	factor, err := matrix.InvertMatrix(alpha)
	if (err != nil) {
		panic(err)
	}

	// factor * (VT * y)
	yMatrix := matrix.GenerateZeroMatrix(len(y), 1)
	for i := range len(y) {
		yMatrix[i][0] = y[i]
	}

	gamma, err := matrix.MatrixMultiplication(Vt, yMatrix)
	if (err != nil) {
		panic(err)
	}

	final_parameters, err := matrix.MatrixMultiplication(factor, gamma)
	if (err != nil) {
		panic(err)
	}

	return matrix.FlattenMatrix(final_parameters), nil

}
