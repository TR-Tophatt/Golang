package main

import "fmt"

func main() {
	var potato int = 2

	switch potato {
	case 1:
		fmt.Println("Wassup Daniel")
		fallthrough
	case 2:
		fmt.Println("Wassup Medhi")
		fallthrough
	case 3:
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("have you no friends?")
	}
}
