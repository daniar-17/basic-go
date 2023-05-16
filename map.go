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

	//Declare Map
	secondMap := map[string]string{
		"Apple" : "is Apple",
		"Samsung" : "is Samsung",
		"Xiaomi" : "is Xiaomi",
	}
	fmt.Println("Looping with For")
	for key, value := range secondMap {
		fmt.Println("Key :", key,", Value :", value)
	}

	//Add Map
	secondMap["Asus"] = "is Asus"
	fmt.Println(secondMap)

	//Delete Map
	delete(secondMap, "Xiaomi")
	fmt.Println(secondMap)

	//Check Value Map
	value, isAvailable := secondMap["Nokia"]
	fmt.Println(value)
	fmt.Println(isAvailable)

	//Slice Map
	students := []map[string]string{
		{"name": "Agung", "Score":"A", "Study":"Math"},
		{"name": "Sarif", "Score":"B", "Study":"Algoritma"},
		{"name": "Agung", "Score":"C", "Study":"English"},
	}

	for _, student := range students{
		fmt.Println(student)
		fmt.Println("Name : ",student["name"], "-", "Study : ", student["Study"])
	}
}