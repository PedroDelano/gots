package main

import (
	"fmt"
	"example.com/matrix"
)


func main() {
	fmt.Println("Hello!");
	m := matrix.New(4,3)
	m.ScalarSum(10)
	m.Print()
	tm := m.Transpose()
	tm.Print()
}
