package main

import "fmt"

func main() {
	var potato = string(15)
	switch potato {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case string(15):
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("have you no friends?")
	}
}
