package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 3)
	userNames[0] = "Julia"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")
	//fmt.Println(userNames)

	//courseRating := make(map[string]float64, 3)
	courseRating := make(floatMap, 3)
	courseRating["Go"] = 4.7
	courseRating["React"] = 4.8
	courseRating["Angular"] = 5.2

	//courseRating.output()

	//fmt.Println(courseRating)

	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for key, value := range courseRating {
		fmt.Println(key)
		fmt.Println(value)
	}
}
