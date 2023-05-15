package main
import (
	"fmt"
)

func main()  {
	languages := [...]string{"Ruby", "Go", "Php", "Java", "C#", "Kontlin", "Swift", "Dart"}
	fmt.Println(languages)
	fmt.Println("Count Language : ",len(languages))
	for index, lang := range languages {
		fmt.Println("Index : ", index, " Language : ", lang)
	}
}