package main

import (
	"fmt"
	"example.com/matrix"
)


func main(){
	fmt.Println("Hello!");
	m := matrix.GenerateRandomMatrix(1000, 1000);
	fmt.Println(m);

	tm := matrix.TransposeMatrix(m);
	fmt.Println(tm);
}
