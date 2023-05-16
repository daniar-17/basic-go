package main
import (
	"fmt"
)

func main()  {
	fmt.Println("Learning Map")
	var myMap map[string]int
	myMap = map[string]int{}
	myMap["Ruby"] = 5
	myMap["JavaScript"] = 3
	myMap["Php"] = 2
	myMap["JavaScript"] = 6
	fmt.Println("All Result : ", myMap)
	fmt.Println("Spesifc Map : ",myMap["Php"])
}