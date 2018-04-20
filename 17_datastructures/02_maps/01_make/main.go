package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Bonjour!"
	myGreeting["Jimmy"] = "G'day!"

	fmt.Println(myGreeting)

	for key, val := range myGreeting {
		fmt.Println(key, "-", val)
	}
}
