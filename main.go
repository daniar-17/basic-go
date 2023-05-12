package main

import (
	"fmt"
	"pertama/calculation"
)

func main(){
	var a, b int = 10, 12
	a = 25

	var name string = "Who"
	name = "Daniar"

	fmt.Println("Hello I am ", name)
	resultMultiply := calculation.Multiply(a, b)
	resultAdd := calculation.Add(a, b)
	
	fmt.Println("Hasil Tambah : ", resultAdd)
	fmt.Println("Hasil Kali : ", resultMultiply)
}