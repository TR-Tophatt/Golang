package main

import "fmt"

func main() {
	greet("This is first argument", "This is Second argument")
}

func greet(args string, args2 string) {
	fmt.Println(args)
	fmt.Println(args2)
}
