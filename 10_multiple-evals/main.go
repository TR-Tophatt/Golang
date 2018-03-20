package main

import "fmt"

func main() {
	var potato int = 4

	switch potato {
	case 1, 5:
		fmt.Println("Wassup Daniel")
		fallthrough
	case 2, 7:
		fmt.Println("Wassup Medhi")
		fallthrough
	case 3, 4:
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("have you no friends?")
	}
}
