package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Hello World"
	}
}

func main() {

	greeting := makeGreeter()
	fmt.Println(greeting())
}
