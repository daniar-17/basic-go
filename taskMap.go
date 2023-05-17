package main
import "fmt"

func main(){
	fmt.Println("Task Map in Golang")
	fmt.Println("------------------")
	scores := []int{100, 80, 75, 92, 70, 93, 88, 67, 84, 24, 97}
	tempVal := 0;
	var goodScores []int
	for _, value := range scores {
		fmt.Println("Value :", value)
		tempVal = tempVal + value
		if(value >= 90){
			goodScores = append(goodScores, value)
		}
	}
	hasil := float64(tempVal/len(scores))
	fmt.Println("Average Value : ", hasil)
	fmt.Println("------------------")
	fmt.Println("Value >= 90 :")
	for _, value := range goodScores{
		fmt.Println(value)
	}
}