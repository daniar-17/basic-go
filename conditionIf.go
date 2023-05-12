package main

import (
	"fmt"
)

func main(){
	fmt.Println("Learning Condition")
	grade := 74
	if grade > 80 {
		fmt.Println("Grade : A")
	} else if (grade > 70 && grade <= 80) {
		fmt.Println("Grade : B")
	} else if (grade > 60 && grade <= 70) {
		fmt.Println("Grade : C")
	} else if (grade > 50 && grade <= 60) {
		fmt.Println("Grade : D")
	} else if (grade <= 50){
		fmt.Println("Grade : E")
	}
}