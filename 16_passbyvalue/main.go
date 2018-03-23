//everything in Go is pass by value
//you always pass by value

package main

import "fmt"

func main() {
	age := 44
	fmt.Println(&age)

	changeMe(&age)

	fmt.Println(&age)
	fmt.Println(age)
}

func changeMe(z *int) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = 24 //assign the value 24 to the value at the address *z
	fmt.Println(z)
	fmt.Println(*z)

}
