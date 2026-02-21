package main

import (
	"fmt"
	"example.com/matrix"
	"example.com/polyfit"
)


func main(){
	fmt.Println("Hello!");

	m := matrix.GenerateRandomMatrix(2, 2);
	m[0][0] = 1
	m[0][1] = 2
	m[1][0] = 3
	m[1][1] = 4

	n := matrix.GenerateRandomMatrix(2, 2);
	n[0][0] = 1
	n[0][1] = 2
	n[1][0] = 3
	n[1][1] = 4

	_, err := matrix.MatrixMultiplication(m,n)

	if (err != nil) {
		panic(err)
	}

	values := matrix.GenerateSequencialVector(10)
	params, err := polyfit.PolynomialFit(values, 4)

	if (err != nil) {
		panic(err)
	}

	fmt.Println(values)
	fmt.Println(params)

}
