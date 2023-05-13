package main

import "fmt"

func main() {
	// var n int = 10
	// for i:=0; i < n; i++ {
	//   fmt.Print(i, ", ")
	// }

	// i:=0
	// for i <= n{
	//   fmt.Println("Learning Go : ",i)
	//   i++
	// }

	title := "Golang the best language"
  fmt.Println("Print Even")
	for index, letter := range title {
    if index % 2 == 0 {
      fmt.Println("Index : ", index," Letter : ",string(letter))
    }
	}
  fmt.Println()

  fmt.Println("Print Consonant Alfabet")
  for index, letter := range title {
    if (string(letter) == "a" || string(letter) == "i" || string(letter) == "u" || string(letter) == "e" || string(letter) == "o") {
      fmt.Println("Index : ", index," Letter : ",string(letter))
    }
	}
}