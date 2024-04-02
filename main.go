package main

import "fmt"

func main() {
	userNames := make([]string, 2, 3)
	userNames[0] = "Julia"
	userNames[1] = "Mary"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Leo")
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Veronica")
	fmt.Println(userNames)
}
