package main

import (
	"fmt"
	"pertama/calculation"
)

func main(){
	var a, b int = 10, 12
	fmt.Println("Hello I am Daniar Nur Amin")
	resultMultiply := calculation.Multiply(a, b)
	resultAdd := calculation.Add(a, b)
	fmt.Println("Hasil Tambah : ", resultAdd)
	fmt.Println("Hasil Kali : ", resultMultiply)
}