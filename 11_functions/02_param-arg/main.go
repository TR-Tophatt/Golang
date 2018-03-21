package main

import "fmt"

func main() {
	greet("Jane")
	greet("John")
}

func greet(name string) {
	fmt.Println(name)
}

//greet is declared with a parameter of type string
//greet is called by passing in an argument "name"
