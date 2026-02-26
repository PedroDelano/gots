package main

import (
	"example.com/matrix"
	"fmt"
)

func main() {
	fmt.Println("Hello!")
	m := matrix.New(4, 3)
	m.ScalarSum(10)
	m.Print()
	tm := m.Transpose()
	tm.Print()

	m3, _ := m.MatrixMult(tm)
	m3.Print()

}
